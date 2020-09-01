// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi

import (
	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// AzureProviders maps provider names (e.g. Compute) to versions in that providers and resources therein.
type AzureProviders = map[string]ProviderVersions

// ProviderVersions maps API Versions (e.g. 2020-08-01) to resources and invokes in that version.
type ProviderVersions = map[string]VersionResources

// VersionResources contains all resources and invokes in a given API version.
type VersionResources struct {
	Resources map[string]*ResourceSpec
	Invokes   map[string]*ResourceSpec
}

// ResourceSpec contains a pointer in an Open API Spec that defines a resource and related metadata.
type ResourceSpec struct {
	Path               string
	PathItem           *spec.PathItem
	Swagger            *Spec
	CompatibleVersions []string
}

// Providers finds all Azure Open API specs on disk, parses them, and creates in-memory representation of resources,
// collected per Azure Provider and API Version.
func Providers() AzureProviders {
	swaggerSpecLocations, err := swaggerLocations()
	if err != nil {
		panic(err)
	}

	var specs []*Spec
	for _, location := range swaggerSpecLocations {
		swagger, err := NewSpec(location)
		if err != nil {
			panic(err)
		}

		specs = append(specs, swagger)
	}

	providers := AzureProviders{}

	// Collect all versions for each path in the API across all Swagger files.
	for _, swagger := range specs {
		for path := range swagger.Paths.Paths {
			addApiPath(providers, path, swagger)
		}
	}

	for _, versionMap := range providers {
		// Add a `latest` version for each resource and invoke.
		latestResources := calculateLatestVersions(versionMap, false)
		latestInvokes := calculateLatestVersions(versionMap, true)
		versionMap["latest"] = VersionResources {
			Resources: latestResources,
			Invokes:   latestInvokes,
		}

		// Set compatible versions to all other versions of the resource with the same normalized API path.
		pathVersions := calculatePathVersions(versionMap)
		for version, items := range versionMap {
			for _, r := range items.Resources {
				var otherVersions []string
				for _, otherVersion := range pathVersions[normalizePath(r.Path)].SortedValues() {
					if otherVersion != version {
						otherVersions = append(otherVersions, otherVersion)
					}
				}
				r.CompatibleVersions = otherVersions
			}
		}
	}

	return providers
}

// swaggerLocations returns a slice of URLs of all known Azure Resource Manager swagger files.
func swaggerLocations() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	pattern := filepath.Join(dir, "/azure-rest-api-specs/specification/**/resource-manager/Microsoft.*/stable/*/*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	// Sorting alphabetically means the schemas with the latest API version are the last.
	sort.Strings(files)

	return files, nil
}

// addApiPath considers whether an API path contains resources and/or invokes and adds corresponding entries to the
// provider map. `providers` are mutated in-place.
func addApiPath(providers AzureProviders, path string, spec *Spec) {
	prov := provider.ResourceProvider(path)
	if prov == "" {
		return
	}

	// Find (or create) the version map with this name.
	versionMap, ok := providers[prov]
	if !ok {
		versionMap = map[string]VersionResources{}
		providers[prov] = versionMap
	}

	// Find (or create) the resource map with this name.
	apiVersion := "v" + strings.ReplaceAll(spec.Info.Version, "-", "")
	version, ok := versionMap[apiVersion]
	if !ok {
		version = VersionResources{
			Resources: map[string]*ResourceSpec{},
			Invokes:   map[string]*ResourceSpec{},
		}
		versionMap[apiVersion] = version
	}

	pathItem := spec.Paths.Paths[path]

	// Add a resource entry.
	if pathItem.Put != nil && pathItem.Get != nil && pathItem.Delete != nil {
		typeName := provider.ResourceName(pathItem.Get.ID)
		if typeName != "" {
			version.Resources[typeName] = &ResourceSpec{
				Path:     path,
				PathItem: &pathItem,
				Swagger:  spec,
			}
		}
	}

	// Add a list invoke entry.
	if pathItem.Post != nil {
		parts := strings.Split(path, "/")
		listOperation := parts[len(parts)-1]
		if !strings.HasPrefix(listOperation, "list") {
			return
		}
		typeName := provider.ResourceName(pathItem.Post.ID)
		if typeName != "" {
			version.Invokes[typeName] = &ResourceSpec{
				Path:     path,
				PathItem: &pathItem,
				Swagger:  spec,
			}
		}
	}
}

// calculateLatestVersions builds a map of latest versions per API paths from a map of all versions of a resource
// provider. The result is a map from a resource type name to resource specs.
func calculateLatestVersions(versionMap ProviderVersions, invokes bool) (latestResources map[string]*ResourceSpec)  {
	var versions []string
	for version := range versionMap {
		versions = append(versions, version)
	}
	sort.Strings(versions)

	pathTypeNames := map[string]string{}
	latestResources = map[string]*ResourceSpec{}
	for _, version := range versions {
		items := versionMap[version]
		resources := items.Resources
		if invokes {
			resources = items.Invokes
		}
		for typeName, r := range resources {
			normalizedPath := normalizePath(r.Path)
			previousTypeName := pathTypeNames[normalizedPath]
			if previousTypeName != "" && previousTypeName != typeName {
				delete(latestResources, previousTypeName)
			}

			pathTypeNames[normalizedPath] = typeName
			copyResource := *r
			latestResources[typeName] = &copyResource
		}
	}
	return
}

// calculatePathVersions builds a map of all versions defined for an API paths from a map of all versions of a resource
// provider. The result is a map from a normalized path to a set of versions for that path.
func calculatePathVersions(versionMap ProviderVersions) (pathVersions map[string]codegen.StringSet) {
	pathVersions = map[string]codegen.StringSet{}
	for version, items := range versionMap {
		for _, r := range items.Resources {
			normalizedPath := normalizePath(r.Path)
			versions, ok := pathVersions[normalizedPath]
			if !ok {
				versions = codegen.StringSet{}
				pathVersions[normalizedPath] = versions
			}
			versions.Add(version)
		}
	}
	return
}

// normalizePath converts an API path to its canonical form (lowercase, with all placeholders removed). The paths that
// convert to the same canonical path are considered to represent the same resource.
func normalizePath(path string) string {
	lowerPath := strings.ReplaceAll(strings.ToLower(path), "-", "")
	parts := strings.Split(lowerPath, "/")
	newParts := make([]string, len(parts))
	for i, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			newParts[i] = "{}"
		} else {
			newParts[i] = part
		}
	}
	return strings.Join(newParts, "/")
}
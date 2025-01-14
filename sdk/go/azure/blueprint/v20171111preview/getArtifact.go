// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Blueprint artifact.
//
// Deprecated: Please use one of the variants: PolicyAssignmentArtifact, RoleAssignmentArtifact, TemplateArtifact.
func LookupArtifact(ctx *pulumi.Context, args *LookupArtifactArgs, opts ...pulumi.InvokeOption) (*LookupArtifactResult, error) {
	var rv LookupArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactArgs struct {
	// name of the artifact.
	ArtifactName string `pulumi:"artifactName"`
	// name of the blueprint.
	BlueprintName string `pulumi:"blueprintName"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName string `pulumi:"managementGroupName"`
}

// Represents a Blueprint artifact.
type LookupArtifactResult struct {
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Specifies the kind of Blueprint artifact.
	Kind string `pulumi:"kind"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

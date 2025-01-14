// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a published Blueprint.
func LookupPublishedBlueprint(ctx *pulumi.Context, args *LookupPublishedBlueprintArgs, opts ...pulumi.InvokeOption) (*LookupPublishedBlueprintResult, error) {
	var rv LookupPublishedBlueprintResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getPublishedBlueprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublishedBlueprintArgs struct {
	// name of the blueprint.
	BlueprintName string `pulumi:"blueprintName"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName string `pulumi:"managementGroupName"`
	// version of the published blueprint.
	VersionId string `pulumi:"versionId"`
}

// Represents a published Blueprint.
type LookupPublishedBlueprintResult struct {
	// Name of the Blueprint definition.
	BlueprintName *string `pulumi:"blueprintName"`
	// Version-specific change notes
	ChangeNotes *string `pulumi:"changeNotes"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Parameters required by this Blueprint definition.
	Parameters map[string]ParameterDefinitionResponse `pulumi:"parameters"`
	// Resource group placeholders defined by this Blueprint definition.
	ResourceGroups map[string]ResourceGroupDefinitionResponse `pulumi:"resourceGroups"`
	// Status of the Blueprint. This field is readonly.
	Status BlueprintStatusResponse `pulumi:"status"`
	// The scope where this Blueprint can be applied.
	TargetScope *string `pulumi:"targetScope"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

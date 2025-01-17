// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Blueprint artifact applies Policy assignments.
func LookupPolicyAssignmentArtifact(ctx *pulumi.Context, args *LookupPolicyAssignmentArtifactArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentArtifactResult, error) {
	var rv LookupPolicyAssignmentArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getPolicyAssignmentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyAssignmentArtifactArgs struct {
	// name of the artifact.
	ArtifactName string `pulumi:"artifactName"`
	// name of the blueprint.
	BlueprintName string `pulumi:"blueprintName"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName string `pulumi:"managementGroupName"`
}

// Blueprint artifact applies Policy assignments.
type LookupPolicyAssignmentArtifactResult struct {
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []string `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Specifies the kind of Blueprint artifact.
	// Expected value is 'policyAssignment'.
	Kind string `pulumi:"kind"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Parameter values for the policy definition.
	Parameters map[string]ParameterValueBaseResponse `pulumi:"parameters"`
	// Azure resource ID of the policy definition.
	PolicyDefinitionId string `pulumi:"policyDefinitionId"`
	// Name of the resource group placeholder to which the policy will be assigned.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

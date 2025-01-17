// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Build step resource properties
func LookupBuildStep(ctx *pulumi.Context, args *LookupBuildStepArgs, opts ...pulumi.InvokeOption) (*LookupBuildStepResult, error) {
	var rv LookupBuildStepResult
	err := ctx.Invoke("azure-native:containerregistry/v20180201preview:getBuildStep", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBuildStepArgs struct {
	// The name of the container registry build task.
	BuildTaskName string `pulumi:"buildTaskName"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of a build step for a container registry build task.
	StepName string `pulumi:"stepName"`
}

// Build step resource properties
type LookupBuildStepResult struct {
	// The resource ID.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties of a build step.
	Properties DockerBuildStepResponse `pulumi:"properties"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

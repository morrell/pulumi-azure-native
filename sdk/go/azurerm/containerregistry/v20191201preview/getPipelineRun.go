// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPipelineRun(ctx *pulumi.Context, args *LookupPipelineRunArgs, opts ...pulumi.InvokeOption) (*LookupPipelineRunResult, error) {
	var rv LookupPipelineRunResult
	err := ctx.Invoke("azurerm:containerregistry/v20191201preview:getPipelineRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineRunArgs struct {
	// The name of the pipeline run.
	PipelineRunName string `pulumi:"pipelineRunName"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a pipeline run for a container registry.
type LookupPipelineRunResult struct {
	// How the pipeline run should be forced to recreate even if the pipeline run configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning state of a pipeline run.
	ProvisioningState string `pulumi:"provisioningState"`
	// The request parameters for a pipeline run.
	Request *PipelineRunRequestResponse `pulumi:"request"`
	// The response of a pipeline run.
	Response PipelineRunResponseResponse `pulumi:"response"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
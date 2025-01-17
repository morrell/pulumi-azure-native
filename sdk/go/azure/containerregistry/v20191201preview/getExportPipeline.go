// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents an export pipeline for a container registry.
func LookupExportPipeline(ctx *pulumi.Context, args *LookupExportPipelineArgs, opts ...pulumi.InvokeOption) (*LookupExportPipelineResult, error) {
	var rv LookupExportPipelineResult
	err := ctx.Invoke("azure-native:containerregistry/v20191201preview:getExportPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportPipelineArgs struct {
	// The name of the export pipeline.
	ExportPipelineName string `pulumi:"exportPipelineName"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents an export pipeline for a container registry.
type LookupExportPipelineResult struct {
	// The resource ID.
	Id string `pulumi:"id"`
	// The identity of the export pipeline.
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	// The location of the export pipeline.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The list of all options configured for the pipeline.
	Options []string `pulumi:"options"`
	// The provisioning state of the pipeline at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The target properties of the export pipeline.
	Target ExportPipelineTargetPropertiesResponse `pulumi:"target"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

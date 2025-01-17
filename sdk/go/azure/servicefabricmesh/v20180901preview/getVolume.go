// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This type describes a volume resource.
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:servicefabricmesh/v20180901preview:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The identity of the volume.
	VolumeResourceName string `pulumi:"volumeResourceName"`
}

// This type describes a volume resource.
type LookupVolumeResult struct {
	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters *VolumeProviderParametersAzureFileResponse `pulumi:"azureFileParameters"`
	// User readable description of the volume.
	Description *string `pulumi:"description"`
	// Fully qualified identifier for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provider of the volume.
	Provider string `pulumi:"provider"`
	// State of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Status of the volume.
	Status string `pulumi:"status"`
	// Gives additional information about the current status of the volume.
	StatusDetails string `pulumi:"statusDetails"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

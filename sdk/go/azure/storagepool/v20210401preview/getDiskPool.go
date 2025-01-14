// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response for Disk Pool request.
func LookupDiskPool(ctx *pulumi.Context, args *LookupDiskPoolArgs, opts ...pulumi.InvokeOption) (*LookupDiskPoolResult, error) {
	var rv LookupDiskPoolResult
	err := ctx.Invoke("azure-native:storagepool/v20210401preview:getDiskPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskPoolArgs struct {
	// The name of the Disk Pool.
	DiskPoolName string `pulumi:"diskPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for Disk Pool request.
type LookupDiskPoolResult struct {
	// List of additional capabilities for Disk Pool.
	AdditionalCapabilities []string `pulumi:"additionalCapabilities"`
	// Logical zone for Disk Pool resource; example: ["1"].
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// List of Azure Managed Disks to attach to a Disk Pool.
	Disks []DiskResponse `pulumi:"disks"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives.
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// State of the operation on the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Operational status of the Disk Pool.
	Status string `pulumi:"status"`
	// Azure Resource ID of a Subnet for the Disk Pool.
	SubnetId string `pulumi:"subnetId"`
	// Resource metadata required by ARM RPC
	SystemData SystemMetadataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Sku tier
	Tier *string `pulumi:"tier"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managednetwork

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Managed Network Group resource
// API Version: 2019-06-01-preview.
func LookupManagedNetworkGroup(ctx *pulumi.Context, args *LookupManagedNetworkGroupArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkGroupResult, error) {
	var rv LookupManagedNetworkGroupResult
	err := ctx.Invoke("azure-native:managednetwork:getManagedNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkGroupArgs struct {
	// The name of the Managed Network Group.
	ManagedNetworkGroupName string `pulumi:"managedNetworkGroupName"`
	// The name of the Managed Network.
	ManagedNetworkName string `pulumi:"managedNetworkName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Managed Network Group resource
type LookupManagedNetworkGroupResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Responsibility role under which this Managed Network Group will be created
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The collection of management groups covered by the Managed Network
	ManagementGroups []ResourceIdResponse `pulumi:"managementGroups"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The collection of  subnets covered by the Managed Network
	Subnets []ResourceIdResponse `pulumi:"subnets"`
	// The collection of subscriptions covered by the Managed Network
	Subscriptions []ResourceIdResponse `pulumi:"subscriptions"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// The collection of virtual nets covered by the Managed Network
	VirtualNetworks []ResourceIdResponse `pulumi:"virtualNetworks"`
}

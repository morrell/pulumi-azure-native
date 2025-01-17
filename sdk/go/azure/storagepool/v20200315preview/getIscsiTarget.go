// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200315preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response for iSCSI target requests.
func LookupIscsiTarget(ctx *pulumi.Context, args *LookupIscsiTargetArgs, opts ...pulumi.InvokeOption) (*LookupIscsiTargetResult, error) {
	var rv LookupIscsiTargetResult
	err := ctx.Invoke("azure-native:storagepool/v20200315preview:getIscsiTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiTargetArgs struct {
	// The name of the Disk pool.
	DiskPoolName string `pulumi:"diskPoolName"`
	// The name of the iSCSI target.
	IscsiTargetName string `pulumi:"iscsiTargetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for iSCSI target requests.
type LookupIscsiTargetResult struct {
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// State of the operation on the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Operational status of the iSCSI target.
	Status string `pulumi:"status"`
	// iSCSI target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn string `pulumi:"targetIqn"`
	// List of iSCSI target portal groups. Can have 1 portal group at most.
	Tpgs []TargetPortalGroupResponse `pulumi:"tpgs"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

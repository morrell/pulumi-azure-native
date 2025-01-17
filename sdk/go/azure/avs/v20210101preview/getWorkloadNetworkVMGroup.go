// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX VM Group
func LookupWorkloadNetworkVMGroup(ctx *pulumi.Context, args *LookupWorkloadNetworkVMGroupArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkVMGroupResult, error) {
	var rv LookupWorkloadNetworkVMGroupResult
	err := ctx.Invoke("azure-native:avs/v20210101preview:getWorkloadNetworkVMGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkVMGroupArgs struct {
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX VM Group identifier. Generally the same as the VM Group's display name
	VmGroupId string `pulumi:"vmGroupId"`
}

// NSX VM Group
type LookupWorkloadNetworkVMGroupResult struct {
	// Display name of the VM group.
	DisplayName *string `pulumi:"displayName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Virtual machine members of this group.
	Members []string `pulumi:"members"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// VM Group status.
	Status string `pulumi:"status"`
	// Resource type.
	Type string `pulumi:"type"`
}

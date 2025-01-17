// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Vendor resource.
func LookupVendor(ctx *pulumi.Context, args *LookupVendorArgs, opts ...pulumi.InvokeOption) (*LookupVendorResult, error) {
	var rv LookupVendorResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20210501:getVendor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVendorArgs struct {
	// The name of the vendor.
	VendorName string `pulumi:"vendorName"`
}

// Vendor resource.
type LookupVendorResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the vendor resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// A list of IDs of the vendor skus offered by the vendor.
	Skus []SubResourceResponse `pulumi:"skus"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

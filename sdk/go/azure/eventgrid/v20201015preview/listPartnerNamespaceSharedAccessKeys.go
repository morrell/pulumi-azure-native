// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Shared access keys of the partner namespace.
func ListPartnerNamespaceSharedAccessKeys(ctx *pulumi.Context, args *ListPartnerNamespaceSharedAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListPartnerNamespaceSharedAccessKeysResult, error) {
	var rv ListPartnerNamespaceSharedAccessKeysResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:listPartnerNamespaceSharedAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPartnerNamespaceSharedAccessKeysArgs struct {
	// Name of the partner namespace.
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Shared access keys of the partner namespace.
type ListPartnerNamespaceSharedAccessKeysResult struct {
	// Shared access key1 for the partner namespace.
	Key1 *string `pulumi:"key1"`
	// Shared access key2 for the partner namespace.
	Key2 *string `pulumi:"key2"`
}

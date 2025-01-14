// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Collection of the API key payload which is exposed in the response of the resource provider.
func ListTransactionNodeApiKeys(ctx *pulumi.Context, args *ListTransactionNodeApiKeysArgs, opts ...pulumi.InvokeOption) (*ListTransactionNodeApiKeysResult, error) {
	var rv ListTransactionNodeApiKeysResult
	err := ctx.Invoke("azure-native:blockchain/v20180601preview:listTransactionNodeApiKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTransactionNodeApiKeysArgs struct {
	// Blockchain member name.
	BlockchainMemberName string `pulumi:"blockchainMemberName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Transaction node name.
	TransactionNodeName string `pulumi:"transactionNodeName"`
}

// Collection of the API key payload which is exposed in the response of the resource provider.
type ListTransactionNodeApiKeysResult struct {
	// Gets or sets the collection of API key.
	Keys []ApiKeyResponse `pulumi:"keys"`
}

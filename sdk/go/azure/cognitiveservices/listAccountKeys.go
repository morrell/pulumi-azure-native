// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognitiveservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The access keys for the cognitive services account.
// API Version: 2017-04-18.
func ListAccountKeys(ctx *pulumi.Context, args *ListAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListAccountKeysResult, error) {
	var rv ListAccountKeysResult
	err := ctx.Invoke("azure-native:cognitiveservices:listAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountKeysArgs struct {
	// The name of Cognitive Services account.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The access keys for the cognitive services account.
type ListAccountKeysResult struct {
	// Gets the value of key 1.
	Key1 *string `pulumi:"key1"`
	// Gets the value of key 2.
	Key2 *string `pulumi:"key2"`
}

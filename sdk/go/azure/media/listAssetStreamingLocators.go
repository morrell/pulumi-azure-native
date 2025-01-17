// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Streaming Locators associated with this Asset.
// API Version: 2020-05-01.
func ListAssetStreamingLocators(ctx *pulumi.Context, args *ListAssetStreamingLocatorsArgs, opts ...pulumi.InvokeOption) (*ListAssetStreamingLocatorsResult, error) {
	var rv ListAssetStreamingLocatorsResult
	err := ctx.Invoke("azure-native:media:listAssetStreamingLocators", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAssetStreamingLocatorsArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Streaming Locators associated with this Asset.
type ListAssetStreamingLocatorsResult struct {
	// The list of Streaming Locators.
	StreamingLocators []AssetStreamingLocatorResponse `pulumi:"streamingLocators"`
}

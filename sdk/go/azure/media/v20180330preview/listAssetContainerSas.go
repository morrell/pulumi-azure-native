// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180330preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Asset Storage container SAS URLs.
func ListAssetContainerSas(ctx *pulumi.Context, args *ListAssetContainerSasArgs, opts ...pulumi.InvokeOption) (*ListAssetContainerSasResult, error) {
	var rv ListAssetContainerSasResult
	err := ctx.Invoke("azure-native:media/v20180330preview:listAssetContainerSas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAssetContainerSasArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The SAS URL expiration time.  This must be less than 24 hours from the current time.
	ExpiryTime *string `pulumi:"expiryTime"`
	// The permissions to set on the SAS URL.
	Permissions *string `pulumi:"permissions"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Asset Storage container SAS URLs.
type ListAssetContainerSasResult struct {
	// The list of Asset container SAS URLs.
	AssetContainerSasUrls []string `pulumi:"assetContainerSasUrls"`
}

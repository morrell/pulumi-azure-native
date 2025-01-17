// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The SAS response that contains the storage account, container and associated SAS token for connection use.
func ListAccountSasTokens(ctx *pulumi.Context, args *ListAccountSasTokensArgs, opts ...pulumi.InvokeOption) (*ListAccountSasTokensResult, error) {
	var rv ListAccountSasTokensResult
	err := ctx.Invoke("azure-native:datalakeanalytics/v20151001preview:listAccountSasTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountSasTokensArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName string `pulumi:"accountName"`
	// The name of the Azure storage container for which the SAS token is being requested.
	ContainerName string `pulumi:"containerName"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure storage account for which the SAS token is being requested.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// The SAS response that contains the storage account, container and associated SAS token for connection use.
type ListAccountSasTokensResult struct {
	// The link (url) to the next page of results.
	NextLink string `pulumi:"nextLink"`
	// The results of the list operation.
	Value []SasTokenInformationResponse `pulumi:"value"`
}

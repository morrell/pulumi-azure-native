// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An ADLS Gen 2 file data set.
func LookupADLSGen2FileDataSet(ctx *pulumi.Context, args *LookupADLSGen2FileDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileDataSetResult, error) {
	var rv LookupADLSGen2FileDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20200901:getADLSGen2FileDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName string `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// An ADLS Gen 2 file data set.
type LookupADLSGen2FileDataSetResult struct {
	// Unique id for identifying a data set resource
	DataSetId string `pulumi:"dataSetId"`
	// File path within the file system.
	FilePath string `pulumi:"filePath"`
	// File system to which the file belongs.
	FileSystem string `pulumi:"fileSystem"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set.
	// Expected value is 'AdlsGen2File'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Resource group of storage account
	ResourceGroup string `pulumi:"resourceGroup"`
	// Storage account name of the source data set
	StorageAccountName string `pulumi:"storageAccountName"`
	// Subscription id of storage account
	SubscriptionId string `pulumi:"subscriptionId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

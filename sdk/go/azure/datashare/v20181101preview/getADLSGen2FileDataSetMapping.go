// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An ADLS Gen2 file data set mapping.
func LookupADLSGen2FileDataSetMapping(ctx *pulumi.Context, args *LookupADLSGen2FileDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileDataSetMappingResult, error) {
	var rv LookupADLSGen2FileDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getADLSGen2FileDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName string `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// An ADLS Gen2 file data set mapping.
type LookupADLSGen2FileDataSetMappingResult struct {
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	// File path within the file system.
	FilePath string `pulumi:"filePath"`
	// File system to which the file belongs.
	FileSystem string `pulumi:"fileSystem"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set mapping.
	// Expected value is 'AdlsGen2File'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Type of output file
	OutputType *string `pulumi:"outputType"`
	// Provisioning state of the data set mapping.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource group of storage account.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Storage account name of the source data set.
	StorageAccountName string `pulumi:"storageAccountName"`
	// Subscription id of storage account.
	SubscriptionId string `pulumi:"subscriptionId"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

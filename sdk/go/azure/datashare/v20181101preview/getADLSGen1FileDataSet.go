// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An ADLS Gen 1 file data set.
func LookupADLSGen1FileDataSet(ctx *pulumi.Context, args *LookupADLSGen1FileDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen1FileDataSetResult, error) {
	var rv LookupADLSGen1FileDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getADLSGen1FileDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen1FileDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName string `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// An ADLS Gen 1 file data set.
type LookupADLSGen1FileDataSetResult struct {
	// The ADLS account name.
	AccountName string `pulumi:"accountName"`
	// Unique id for identifying a data set resource
	DataSetId string `pulumi:"dataSetId"`
	// The file name in the ADLS account.
	FileName string `pulumi:"fileName"`
	// The folder path within the ADLS account.
	FolderPath string `pulumi:"folderPath"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set.
	// Expected value is 'AdlsGen1File'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Resource group of ADLS account.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Subscription id of ADLS account.
	SubscriptionId string `pulumi:"subscriptionId"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

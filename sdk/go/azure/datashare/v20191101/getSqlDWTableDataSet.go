// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A SQL DW table data set.
func LookupSqlDWTableDataSet(ctx *pulumi.Context, args *LookupSqlDWTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSqlDWTableDataSetResult, error) {
	var rv LookupSqlDWTableDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getSqlDWTableDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDWTableDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName string `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// A SQL DW table data set.
type LookupSqlDWTableDataSetResult struct {
	// Unique id for identifying a data set resource
	DataSetId string `pulumi:"dataSetId"`
	// DataWarehouse name of the source data set
	DataWarehouseName string `pulumi:"dataWarehouseName"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set.
	// Expected value is 'SqlDWTable'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Schema of the table. Default value is dbo.
	SchemaName string `pulumi:"schemaName"`
	// Resource id of SQL server
	SqlServerResourceId string `pulumi:"sqlServerResourceId"`
	// SQL DW table name.
	TableName string `pulumi:"tableName"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

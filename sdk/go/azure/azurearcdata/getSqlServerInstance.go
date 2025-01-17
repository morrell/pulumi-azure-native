// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azurearcdata

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A SqlServerInstance.
// API Version: 2021-06-01-preview.
func LookupSqlServerInstance(ctx *pulumi.Context, args *LookupSqlServerInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerInstanceResult, error) {
	var rv LookupSqlServerInstanceResult
	err := ctx.Invoke("azure-native:azurearcdata:getSqlServerInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerInstanceArgs struct {
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of SQL Server Instance
	SqlServerInstanceName string `pulumi:"sqlServerInstanceName"`
}

// A SqlServerInstance.
type LookupSqlServerInstanceResult struct {
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// null
	Properties SqlServerInstancePropertiesResponse `pulumi:"properties"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

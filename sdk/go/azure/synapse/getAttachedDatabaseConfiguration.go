// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing an attached database configuration.
// API Version: 2021-06-01-preview.
func LookupAttachedDatabaseConfiguration(ctx *pulumi.Context, args *LookupAttachedDatabaseConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAttachedDatabaseConfigurationResult, error) {
	var rv LookupAttachedDatabaseConfigurationResult
	err := ctx.Invoke("azure-native:synapse:getAttachedDatabaseConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttachedDatabaseConfigurationArgs struct {
	// The name of the attached database configuration.
	AttachedDatabaseConfigurationName string `pulumi:"attachedDatabaseConfigurationName"`
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// Class representing an attached database configuration.
type LookupAttachedDatabaseConfigurationResult struct {
	// The list of databases from the clusterResourceId which are currently attached to the kusto pool.
	AttachedDatabaseNames []string `pulumi:"attachedDatabaseNames"`
	// The resource id of the kusto pool where the databases you would like to attach reside.
	ClusterResourceId string `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName string `pulumi:"databaseName"`
	// The default principals modification kind
	DefaultPrincipalsModificationKind string `pulumi:"defaultPrincipalsModificationKind"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Table level sharing specifications
	TableLevelSharingProperties *TableLevelSharingPropertiesResponse `pulumi:"tableLevelSharingProperties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

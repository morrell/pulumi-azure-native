// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing an Event Grid data connection.
func LookupEventGridDataConnection(ctx *pulumi.Context, args *LookupEventGridDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventGridDataConnectionResult, error) {
	var rv LookupEventGridDataConnectionResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getEventGridDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventGridDataConnectionArgs struct {
	// The name of the data connection.
	DataConnectionName string `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto pool.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// Class representing an Event Grid data connection.
type LookupEventGridDataConnectionResult struct {
	// The name of blob storage event type to process.
	BlobStorageEventType *string `pulumi:"blobStorageEventType"`
	// The event hub consumer group.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat *string `pulumi:"dataFormat"`
	// The resource ID where the event grid is configured to send events.
	EventHubResourceId string `pulumi:"eventHubResourceId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// A Boolean value that, if set to true, indicates that ingestion should ignore the first record of every file
	IgnoreFirstRecord *bool `pulumi:"ignoreFirstRecord"`
	// Kind of the endpoint for the data connection
	// Expected value is 'EventGrid'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource ID of the storage account where the data resides.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

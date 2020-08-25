// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupEventHub(ctx *pulumi.Context, args *LookupEventHubArgs, opts ...pulumi.InvokeOption) (*LookupEventHubResult, error) {
	var rv LookupEventHubResult
	err := ctx.Invoke("azurerm:eventhub/v20170401:getEventHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubArgs struct {
	// The Event Hub name
	Name string `pulumi:"name"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in List or Get Event Hub operation
type LookupEventHubResult struct {
	// Properties of capture description
	CaptureDescription *CaptureDescriptionResponse `pulumi:"captureDescription"`
	// Exact time the Event Hub was created.
	CreatedAt string `pulumi:"createdAt"`
	// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *int `pulumi:"messageRetentionInDays"`
	// Resource name.
	Name string `pulumi:"name"`
	// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *int `pulumi:"partitionCount"`
	// Current number of shards on the Event Hub.
	PartitionIds []string `pulumi:"partitionIds"`
	// Enumerates the possible values for the status of the Event Hub.
	Status *string `pulumi:"status"`
	// Resource type.
	Type string `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt string `pulumi:"updatedAt"`
}
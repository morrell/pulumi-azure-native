// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The result of a request to list events for a webhook.
func ListWebhookEvents(ctx *pulumi.Context, args *ListWebhookEventsArgs, opts ...pulumi.InvokeOption) (*ListWebhookEventsResult, error) {
	var rv ListWebhookEventsResult
	err := ctx.Invoke("azure-native:containerregistry/v20210601preview:listWebhookEvents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebhookEventsArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the webhook.
	WebhookName string `pulumi:"webhookName"`
}

// The result of a request to list events for a webhook.
type ListWebhookEventsResult struct {
	// The URI that can be used to request the next list of events.
	NextLink *string `pulumi:"nextLink"`
	// The list of events. Since this list may be incomplete, the nextLink field should be used to request the next list of events.
	Value []EventResponse `pulumi:"value"`
}

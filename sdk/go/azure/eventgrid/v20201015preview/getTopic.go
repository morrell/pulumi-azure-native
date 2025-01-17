// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EventGrid Topic
func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicArgs struct {
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the topic.
	TopicName string `pulumi:"topicName"`
}

// EventGrid Topic
type LookupTopicResult struct {
	// Endpoint for the topic.
	Endpoint string `pulumi:"endpoint"`
	// Extended location of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Identity information for the resource.
	Identity *IdentityInfoResponse `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRuleResponse `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Metric resource id for the topic.
	MetricResourceId string `pulumi:"metricResourceId"`
	// Name of the resource.
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the topic.
	ProvisioningState string `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The Sku pricing tier for the topic.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// The system metadata relating to Topic resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a namespace resource.
// API Version: 2017-04-01.
func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:servicebus:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a namespace resource.
type LookupNamespaceResult struct {
	// The time the namespace was created.
	CreatedAt string `pulumi:"createdAt"`
	// Resource Id
	Id string `pulumi:"id"`
	// The Geo-location where the resource lives
	Location string `pulumi:"location"`
	// Identifier for Azure Insights metrics
	MetricId string `pulumi:"metricId"`
	// Resource name
	Name string `pulumi:"name"`
	// Provisioning state of the namespace.
	ProvisioningState string `pulumi:"provisioningState"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint string `pulumi:"serviceBusEndpoint"`
	// Properties of Sku
	Sku *SBSkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

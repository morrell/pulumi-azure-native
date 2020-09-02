// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPartnerNamespace(ctx *pulumi.Context, args *LookupPartnerNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupPartnerNamespaceResult, error) {
	var rv LookupPartnerNamespaceResult
	err := ctx.Invoke("azurerm:eventgrid/v20200401preview:getPartnerNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerNamespaceArgs struct {
	// Name of the partner namespace.
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// EventGrid Partner Namespace.
type LookupPartnerNamespaceResult struct {
	// Endpoint for the partner namespace.
	Endpoint string `pulumi:"endpoint"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Name of the resource
	Name string `pulumi:"name"`
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId *string `pulumi:"partnerRegistrationFullyQualifiedId"`
	// Provisioning state of the partner namespace.
	ProvisioningState string `pulumi:"provisioningState"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type string `pulumi:"type"`
}
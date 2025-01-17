// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hub resource.
func LookupHub(ctx *pulumi.Context, args *LookupHubArgs, opts ...pulumi.InvokeOption) (*LookupHubResult, error) {
	var rv LookupHubResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Hub resource.
type LookupHubResult struct {
	// API endpoint URL of the hub.
	ApiEndpoint string `pulumi:"apiEndpoint"`
	// Billing settings of the hub.
	HubBillingInfo *HubBillingInfoFormatResponse `pulumi:"hubBillingInfo"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Provisioning state of the hub.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The bit flags for enabled hub features. Bit 0 is set to 1 indicates graph is enabled, or disabled if set to 0. Bit 1 is set to 1 indicates the hub is disabled, or enabled if set to 0.
	TenantFeatures *int `pulumi:"tenantFeatures"`
	// Resource type.
	Type string `pulumi:"type"`
	// Web endpoint URL of the hub.
	WebEndpoint string `pulumi:"webEndpoint"`
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIntegrationAccount(ctx *pulumi.Context, args *LookupIntegrationAccountArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountResult, error) {
	var rv LookupIntegrationAccountResult
	err := ctx.Invoke("azurerm:logic/v20180701preview:getIntegrationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountArgs struct {
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The integration account.
type LookupIntegrationAccountResult struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// The sku.
	Sku *IntegrationAccountSkuResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}
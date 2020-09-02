// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDigitalTwin(ctx *pulumi.Context, args *LookupDigitalTwinArgs, opts ...pulumi.InvokeOption) (*LookupDigitalTwinResult, error) {
	var rv LookupDigitalTwinResult
	err := ctx.Invoke("azurerm:digitaltwins/v20200301preview:getDigitalTwin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDigitalTwinArgs struct {
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName string `pulumi:"resourceName"`
}

// The description of the DigitalTwins service.
type LookupDigitalTwinResult struct {
	// Time when DigitalTwinsInstance was created.
	CreatedTime string `pulumi:"createdTime"`
	// Api endpoint to work with DigitalTwinsInstance.
	HostName string `pulumi:"hostName"`
	// Time when DigitalTwinsInstance was created.
	LastUpdatedTime string `pulumi:"lastUpdatedTime"`
	// The resource location.
	Location string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// This property is reserved for future use, and will be ignored/omitted
	Sku *DigitalTwinsSkuInfoResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}
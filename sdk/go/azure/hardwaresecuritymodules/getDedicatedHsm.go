// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hardwaresecuritymodules

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource information with extended details.
// API Version: 2018-10-31-preview.
func LookupDedicatedHsm(ctx *pulumi.Context, args *LookupDedicatedHsmArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHsmResult, error) {
	var rv LookupDedicatedHsmResult
	err := ctx.Invoke("azure-native:hardwaresecuritymodules:getDedicatedHsm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHsmArgs struct {
	// The name of the dedicated HSM.
	Name string `pulumi:"name"`
	// The name of the Resource Group to which the dedicated hsm belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Resource information with extended details.
type LookupDedicatedHsmResult struct {
	// The Azure Resource Manager resource ID for the dedicated HSM.
	Id string `pulumi:"id"`
	// The supported Azure location where the dedicated HSM should be created.
	Location string `pulumi:"location"`
	// The name of the dedicated HSM.
	Name string `pulumi:"name"`
	// Specifies the network interfaces of the dedicated hsm.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// SKU details
	Sku SkuResponse `pulumi:"sku"`
	// This field will be used when RP does not support Availability zones.
	StampId *string `pulumi:"stampId"`
	// Resource Status Message.
	StatusMessage string `pulumi:"statusMessage"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The resource type of the dedicated HSM.
	Type string `pulumi:"type"`
	// The Dedicated Hsm zones.
	Zones []string `pulumi:"zones"`
}

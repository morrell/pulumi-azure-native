// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// String dictionary resource
func ListSiteAppSettingsSlot(ctx *pulumi.Context, args *ListSiteAppSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteAppSettingsSlotResult, error) {
	var rv ListSiteAppSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteAppSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteAppSettingsSlotArgs struct {
	// Name of web app
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
}

// String dictionary resource
type ListSiteAppSettingsSlotResult struct {
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Settings
	Properties map[string]string `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

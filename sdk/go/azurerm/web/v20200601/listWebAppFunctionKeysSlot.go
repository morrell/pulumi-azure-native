// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWebAppFunctionKeysSlot(ctx *pulumi.Context, args *ListWebAppFunctionKeysSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionKeysSlotResult, error) {
	var rv ListWebAppFunctionKeysSlotResult
	err := ctx.Invoke("azurerm:web/v20200601:listWebAppFunctionKeysSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionKeysSlotArgs struct {
	// Function name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot.
	Slot string `pulumi:"slot"`
}

// String dictionary resource.
type ListWebAppFunctionKeysSlotResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}
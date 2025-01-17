// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Function secrets.
// API Version: 2020-12-01.
func ListWebAppSyncFunctionTriggers(ctx *pulumi.Context, args *ListWebAppSyncFunctionTriggersArgs, opts ...pulumi.InvokeOption) (*ListWebAppSyncFunctionTriggersResult, error) {
	var rv ListWebAppSyncFunctionTriggersResult
	err := ctx.Invoke("azure-native:web:listWebAppSyncFunctionTriggers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSyncFunctionTriggersArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Function secrets.
type ListWebAppSyncFunctionTriggersResult struct {
	// Secret key.
	Key *string `pulumi:"key"`
	// Trigger URL.
	TriggerUrl *string `pulumi:"triggerUrl"`
}

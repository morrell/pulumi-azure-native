// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package powerbi

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version: 2016-01-29.
func ListWorkspaceCollectionAccessKeys(ctx *pulumi.Context, args *ListWorkspaceCollectionAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceCollectionAccessKeysResult, error) {
	var rv ListWorkspaceCollectionAccessKeysResult
	err := ctx.Invoke("azure-native:powerbi:listWorkspaceCollectionAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceCollectionAccessKeysArgs struct {
	// Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Power BI Embedded Workspace Collection name
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

type ListWorkspaceCollectionAccessKeysResult struct {
	// Access key 1
	Key1 *string `pulumi:"key1"`
	// Access key 2
	Key2 *string `pulumi:"key2"`
}

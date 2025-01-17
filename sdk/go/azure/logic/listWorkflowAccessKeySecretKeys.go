// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version: 2015-02-01-preview.
func ListWorkflowAccessKeySecretKeys(ctx *pulumi.Context, args *ListWorkflowAccessKeySecretKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkflowAccessKeySecretKeysResult, error) {
	var rv ListWorkflowAccessKeySecretKeysResult
	err := ctx.Invoke("azure-native:logic:listWorkflowAccessKeySecretKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowAccessKeySecretKeysArgs struct {
	// The workflow access key name.
	AccessKeyName string `pulumi:"accessKeyName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

type ListWorkflowAccessKeySecretKeysResult struct {
	// Gets the primary secret key.
	PrimarySecretKey string `pulumi:"primarySecretKey"`
	// Gets the secondary secret key.
	SecondarySecretKey string `pulumi:"secondarySecretKey"`
}

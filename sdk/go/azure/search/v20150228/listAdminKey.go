// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150228

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response containing the primary and secondary API keys for a given Azure Search service.
func ListAdminKey(ctx *pulumi.Context, args *ListAdminKeyArgs, opts ...pulumi.InvokeOption) (*ListAdminKeyResult, error) {
	var rv ListAdminKeyResult
	err := ctx.Invoke("azure-native:search/v20150228:listAdminKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAdminKeyArgs struct {
	// The name of the resource group within the current subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Search service for which to list admin keys.
	ServiceName string `pulumi:"serviceName"`
}

// Response containing the primary and secondary API keys for a given Azure Search service.
type ListAdminKeyResult struct {
	// The primary API key of the Search service.
	PrimaryKey string `pulumi:"primaryKey"`
	// The secondary API key of the Search service.
	SecondaryKey string `pulumi:"secondaryKey"`
}

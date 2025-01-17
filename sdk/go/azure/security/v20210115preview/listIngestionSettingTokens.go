// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210115preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures how to correlate scan data and logs with resources associated with the subscription.
func ListIngestionSettingTokens(ctx *pulumi.Context, args *ListIngestionSettingTokensArgs, opts ...pulumi.InvokeOption) (*ListIngestionSettingTokensResult, error) {
	var rv ListIngestionSettingTokensResult
	err := ctx.Invoke("azure-native:security/v20210115preview:listIngestionSettingTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIngestionSettingTokensArgs struct {
	// Name of the ingestion setting
	IngestionSettingName string `pulumi:"ingestionSettingName"`
}

// Configures how to correlate scan data and logs with resources associated with the subscription.
type ListIngestionSettingTokensResult struct {
	// The token is used for correlating security data and logs with the resources in the subscription.
	Token string `pulumi:"token"`
}

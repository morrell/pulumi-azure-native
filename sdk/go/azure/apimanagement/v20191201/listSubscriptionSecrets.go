// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Subscription keys.
func ListSubscriptionSecrets(ctx *pulumi.Context, args *ListSubscriptionSecretsArgs, opts ...pulumi.InvokeOption) (*ListSubscriptionSecretsResult, error) {
	var rv ListSubscriptionSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:listSubscriptionSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubscriptionSecretsArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid string `pulumi:"sid"`
}

// Subscription keys.
type ListSubscriptionSecretsResult struct {
	// Subscription primary key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Subscription secondary key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

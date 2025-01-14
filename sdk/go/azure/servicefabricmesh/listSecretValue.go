// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabricmesh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This type represents the unencrypted value of the secret.
// API Version: 2018-09-01-preview.
func ListSecretValue(ctx *pulumi.Context, args *ListSecretValueArgs, opts ...pulumi.InvokeOption) (*ListSecretValueResult, error) {
	var rv ListSecretValueResult
	err := ctx.Invoke("azure-native:servicefabricmesh:listSecretValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSecretValueArgs struct {
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the secret resource.
	SecretResourceName string `pulumi:"secretResourceName"`
	// The name of the secret resource value which is typically the version identifier for the value.
	SecretValueResourceName string `pulumi:"secretValueResourceName"`
}

// This type represents the unencrypted value of the secret.
type ListSecretValueResult struct {
	// The actual value of the secret.
	Value *string `pulumi:"value"`
}

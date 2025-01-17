// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Keys for endpoint authentication.
func ListBatchEndpointKeys(ctx *pulumi.Context, args *ListBatchEndpointKeysArgs, opts ...pulumi.InvokeOption) (*ListBatchEndpointKeysResult, error) {
	var rv ListBatchEndpointKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:listBatchEndpointKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBatchEndpointKeysArgs struct {
	// Inference Endpoint name.
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Keys for endpoint authentication.
type ListBatchEndpointKeysResult struct {
	// The primary key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The secondary key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

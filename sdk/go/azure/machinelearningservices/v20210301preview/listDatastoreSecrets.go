// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Base definition for datastore secrets.
func ListDatastoreSecrets(ctx *pulumi.Context, args *ListDatastoreSecretsArgs, opts ...pulumi.InvokeOption) (*ListDatastoreSecretsResult, error) {
	var rv ListDatastoreSecretsResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:listDatastoreSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatastoreSecretsArgs struct {
	// Datastore name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Base definition for datastore secrets.
type ListDatastoreSecretsResult struct {
	// Credential type used to authentication with storage.
	SecretsType string `pulumi:"secretsType"`
}

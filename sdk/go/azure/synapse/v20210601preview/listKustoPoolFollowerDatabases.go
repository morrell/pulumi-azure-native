// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The list Kusto database principals operation response.
func ListKustoPoolFollowerDatabases(ctx *pulumi.Context, args *ListKustoPoolFollowerDatabasesArgs, opts ...pulumi.InvokeOption) (*ListKustoPoolFollowerDatabasesResult, error) {
	var rv ListKustoPoolFollowerDatabasesResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:listKustoPoolFollowerDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKustoPoolFollowerDatabasesArgs struct {
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// The list Kusto database principals operation response.
type ListKustoPoolFollowerDatabasesResult struct {
	// The list of follower database result.
	Value []FollowerDatabaseDefinitionResponse `pulumi:"value"`
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150408

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDatabaseAccountCassandraTable(ctx *pulumi.Context, args *LookupDatabaseAccountCassandraTableArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountCassandraTableResult, error) {
	var rv LookupDatabaseAccountCassandraTableResult
	err := ctx.Invoke("azurerm:documentdb/v20150408:getDatabaseAccountCassandraTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName string `pulumi:"keyspaceName"`
	// Cosmos DB table name.
	Name string `pulumi:"name"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB Cassandra table.
type LookupDatabaseAccountCassandraTableResult struct {
	// Time to live of the Cosmos DB Cassandra table
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name string `pulumi:"name"`
	// Schema of the Cosmos DB Cassandra table
	Schema *CassandraSchemaResponse `pulumi:"schema"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}
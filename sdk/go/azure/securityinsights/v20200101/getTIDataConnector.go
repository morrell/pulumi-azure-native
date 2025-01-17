// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents threat intelligence data connector.
func LookupTIDataConnector(ctx *pulumi.Context, args *LookupTIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupTIDataConnectorResult, error) {
	var rv LookupTIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getTIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTIDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents threat intelligence data connector.
type LookupTIDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes *TIDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'ThreatIntelligence'.
	Kind string `pulumi:"kind"`
	// Azure resource name
	Name string `pulumi:"name"`
	// The tenant id to connect to, and get the data from.
	TenantId *string `pulumi:"tenantId"`
	// The lookback period for the feed to be imported.
	TipLookbackPeriod *string `pulumi:"tipLookbackPeriod"`
	// Azure resource type
	Type string `pulumi:"type"`
}

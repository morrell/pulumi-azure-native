// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data connector.
//
// Deprecated: Please use one of the variants: AADDataConnector, AATPDataConnector, ASCDataConnector, AwsCloudTrailDataConnector, Dynamics365DataConnector, MCASDataConnector, MDATPDataConnector, MSTIDataConnector, MTPDataConnector, OfficeATPDataConnector, OfficeDataConnector, TIDataConnector, TiTaxiiDataConnector.
func LookupDataConnector(ctx *pulumi.Context, args *LookupDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupDataConnectorResult, error) {
	var rv LookupDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Data connector.
type LookupDataConnectorResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// The kind of the data connector
	Kind string `pulumi:"kind"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Azure resource type
	Type string `pulumi:"type"`
}

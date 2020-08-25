// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160129

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWorkspaceCollection(ctx *pulumi.Context, args *LookupWorkspaceCollectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceCollectionResult, error) {
	var rv LookupWorkspaceCollectionResult
	err := ctx.Invoke("azurerm:powerbi/v20160129:getWorkspaceCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceCollectionArgs struct {
	// Power BI Embedded Workspace Collection name
	Name string `pulumi:"name"`
	// Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupWorkspaceCollectionResult struct {
	// Azure location
	Location *string `pulumi:"location"`
	// Workspace collection name
	Name *string `pulumi:"name"`
	// Properties
	Properties map[string]interface{} `pulumi:"properties"`
	Sku        *AzureSkuResponse      `pulumi:"sku"`
	Tags       map[string]string      `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}
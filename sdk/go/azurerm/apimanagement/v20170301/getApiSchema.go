// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupApiSchema(ctx *pulumi.Context, args *LookupApiSchemaArgs, opts ...pulumi.InvokeOption) (*LookupApiSchemaResult, error) {
	var rv LookupApiSchemaResult
	err := ctx.Invoke("azurerm:apimanagement/v20170301:getApiSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiSchemaArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Schema identifier within an API. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Schema Contract details.
type LookupApiSchemaResult struct {
	// Must be a valid a media type used in a Content-Type header as defined in the RFC 2616. Media type of the schema document (e.g. application/json, application/xml).
	ContentType string `pulumi:"contentType"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
	// Json escaped string defining the document representing the Schema.
	Value *string `pulumi:"value"`
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Instance of an Azure ML web service resource.
func LookupWebService(ctx *pulumi.Context, args *LookupWebServiceArgs, opts ...pulumi.InvokeOption) (*LookupWebServiceResult, error) {
	var rv LookupWebServiceResult
	err := ctx.Invoke("azure-native:machinelearning/v20160501preview:getWebService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebServiceArgs struct {
	// Name of the resource group in which the web service is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the web service.
	WebServiceName string `pulumi:"webServiceName"`
}

// Instance of an Azure ML web service resource.
type LookupWebServiceResult struct {
	// Specifies the resource ID.
	Id string `pulumi:"id"`
	// Specifies the location of the resource.
	Location string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesForGraphResponse `pulumi:"properties"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}

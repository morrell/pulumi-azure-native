// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A manifest file that defines the custom resource provider resources.
func LookupCustomResourceProvider(ctx *pulumi.Context, args *LookupCustomResourceProviderArgs, opts ...pulumi.InvokeOption) (*LookupCustomResourceProviderResult, error) {
	var rv LookupCustomResourceProviderResult
	err := ctx.Invoke("azure-native:customproviders/v20180901preview:getCustomResourceProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomResourceProviderArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource provider.
	ResourceProviderName string `pulumi:"resourceProviderName"`
}

// A manifest file that defines the custom resource provider resources.
type LookupCustomResourceProviderResult struct {
	// A list of actions that the custom resource provider implements.
	Actions []CustomRPActionRouteDefinitionResponse `pulumi:"actions"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The provisioning state of the resource provider.
	ProvisioningState string `pulumi:"provisioningState"`
	// A list of resource types that the custom resource provider implements.
	ResourceTypes []CustomRPResourceTypeRouteDefinitionResponse `pulumi:"resourceTypes"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// A list of validations to run on the custom resource provider's requests.
	Validations []CustomRPValidationsResponse `pulumi:"validations"`
}

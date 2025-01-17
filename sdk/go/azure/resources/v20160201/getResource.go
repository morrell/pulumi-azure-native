// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource information.
func LookupResource(ctx *pulumi.Context, args *LookupResourceArgs, opts ...pulumi.InvokeOption) (*LookupResourceResult, error) {
	var rv LookupResourceResult
	err := ctx.Invoke("azure-native:resources/v20160201:getResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceArgs struct {
	// Resource identity.
	ParentResourcePath string `pulumi:"parentResourcePath"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource identity.
	ResourceName string `pulumi:"resourceName"`
	// Resource identity.
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	// Resource identity.
	ResourceType string `pulumi:"resourceType"`
}

// Resource information.
type LookupResourceResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The kind of the resource.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// Id of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// Resource name
	Name string `pulumi:"name"`
	// The plan of the resource.
	Plan *PlanResponse `pulumi:"plan"`
	// The resource properties.
	Properties interface{} `pulumi:"properties"`
	// The sku of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

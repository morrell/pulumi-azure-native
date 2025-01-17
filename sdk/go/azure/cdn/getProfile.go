// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CDN profile is a logical grouping of endpoints that share the same settings, such as CDN provider and pricing tier.
// API Version: 2020-09-01.
func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azure-native:cdn:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// CDN profile is a logical grouping of endpoints that share the same settings, such as CDN provider and pricing tier.
type LookupProfileResult struct {
	// The Id of the frontdoor.
	FrontdoorId string `pulumi:"frontdoorId"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Provisioning status of the profile.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource status of the profile.
	ResourceState string `pulumi:"resourceState"`
	// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
	Sku SkuResponse `pulumi:"sku"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

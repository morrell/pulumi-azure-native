// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azurerm:network/v20170301:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	// The name of the Traffic Manager profile.
	Name string `pulumi:"name"`
	// The name of the resource group containing the Traffic Manager profile.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a Traffic Manager profile.
type LookupProfileResult struct {
	// Gets or sets the DNS settings of the Traffic Manager profile.
	DnsConfig *DnsConfigResponse `pulumi:"dnsConfig"`
	// Gets or sets the list of endpoints in the Traffic Manager profile.
	Endpoints []EndpointResponse `pulumi:"endpoints"`
	// Resource location
	Location *string `pulumi:"location"`
	// Gets or sets the endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfigResponse `pulumi:"monitorConfig"`
	// Resource name
	Name string `pulumi:"name"`
	// Gets or sets the status of the Traffic Manager profile.  Possible values are 'Enabled' and 'Disabled'.
	ProfileStatus *string `pulumi:"profileStatus"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the traffic routing method of the Traffic Manager profile.  Possible values are 'Performance', 'Weighted', 'Priority' or 'Geographic'.
	TrafficRoutingMethod *string `pulumi:"trafficRoutingMethod"`
	// Resource type
	Type string `pulumi:"type"`
}
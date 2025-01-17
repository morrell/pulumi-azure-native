// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PublicIPAddress resource
func LookupPublicIpAddress(ctx *pulumi.Context, args *LookupPublicIpAddressArgs, opts ...pulumi.InvokeOption) (*LookupPublicIpAddressResult, error) {
	var rv LookupPublicIpAddressResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getPublicIpAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicIpAddressArgs struct {
	// The name of the subnet.
	PublicIpAddressName string `pulumi:"publicIpAddressName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// PublicIPAddress resource
type LookupPublicIpAddressResult struct {
	// Gets or sets FQDN of the DNS record associated with the public IP address
	DnsSettings *PublicIpAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id string `pulumi:"id"`
	// Gets or sets the idle timeout of the public IP address
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// Gets the assigned public IP address
	IpAddress *string `pulumi:"ipAddress"`
	// Gets a reference to the network interface IP configurations using this public IP address
	IpConfiguration *SubResourceResponse `pulumi:"ipConfiguration"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets PublicIP allocation method (Static/Dynamic)
	PublicIPAllocationMethod string `pulumi:"publicIPAllocationMethod"`
	// Gets or sets resource guid property of the PublicIP resource
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

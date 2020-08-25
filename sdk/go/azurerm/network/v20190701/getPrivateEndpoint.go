// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azurerm:network/v20190701:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the private endpoint.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Private endpoint resource.
type LookupPrivateEndpointResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"manualPrivateLinkServiceConnections"`
	// Resource name.
	Name string `pulumi:"name"`
	// An array of references to the network interfaces created for this private endpoint.
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	// A grouping of information about the connection to the remote resource.
	PrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"privateLinkServiceConnections"`
	// The provisioning state of the private endpoint resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The ID of the subnet from which the private IP will be allocated.
	Subnet *SubnetResponse `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}
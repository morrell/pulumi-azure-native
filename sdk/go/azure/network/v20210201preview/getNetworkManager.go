// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Managed Network resource
func LookupNetworkManager(ctx *pulumi.Context, args *LookupNetworkManagerArgs, opts ...pulumi.InvokeOption) (*LookupNetworkManagerResult, error) {
	var rv LookupNetworkManagerResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getNetworkManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkManagerArgs struct {
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Managed Network resource
type LookupNetworkManagerResult struct {
	// A description of the network manager.
	Description *string `pulumi:"description"`
	// A friendly name for the network manager.
	DisplayName *string `pulumi:"displayName"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Scope Access.
	NetworkManagerScopeAccesses []string `pulumi:"networkManagerScopeAccesses"`
	// Scope of Network Manager.
	NetworkManagerScopes *NetworkManagerPropertiesResponseNetworkManagerScopes `pulumi:"networkManagerScopes"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

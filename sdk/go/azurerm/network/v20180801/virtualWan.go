// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VirtualWAN Resource.
type VirtualWan struct {
	pulumi.CustomResourceState

	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic pulumi.BoolPtrOutput `pulumi:"allowBranchToBranchTraffic"`
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic pulumi.BoolPtrOutput `pulumi:"allowVnetToVnetTraffic"`
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption pulumi.BoolPtrOutput `pulumi:"disableVpnEncryption"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The office local breakout category.
	Office365LocalBreakoutCategory pulumi.StringPtrOutput `pulumi:"office365LocalBreakoutCategory"`
	// list of all P2SVpnServerConfigurations associated with the virtual wan.
	P2SVpnServerConfigurations P2SVpnServerConfigurationResponseArrayOutput `pulumi:"p2SVpnServerConfigurations"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The Security Provider name.
	SecurityProviderName pulumi.StringPtrOutput `pulumi:"securityProviderName"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of VirtualHubs in the VirtualWAN.
	VirtualHubs SubResourceResponseArrayOutput `pulumi:"virtualHubs"`
	VpnSites    SubResourceResponseArrayOutput `pulumi:"vpnSites"`
}

// NewVirtualWan registers a new resource with the given unique name, arguments, and options.
func NewVirtualWan(ctx *pulumi.Context,
	name string, args *VirtualWanArgs, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VirtualWanArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/v20180401:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:VirtualWan"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:VirtualWan"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualWan
	err := ctx.RegisterResource("azurerm:network/v20180801:VirtualWan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualWan gets an existing VirtualWan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualWan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualWanState, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	var resource VirtualWan
	err := ctx.ReadResource("azurerm:network/v20180801:VirtualWan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualWan resources.
type virtualWanState struct {
	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic *bool `pulumi:"allowVnetToVnetTraffic"`
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The office local breakout category.
	Office365LocalBreakoutCategory *string `pulumi:"office365LocalBreakoutCategory"`
	// list of all P2SVpnServerConfigurations associated with the virtual wan.
	P2SVpnServerConfigurations []P2SVpnServerConfigurationResponse `pulumi:"p2SVpnServerConfigurations"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The Security Provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// List of VirtualHubs in the VirtualWAN.
	VirtualHubs []SubResourceResponse `pulumi:"virtualHubs"`
	VpnSites    []SubResourceResponse `pulumi:"vpnSites"`
}

type VirtualWanState struct {
	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic pulumi.BoolPtrInput
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption pulumi.BoolPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The office local breakout category.
	Office365LocalBreakoutCategory pulumi.StringPtrInput
	// list of all P2SVpnServerConfigurations associated with the virtual wan.
	P2SVpnServerConfigurations P2SVpnServerConfigurationResponseArrayInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The Security Provider name.
	SecurityProviderName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// List of VirtualHubs in the VirtualWAN.
	VirtualHubs SubResourceResponseArrayInput
	VpnSites    SubResourceResponseArrayInput
}

func (VirtualWanState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanState)(nil)).Elem()
}

type virtualWanArgs struct {
	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic *bool `pulumi:"allowVnetToVnetTraffic"`
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the VirtualWAN being created or updated.
	Name string `pulumi:"name"`
	// The office local breakout category.
	Office365LocalBreakoutCategory *string `pulumi:"office365LocalBreakoutCategory"`
	// list of all P2SVpnServerConfigurations associated with the virtual wan.
	P2SVpnServerConfigurations []P2SVpnServerConfiguration `pulumi:"p2SVpnServerConfigurations"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource group name of the VirtualWan.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Security Provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualWan resource.
type VirtualWanArgs struct {
	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic pulumi.BoolPtrInput
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The name of the VirtualWAN being created or updated.
	Name pulumi.StringInput
	// The office local breakout category.
	Office365LocalBreakoutCategory pulumi.StringPtrInput
	// list of all P2SVpnServerConfigurations associated with the virtual wan.
	P2SVpnServerConfigurations P2SVpnServerConfigurationArrayInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The resource group name of the VirtualWan.
	ResourceGroupName pulumi.StringInput
	// The Security Provider name.
	SecurityProviderName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (VirtualWanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanArgs)(nil)).Elem()
}
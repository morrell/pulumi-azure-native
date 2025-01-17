// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160430preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Virtual Machine Scale Set.
type VirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	// The identity of the virtual machine scale set, if configured.
	Identity VirtualMachineScaleSetIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	OverProvision pulumi.BoolPtrOutput `pulumi:"overProvision"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	SinglePlacementGroup pulumi.BoolPtrOutput `pulumi:"singlePlacementGroup"`
	// The virtual machine scale set sku.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The upgrade policy.
	UpgradePolicy UpgradePolicyResponsePtrOutput `pulumi:"upgradePolicy"`
	// The virtual machine profile.
	VirtualMachineProfile VirtualMachineScaleSetVMProfileResponsePtrOutput `pulumi:"virtualMachineProfile"`
}

// NewVirtualMachineScaleSet registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20150615:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20150615:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20160330:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20171201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180401:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20181001:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200601:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20201201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210301:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210401:VirtualMachineScaleSet"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSet
	err := ctx.RegisterResource("azure-native:compute/v20160430preview:VirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSet gets an existing VirtualMachineScaleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	var resource VirtualMachineScaleSet
	err := ctx.ReadResource("azure-native:compute/v20160430preview:VirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSet resources.
type virtualMachineScaleSetState struct {
}

type VirtualMachineScaleSetState struct {
}

func (VirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetState)(nil)).Elem()
}

type virtualMachineScaleSetArgs struct {
	// The identity of the virtual machine scale set, if configured.
	Identity *VirtualMachineScaleSetIdentity `pulumi:"identity"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the VM scale set to create or update.
	Name *string `pulumi:"name"`
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	OverProvision *bool `pulumi:"overProvision"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *Plan `pulumi:"plan"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// The virtual machine scale set sku.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The upgrade policy.
	UpgradePolicy *UpgradePolicy `pulumi:"upgradePolicy"`
	// The virtual machine profile.
	VirtualMachineProfile *VirtualMachineScaleSetVMProfile `pulumi:"virtualMachineProfile"`
}

// The set of arguments for constructing a VirtualMachineScaleSet resource.
type VirtualMachineScaleSetArgs struct {
	// The identity of the virtual machine scale set, if configured.
	Identity VirtualMachineScaleSetIdentityPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the VM scale set to create or update.
	Name pulumi.StringPtrInput
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	OverProvision pulumi.BoolPtrInput
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	SinglePlacementGroup pulumi.BoolPtrInput
	// The virtual machine scale set sku.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The upgrade policy.
	UpgradePolicy UpgradePolicyPtrInput
	// The virtual machine profile.
	VirtualMachineProfile VirtualMachineScaleSetVMProfilePtrInput
}

func (VirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetArgs)(nil)).Elem()
}

type VirtualMachineScaleSetInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput
	ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput
}

func (*VirtualMachineScaleSet) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSet)(nil))
}

func (i *VirtualMachineScaleSet) ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput {
	return i.ToVirtualMachineScaleSetOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSet) ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOutput)
}

type VirtualMachineScaleSetOutput struct {
	*pulumi.OutputState
}

func (VirtualMachineScaleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSet)(nil))
}

func (o VirtualMachineScaleSetOutput) ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput {
	return o
}

func (o VirtualMachineScaleSetOutput) ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetOutput{})
}

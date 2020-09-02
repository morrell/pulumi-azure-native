// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180630preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Guest configuration assignment is an association between a VM and guest configuration.
type GuestConfigurationAssignment struct {
	pulumi.CustomResourceState

	// Region where the VM is located.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGuestConfigurationAssignment registers a new resource with the given unique name, arguments, and options.
func NewGuestConfigurationAssignment(ctx *pulumi.Context,
	name string, args *GuestConfigurationAssignmentArgs, opts ...pulumi.ResourceOption) (*GuestConfigurationAssignment, error) {
	if args == nil || args.GuestConfigurationAssignmentName == nil {
		return nil, errors.New("missing required argument 'GuestConfigurationAssignmentName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VmName == nil {
		return nil, errors.New("missing required argument 'VmName'")
	}
	if args == nil {
		args = &GuestConfigurationAssignmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:compute/latest:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20181120:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20200625:GuestConfigurationAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestConfigurationAssignment
	err := ctx.RegisterResource("azurerm:compute/v20180630preview:GuestConfigurationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestConfigurationAssignment gets an existing GuestConfigurationAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestConfigurationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestConfigurationAssignmentState, opts ...pulumi.ResourceOption) (*GuestConfigurationAssignment, error) {
	var resource GuestConfigurationAssignment
	err := ctx.ReadResource("azurerm:compute/v20180630preview:GuestConfigurationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestConfigurationAssignment resources.
type guestConfigurationAssignmentState struct {
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties *GuestConfigurationAssignmentPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type GuestConfigurationAssignmentState struct {
	// Region where the VM is located.
	Location pulumi.StringPtrInput
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrInput
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (GuestConfigurationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationAssignmentState)(nil)).Elem()
}

type guestConfigurationAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName string `pulumi:"guestConfigurationAssignmentName"`
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties *GuestConfigurationAssignmentProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName string `pulumi:"vmName"`
}

// The set of arguments for constructing a GuestConfigurationAssignment resource.
type GuestConfigurationAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName pulumi.StringInput
	// Region where the VM is located.
	Location pulumi.StringPtrInput
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrInput
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the virtual machine.
	VmName pulumi.StringInput
}

func (GuestConfigurationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationAssignmentArgs)(nil)).Elem()
}
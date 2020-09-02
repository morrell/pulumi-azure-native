// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200630preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Configuration profile assignment is an association between a VM and automanage profile configuration.
type ConfigurationProfileAssignment struct {
	pulumi.CustomResourceState

	// Region where the VM is located.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the Automanage assignment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the configuration profile assignment.
	Properties ConfigurationProfileAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationProfileAssignment registers a new resource with the given unique name, arguments, and options.
func NewConfigurationProfileAssignment(ctx *pulumi.Context,
	name string, args *ConfigurationProfileAssignmentArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfileAssignment, error) {
	if args == nil || args.ConfigurationProfileAssignmentName == nil {
		return nil, errors.New("missing required argument 'ConfigurationProfileAssignmentName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VmName == nil {
		return nil, errors.New("missing required argument 'VmName'")
	}
	if args == nil {
		args = &ConfigurationProfileAssignmentArgs{}
	}
	var resource ConfigurationProfileAssignment
	err := ctx.RegisterResource("azurerm:compute/v20200630preview:ConfigurationProfileAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationProfileAssignment gets an existing ConfigurationProfileAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationProfileAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileAssignmentState, opts ...pulumi.ResourceOption) (*ConfigurationProfileAssignment, error) {
	var resource ConfigurationProfileAssignment
	err := ctx.ReadResource("azurerm:compute/v20200630preview:ConfigurationProfileAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationProfileAssignment resources.
type configurationProfileAssignmentState struct {
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the Automanage assignment.
	Name *string `pulumi:"name"`
	// Properties of the configuration profile assignment.
	Properties *ConfigurationProfileAssignmentPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ConfigurationProfileAssignmentState struct {
	// Region where the VM is located.
	Location pulumi.StringPtrInput
	// Name of the Automanage assignment.
	Name pulumi.StringPtrInput
	// Properties of the configuration profile assignment.
	Properties ConfigurationProfileAssignmentPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ConfigurationProfileAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileAssignmentState)(nil)).Elem()
}

type configurationProfileAssignmentArgs struct {
	// Name of the configuration profile assignment. Only default is supported.
	ConfigurationProfileAssignmentName string `pulumi:"configurationProfileAssignmentName"`
	// Properties of the configuration profile assignment.
	Properties *ConfigurationProfileAssignmentProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName string `pulumi:"vmName"`
}

// The set of arguments for constructing a ConfigurationProfileAssignment resource.
type ConfigurationProfileAssignmentArgs struct {
	// Name of the configuration profile assignment. Only default is supported.
	ConfigurationProfileAssignmentName pulumi.StringInput
	// Properties of the configuration profile assignment.
	Properties ConfigurationProfileAssignmentPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the virtual machine.
	VmName pulumi.StringInput
}

func (ConfigurationProfileAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileAssignmentArgs)(nil)).Elem()
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Management lock information.
type ManagementLockAtResourceGroupLevel struct {
	pulumi.CustomResourceState

	// The lock level of the management lock.
	Level pulumi.StringPtrOutput `pulumi:"level"`
	// The name of the lock.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The notes of the management lock.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The type of the lock.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementLockAtResourceGroupLevel registers a new resource with the given unique name, arguments, and options.
func NewManagementLockAtResourceGroupLevel(ctx *pulumi.Context,
	name string, args *ManagementLockAtResourceGroupLevelArgs, opts ...pulumi.ResourceOption) (*ManagementLockAtResourceGroupLevel, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ManagementLockAtResourceGroupLevelArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:authorization/v20160901:ManagementLockAtResourceGroupLevel"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementLockAtResourceGroupLevel
	err := ctx.RegisterResource("azurerm:authorization/v20150101:ManagementLockAtResourceGroupLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementLockAtResourceGroupLevel gets an existing ManagementLockAtResourceGroupLevel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementLockAtResourceGroupLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementLockAtResourceGroupLevelState, opts ...pulumi.ResourceOption) (*ManagementLockAtResourceGroupLevel, error) {
	var resource ManagementLockAtResourceGroupLevel
	err := ctx.ReadResource("azurerm:authorization/v20150101:ManagementLockAtResourceGroupLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementLockAtResourceGroupLevel resources.
type managementLockAtResourceGroupLevelState struct {
	// The lock level of the management lock.
	Level *string `pulumi:"level"`
	// The name of the lock.
	Name *string `pulumi:"name"`
	// The notes of the management lock.
	Notes *string `pulumi:"notes"`
	// The type of the lock.
	Type *string `pulumi:"type"`
}

type ManagementLockAtResourceGroupLevelState struct {
	// The lock level of the management lock.
	Level pulumi.StringPtrInput
	// The name of the lock.
	Name pulumi.StringPtrInput
	// The notes of the management lock.
	Notes pulumi.StringPtrInput
	// The type of the lock.
	Type pulumi.StringPtrInput
}

func (ManagementLockAtResourceGroupLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtResourceGroupLevelState)(nil)).Elem()
}

type managementLockAtResourceGroupLevelArgs struct {
	// The lock level of the management lock.
	Level *string `pulumi:"level"`
	// The lock name.
	Name string `pulumi:"name"`
	// The notes of the management lock.
	Notes *string `pulumi:"notes"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagementLockAtResourceGroupLevel resource.
type ManagementLockAtResourceGroupLevelArgs struct {
	// The lock level of the management lock.
	Level pulumi.StringPtrInput
	// The lock name.
	Name pulumi.StringInput
	// The notes of the management lock.
	Notes pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (ManagementLockAtResourceGroupLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtResourceGroupLevelArgs)(nil)).Elem()
}
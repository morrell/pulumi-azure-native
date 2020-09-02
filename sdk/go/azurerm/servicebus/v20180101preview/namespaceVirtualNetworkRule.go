// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Single item in a List or Get VirtualNetworkRules operation
type NamespaceVirtualNetworkRule struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Resource ID of Virtual Network Subnet
	VirtualNetworkSubnetId pulumi.StringPtrOutput `pulumi:"virtualNetworkSubnetId"`
}

// NewNamespaceVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewNamespaceVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *NamespaceVirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceVirtualNetworkRule, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualNetworkRuleName == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkRuleName'")
	}
	if args == nil {
		args = &NamespaceVirtualNetworkRuleArgs{}
	}
	var resource NamespaceVirtualNetworkRule
	err := ctx.RegisterResource("azurerm:servicebus/v20180101preview:NamespaceVirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceVirtualNetworkRule gets an existing NamespaceVirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceVirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*NamespaceVirtualNetworkRule, error) {
	var resource NamespaceVirtualNetworkRule
	err := ctx.ReadResource("azurerm:servicebus/v20180101preview:NamespaceVirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceVirtualNetworkRule resources.
type namespaceVirtualNetworkRuleState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// Resource type
	Type *string `pulumi:"type"`
	// Resource ID of Virtual Network Subnet
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}

type NamespaceVirtualNetworkRuleState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// Resource ID of Virtual Network Subnet
	VirtualNetworkSubnetId pulumi.StringPtrInput
}

func (NamespaceVirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceVirtualNetworkRuleState)(nil)).Elem()
}

type namespaceVirtualNetworkRuleArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Virtual Network Rule name.
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
	// Resource ID of Virtual Network Subnet
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}

// The set of arguments for constructing a NamespaceVirtualNetworkRule resource.
type NamespaceVirtualNetworkRuleArgs struct {
	// The namespace name
	NamespaceName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The Virtual Network Rule name.
	VirtualNetworkRuleName pulumi.StringInput
	// Resource ID of Virtual Network Subnet
	VirtualNetworkSubnetId pulumi.StringPtrInput
}

func (NamespaceVirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceVirtualNetworkRuleArgs)(nil)).Elem()
}
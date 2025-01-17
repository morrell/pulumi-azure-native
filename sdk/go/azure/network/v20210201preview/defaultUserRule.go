// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network security default user rule.
type DefaultUserRule struct {
	pulumi.CustomResourceState

	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringOutput `pulumi:"description"`
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayOutput `pulumi:"destinationPortRanges"`
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations AddressPrefixItemResponseArrayOutput `pulumi:"destinations"`
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// A friendly name for the rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Default rule flag.
	Flag pulumi.StringPtrOutput `pulumi:"flag"`
	// Whether the rule is custom or default.
	// Expected value is 'Default'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network protocol this rule applies to.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The provisioning state of the security configuration user rule resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayOutput `pulumi:"sourcePortRanges"`
	// The CIDR or source IP ranges.
	Sources AddressPrefixItemResponseArrayOutput `pulumi:"sources"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDefaultUserRule registers a new resource with the given unique name, arguments, and options.
func NewDefaultUserRule(ctx *pulumi.Context,
	name string, args *DefaultUserRuleArgs, opts ...pulumi.ResourceOption) (*DefaultUserRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'RuleCollectionName'")
	}
	args.Kind = pulumi.String("Default")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20210201preview:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:DefaultUserRule"),
		},
	})
	opts = append(opts, aliases)
	var resource DefaultUserRule
	err := ctx.RegisterResource("azure-native:network/v20210201preview:DefaultUserRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultUserRule gets an existing DefaultUserRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultUserRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultUserRuleState, opts ...pulumi.ResourceOption) (*DefaultUserRule, error) {
	var resource DefaultUserRule
	err := ctx.ReadResource("azure-native:network/v20210201preview:DefaultUserRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultUserRule resources.
type defaultUserRuleState struct {
}

type DefaultUserRuleState struct {
}

func (DefaultUserRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultUserRuleState)(nil)).Elem()
}

type defaultUserRuleArgs struct {
	// The name of the network manager security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// Default rule flag.
	Flag *string `pulumi:"flag"`
	// Whether the rule is custom or default.
	// Expected value is 'Default'.
	Kind string `pulumi:"kind"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName *string `pulumi:"ruleName"`
}

// The set of arguments for constructing a DefaultUserRule resource.
type DefaultUserRuleArgs struct {
	// The name of the network manager security Configuration.
	ConfigurationName pulumi.StringInput
	// Default rule flag.
	Flag pulumi.StringPtrInput
	// Whether the rule is custom or default.
	// Expected value is 'Default'.
	Kind pulumi.StringInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringInput
	// The name of the rule.
	RuleName pulumi.StringPtrInput
}

func (DefaultUserRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultUserRuleArgs)(nil)).Elem()
}

type DefaultUserRuleInput interface {
	pulumi.Input

	ToDefaultUserRuleOutput() DefaultUserRuleOutput
	ToDefaultUserRuleOutputWithContext(ctx context.Context) DefaultUserRuleOutput
}

func (*DefaultUserRule) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultUserRule)(nil))
}

func (i *DefaultUserRule) ToDefaultUserRuleOutput() DefaultUserRuleOutput {
	return i.ToDefaultUserRuleOutputWithContext(context.Background())
}

func (i *DefaultUserRule) ToDefaultUserRuleOutputWithContext(ctx context.Context) DefaultUserRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultUserRuleOutput)
}

type DefaultUserRuleOutput struct {
	*pulumi.OutputState
}

func (DefaultUserRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultUserRule)(nil))
}

func (o DefaultUserRuleOutput) ToDefaultUserRuleOutput() DefaultUserRuleOutput {
	return o
}

func (o DefaultUserRuleOutput) ToDefaultUserRuleOutputWithContext(ctx context.Context) DefaultUserRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DefaultUserRuleOutput{})
}

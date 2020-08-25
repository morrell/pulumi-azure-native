// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The alert rule resource.
type AlertRule struct {
	pulumi.CustomResourceState

	// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions RuleActionResponseArrayOutput `pulumi:"actions"`
	// the condition that results in the alert rule being activated.
	Condition RuleConditionResponseOutput `pulumi:"condition"`
	// the description of the alert rule that will be included in the alert email.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// the flag that indicates whether the alert rule is enabled.
	IsEnabled pulumi.BoolOutput `pulumi:"isEnabled"`
	// Last time the rule was updated in ISO8601 format.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlertRule registers a new resource with the given unique name, arguments, and options.
func NewAlertRule(ctx *pulumi.Context,
	name string, args *AlertRuleArgs, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	if args == nil || args.Condition == nil {
		return nil, errors.New("missing required argument 'Condition'")
	}
	if args == nil || args.IsEnabled == nil {
		return nil, errors.New("missing required argument 'IsEnabled'")
	}
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
		args = &AlertRuleArgs{}
	}
	var resource AlertRule
	err := ctx.RegisterResource("azurerm:insights/v20160301:AlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertRule gets an existing AlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleState, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	var resource AlertRule
	err := ctx.ReadResource("azurerm:insights/v20160301:AlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertRule resources.
type alertRuleState struct {
	// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions []RuleActionResponse `pulumi:"actions"`
	// the condition that results in the alert rule being activated.
	Condition *RuleConditionResponse `pulumi:"condition"`
	// the description of the alert rule that will be included in the alert email.
	Description *string `pulumi:"description"`
	// the flag that indicates whether the alert rule is enabled.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Last time the rule was updated in ISO8601 format.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type AlertRuleState struct {
	// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions RuleActionResponseArrayInput
	// the condition that results in the alert rule being activated.
	Condition RuleConditionResponsePtrInput
	// the description of the alert rule that will be included in the alert email.
	Description pulumi.StringPtrInput
	// the flag that indicates whether the alert rule is enabled.
	IsEnabled pulumi.BoolPtrInput
	// Last time the rule was updated in ISO8601 format.
	LastUpdatedTime pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (AlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleState)(nil)).Elem()
}

type alertRuleArgs struct {
	// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions []RuleAction `pulumi:"actions"`
	// the condition that results in the alert rule being activated.
	Condition RuleCondition `pulumi:"condition"`
	// the description of the alert rule that will be included in the alert email.
	Description *string `pulumi:"description"`
	// the flag that indicates whether the alert rule is enabled.
	IsEnabled bool `pulumi:"isEnabled"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the rule.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AlertRule resource.
type AlertRuleArgs struct {
	// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions RuleActionArrayInput
	// the condition that results in the alert rule being activated.
	Condition RuleConditionInput
	// the description of the alert rule that will be included in the alert email.
	Description pulumi.StringPtrInput
	// the flag that indicates whether the alert rule is enabled.
	IsEnabled pulumi.BoolInput
	// Resource location
	Location pulumi.StringInput
	// The name of the rule.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (AlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleArgs)(nil)).Elem()
}
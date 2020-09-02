// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes the suppression rule
type AlertsSuppressionRule struct {
	pulumi.CustomResourceState

	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType pulumi.StringOutput `pulumi:"alertType"`
	// Any comment regarding the rule
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc pulumi.StringPtrOutput `pulumi:"expirationDateUtc"`
	// The last time this rule was modified
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The reason for dismissing the alert
	Reason pulumi.StringOutput `pulumi:"reason"`
	// Possible states of the rule
	State pulumi.StringOutput `pulumi:"state"`
	// The suppression conditions
	SuppressionAlertsScope SuppressionAlertsScopeResponsePtrOutput `pulumi:"suppressionAlertsScope"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlertsSuppressionRule registers a new resource with the given unique name, arguments, and options.
func NewAlertsSuppressionRule(ctx *pulumi.Context,
	name string, args *AlertsSuppressionRuleArgs, opts ...pulumi.ResourceOption) (*AlertsSuppressionRule, error) {
	if args == nil || args.AlertType == nil {
		return nil, errors.New("missing required argument 'AlertType'")
	}
	if args == nil || args.AlertsSuppressionRuleName == nil {
		return nil, errors.New("missing required argument 'AlertsSuppressionRuleName'")
	}
	if args == nil || args.Reason == nil {
		return nil, errors.New("missing required argument 'Reason'")
	}
	if args == nil || args.State == nil {
		return nil, errors.New("missing required argument 'State'")
	}
	if args == nil {
		args = &AlertsSuppressionRuleArgs{}
	}
	var resource AlertsSuppressionRule
	err := ctx.RegisterResource("azurerm:security/v20190101preview:AlertsSuppressionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertsSuppressionRule gets an existing AlertsSuppressionRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertsSuppressionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertsSuppressionRuleState, opts ...pulumi.ResourceOption) (*AlertsSuppressionRule, error) {
	var resource AlertsSuppressionRule
	err := ctx.ReadResource("azurerm:security/v20190101preview:AlertsSuppressionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertsSuppressionRule resources.
type alertsSuppressionRuleState struct {
	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType *string `pulumi:"alertType"`
	// Any comment regarding the rule
	Comment *string `pulumi:"comment"`
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc *string `pulumi:"expirationDateUtc"`
	// The last time this rule was modified
	LastModifiedUtc *string `pulumi:"lastModifiedUtc"`
	// Resource name
	Name *string `pulumi:"name"`
	// The reason for dismissing the alert
	Reason *string `pulumi:"reason"`
	// Possible states of the rule
	State *string `pulumi:"state"`
	// The suppression conditions
	SuppressionAlertsScope *SuppressionAlertsScopeResponse `pulumi:"suppressionAlertsScope"`
	// Resource type
	Type *string `pulumi:"type"`
}

type AlertsSuppressionRuleState struct {
	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType pulumi.StringPtrInput
	// Any comment regarding the rule
	Comment pulumi.StringPtrInput
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc pulumi.StringPtrInput
	// The last time this rule was modified
	LastModifiedUtc pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The reason for dismissing the alert
	Reason pulumi.StringPtrInput
	// Possible states of the rule
	State pulumi.StringPtrInput
	// The suppression conditions
	SuppressionAlertsScope SuppressionAlertsScopeResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (AlertsSuppressionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertsSuppressionRuleState)(nil)).Elem()
}

type alertsSuppressionRuleArgs struct {
	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType string `pulumi:"alertType"`
	// The unique name of the suppression alert rule
	AlertsSuppressionRuleName string `pulumi:"alertsSuppressionRuleName"`
	// Any comment regarding the rule
	Comment *string `pulumi:"comment"`
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc *string `pulumi:"expirationDateUtc"`
	// The reason for dismissing the alert
	Reason string `pulumi:"reason"`
	// Possible states of the rule
	State string `pulumi:"state"`
	// The suppression conditions
	SuppressionAlertsScope *SuppressionAlertsScope `pulumi:"suppressionAlertsScope"`
}

// The set of arguments for constructing a AlertsSuppressionRule resource.
type AlertsSuppressionRuleArgs struct {
	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType pulumi.StringInput
	// The unique name of the suppression alert rule
	AlertsSuppressionRuleName pulumi.StringInput
	// Any comment regarding the rule
	Comment pulumi.StringPtrInput
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc pulumi.StringPtrInput
	// The reason for dismissing the alert
	Reason pulumi.StringInput
	// Possible states of the rule
	State pulumi.StringInput
	// The suppression conditions
	SuppressionAlertsScope SuppressionAlertsScopePtrInput
}

func (AlertsSuppressionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertsSuppressionRuleArgs)(nil)).Elem()
}
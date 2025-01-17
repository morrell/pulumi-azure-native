// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Friendly Rules name mapping to the any Rules or secret related information.
// API Version: 2020-09-01.
func LookupRule(ctx *pulumi.Context, args *LookupRuleArgs, opts ...pulumi.InvokeOption) (*LookupRuleResult, error) {
	var rv LookupRuleResult
	err := ctx.Invoke("azure-native:cdn:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleArgs struct {
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the delivery rule which is unique within the endpoint.
	RuleName string `pulumi:"ruleName"`
	// Name of the rule set under the profile.
	RuleSetName string `pulumi:"ruleSetName"`
}

// Friendly Rules name mapping to the any Rules or secret related information.
type LookupRuleResult struct {
	// A list of actions that are executed when all the conditions of a rule are satisfied.
	Actions []interface{} `pulumi:"actions"`
	// A list of conditions that must be matched for the actions to be executed
	Conditions       []interface{} `pulumi:"conditions"`
	DeploymentStatus string        `pulumi:"deploymentStatus"`
	// Resource ID.
	Id string `pulumi:"id"`
	// If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
	MatchProcessingBehavior *string `pulumi:"matchProcessingBehavior"`
	// Resource name.
	Name string `pulumi:"name"`
	// The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
	Order int `pulumi:"order"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

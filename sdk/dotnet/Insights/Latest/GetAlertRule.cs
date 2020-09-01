// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.Latest
{
    public static class GetAlertRule
    {
        public static Task<GetAlertRuleResult> InvokeAsync(GetAlertRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlertRuleResult>("azurerm:insights/latest:getAlertRule", args ?? new GetAlertRuleArgs(), options.WithVersion());
    }


    public sealed class GetAlertRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("ruleName", required: true)]
        public string RuleName { get; set; } = null!;

        public GetAlertRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlertRuleResult
    {
        /// <summary>
        /// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleActionResponseResult> Actions;
        /// <summary>
        /// the condition that results in the alert rule being activated.
        /// </summary>
        public readonly Outputs.RuleConditionResponseResult Condition;
        /// <summary>
        /// the description of the alert rule that will be included in the alert email.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// the flag that indicates whether the alert rule is enabled.
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// Last time the rule was updated in ISO8601 format.
        /// </summary>
        public readonly string LastUpdatedTime;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAlertRuleResult(
            ImmutableArray<Outputs.RuleActionResponseResult> actions,

            Outputs.RuleConditionResponseResult condition,

            string? description,

            bool isEnabled,

            string lastUpdatedTime,

            string location,

            string name,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Actions = actions;
            Condition = condition;
            Description = description;
            IsEnabled = isEnabled;
            LastUpdatedTime = lastUpdatedTime;
            Location = location;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
}
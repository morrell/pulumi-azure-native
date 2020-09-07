// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AlertsManagement.V20181102PrivatePreview
{
    public static class GetActionRuleByName
    {
        public static Task<GetActionRuleByNameResult> InvokeAsync(GetActionRuleByNameArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetActionRuleByNameResult>("azurerm:alertsmanagement/v20181102privatepreview:getActionRuleByName", args ?? new GetActionRuleByNameArgs(), options.WithVersion());
    }


    public sealed class GetActionRuleByNameArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of action rule that needs to be fetched
        /// </summary>
        [Input("actionRuleName", required: true)]
        public string ActionRuleName { get; set; } = null!;

        /// <summary>
        /// Resource group name where the resource is created.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public string ResourceGroup { get; set; } = null!;

        public GetActionRuleByNameArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetActionRuleByNameResult
    {
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Action rule properties defining scope, conditions, suppression logic for action rule
        /// </summary>
        public readonly Outputs.ActionRulePropertiesResponseResult Properties;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetActionRuleByNameResult(
            string location,

            string name,

            Outputs.ActionRulePropertiesResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
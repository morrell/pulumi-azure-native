// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Management.V20190101
{
    public static class GetPolicyDefinitionAtManagementGroup
    {
        public static Task<GetPolicyDefinitionAtManagementGroupResult> InvokeAsync(GetPolicyDefinitionAtManagementGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyDefinitionAtManagementGroupResult>("azurerm:management/v20190101:getPolicyDefinitionAtManagementGroup", args ?? new GetPolicyDefinitionAtManagementGroupArgs(), options.WithVersion());
    }


    public sealed class GetPolicyDefinitionAtManagementGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the management group.
        /// </summary>
        [Input("managementGroupId", required: true)]
        public string ManagementGroupId { get; set; } = null!;

        /// <summary>
        /// The name of the policy definition to get.
        /// </summary>
        [Input("policyDefinitionName", required: true)]
        public string PolicyDefinitionName { get; set; } = null!;

        public GetPolicyDefinitionAtManagementGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyDefinitionAtManagementGroupResult
    {
        /// <summary>
        /// The policy definition description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The display name of the policy definition.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The policy definition metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Metadata;
        /// <summary>
        /// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// The name of the policy definition.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Required if a parameter is used in policy rule.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Parameters;
        /// <summary>
        /// The policy rule.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? PolicyRule;
        /// <summary>
        /// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
        /// </summary>
        public readonly string? PolicyType;
        /// <summary>
        /// The type of the resource (Microsoft.Authorization/policyDefinitions).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyDefinitionAtManagementGroupResult(
            string? description,

            string? displayName,

            ImmutableDictionary<string, object>? metadata,

            string? mode,

            string name,

            ImmutableDictionary<string, object>? parameters,

            ImmutableDictionary<string, object>? policyRule,

            string? policyType,

            string type)
        {
            Description = description;
            DisplayName = displayName;
            Metadata = metadata;
            Mode = mode;
            Name = name;
            Parameters = parameters;
            PolicyRule = policyRule;
            PolicyType = policyType;
            Type = type;
        }
    }
}
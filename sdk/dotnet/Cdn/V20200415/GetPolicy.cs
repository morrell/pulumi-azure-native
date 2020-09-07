// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20200415
{
    public static class GetPolicy
    {
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("azurerm:cdn/v20200415:getPolicy", args ?? new GetPolicyArgs(), options.WithVersion());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the CdnWebApplicationFirewallPolicy.
        /// </summary>
        [Input("policyName", required: true)]
        public string PolicyName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// Describes custom rules inside the policy.
        /// </summary>
        public readonly Outputs.CustomRuleListResponseResult? CustomRules;
        /// <summary>
        /// Describes Azure CDN endpoints associated with this Web Application Firewall policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.CdnEndpointResponseResult> EndpointLinks;
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Describes managed rules inside the policy.
        /// </summary>
        public readonly Outputs.ManagedRuleSetListResponseResult? ManagedRules;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Describes  policySettings for policy
        /// </summary>
        public readonly Outputs.PolicySettingsResponseResult? PolicySettings;
        /// <summary>
        /// Provisioning state of the WebApplicationFirewallPolicy.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Describes rate limit rules inside the policy.
        /// </summary>
        public readonly Outputs.RateLimitRuleListResponseResult? RateLimitRules;
        public readonly string ResourceState;
        /// <summary>
        /// The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
        /// </summary>
        public readonly Outputs.SkuResponseResult Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyResult(
            Outputs.CustomRuleListResponseResult? customRules,

            ImmutableArray<Outputs.CdnEndpointResponseResult> endpointLinks,

            string? etag,

            string location,

            Outputs.ManagedRuleSetListResponseResult? managedRules,

            string name,

            Outputs.PolicySettingsResponseResult? policySettings,

            string provisioningState,

            Outputs.RateLimitRuleListResponseResult? rateLimitRules,

            string resourceState,

            Outputs.SkuResponseResult sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            CustomRules = customRules;
            EndpointLinks = endpointLinks;
            Etag = etag;
            Location = location;
            ManagedRules = managedRules;
            Name = name;
            PolicySettings = policySettings;
            ProvisioningState = provisioningState;
            RateLimitRules = rateLimitRules;
            ResourceState = resourceState;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
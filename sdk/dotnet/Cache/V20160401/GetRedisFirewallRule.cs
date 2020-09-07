// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache.V20160401
{
    public static class GetRedisFirewallRule
    {
        public static Task<GetRedisFirewallRuleResult> InvokeAsync(GetRedisFirewallRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRedisFirewallRuleResult>("azurerm:cache/v20160401:getRedisFirewallRule", args ?? new GetRedisFirewallRuleArgs(), options.WithVersion());
    }


    public sealed class GetRedisFirewallRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("cacheName", required: true)]
        public string CacheName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the firewall rule.
        /// </summary>
        [Input("ruleName", required: true)]
        public string RuleName { get; set; } = null!;

        public GetRedisFirewallRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRedisFirewallRuleResult
    {
        /// <summary>
        /// highest IP address included in the range
        /// </summary>
        public readonly string EndIP;
        /// <summary>
        /// name of the firewall rule
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// lowest IP address included in the range
        /// </summary>
        public readonly string StartIP;
        /// <summary>
        /// type (of the firewall rule resource = 'Microsoft.Cache/redis/firewallRule')
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRedisFirewallRuleResult(
            string endIP,

            string name,

            string startIP,

            string type)
        {
            EndIP = endIP;
            Name = name;
            StartIP = startIP;
            Type = type;
        }
    }
}
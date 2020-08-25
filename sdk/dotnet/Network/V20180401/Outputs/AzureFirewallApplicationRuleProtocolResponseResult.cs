// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180401.Outputs
{

    [OutputType]
    public sealed class AzureFirewallApplicationRuleProtocolResponseResult
    {
        /// <summary>
        /// Port number for the protocol, cannot be greater than 64000. This field is optional.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Protocol type
        /// </summary>
        public readonly string? ProtocolType;

        [OutputConstructor]
        private AzureFirewallApplicationRuleProtocolResponseResult(
            int? port,

            string? protocolType)
        {
            Port = port;
            ProtocolType = protocolType;
        }
    }
}
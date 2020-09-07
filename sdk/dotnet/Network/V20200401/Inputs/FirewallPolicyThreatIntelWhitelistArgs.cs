// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401.Inputs
{

    /// <summary>
    /// ThreatIntel Whitelist for Firewall Policy.
    /// </summary>
    public sealed class FirewallPolicyThreatIntelWhitelistArgs : Pulumi.ResourceArgs
    {
        [Input("fqdns")]
        private InputList<string>? _fqdns;

        /// <summary>
        /// List of FQDNs for the ThreatIntel Whitelist.
        /// </summary>
        public InputList<string> Fqdns
        {
            get => _fqdns ?? (_fqdns = new InputList<string>());
            set => _fqdns = value;
        }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for the ThreatIntel Whitelist.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        public FirewallPolicyThreatIntelWhitelistArgs()
        {
        }
    }
}
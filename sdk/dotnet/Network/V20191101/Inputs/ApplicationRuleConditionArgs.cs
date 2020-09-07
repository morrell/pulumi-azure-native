// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101.Inputs
{

    /// <summary>
    /// Rule condition of type application.
    /// </summary>
    public sealed class ApplicationRuleConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the rule condition.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationAddresses")]
        private InputList<string>? _destinationAddresses;

        /// <summary>
        /// List of destination IP addresses or Service Tags.
        /// </summary>
        public InputList<string> DestinationAddresses
        {
            get => _destinationAddresses ?? (_destinationAddresses = new InputList<string>());
            set => _destinationAddresses = value;
        }

        [Input("fqdnTags")]
        private InputList<string>? _fqdnTags;

        /// <summary>
        /// List of FQDN Tags for this rule condition.
        /// </summary>
        public InputList<string> FqdnTags
        {
            get => _fqdnTags ?? (_fqdnTags = new InputList<string>());
            set => _fqdnTags = value;
        }

        /// <summary>
        /// Name of the rule condition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("protocols")]
        private InputList<Inputs.FirewallPolicyRuleConditionApplicationProtocolArgs>? _protocols;

        /// <summary>
        /// Array of Application Protocols.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleConditionApplicationProtocolArgs> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<Inputs.FirewallPolicyRuleConditionApplicationProtocolArgs>());
            set => _protocols = value;
        }

        /// <summary>
        /// Rule Condition Type.
        /// </summary>
        [Input("ruleConditionType", required: true)]
        public Input<string> RuleConditionType { get; set; } = null!;

        [Input("sourceAddresses")]
        private InputList<string>? _sourceAddresses;

        /// <summary>
        /// List of source IP addresses for this rule.
        /// </summary>
        public InputList<string> SourceAddresses
        {
            get => _sourceAddresses ?? (_sourceAddresses = new InputList<string>());
            set => _sourceAddresses = value;
        }

        [Input("targetFqdns")]
        private InputList<string>? _targetFqdns;

        /// <summary>
        /// List of FQDNs for this rule condition.
        /// </summary>
        public InputList<string> TargetFqdns
        {
            get => _targetFqdns ?? (_targetFqdns = new InputList<string>());
            set => _targetFqdns = value;
        }

        public ApplicationRuleConditionArgs()
        {
        }
    }
}
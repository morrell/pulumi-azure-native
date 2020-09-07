// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Latest.Inputs
{

    /// <summary>
    /// Rewrite rule set of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayRewriteRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the rewrite rule set that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rewriteRules")]
        private InputList<Inputs.ApplicationGatewayRewriteRuleArgs>? _rewriteRules;

        /// <summary>
        /// Rewrite rules in the rewrite rule set.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayRewriteRuleArgs> RewriteRules
        {
            get => _rewriteRules ?? (_rewriteRules = new InputList<Inputs.ApplicationGatewayRewriteRuleArgs>());
            set => _rewriteRules = value;
        }

        public ApplicationGatewayRewriteRuleSetArgs()
        {
        }
    }
}
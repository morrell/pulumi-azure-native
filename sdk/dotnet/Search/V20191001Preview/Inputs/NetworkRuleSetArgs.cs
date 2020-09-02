// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search.V20191001Preview.Inputs
{

    /// <summary>
    /// Network specific rules that determine how the Azure Cognitive Search service may be reached.
    /// </summary>
    public sealed class NetworkRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The level of access to the search service endpoint. Public, the search service endpoint is reachable from the internet. Private, the search service endpoint can only be accessed via private endpoints. Default is Public.
        /// </summary>
        [Input("endpointAccess")]
        public Input<string>? EndpointAccess { get; set; }

        [Input("ipRules")]
        private InputList<Inputs.IpRuleArgs>? _ipRules;

        /// <summary>
        /// A list of IP restriction rules that defines the inbound network access to the search service endpoint. These restriction rules are applied only when the EndpointAccess of the search service is Public.
        /// </summary>
        public InputList<Inputs.IpRuleArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.IpRuleArgs>());
            set => _ipRules = value;
        }

        public NetworkRuleSetArgs()
        {
        }
    }
}
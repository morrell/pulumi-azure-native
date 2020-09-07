// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.V20190801Preview.Inputs
{

    /// <summary>
    /// The properties that define a direct peering.
    /// </summary>
    public sealed class PeeringPropertiesDirectArgs : Pulumi.ResourceArgs
    {
        [Input("connections")]
        private InputList<Inputs.DirectConnectionArgs>? _connections;

        /// <summary>
        /// The set of connections that constitute a direct peering.
        /// </summary>
        public InputList<Inputs.DirectConnectionArgs> Connections
        {
            get => _connections ?? (_connections = new InputList<Inputs.DirectConnectionArgs>());
            set => _connections = value;
        }

        /// <summary>
        /// The type of direct peering.
        /// </summary>
        [Input("directPeeringType")]
        public Input<string>? DirectPeeringType { get; set; }

        /// <summary>
        /// The reference of the peer ASN.
        /// </summary>
        [Input("peerAsn")]
        public Input<Inputs.SubResourceArgs>? PeerAsn { get; set; }

        /// <summary>
        /// The flag that indicates whether or not the peering is used for peering service.
        /// </summary>
        [Input("useForPeeringService")]
        public Input<bool>? UseForPeeringService { get; set; }

        public PeeringPropertiesDirectArgs()
        {
        }
    }
}
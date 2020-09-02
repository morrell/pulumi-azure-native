// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.V20190801Preview.Outputs
{

    [OutputType]
    public sealed class PeeringPropertiesExchangeResponseResult
    {
        /// <summary>
        /// The set of connections that constitute an exchange peering.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExchangeConnectionResponseResult> Connections;
        /// <summary>
        /// The reference of the peer ASN.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? PeerAsn;

        [OutputConstructor]
        private PeeringPropertiesExchangeResponseResult(
            ImmutableArray<Outputs.ExchangeConnectionResponseResult> connections,

            Outputs.SubResourceResponseResult? peerAsn)
        {
            Connections = connections;
            PeerAsn = peerAsn;
        }
    }
}
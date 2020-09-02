// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class DirectConnectionResponseResult
    {
        /// <summary>
        /// The bandwidth of the connection.
        /// </summary>
        public readonly int? BandwidthInMbps;
        /// <summary>
        /// The BGP session associated with the connection.
        /// </summary>
        public readonly Outputs.BgpSessionResponseResult? BgpSession;
        /// <summary>
        /// The unique identifier (GUID) for the connection.
        /// </summary>
        public readonly string? ConnectionIdentifier;
        /// <summary>
        /// The state of the connection.
        /// </summary>
        public readonly string ConnectionState;
        /// <summary>
        /// The error message related to the connection state, if any.
        /// </summary>
        public readonly string ErrorMessage;
        /// <summary>
        /// The PeeringDB.com ID of the facility at which the connection has to be set up.
        /// </summary>
        public readonly int? PeeringDBFacilityId;
        /// <summary>
        /// The bandwidth that is actually provisioned.
        /// </summary>
        public readonly int ProvisionedBandwidthInMbps;
        /// <summary>
        /// The field indicating if Microsoft provides session ip addresses.
        /// </summary>
        public readonly string? SessionAddressProvider;
        /// <summary>
        /// The flag that indicates whether or not the connection is used for peering service.
        /// </summary>
        public readonly bool? UseForPeeringService;

        [OutputConstructor]
        private DirectConnectionResponseResult(
            int? bandwidthInMbps,

            Outputs.BgpSessionResponseResult? bgpSession,

            string? connectionIdentifier,

            string connectionState,

            string errorMessage,

            int? peeringDBFacilityId,

            int provisionedBandwidthInMbps,

            string? sessionAddressProvider,

            bool? useForPeeringService)
        {
            BandwidthInMbps = bandwidthInMbps;
            BgpSession = bgpSession;
            ConnectionIdentifier = connectionIdentifier;
            ConnectionState = connectionState;
            ErrorMessage = errorMessage;
            PeeringDBFacilityId = peeringDBFacilityId;
            ProvisionedBandwidthInMbps = provisionedBandwidthInMbps;
            SessionAddressProvider = sessionAddressProvider;
            UseForPeeringService = useForPeeringService;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20200201Preview.Outputs
{

    [OutputType]
    public sealed class MediaGraphRtspSourceResponseResult
    {
        /// <summary>
        /// RTSP endpoint of the stream being connected to.
        /// </summary>
        public readonly Union<Outputs.MediaGraphClearEndpointResponseResult, Outputs.MediaGraphTlsEndpointResponseResult> Endpoint;
        /// <summary>
        /// Source name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The discriminator for derived types.
        /// </summary>
        public readonly string OdataType;
        /// <summary>
        /// Underlying RTSP transport. This can be used to enable or disable HTTP tunneling.
        /// </summary>
        public readonly string Transport;

        [OutputConstructor]
        private MediaGraphRtspSourceResponseResult(
            Union<Outputs.MediaGraphClearEndpointResponseResult, Outputs.MediaGraphTlsEndpointResponseResult> endpoint,

            string name,

            string odataType,

            string transport)
        {
            Endpoint = endpoint;
            Name = name;
            OdataType = odataType;
            Transport = transport;
        }
    }
}
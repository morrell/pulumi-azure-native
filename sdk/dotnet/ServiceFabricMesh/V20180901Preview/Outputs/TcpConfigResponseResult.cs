// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180901Preview.Outputs
{

    [OutputType]
    public sealed class TcpConfigResponseResult
    {
        /// <summary>
        /// Describes destination endpoint for routing traffic.
        /// </summary>
        public readonly Outputs.GatewayDestinationResponseResult Destination;
        /// <summary>
        /// tcp gateway config name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the port at which the service endpoint below needs to be exposed.
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private TcpConfigResponseResult(
            Outputs.GatewayDestinationResponseResult destination,

            string name,

            int port)
        {
            Destination = destination;
            Name = name;
            Port = port;
        }
    }
}
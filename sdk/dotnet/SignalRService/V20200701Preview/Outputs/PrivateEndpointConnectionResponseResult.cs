// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SignalRService.V20200701Preview.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointConnectionResponseResult
    {
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private endpoint associated with the private endpoint connection
        /// </summary>
        public readonly Outputs.PrivateEndpointResponseResult? PrivateEndpoint;
        /// <summary>
        /// Connection state
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStateResponseResult? PrivateLinkServiceConnectionState;
        /// <summary>
        /// Provisioning state of the private endpoint connection
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PrivateEndpointConnectionResponseResult(
            string id,

            string name,

            Outputs.PrivateEndpointResponseResult? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStateResponseResult? privateLinkServiceConnectionState,

            string provisioningState,

            string type)
        {
            Id = id;
            Name = name;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
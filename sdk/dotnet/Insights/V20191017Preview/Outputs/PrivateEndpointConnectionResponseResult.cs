// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20191017Preview.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointConnectionResponseResult
    {
        /// <summary>
        /// Azure resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private endpoint which the connection belongs to.
        /// </summary>
        public readonly Outputs.PrivateEndpointPropertyResponseResult? PrivateEndpoint;
        /// <summary>
        /// Connection state of the private endpoint connection.
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStatePropertyResponseResult? PrivateLinkServiceConnectionState;
        /// <summary>
        /// State of the private endpoint connection.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PrivateEndpointConnectionResponseResult(
            string id,

            string name,

            Outputs.PrivateEndpointPropertyResponseResult? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStatePropertyResponseResult? privateLinkServiceConnectionState,

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
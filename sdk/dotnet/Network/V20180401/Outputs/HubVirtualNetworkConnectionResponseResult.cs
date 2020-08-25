// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180401.Outputs
{

    [OutputType]
    public sealed class HubVirtualNetworkConnectionResponseResult
    {
        /// <summary>
        /// VirtualHub to RemoteVnet transit to enabled or not.
        /// </summary>
        public readonly bool? AllowHubToRemoteVnetTransit;
        /// <summary>
        /// Allow RemoteVnet to use Virtual Hub's gateways.
        /// </summary>
        public readonly bool? AllowRemoteVnetToUseHubVnetGateways;
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Reference to the remote virtual network.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? RemoteVirtualNetwork;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private HubVirtualNetworkConnectionResponseResult(
            bool? allowHubToRemoteVnetTransit,

            bool? allowRemoteVnetToUseHubVnetGateways,

            string etag,

            string? id,

            string? location,

            string name,

            string? provisioningState,

            Outputs.SubResourceResponseResult? remoteVirtualNetwork,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AllowHubToRemoteVnetTransit = allowHubToRemoteVnetTransit;
            AllowRemoteVnetToUseHubVnetGateways = allowRemoteVnetToUseHubVnetGateways;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            RemoteVirtualNetwork = remoteVirtualNetwork;
            Tags = tags;
            Type = type;
        }
    }
}
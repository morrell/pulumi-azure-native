// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190901
{
    public static class GetNatGateway
    {
        public static Task<GetNatGatewayResult> InvokeAsync(GetNatGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNatGatewayResult>("azurerm:network/v20190901:getNatGateway", args ?? new GetNatGatewayArgs(), options.WithVersion());
    }


    public sealed class GetNatGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the nat gateway.
        /// </summary>
        [Input("natGatewayName", required: true)]
        public string NatGatewayName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNatGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNatGatewayResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The idle timeout of the nat gateway.
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the NAT gateway resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// An array of public ip addresses associated with the nat gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> PublicIpAddresses;
        /// <summary>
        /// An array of public ip prefixes associated with the nat gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> PublicIpPrefixes;
        /// <summary>
        /// The resource GUID property of the NAT gateway resource.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// The nat gateway SKU.
        /// </summary>
        public readonly Outputs.NatGatewaySkuResponseResult? Sku;
        /// <summary>
        /// An array of references to the subnets using this nat gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> Subnets;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetNatGatewayResult(
            string etag,

            int? idleTimeoutInMinutes,

            string? location,

            string name,

            string provisioningState,

            ImmutableArray<Outputs.SubResourceResponseResult> publicIpAddresses,

            ImmutableArray<Outputs.SubResourceResponseResult> publicIpPrefixes,

            string resourceGuid,

            Outputs.NatGatewaySkuResponseResult? sku,

            ImmutableArray<Outputs.SubResourceResponseResult> subnets,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<string> zones)
        {
            Etag = etag;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            PublicIpAddresses = publicIpAddresses;
            PublicIpPrefixes = publicIpPrefixes;
            ResourceGuid = resourceGuid;
            Sku = sku;
            Subnets = subnets;
            Tags = tags;
            Type = type;
            Zones = zones;
        }
    }
}
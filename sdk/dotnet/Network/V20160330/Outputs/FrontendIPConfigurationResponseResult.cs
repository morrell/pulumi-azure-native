// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330.Outputs
{

    [OutputType]
    public sealed class FrontendIPConfigurationResponseResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Read only.Inbound pools URIs that use this frontend IP
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> InboundNatPools;
        /// <summary>
        /// Read only.Inbound rules URIs that use this frontend IP
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> InboundNatRules;
        /// <summary>
        /// Gets Load Balancing rules URIs that use this frontend IP
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> LoadBalancingRules;
        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Read only.Outbound rules URIs that use this frontend IP
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> OutboundNatRules;
        /// <summary>
        /// Gets or sets the privateIPAddress of the IP Configuration
        /// </summary>
        public readonly string? PrivateIPAddress;
        /// <summary>
        /// Gets or sets PrivateIP allocation method (Static/Dynamic)
        /// </summary>
        public readonly string? PrivateIPAllocationMethod;
        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets the reference of the PublicIP resource
        /// </summary>
        public readonly Outputs.PublicIPAddressResponseResult? PublicIPAddress;
        /// <summary>
        /// Gets or sets the reference of the subnet resource
        /// </summary>
        public readonly Outputs.SubnetResponseResult? Subnet;

        [OutputConstructor]
        private FrontendIPConfigurationResponseResult(
            string? etag,

            string? id,

            ImmutableArray<Outputs.SubResourceResponseResult> inboundNatPools,

            ImmutableArray<Outputs.SubResourceResponseResult> inboundNatRules,

            ImmutableArray<Outputs.SubResourceResponseResult> loadBalancingRules,

            string? name,

            ImmutableArray<Outputs.SubResourceResponseResult> outboundNatRules,

            string? privateIPAddress,

            string? privateIPAllocationMethod,

            string? provisioningState,

            Outputs.PublicIPAddressResponseResult? publicIPAddress,

            Outputs.SubnetResponseResult? subnet)
        {
            Etag = etag;
            Id = id;
            InboundNatPools = inboundNatPools;
            InboundNatRules = inboundNatRules;
            LoadBalancingRules = loadBalancingRules;
            Name = name;
            OutboundNatRules = outboundNatRules;
            PrivateIPAddress = privateIPAddress;
            PrivateIPAllocationMethod = privateIPAllocationMethod;
            ProvisioningState = provisioningState;
            PublicIPAddress = publicIPAddress;
            Subnet = subnet;
        }
    }
}
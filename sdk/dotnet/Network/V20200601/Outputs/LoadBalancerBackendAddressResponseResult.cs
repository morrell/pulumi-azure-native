// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200601.Outputs
{

    [OutputType]
    public sealed class LoadBalancerBackendAddressResponseResult
    {
        /// <summary>
        /// IP Address belonging to the referenced virtual network.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// Name of the backend address.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Reference to IP address defined in network interfaces.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult NetworkInterfaceIPConfiguration;
        /// <summary>
        /// Reference to an existing virtual network.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? VirtualNetwork;

        [OutputConstructor]
        private LoadBalancerBackendAddressResponseResult(
            string? ipAddress,

            string? name,

            Outputs.SubResourceResponseResult networkInterfaceIPConfiguration,

            Outputs.SubResourceResponseResult? virtualNetwork)
        {
            IpAddress = ipAddress;
            Name = name;
            NetworkInterfaceIPConfiguration = networkInterfaceIPConfiguration;
            VirtualNetwork = virtualNetwork;
        }
    }
}
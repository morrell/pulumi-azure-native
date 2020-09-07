// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20161201.Outputs
{

    [OutputType]
    public sealed class NetworkInterfaceResponseResult
    {
        /// <summary>
        /// The DNS settings in network interface.
        /// </summary>
        public readonly Outputs.NetworkInterfaceDnsSettingsResponseResult? DnsSettings;
        /// <summary>
        /// If the network interface is accelerated networking enabled.
        /// </summary>
        public readonly bool? EnableAcceleratedNetworking;
        /// <summary>
        /// Indicates whether IP forwarding is enabled on this network interface.
        /// </summary>
        public readonly bool? EnableIPForwarding;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// A list of IPConfigurations of the network interface.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The MAC address of the network interface.
        /// </summary>
        public readonly string? MacAddress;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        public readonly Outputs.NetworkSecurityGroupResponseResult? NetworkSecurityGroup;
        /// <summary>
        /// Gets whether this is a primary network interface on a virtual machine.
        /// </summary>
        public readonly bool? Primary;
        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The resource GUID property of the network interface resource.
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The reference of a virtual machine.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? VirtualMachine;

        [OutputConstructor]
        private NetworkInterfaceResponseResult(
            Outputs.NetworkInterfaceDnsSettingsResponseResult? dnsSettings,

            bool? enableAcceleratedNetworking,

            bool? enableIPForwarding,

            string? etag,

            string? id,

            ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponseResult> ipConfigurations,

            string? location,

            string? macAddress,

            string name,

            Outputs.NetworkSecurityGroupResponseResult? networkSecurityGroup,

            bool? primary,

            string? provisioningState,

            string? resourceGuid,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.SubResourceResponseResult? virtualMachine)
        {
            DnsSettings = dnsSettings;
            EnableAcceleratedNetworking = enableAcceleratedNetworking;
            EnableIPForwarding = enableIPForwarding;
            Etag = etag;
            Id = id;
            IpConfigurations = ipConfigurations;
            Location = location;
            MacAddress = macAddress;
            Name = name;
            NetworkSecurityGroup = networkSecurityGroup;
            Primary = primary;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            Tags = tags;
            Type = type;
            VirtualMachine = virtualMachine;
        }
    }
}
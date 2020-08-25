// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Outputs
{

    [OutputType]
    public sealed class SubnetResponseResult
    {
        /// <summary>
        /// The address prefix for the subnet.
        /// </summary>
        public readonly string? AddressPrefix;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Gets an array of references to the network interface IP configurations using subnet.
        /// </summary>
        public readonly ImmutableArray<Outputs.IPConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        public readonly Outputs.NetworkSecurityGroupResponseResult? NetworkSecurityGroup;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets an array of references to the external resources using subnet.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceNavigationLinkResponseResult> ResourceNavigationLinks;
        /// <summary>
        /// The reference of the RouteTable resource.
        /// </summary>
        public readonly Outputs.RouteTableResponseResult? RouteTable;
        /// <summary>
        /// An array of service endpoints.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceEndpointPropertiesFormatResponseResult> ServiceEndpoints;

        [OutputConstructor]
        private SubnetResponseResult(
            string? addressPrefix,

            string? etag,

            string? id,

            ImmutableArray<Outputs.IPConfigurationResponseResult> ipConfigurations,

            string? name,

            Outputs.NetworkSecurityGroupResponseResult? networkSecurityGroup,

            string? provisioningState,

            ImmutableArray<Outputs.ResourceNavigationLinkResponseResult> resourceNavigationLinks,

            Outputs.RouteTableResponseResult? routeTable,

            ImmutableArray<Outputs.ServiceEndpointPropertiesFormatResponseResult> serviceEndpoints)
        {
            AddressPrefix = addressPrefix;
            Etag = etag;
            Id = id;
            IpConfigurations = ipConfigurations;
            Name = name;
            NetworkSecurityGroup = networkSecurityGroup;
            ProvisioningState = provisioningState;
            ResourceNavigationLinks = resourceNavigationLinks;
            RouteTable = routeTable;
            ServiceEndpoints = serviceEndpoints;
        }
    }
}
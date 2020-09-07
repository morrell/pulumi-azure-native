// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330
{
    public static class GetSubnet
    {
        public static Task<GetSubnetResult> InvokeAsync(GetSubnetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetResult>("azurerm:network/v20160330:getSubnet", args ?? new GetSubnetArgs(), options.WithVersion());
    }


    public sealed class GetSubnetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// expand references resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the subnet.
        /// </summary>
        [Input("subnetName", required: true)]
        public string SubnetName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network.
        /// </summary>
        [Input("virtualNetworkName", required: true)]
        public string VirtualNetworkName { get; set; } = null!;

        public GetSubnetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubnetResult
    {
        /// <summary>
        /// Gets or sets Address prefix for the subnet.
        /// </summary>
        public readonly string? AddressPrefix;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Gets array of references to the network interface IP configurations using subnet
        /// </summary>
        public readonly ImmutableArray<Outputs.IPConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Gets or sets the reference of the NetworkSecurityGroup resource
        /// </summary>
        public readonly Outputs.NetworkSecurityGroupResponseResult? NetworkSecurityGroup;
        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets the reference of the RouteTable resource
        /// </summary>
        public readonly Outputs.RouteTableResponseResult? RouteTable;

        [OutputConstructor]
        private GetSubnetResult(
            string? addressPrefix,

            string? etag,

            ImmutableArray<Outputs.IPConfigurationResponseResult> ipConfigurations,

            string? name,

            Outputs.NetworkSecurityGroupResponseResult? networkSecurityGroup,

            string? provisioningState,

            Outputs.RouteTableResponseResult? routeTable)
        {
            AddressPrefix = addressPrefix;
            Etag = etag;
            IpConfigurations = ipConfigurations;
            Name = name;
            NetworkSecurityGroup = networkSecurityGroup;
            ProvisioningState = provisioningState;
            RouteTable = routeTable;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridNetwork.V20200101Preview
{
    public static class GetDevice
    {
        public static Task<GetDeviceResult> InvokeAsync(GetDeviceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeviceResult>("azurerm:hybridnetwork/v20200101preview:getDevice", args ?? new GetDeviceArgs(), options.WithVersion());
    }


    public sealed class GetDeviceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of hybrid network device.
        /// </summary>
        [Input("deviceName", required: true)]
        public string DeviceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDeviceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDeviceResult
    {
        /// <summary>
        /// The reference to the azure stack edge device.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? AzureStackEdge;
        /// <summary>
        /// The type of the hybrid network device.
        /// </summary>
        public readonly string DeviceType;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the hybrid network device resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The current device status.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The list of virtual network functions deployed on the hybrid network device.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> VirtualNetworkFunctions;

        [OutputConstructor]
        private GetDeviceResult(
            Outputs.SubResourceResponseResult? azureStackEdge,

            string deviceType,

            string? location,

            string name,

            string provisioningState,

            string status,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.SubResourceResponseResult> virtualNetworkFunctions)
        {
            AzureStackEdge = azureStackEdge;
            DeviceType = deviceType;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Status = status;
            Tags = tags;
            Type = type;
            VirtualNetworkFunctions = virtualNetworkFunctions;
        }
    }
}
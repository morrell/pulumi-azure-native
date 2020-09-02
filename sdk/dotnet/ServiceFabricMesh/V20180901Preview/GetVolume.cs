// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180901Preview
{
    public static class GetVolume
    {
        public static Task<GetVolumeResult> InvokeAsync(GetVolumeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("azurerm:servicefabricmesh/v20180901preview:getVolume", args ?? new GetVolumeArgs(), options.WithVersion());
    }


    public sealed class GetVolumeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Azure resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The identity of the volume.
        /// </summary>
        [Input("volumeResourceName", required: true)]
        public string VolumeResourceName { get; set; } = null!;

        public GetVolumeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVolumeResult
    {
        /// <summary>
        /// This type describes a volume provided by an Azure Files file share.
        /// </summary>
        public readonly Outputs.VolumeProviderParametersAzureFileResponseResult? AzureFileParameters;
        /// <summary>
        /// User readable description of the volume.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provider of the volume.
        /// </summary>
        public readonly string Provider;
        /// <summary>
        /// State of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Status of the volume.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Gives additional information about the current status of the volume.
        /// </summary>
        public readonly string StatusDetails;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVolumeResult(
            Outputs.VolumeProviderParametersAzureFileResponseResult? azureFileParameters,

            string? description,

            string location,

            string name,

            string provider,

            string provisioningState,

            string status,

            string statusDetails,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureFileParameters = azureFileParameters;
            Description = description;
            Location = location;
            Name = name;
            Provider = provider;
            ProvisioningState = provisioningState;
            Status = status;
            StatusDetails = statusDetails;
            Tags = tags;
            Type = type;
        }
    }
}
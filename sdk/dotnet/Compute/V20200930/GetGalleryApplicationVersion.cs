// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20200930
{
    public static class GetGalleryApplicationVersion
    {
        public static Task<GetGalleryApplicationVersionResult> InvokeAsync(GetGalleryApplicationVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGalleryApplicationVersionResult>("azurerm:compute/v20200930:getGalleryApplicationVersion", args ?? new GetGalleryApplicationVersionArgs(), options.WithVersion());
    }


    public sealed class GetGalleryApplicationVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The expand expression to apply on the operation.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the gallery Application Definition in which the Application Version resides.
        /// </summary>
        [Input("galleryApplicationName", required: true)]
        public string GalleryApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the gallery Application Version to be retrieved.
        /// </summary>
        [Input("galleryApplicationVersionName", required: true)]
        public string GalleryApplicationVersionName { get; set; } = null!;

        /// <summary>
        /// The name of the Shared Application Gallery in which the Application Definition resides.
        /// </summary>
        [Input("galleryName", required: true)]
        public string GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetGalleryApplicationVersionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGalleryApplicationVersionResult
    {
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The publishing profile of a gallery image version.
        /// </summary>
        public readonly Outputs.GalleryApplicationVersionPublishingProfileResponseResult PublishingProfile;
        /// <summary>
        /// This is the replication status of the gallery image version.
        /// </summary>
        public readonly Outputs.ReplicationStatusResponseResult ReplicationStatus;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGalleryApplicationVersionResult(
            string location,

            string name,

            string provisioningState,

            Outputs.GalleryApplicationVersionPublishingProfileResponseResult publishingProfile,

            Outputs.ReplicationStatusResponseResult replicationStatus,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            PublishingProfile = publishingProfile;
            ReplicationStatus = replicationStatus;
            Tags = tags;
            Type = type;
        }
    }
}
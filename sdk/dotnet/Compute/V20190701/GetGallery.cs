// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190701
{
    public static class GetGallery
    {
        public static Task<GetGalleryResult> InvokeAsync(GetGalleryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGalleryResult>("azurerm:compute/v20190701:getGallery", args ?? new GetGalleryArgs(), options.WithVersion());
    }


    public sealed class GetGalleryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Image Gallery.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetGalleryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGalleryResult
    {
        /// <summary>
        /// The description of this Shared Image Gallery resource. This property is updatable.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Describes the gallery unique name.
        /// </summary>
        public readonly Outputs.GalleryIdentifierResponseResult? Identifier;
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
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGalleryResult(
            string? description,

            Outputs.GalleryIdentifierResponseResult? identifier,

            string location,

            string name,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Description = description;
            Identifier = identifier;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
        }
    }
}
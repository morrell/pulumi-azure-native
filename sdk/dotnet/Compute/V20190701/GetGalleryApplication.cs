// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190701
{
    public static class GetGalleryApplication
    {
        public static Task<GetGalleryApplicationResult> InvokeAsync(GetGalleryApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGalleryApplicationResult>("azurerm:compute/v20190701:getGalleryApplication", args ?? new GetGalleryApplicationArgs(), options.WithVersion());
    }


    public sealed class GetGalleryApplicationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Application Gallery from which the Application Definitions are to be retrieved.
        /// </summary>
        [Input("galleryName", required: true)]
        public string GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the gallery Application Definition to be retrieved.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetGalleryApplicationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGalleryApplicationResult
    {
        /// <summary>
        /// The description of this gallery Application Definition resource. This property is updatable.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
        /// </summary>
        public readonly string? EndOfLifeDate;
        /// <summary>
        /// The Eula agreement for the gallery Application Definition.
        /// </summary>
        public readonly string? Eula;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The privacy statement uri.
        /// </summary>
        public readonly string? PrivacyStatementUri;
        /// <summary>
        /// The release note uri.
        /// </summary>
        public readonly string? ReleaseNoteUri;
        /// <summary>
        /// This property allows you to specify the supported type of the OS that application is built for. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **Windows** &lt;br&gt;&lt;br&gt; **Linux**
        /// </summary>
        public readonly string SupportedOSType;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGalleryApplicationResult(
            string? description,

            string? endOfLifeDate,

            string? eula,

            string location,

            string name,

            string? privacyStatementUri,

            string? releaseNoteUri,

            string supportedOSType,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Description = description;
            EndOfLifeDate = endOfLifeDate;
            Eula = eula;
            Location = location;
            Name = name;
            PrivacyStatementUri = privacyStatementUri;
            ReleaseNoteUri = releaseNoteUri;
            SupportedOSType = supportedOSType;
            Tags = tags;
            Type = type;
        }
    }
}
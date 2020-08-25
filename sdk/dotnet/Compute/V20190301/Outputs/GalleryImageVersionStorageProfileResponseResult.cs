// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190301.Outputs
{

    [OutputType]
    public sealed class GalleryImageVersionStorageProfileResponseResult
    {
        /// <summary>
        /// A list of data disk images.
        /// </summary>
        public readonly ImmutableArray<Outputs.GalleryDataDiskImageResponseResult> DataDiskImages;
        /// <summary>
        /// This is the OS disk image.
        /// </summary>
        public readonly Outputs.GalleryOSDiskImageResponseResult OsDiskImage;

        [OutputConstructor]
        private GalleryImageVersionStorageProfileResponseResult(
            ImmutableArray<Outputs.GalleryDataDiskImageResponseResult> dataDiskImages,

            Outputs.GalleryOSDiskImageResponseResult osDiskImage)
        {
            DataDiskImages = dataDiskImages;
            OsDiskImage = osDiskImage;
        }
    }
}
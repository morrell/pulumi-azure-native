// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Outputs
{

    [OutputType]
    public sealed class VirtualMachineScaleSetStorageProfileResponseResult
    {
        /// <summary>
        /// The data disks.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineScaleSetDataDiskResponseResult> DataDisks;
        /// <summary>
        /// The image reference.
        /// </summary>
        public readonly Outputs.ImageReferenceResponseResult? ImageReference;
        /// <summary>
        /// The OS disk.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetOSDiskResponseResult? OsDisk;

        [OutputConstructor]
        private VirtualMachineScaleSetStorageProfileResponseResult(
            ImmutableArray<Outputs.VirtualMachineScaleSetDataDiskResponseResult> dataDisks,

            Outputs.ImageReferenceResponseResult? imageReference,

            Outputs.VirtualMachineScaleSetOSDiskResponseResult? osDisk)
        {
            DataDisks = dataDisks;
            ImageReference = imageReference;
            OsDisk = osDisk;
        }
    }
}
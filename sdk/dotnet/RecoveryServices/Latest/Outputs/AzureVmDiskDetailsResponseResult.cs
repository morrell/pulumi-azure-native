// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.Latest.Outputs
{

    [OutputType]
    public sealed class AzureVmDiskDetailsResponseResult
    {
        /// <summary>
        /// The DiskEncryptionSet ARM ID.
        /// </summary>
        public readonly string? DiskEncryptionSetId;
        /// <summary>
        /// The disk resource id.
        /// </summary>
        public readonly string? DiskId;
        /// <summary>
        /// Ordinal\LunId of the disk for the Azure VM.
        /// </summary>
        public readonly string? LunId;
        /// <summary>
        /// Max side in MB.
        /// </summary>
        public readonly string? MaxSizeMB;
        /// <summary>
        /// Blob uri of the Azure disk.
        /// </summary>
        public readonly string? TargetDiskLocation;
        /// <summary>
        /// The target Azure disk name.
        /// </summary>
        public readonly string? TargetDiskName;
        /// <summary>
        /// The VHD id.
        /// </summary>
        public readonly string? VhdId;
        /// <summary>
        /// VHD name.
        /// </summary>
        public readonly string? VhdName;
        /// <summary>
        /// VHD type.
        /// </summary>
        public readonly string? VhdType;

        [OutputConstructor]
        private AzureVmDiskDetailsResponseResult(
            string? diskEncryptionSetId,

            string? diskId,

            string? lunId,

            string? maxSizeMB,

            string? targetDiskLocation,

            string? targetDiskName,

            string? vhdId,

            string? vhdName,

            string? vhdType)
        {
            DiskEncryptionSetId = diskEncryptionSetId;
            DiskId = diskId;
            LunId = lunId;
            MaxSizeMB = maxSizeMB;
            TargetDiskLocation = targetDiskLocation;
            TargetDiskName = targetDiskName;
            VhdId = vhdId;
            VhdName = vhdName;
            VhdType = vhdType;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20160810.Outputs
{

    [OutputType]
    public sealed class InMageProtectedDiskDetailsResponseResult
    {
        /// <summary>
        /// The disk capacity in bytes.
        /// </summary>
        public readonly int? DiskCapacityInBytes;
        /// <summary>
        /// The disk id.
        /// </summary>
        public readonly string? DiskId;
        /// <summary>
        /// The disk name.
        /// </summary>
        public readonly string? DiskName;
        /// <summary>
        /// A value indicating whether disk is resized.
        /// </summary>
        public readonly string? DiskResized;
        /// <summary>
        /// The file system capacity in bytes.
        /// </summary>
        public readonly int? FileSystemCapacityInBytes;
        /// <summary>
        /// The health error code for the disk.
        /// </summary>
        public readonly string? HealthErrorCode;
        /// <summary>
        /// The last RPO calculated time.
        /// </summary>
        public readonly string? LastRpoCalculatedTime;
        /// <summary>
        /// The protection stage.
        /// </summary>
        public readonly string? ProtectionStage;
        /// <summary>
        /// The PS data transit in MB.
        /// </summary>
        public readonly double? PsDataInMB;
        /// <summary>
        /// The resync duration in seconds.
        /// </summary>
        public readonly int? ResyncDurationInSeconds;
        /// <summary>
        /// The resync progress percentage.
        /// </summary>
        public readonly int? ResyncProgressPercentage;
        /// <summary>
        /// A value indicating whether resync is required for this disk.
        /// </summary>
        public readonly string? ResyncRequired;
        /// <summary>
        /// The RPO in seconds.
        /// </summary>
        public readonly int? RpoInSeconds;
        /// <summary>
        /// The source data transit in MB.
        /// </summary>
        public readonly double? SourceDataInMB;
        /// <summary>
        /// The target data transit in MB.
        /// </summary>
        public readonly double? TargetDataInMB;

        [OutputConstructor]
        private InMageProtectedDiskDetailsResponseResult(
            int? diskCapacityInBytes,

            string? diskId,

            string? diskName,

            string? diskResized,

            int? fileSystemCapacityInBytes,

            string? healthErrorCode,

            string? lastRpoCalculatedTime,

            string? protectionStage,

            double? psDataInMB,

            int? resyncDurationInSeconds,

            int? resyncProgressPercentage,

            string? resyncRequired,

            int? rpoInSeconds,

            double? sourceDataInMB,

            double? targetDataInMB)
        {
            DiskCapacityInBytes = diskCapacityInBytes;
            DiskId = diskId;
            DiskName = diskName;
            DiskResized = diskResized;
            FileSystemCapacityInBytes = fileSystemCapacityInBytes;
            HealthErrorCode = healthErrorCode;
            LastRpoCalculatedTime = lastRpoCalculatedTime;
            ProtectionStage = protectionStage;
            PsDataInMB = psDataInMB;
            ResyncDurationInSeconds = resyncDurationInSeconds;
            ResyncProgressPercentage = resyncProgressPercentage;
            ResyncRequired = resyncRequired;
            RpoInSeconds = rpoInSeconds;
            SourceDataInMB = sourceDataInMB;
            TargetDataInMB = targetDataInMB;
        }
    }
}
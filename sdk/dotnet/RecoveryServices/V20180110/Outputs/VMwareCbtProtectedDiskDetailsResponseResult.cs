// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20180110.Outputs
{

    [OutputType]
    public sealed class VMwareCbtProtectedDiskDetailsResponseResult
    {
        /// <summary>
        /// The disk capacity in bytes.
        /// </summary>
        public readonly int CapacityInBytes;
        /// <summary>
        /// The DiskEncryptionSet ARM Id.
        /// </summary>
        public readonly string DiskEncryptionSetId;
        /// <summary>
        /// The disk id.
        /// </summary>
        public readonly string DiskId;
        /// <summary>
        /// The disk name.
        /// </summary>
        public readonly string DiskName;
        /// <summary>
        /// The disk path.
        /// </summary>
        public readonly string DiskPath;
        /// <summary>
        /// The disk type.
        /// </summary>
        public readonly string? DiskType;
        /// <summary>
        /// A value indicating whether the disk is the OS disk.
        /// </summary>
        public readonly string IsOSDisk;
        /// <summary>
        /// The log storage account ARM Id.
        /// </summary>
        public readonly string LogStorageAccountId;
        /// <summary>
        /// The key vault secret name of the log storage account.
        /// </summary>
        public readonly string LogStorageAccountSasSecretName;
        /// <summary>
        /// The ARM Id of the seed managed disk.
        /// </summary>
        public readonly string SeedManagedDiskId;
        /// <summary>
        /// The ARM Id of the target managed disk.
        /// </summary>
        public readonly string TargetManagedDiskId;

        [OutputConstructor]
        private VMwareCbtProtectedDiskDetailsResponseResult(
            int capacityInBytes,

            string diskEncryptionSetId,

            string diskId,

            string diskName,

            string diskPath,

            string? diskType,

            string isOSDisk,

            string logStorageAccountId,

            string logStorageAccountSasSecretName,

            string seedManagedDiskId,

            string targetManagedDiskId)
        {
            CapacityInBytes = capacityInBytes;
            DiskEncryptionSetId = diskEncryptionSetId;
            DiskId = diskId;
            DiskName = diskName;
            DiskPath = diskPath;
            DiskType = diskType;
            IsOSDisk = isOSDisk;
            LogStorageAccountId = logStorageAccountId;
            LogStorageAccountSasSecretName = logStorageAccountSasSecretName;
            SeedManagedDiskId = seedManagedDiskId;
            TargetManagedDiskId = targetManagedDiskId;
        }
    }
}
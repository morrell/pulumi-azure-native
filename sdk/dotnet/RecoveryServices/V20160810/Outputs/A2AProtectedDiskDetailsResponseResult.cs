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
    public sealed class A2AProtectedDiskDetailsResponseResult
    {
        /// <summary>
        /// The data pending at source virtual machine in MB.
        /// </summary>
        public readonly double? DataPendingAtSourceAgentInMB;
        /// <summary>
        /// The data pending for replication in MB at staging account.
        /// </summary>
        public readonly double? DataPendingInStagingStorageAccountInMB;
        /// <summary>
        /// The disk capacity in bytes.
        /// </summary>
        public readonly int? DiskCapacityInBytes;
        /// <summary>
        /// The disk name.
        /// </summary>
        public readonly string? DiskName;
        /// <summary>
        /// The type of disk.
        /// </summary>
        public readonly string? DiskType;
        /// <summary>
        /// The disk uri.
        /// </summary>
        public readonly string? DiskUri;
        /// <summary>
        /// The type of the monitoring job. The progress is contained in MonitoringPercentageCompletion property.
        /// </summary>
        public readonly string? MonitoringJobType;
        /// <summary>
        /// The percentage of the monitoring job. The type of the monitoring job is defined by MonitoringJobType property.
        /// </summary>
        public readonly int? MonitoringPercentageCompletion;
        /// <summary>
        /// The primary disk storage account.
        /// </summary>
        public readonly string? PrimaryDiskAzureStorageAccountId;
        /// <summary>
        /// The primary staging storage account.
        /// </summary>
        public readonly string? PrimaryStagingAzureStorageAccountId;
        /// <summary>
        /// The recovery disk storage account.
        /// </summary>
        public readonly string? RecoveryAzureStorageAccountId;
        /// <summary>
        /// Recovery disk uri.
        /// </summary>
        public readonly string? RecoveryDiskUri;
        /// <summary>
        /// A value indicating whether resync is required for this disk.
        /// </summary>
        public readonly bool? ResyncRequired;

        [OutputConstructor]
        private A2AProtectedDiskDetailsResponseResult(
            double? dataPendingAtSourceAgentInMB,

            double? dataPendingInStagingStorageAccountInMB,

            int? diskCapacityInBytes,

            string? diskName,

            string? diskType,

            string? diskUri,

            string? monitoringJobType,

            int? monitoringPercentageCompletion,

            string? primaryDiskAzureStorageAccountId,

            string? primaryStagingAzureStorageAccountId,

            string? recoveryAzureStorageAccountId,

            string? recoveryDiskUri,

            bool? resyncRequired)
        {
            DataPendingAtSourceAgentInMB = dataPendingAtSourceAgentInMB;
            DataPendingInStagingStorageAccountInMB = dataPendingInStagingStorageAccountInMB;
            DiskCapacityInBytes = diskCapacityInBytes;
            DiskName = diskName;
            DiskType = diskType;
            DiskUri = diskUri;
            MonitoringJobType = monitoringJobType;
            MonitoringPercentageCompletion = monitoringPercentageCompletion;
            PrimaryDiskAzureStorageAccountId = primaryDiskAzureStorageAccountId;
            PrimaryStagingAzureStorageAccountId = primaryStagingAzureStorageAccountId;
            RecoveryAzureStorageAccountId = recoveryAzureStorageAccountId;
            RecoveryDiskUri = recoveryDiskUri;
            ResyncRequired = resyncRequired;
        }
    }
}
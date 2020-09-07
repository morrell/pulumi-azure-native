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
    public sealed class VMwareCbtMigrationDetailsResponseResult
    {
        /// <summary>
        /// The data mover RunAs account Id.
        /// </summary>
        public readonly string DataMoverRunAsAccountId;
        /// <summary>
        /// Gets the instance type.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The last recovery point received time.
        /// </summary>
        public readonly string LastRecoveryPointReceived;
        /// <summary>
        /// License Type of the VM to be used.
        /// </summary>
        public readonly string? LicenseType;
        /// <summary>
        /// The recovery point Id to which the VM was migrated.
        /// </summary>
        public readonly string MigrationRecoveryPointId;
        /// <summary>
        /// The type of the OS on the VM.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// The list of protected disks.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMwareCbtProtectedDiskDetailsResponseResult> ProtectedDisks;
        /// <summary>
        /// The snapshot RunAs account Id.
        /// </summary>
        public readonly string SnapshotRunAsAccountId;
        /// <summary>
        /// The target availability set Id.
        /// </summary>
        public readonly string? TargetAvailabilitySetId;
        /// <summary>
        /// The target boot diagnostics storage account ARM Id.
        /// </summary>
        public readonly string? TargetBootDiagnosticsStorageAccountId;
        /// <summary>
        /// The target location.
        /// </summary>
        public readonly string TargetLocation;
        /// <summary>
        /// The target network Id.
        /// </summary>
        public readonly string? TargetNetworkId;
        /// <summary>
        /// The target resource group Id.
        /// </summary>
        public readonly string? TargetResourceGroupId;
        /// <summary>
        /// Target VM name.
        /// </summary>
        public readonly string? TargetVmName;
        /// <summary>
        /// The target VM size.
        /// </summary>
        public readonly string? TargetVmSize;
        /// <summary>
        /// The network details.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMwareCbtNicDetailsResponseResult> VmNics;
        /// <summary>
        /// The ARM Id of the VM discovered in VMware.
        /// </summary>
        public readonly string VmwareMachineId;

        [OutputConstructor]
        private VMwareCbtMigrationDetailsResponseResult(
            string dataMoverRunAsAccountId,

            string instanceType,

            string lastRecoveryPointReceived,

            string? licenseType,

            string migrationRecoveryPointId,

            string osType,

            ImmutableArray<Outputs.VMwareCbtProtectedDiskDetailsResponseResult> protectedDisks,

            string snapshotRunAsAccountId,

            string? targetAvailabilitySetId,

            string? targetBootDiagnosticsStorageAccountId,

            string targetLocation,

            string? targetNetworkId,

            string? targetResourceGroupId,

            string? targetVmName,

            string? targetVmSize,

            ImmutableArray<Outputs.VMwareCbtNicDetailsResponseResult> vmNics,

            string vmwareMachineId)
        {
            DataMoverRunAsAccountId = dataMoverRunAsAccountId;
            InstanceType = instanceType;
            LastRecoveryPointReceived = lastRecoveryPointReceived;
            LicenseType = licenseType;
            MigrationRecoveryPointId = migrationRecoveryPointId;
            OsType = osType;
            ProtectedDisks = protectedDisks;
            SnapshotRunAsAccountId = snapshotRunAsAccountId;
            TargetAvailabilitySetId = targetAvailabilitySetId;
            TargetBootDiagnosticsStorageAccountId = targetBootDiagnosticsStorageAccountId;
            TargetLocation = targetLocation;
            TargetNetworkId = targetNetworkId;
            TargetResourceGroupId = targetResourceGroupId;
            TargetVmName = targetVmName;
            TargetVmSize = targetVmSize;
            VmNics = vmNics;
            VmwareMachineId = vmwareMachineId;
        }
    }
}
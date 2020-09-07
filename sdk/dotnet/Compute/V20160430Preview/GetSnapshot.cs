// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview
{
    public static class GetSnapshot
    {
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("azurerm:compute/v20160430preview:getSnapshot", args ?? new GetSnapshotArgs(), options.WithVersion());
    }


    public sealed class GetSnapshotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the snapshot within the given subscription and resource group.
        /// </summary>
        [Input("snapshotName", required: true)]
        public string SnapshotName { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        /// <summary>
        /// the storage account type of the disk.
        /// </summary>
        public readonly string? AccountType;
        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        public readonly Outputs.CreationDataResponseResult CreationData;
        /// <summary>
        /// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        /// </summary>
        public readonly int? DiskSizeGB;
        /// <summary>
        /// Encryption settings for disk or snapshot
        /// </summary>
        public readonly Outputs.EncryptionSettingsResponseResult? EncryptionSettings;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Operating System type.
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// A relative URI containing the VM id that has the disk attached.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// The disk provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The time when the disk was created.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSnapshotResult(
            string? accountType,

            Outputs.CreationDataResponseResult creationData,

            int? diskSizeGB,

            Outputs.EncryptionSettingsResponseResult? encryptionSettings,

            string location,

            string name,

            string? osType,

            string ownerId,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string timeCreated,

            string type)
        {
            AccountType = accountType;
            CreationData = creationData;
            DiskSizeGB = diskSizeGB;
            EncryptionSettings = encryptionSettings;
            Location = location;
            Name = name;
            OsType = osType;
            OwnerId = ownerId;
            ProvisioningState = provisioningState;
            Tags = tags;
            TimeCreated = timeCreated;
            Type = type;
        }
    }
}
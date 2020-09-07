// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180401
{
    public static class GetSnapshot
    {
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("azurerm:compute/v20180401:getSnapshot", args ?? new GetSnapshotArgs(), options.WithVersion());
    }


    public sealed class GetSnapshotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
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
        /// Unused. Always Null.
        /// </summary>
        public readonly string ManagedBy;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Operating System type.
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// The disk provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
        /// </summary>
        public readonly Outputs.SnapshotSkuResponseResult? Sku;
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
            Outputs.CreationDataResponseResult creationData,

            int? diskSizeGB,

            Outputs.EncryptionSettingsResponseResult? encryptionSettings,

            string location,

            string managedBy,

            string name,

            string? osType,

            string provisioningState,

            Outputs.SnapshotSkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string timeCreated,

            string type)
        {
            CreationData = creationData;
            DiskSizeGB = diskSizeGB;
            EncryptionSettings = encryptionSettings;
            Location = location;
            ManagedBy = managedBy;
            Name = name;
            OsType = osType;
            ProvisioningState = provisioningState;
            Sku = sku;
            Tags = tags;
            TimeCreated = timeCreated;
            Type = type;
        }
    }
}
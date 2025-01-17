// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp.V20191101
{
    public static class GetSnapshot
    {
        /// <summary>
        /// Snapshot of a Volume
        /// </summary>
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("azure-native:netapp/v20191101:getSnapshot", args ?? new GetSnapshotArgs(), options.WithVersion());
    }


    public sealed class GetSnapshotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the capacity pool
        /// </summary>
        [Input("poolName", required: true)]
        public string PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the mount target
        /// </summary>
        [Input("snapshotName", required: true)]
        public string SnapshotName { get; set; } = null!;

        /// <summary>
        /// The name of the volume
        /// </summary>
        [Input("volumeName", required: true)]
        public string VolumeName { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        /// <summary>
        /// The creation date of the snapshot
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// UUID v4 used to identify the FileSystem
        /// </summary>
        public readonly string? FileSystemId;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure lifecycle management
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// UUID v4 used to identify the Snapshot
        /// </summary>
        public readonly string SnapshotId;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSnapshotResult(
            string created,

            string? fileSystemId,

            string id,

            string location,

            string name,

            string provisioningState,

            string snapshotId,

            string type)
        {
            Created = created;
            FileSystemId = fileSystemId;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SnapshotId = snapshotId;
            Type = type;
        }
    }
}

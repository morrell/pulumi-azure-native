// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NetApp.V20200201.Outputs
{

    [OutputType]
    public sealed class VolumeSnapshotPropertiesResponseResult
    {
        /// <summary>
        /// Snapshot Policy ResourceId
        /// </summary>
        public readonly string? SnapshotPolicyId;

        [OutputConstructor]
        private VolumeSnapshotPropertiesResponseResult(string? snapshotPolicyId)
        {
            SnapshotPolicyId = snapshotPolicyId;
        }
    }
}
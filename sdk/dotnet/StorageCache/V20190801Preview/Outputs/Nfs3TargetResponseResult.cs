// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageCache.V20190801Preview.Outputs
{

    [OutputType]
    public sealed class Nfs3TargetResponseResult
    {
        /// <summary>
        /// IP or name of an NFS Storage Target host, ie: 10.0.44.44
        /// </summary>
        public readonly string? Target;
        /// <summary>
        /// Identifies the primary usage model to be used for this storage target.   GET choices from .../usageModels
        /// </summary>
        public readonly string? UsageModel;

        [OutputConstructor]
        private Nfs3TargetResponseResult(
            string? target,

            string? usageModel)
        {
            Target = target;
            UsageModel = usageModel;
        }
    }
}
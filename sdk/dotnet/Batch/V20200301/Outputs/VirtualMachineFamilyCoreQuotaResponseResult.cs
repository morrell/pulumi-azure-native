// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20200301.Outputs
{

    [OutputType]
    public sealed class VirtualMachineFamilyCoreQuotaResponseResult
    {
        /// <summary>
        /// The core quota for the VM family for the Batch account.
        /// </summary>
        public readonly int CoreQuota;
        /// <summary>
        /// The Virtual Machine family name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private VirtualMachineFamilyCoreQuotaResponseResult(
            int coreQuota,

            string name)
        {
            CoreQuota = coreQuota;
            Name = name;
        }
    }
}
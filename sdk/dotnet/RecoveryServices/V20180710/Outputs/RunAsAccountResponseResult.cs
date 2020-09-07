// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20180710.Outputs
{

    [OutputType]
    public sealed class RunAsAccountResponseResult
    {
        /// <summary>
        /// The CS RunAs account Id.
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// The CS RunAs account name.
        /// </summary>
        public readonly string? AccountName;

        [OutputConstructor]
        private RunAsAccountResponseResult(
            string? accountId,

            string? accountName)
        {
            AccountId = accountId;
            AccountName = accountName;
        }
    }
}
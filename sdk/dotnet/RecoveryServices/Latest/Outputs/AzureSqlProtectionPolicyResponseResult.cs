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
    public sealed class AzureSqlProtectionPolicyResponseResult
    {
        /// <summary>
        /// This property is used as the discriminator for deciding the specific types in the polymorphic chain of types.
        /// </summary>
        public readonly string? BackupManagementType;
        /// <summary>
        /// The number of items associated with this policy.
        /// </summary>
        public readonly int? ProtectedItemsCount;
        /// <summary>
        /// The retention policy details.
        /// </summary>
        public readonly Union<Outputs.LongTermRetentionPolicyResponseResult, Outputs.SimpleRetentionPolicyResponseResult>? RetentionPolicy;

        [OutputConstructor]
        private AzureSqlProtectionPolicyResponseResult(
            string? backupManagementType,

            int? protectedItemsCount,

            Union<Outputs.LongTermRetentionPolicyResponseResult, Outputs.SimpleRetentionPolicyResponseResult>? retentionPolicy)
        {
            BackupManagementType = backupManagementType;
            ProtectedItemsCount = protectedItemsCount;
            RetentionPolicy = retentionPolicy;
        }
    }
}
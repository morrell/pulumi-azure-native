// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20191201Preview.Outputs
{

    [OutputType]
    public sealed class PoliciesResponseResult
    {
        /// <summary>
        /// The quarantine policy for a container registry.
        /// </summary>
        public readonly Outputs.QuarantinePolicyResponseResult? QuarantinePolicy;
        /// <summary>
        /// The retention policy for a container registry.
        /// </summary>
        public readonly Outputs.RetentionPolicyResponseResult? RetentionPolicy;
        /// <summary>
        /// The content trust policy for a container registry.
        /// </summary>
        public readonly Outputs.TrustPolicyResponseResult? TrustPolicy;

        [OutputConstructor]
        private PoliciesResponseResult(
            Outputs.QuarantinePolicyResponseResult? quarantinePolicy,

            Outputs.RetentionPolicyResponseResult? retentionPolicy,

            Outputs.TrustPolicyResponseResult? trustPolicy)
        {
            QuarantinePolicy = quarantinePolicy;
            RetentionPolicy = retentionPolicy;
            TrustPolicy = trustPolicy;
        }
    }
}
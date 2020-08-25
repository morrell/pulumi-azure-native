// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20190501.Outputs
{

    [OutputType]
    public sealed class QuarantinePolicyResponseResult
    {
        /// <summary>
        /// The value that indicates whether the policy is enabled or not.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private QuarantinePolicyResponseResult(string? status)
        {
            Status = status;
        }
    }
}
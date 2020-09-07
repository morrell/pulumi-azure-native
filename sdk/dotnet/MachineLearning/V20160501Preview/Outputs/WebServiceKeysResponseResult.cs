// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearning.V20160501Preview.Outputs
{

    [OutputType]
    public sealed class WebServiceKeysResponseResult
    {
        /// <summary>
        /// The primary access key.
        /// </summary>
        public readonly string? Primary;
        /// <summary>
        /// The secondary access key.
        /// </summary>
        public readonly string? Secondary;

        [OutputConstructor]
        private WebServiceKeysResponseResult(
            string? primary,

            string? secondary)
        {
            Primary = primary;
            Secondary = secondary;
        }
    }
}
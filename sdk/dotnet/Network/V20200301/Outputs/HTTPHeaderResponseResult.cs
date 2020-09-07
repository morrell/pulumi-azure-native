// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200301.Outputs
{

    [OutputType]
    public sealed class HTTPHeaderResponseResult
    {
        /// <summary>
        /// The name in HTTP header.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The value in HTTP header.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private HTTPHeaderResponseResult(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20190801.Outputs
{

    [OutputType]
    public sealed class ResizeErrorResponseResult
    {
        /// <summary>
        /// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
        /// </summary>
        public readonly string Code;
        public readonly ImmutableArray<Outputs.ResizeErrorResponseResult> Details;
        /// <summary>
        /// A message describing the error, intended to be suitable for display in a user interface.
        /// </summary>
        public readonly string Message;

        [OutputConstructor]
        private ResizeErrorResponseResult(
            string code,

            ImmutableArray<Outputs.ResizeErrorResponseResult> details,

            string message)
        {
            Code = code;
            Details = details;
            Message = message;
        }
    }
}
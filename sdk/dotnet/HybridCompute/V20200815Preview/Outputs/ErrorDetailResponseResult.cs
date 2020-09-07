// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridCompute.V20200815Preview.Outputs
{

    [OutputType]
    public sealed class ErrorDetailResponseResult
    {
        /// <summary>
        /// The error's code.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Additional error details.
        /// </summary>
        public readonly ImmutableArray<Outputs.ErrorDetailResponseResult> Details;
        /// <summary>
        /// A human readable error message.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Indicates which property in the request is responsible for the error.
        /// </summary>
        public readonly string? Target;

        [OutputConstructor]
        private ErrorDetailResponseResult(
            string code,

            ImmutableArray<Outputs.ErrorDetailResponseResult> details,

            string message,

            string? target)
        {
            Code = code;
            Details = details;
            Message = message;
            Target = target;
        }
    }
}
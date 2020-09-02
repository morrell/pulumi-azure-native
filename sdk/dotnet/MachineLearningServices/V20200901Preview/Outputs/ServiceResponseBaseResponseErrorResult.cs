// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200901Preview.Outputs
{

    [OutputType]
    public sealed class ServiceResponseBaseResponseErrorResult
    {
        /// <summary>
        /// Error code.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// An array of error detail objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.ErrorDetailResponseResult> Details;
        /// <summary>
        /// Error message.
        /// </summary>
        public readonly string Message;

        [OutputConstructor]
        private ServiceResponseBaseResponseErrorResult(
            string code,

            ImmutableArray<Outputs.ErrorDetailResponseResult> details,

            string message)
        {
            Code = code;
            Details = details;
            Message = message;
        }
    }
}
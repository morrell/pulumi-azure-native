// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20181201.Outputs
{

    [OutputType]
    public sealed class ResizeOperationStatusResponseResult
    {
        /// <summary>
        /// This property is set only if an error occurred during the last pool resize, and only when the pool allocationState is Steady.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResizeErrorResponseResult> Errors;
        /// <summary>
        /// The default value is requeue.
        /// </summary>
        public readonly string? NodeDeallocationOption;
        /// <summary>
        /// The default value is 15 minutes. The minimum value is 5 minutes. If you specify a value less than 5 minutes, the Batch service returns an error; if you are calling the REST API directly, the HTTP status code is 400 (Bad Request).
        /// </summary>
        public readonly string? ResizeTimeout;
        public readonly string? StartTime;
        public readonly int? TargetDedicatedNodes;
        public readonly int? TargetLowPriorityNodes;

        [OutputConstructor]
        private ResizeOperationStatusResponseResult(
            ImmutableArray<Outputs.ResizeErrorResponseResult> errors,

            string? nodeDeallocationOption,

            string? resizeTimeout,

            string? startTime,

            int? targetDedicatedNodes,

            int? targetLowPriorityNodes)
        {
            Errors = errors;
            NodeDeallocationOption = nodeDeallocationOption;
            ResizeTimeout = resizeTimeout;
            StartTime = startTime;
            TargetDedicatedNodes = targetDedicatedNodes;
            TargetLowPriorityNodes = targetLowPriorityNodes;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate.V20191001Preview.Outputs
{

    [OutputType]
    public sealed class MoveResourcePropertiesResponseMoveStatusResult
    {
        /// <summary>
        /// An error response from the azure region move service.
        /// </summary>
        public readonly Outputs.MoveResourceErrorResponseResult? Errors;
        /// <summary>
        /// Defines the job status.
        /// </summary>
        public readonly Outputs.JobStatusResponseResult? JobStatus;
        /// <summary>
        /// Defines the MoveResource states.
        /// </summary>
        public readonly string? MoveState;
        /// <summary>
        /// Gets the Target ARM Id of the resource.
        /// </summary>
        public readonly string TargetId;

        [OutputConstructor]
        private MoveResourcePropertiesResponseMoveStatusResult(
            Outputs.MoveResourceErrorResponseResult? errors,

            Outputs.JobStatusResponseResult? jobStatus,

            string? moveState,

            string targetId)
        {
            Errors = errors;
            JobStatus = jobStatus;
            MoveState = moveState;
            TargetId = targetId;
        }
    }
}
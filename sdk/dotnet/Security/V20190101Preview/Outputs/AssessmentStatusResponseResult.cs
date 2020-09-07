// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20190101Preview.Outputs
{

    [OutputType]
    public sealed class AssessmentStatusResponseResult
    {
        /// <summary>
        /// Programmatic code for the cause of the assessment status
        /// </summary>
        public readonly string? Cause;
        /// <summary>
        /// Programmatic code for the status of the assessment
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Human readable description of the assessment status
        /// </summary>
        public readonly string? Description;

        [OutputConstructor]
        private AssessmentStatusResponseResult(
            string? cause,

            string code,

            string? description)
        {
            Cause = cause;
            Code = code;
            Description = description;
        }
    }
}
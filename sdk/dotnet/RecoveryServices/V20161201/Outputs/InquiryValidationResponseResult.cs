// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20161201.Outputs
{

    [OutputType]
    public sealed class InquiryValidationResponseResult
    {
        /// <summary>
        /// Error Additional Detail in case the status is non-success.
        /// </summary>
        public readonly string AdditionalDetail;
        /// <summary>
        /// Error Detail in case the status is non-success.
        /// </summary>
        public readonly Outputs.ErrorDetailResponseResult? ErrorDetail;
        /// <summary>
        /// Status for the Inquiry Validation.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private InquiryValidationResponseResult(
            string additionalDetail,

            Outputs.ErrorDetailResponseResult? errorDetail,

            string? status)
        {
            AdditionalDetail = additionalDetail;
            ErrorDetail = errorDetail;
            Status = status;
        }
    }
}
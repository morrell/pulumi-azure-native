// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Billing.V20180801Preview.Outputs
{

    [OutputType]
    public sealed class ReportScheduleResponseResult
    {
        /// <summary>
        /// The schedule recurrence.
        /// </summary>
        public readonly string Recurrence;
        /// <summary>
        /// Has start and end date of the recurrence. The start date must be in future. If present, the end date must be greater than start date.
        /// </summary>
        public readonly Outputs.ReportRecurrencePeriodResponseResult? RecurrencePeriod;
        /// <summary>
        /// The status of the schedule. Whether active or not. If inactive, the report's scheduled execution is paused.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ReportScheduleResponseResult(
            string recurrence,

            Outputs.ReportRecurrencePeriodResponseResult? recurrencePeriod,

            string? status)
        {
            Recurrence = recurrence;
            RecurrencePeriod = recurrencePeriod;
            Status = status;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AlertsManagement.V20181102PrivatePreview.Outputs
{

    [OutputType]
    public sealed class SuppressionConfigResponseResult
    {
        /// <summary>
        /// Specifies when the suppression should be applied
        /// </summary>
        public readonly string RecurrenceType;
        /// <summary>
        /// Schedule for a given suppression configuration.
        /// </summary>
        public readonly Outputs.SuppressionScheduleResponseResult? Schedule;

        [OutputConstructor]
        private SuppressionConfigResponseResult(
            string recurrenceType,

            Outputs.SuppressionScheduleResponseResult? schedule)
        {
            RecurrenceType = recurrenceType;
            Schedule = schedule;
        }
    }
}
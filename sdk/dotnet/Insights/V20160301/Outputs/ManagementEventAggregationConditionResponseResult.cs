// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20160301.Outputs
{

    [OutputType]
    public sealed class ManagementEventAggregationConditionResponseResult
    {
        /// <summary>
        /// the condition operator.
        /// </summary>
        public readonly string? Operator;
        /// <summary>
        /// The threshold value that activates the alert.
        /// </summary>
        public readonly double? Threshold;
        /// <summary>
        /// the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the threshold. If specified then it must be between 5 minutes and 1 day.
        /// </summary>
        public readonly string? WindowSize;

        [OutputConstructor]
        private ManagementEventAggregationConditionResponseResult(
            string? @operator,

            double? threshold,

            string? windowSize)
        {
            Operator = @operator;
            Threshold = threshold;
            WindowSize = windowSize;
        }
    }
}
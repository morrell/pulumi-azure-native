// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20180416.Outputs
{

    [OutputType]
    public sealed class LogToMetricActionResponseResult
    {
        /// <summary>
        /// Criteria of Metric
        /// </summary>
        public readonly ImmutableArray<Outputs.CriteriaResponseResult> Criteria;
        /// <summary>
        /// Specifies the action. Supported values - AlertingAction, LogToMetricAction
        /// </summary>
        public readonly string OdataType;

        [OutputConstructor]
        private LogToMetricActionResponseResult(
            ImmutableArray<Outputs.CriteriaResponseResult> criteria,

            string odataType)
        {
            Criteria = criteria;
            OdataType = odataType;
        }
    }
}
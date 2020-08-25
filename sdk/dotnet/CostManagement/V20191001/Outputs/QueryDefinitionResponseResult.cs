// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CostManagement.V20191001.Outputs
{

    [OutputType]
    public sealed class QueryDefinitionResponseResult
    {
        /// <summary>
        /// Has definition for data in this query.
        /// </summary>
        public readonly Outputs.QueryDatasetResponseResult? Dataset;
        /// <summary>
        /// Has time period for pulling data for the query.
        /// </summary>
        public readonly Outputs.QueryTimePeriodResponseResult? TimePeriod;
        /// <summary>
        /// The time frame for pulling data for the query. If custom, then a specific time period must be provided.
        /// </summary>
        public readonly string Timeframe;
        /// <summary>
        /// The type of the query.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private QueryDefinitionResponseResult(
            Outputs.QueryDatasetResponseResult? dataset,

            Outputs.QueryTimePeriodResponseResult? timePeriod,

            string timeframe,

            string type)
        {
            Dataset = dataset;
            TimePeriod = timePeriod;
            Timeframe = timeframe;
            Type = type;
        }
    }
}
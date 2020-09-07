// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.Latest.Outputs
{

    [OutputType]
    public sealed class QueryExecutionResultResponseResult
    {
        /// <summary>
        /// Query text retrieved from the source server
        /// </summary>
        public readonly string QueryText;
        /// <summary>
        /// Query analysis result from the source
        /// </summary>
        public readonly Outputs.ExecutionStatisticsResponseResult SourceResult;
        /// <summary>
        /// Total no. of statements in the batch
        /// </summary>
        public readonly int StatementsInBatch;
        /// <summary>
        /// Query analysis result from the target
        /// </summary>
        public readonly Outputs.ExecutionStatisticsResponseResult TargetResult;

        [OutputConstructor]
        private QueryExecutionResultResponseResult(
            string queryText,

            Outputs.ExecutionStatisticsResponseResult sourceResult,

            int statementsInBatch,

            Outputs.ExecutionStatisticsResponseResult targetResult)
        {
            QueryText = queryText;
            SourceResult = sourceResult;
            StatementsInBatch = statementsInBatch;
            TargetResult = targetResult;
        }
    }
}
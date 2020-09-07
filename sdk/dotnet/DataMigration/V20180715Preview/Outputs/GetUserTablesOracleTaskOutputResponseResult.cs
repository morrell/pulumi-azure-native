// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180715Preview.Outputs
{

    [OutputType]
    public sealed class GetUserTablesOracleTaskOutputResponseResult
    {
        /// <summary>
        /// The schema this result is for
        /// </summary>
        public readonly string SchemaName;
        /// <summary>
        /// List of valid tables found for this schema
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseTableResponseResult> Tables;
        /// <summary>
        /// Validation errors associated with the task
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportableExceptionResponseResult> ValidationErrors;

        [OutputConstructor]
        private GetUserTablesOracleTaskOutputResponseResult(
            string schemaName,

            ImmutableArray<Outputs.DatabaseTableResponseResult> tables,

            ImmutableArray<Outputs.ReportableExceptionResponseResult> validationErrors)
        {
            SchemaName = schemaName;
            Tables = tables;
            ValidationErrors = validationErrors;
        }
    }
}
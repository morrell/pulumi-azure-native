// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180419.Inputs
{

    /// <summary>
    /// Properties for task that validates migration input for SQL to Azure SQL Database Managed Instance sync scenario
    /// </summary>
    public sealed class ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Task input
        /// </summary>
        [Input("input")]
        public Input<Inputs.ValidateMigrationInputSqlServerSqlMISyncTaskInputArgs>? Input { get; set; }

        /// <summary>
        /// Task type.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        public ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesArgs()
        {
        }
    }
}
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
    /// Properties for task that migrates SQL Server databases to Azure SQL Database Managed Instance
    /// </summary>
    public sealed class MigrateSqlServerSqlMITaskPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Task input
        /// </summary>
        [Input("input")]
        public Input<Inputs.MigrateSqlServerSqlMITaskInputArgs>? Input { get; set; }

        /// <summary>
        /// Task type.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        public MigrateSqlServerSqlMITaskPropertiesArgs()
        {
        }
    }
}
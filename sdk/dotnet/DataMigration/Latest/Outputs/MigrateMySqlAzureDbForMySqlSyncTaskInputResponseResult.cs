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
    public sealed class MigrateMySqlAzureDbForMySqlSyncTaskInputResponseResult
    {
        /// <summary>
        /// Databases to migrate
        /// </summary>
        public readonly ImmutableArray<Outputs.MigrateMySqlAzureDbForMySqlSyncDatabaseInputResponseResult> SelectedDatabases;
        /// <summary>
        /// Connection information for source MySQL
        /// </summary>
        public readonly Outputs.MySqlConnectionInfoResponseResult SourceConnectionInfo;
        /// <summary>
        /// Connection information for target Azure Database for MySQL
        /// </summary>
        public readonly Outputs.MySqlConnectionInfoResponseResult TargetConnectionInfo;

        [OutputConstructor]
        private MigrateMySqlAzureDbForMySqlSyncTaskInputResponseResult(
            ImmutableArray<Outputs.MigrateMySqlAzureDbForMySqlSyncDatabaseInputResponseResult> selectedDatabases,

            Outputs.MySqlConnectionInfoResponseResult sourceConnectionInfo,

            Outputs.MySqlConnectionInfoResponseResult targetConnectionInfo)
        {
            SelectedDatabases = selectedDatabases;
            SourceConnectionInfo = sourceConnectionInfo;
            TargetConnectionInfo = targetConnectionInfo;
        }
    }
}
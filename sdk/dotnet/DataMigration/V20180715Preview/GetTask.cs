// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180715Preview
{
    public static class GetTask
    {
        public static Task<GetTaskResult> InvokeAsync(GetTaskArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTaskResult>("azurerm:datamigration/v20180715preview:getTask", args ?? new GetTaskArgs(), options.WithVersion());
    }


    public sealed class GetTaskArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expand the response
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// Name of the resource group
        /// </summary>
        [Input("groupName", required: true)]
        public string GroupName { get; set; } = null!;

        /// <summary>
        /// Name of the project
        /// </summary>
        [Input("projectName", required: true)]
        public string ProjectName { get; set; } = null!;

        /// <summary>
        /// Name of the service
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// Name of the Task
        /// </summary>
        [Input("taskName", required: true)]
        public string TaskName { get; set; } = null!;

        public GetTaskArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTaskResult
    {
        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Custom task properties
        /// </summary>
        public readonly Union<Outputs.ConnectToMongoDbTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceOracleSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourcePostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceSqlServerSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceSqlServerTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetAzureDbForMySqlTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlDbTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlMISyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlMITaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.GetTdeCertificatesSqlTaskPropertiesResponseResult, Union<Outputs.GetUserTablesOracleTaskPropertiesResponseResult, Union<Outputs.GetUserTablesPostgreSqlTaskPropertiesResponseResult, Union<Outputs.GetUserTablesSqlSyncTaskPropertiesResponseResult, Union<Outputs.GetUserTablesSqlTaskPropertiesResponseResult, Union<Outputs.MigrateMongoDbTaskPropertiesResponseResult, Union<Outputs.MigrateMySqlAzureDbForMySqlSyncTaskPropertiesResponseResult, Union<Outputs.MigrateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.MigratePostgreSqlAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlDbTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlMISyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlMITaskPropertiesResponseResult, Union<Outputs.MigrateSsisTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlMITaskPropertiesResponseResult, Union<Outputs.ValidateMongoDbTaskPropertiesResponseResult, Outputs.ValidateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTaskResult(
            string? etag,

            string name,

            Union<Outputs.ConnectToMongoDbTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceOracleSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourcePostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceSqlServerSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceSqlServerTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetAzureDbForMySqlTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlDbTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlMISyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlMITaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.GetTdeCertificatesSqlTaskPropertiesResponseResult, Union<Outputs.GetUserTablesOracleTaskPropertiesResponseResult, Union<Outputs.GetUserTablesPostgreSqlTaskPropertiesResponseResult, Union<Outputs.GetUserTablesSqlSyncTaskPropertiesResponseResult, Union<Outputs.GetUserTablesSqlTaskPropertiesResponseResult, Union<Outputs.MigrateMongoDbTaskPropertiesResponseResult, Union<Outputs.MigrateMySqlAzureDbForMySqlSyncTaskPropertiesResponseResult, Union<Outputs.MigrateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.MigratePostgreSqlAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlDbTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlMISyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlMITaskPropertiesResponseResult, Union<Outputs.MigrateSsisTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlMITaskPropertiesResponseResult, Union<Outputs.ValidateMongoDbTaskPropertiesResponseResult, Outputs.ValidateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> properties,

            string type)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
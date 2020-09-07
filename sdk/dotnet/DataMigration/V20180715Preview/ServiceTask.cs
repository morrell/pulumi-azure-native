// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180715Preview
{
    /// <summary>
    /// A task resource
    /// </summary>
    public partial class ServiceTask : Pulumi.CustomResource
    {
        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Custom task properties
        /// </summary>
        [Output("properties")]
        public Output<Union<Outputs.ConnectToMongoDbTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceOracleSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourcePostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceSqlServerSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToSourceSqlServerTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetAzureDbForMySqlTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlDbTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlMISyncTaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlMITaskPropertiesResponseResult, Union<Outputs.ConnectToTargetSqlSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.GetTdeCertificatesSqlTaskPropertiesResponseResult, Union<Outputs.GetUserTablesOracleTaskPropertiesResponseResult, Union<Outputs.GetUserTablesPostgreSqlTaskPropertiesResponseResult, Union<Outputs.GetUserTablesSqlSyncTaskPropertiesResponseResult, Union<Outputs.GetUserTablesSqlTaskPropertiesResponseResult, Union<Outputs.MigrateMongoDbTaskPropertiesResponseResult, Union<Outputs.MigrateMySqlAzureDbForMySqlSyncTaskPropertiesResponseResult, Union<Outputs.MigrateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.MigratePostgreSqlAzureDbForPostgreSqlSyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlDbTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlMISyncTaskPropertiesResponseResult, Union<Outputs.MigrateSqlServerSqlMITaskPropertiesResponseResult, Union<Outputs.MigrateSsisTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponseResult, Union<Outputs.ValidateMigrationInputSqlServerSqlMITaskPropertiesResponseResult, Union<Outputs.ValidateMongoDbTaskPropertiesResponseResult, Outputs.ValidateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponseResult>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceTask(string name, ServiceTaskArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datamigration/v20180715preview:ServiceTask", name, args ?? new ServiceTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceTask(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datamigration/v20180715preview:ServiceTask", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceTask Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceTask(name, id, options);
        }
    }

    public sealed class ServiceTaskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Name of the resource group
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Custom task properties
        /// </summary>
        [Input("properties")]
        public InputUnion<Inputs.ConnectToMongoDbTaskPropertiesArgs, InputUnion<Inputs.ConnectToSourceOracleSyncTaskPropertiesArgs, InputUnion<Inputs.ConnectToSourcePostgreSqlSyncTaskPropertiesArgs, InputUnion<Inputs.ConnectToSourceSqlServerSyncTaskPropertiesArgs, InputUnion<Inputs.ConnectToSourceSqlServerTaskPropertiesArgs, InputUnion<Inputs.ConnectToTargetAzureDbForMySqlTaskPropertiesArgs, InputUnion<Inputs.ConnectToTargetAzureDbForPostgreSqlSyncTaskPropertiesArgs, InputUnion<Inputs.ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskPropertiesArgs, InputUnion<Inputs.ConnectToTargetSqlDbTaskPropertiesArgs, InputUnion<Inputs.ConnectToTargetSqlMISyncTaskPropertiesArgs, InputUnion<Inputs.ConnectToTargetSqlMITaskPropertiesArgs, InputUnion<Inputs.ConnectToTargetSqlSqlDbSyncTaskPropertiesArgs, InputUnion<Inputs.GetTdeCertificatesSqlTaskPropertiesArgs, InputUnion<Inputs.GetUserTablesOracleTaskPropertiesArgs, InputUnion<Inputs.GetUserTablesPostgreSqlTaskPropertiesArgs, InputUnion<Inputs.GetUserTablesSqlSyncTaskPropertiesArgs, InputUnion<Inputs.GetUserTablesSqlTaskPropertiesArgs, InputUnion<Inputs.MigrateMongoDbTaskPropertiesArgs, InputUnion<Inputs.MigrateMySqlAzureDbForMySqlSyncTaskPropertiesArgs, InputUnion<Inputs.MigrateOracleAzureDbForPostgreSqlSyncTaskPropertiesArgs, InputUnion<Inputs.MigratePostgreSqlAzureDbForPostgreSqlSyncTaskPropertiesArgs, InputUnion<Inputs.MigrateSqlServerSqlDbSyncTaskPropertiesArgs, InputUnion<Inputs.MigrateSqlServerSqlDbTaskPropertiesArgs, InputUnion<Inputs.MigrateSqlServerSqlMISyncTaskPropertiesArgs, InputUnion<Inputs.MigrateSqlServerSqlMITaskPropertiesArgs, InputUnion<Inputs.MigrateSsisTaskPropertiesArgs, InputUnion<Inputs.ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesArgs, InputUnion<Inputs.ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesArgs, InputUnion<Inputs.ValidateMigrationInputSqlServerSqlMITaskPropertiesArgs, InputUnion<Inputs.ValidateMongoDbTaskPropertiesArgs, Inputs.ValidateOracleAzureDbForPostgreSqlSyncTaskPropertiesArgs>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>? Properties { get; set; }

        /// <summary>
        /// Name of the service
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Name of the Task
        /// </summary>
        [Input("taskName", required: true)]
        public Input<string> TaskName { get; set; } = null!;

        public ServiceTaskArgs()
        {
        }
    }
}
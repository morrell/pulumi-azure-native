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
    public sealed class MigrateSyncCompleteCommandInputResponseResult
    {
        /// <summary>
        /// Time stamp to complete
        /// </summary>
        public readonly string? CommitTimeStamp;
        /// <summary>
        /// Name of database
        /// </summary>
        public readonly string DatabaseName;

        [OutputConstructor]
        private MigrateSyncCompleteCommandInputResponseResult(
            string? commitTimeStamp,

            string databaseName)
        {
            CommitTimeStamp = commitTimeStamp;
            DatabaseName = databaseName;
        }
    }
}
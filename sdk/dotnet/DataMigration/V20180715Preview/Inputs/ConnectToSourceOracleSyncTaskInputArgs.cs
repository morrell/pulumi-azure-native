// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180715Preview.Inputs
{

    /// <summary>
    /// Input for the task that validates Oracle database connection
    /// </summary>
    public sealed class ConnectToSourceOracleSyncTaskInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information for connecting to Oracle source
        /// </summary>
        [Input("sourceConnectionInfo", required: true)]
        public Input<Inputs.OracleConnectionInfoArgs> SourceConnectionInfo { get; set; } = null!;

        public ConnectToSourceOracleSyncTaskInputArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180331Preview.Outputs
{

    [OutputType]
    public sealed class ConnectToTargetSqlMITaskOutputResponseResult
    {
        /// <summary>
        /// List of agent jobs on the target server.
        /// </summary>
        public readonly ImmutableArray<string> AgentJobs;
        /// <summary>
        /// Result identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of logins on the target server.
        /// </summary>
        public readonly ImmutableArray<string> Logins;
        /// <summary>
        /// Target server brand version
        /// </summary>
        public readonly string TargetServerBrandVersion;
        /// <summary>
        /// Target server version
        /// </summary>
        public readonly string TargetServerVersion;
        /// <summary>
        /// Validation errors
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportableExceptionResponseResult> ValidationErrors;

        [OutputConstructor]
        private ConnectToTargetSqlMITaskOutputResponseResult(
            ImmutableArray<string> agentJobs,

            string id,

            ImmutableArray<string> logins,

            string targetServerBrandVersion,

            string targetServerVersion,

            ImmutableArray<Outputs.ReportableExceptionResponseResult> validationErrors)
        {
            AgentJobs = agentJobs;
            Id = id;
            Logins = logins;
            TargetServerBrandVersion = targetServerBrandVersion;
            TargetServerVersion = targetServerVersion;
            ValidationErrors = validationErrors;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SqlVirtualMachine.V20170301Preview.Outputs
{

    [OutputType]
    public sealed class SqlWorkloadTypeUpdateSettingsResponseResult
    {
        /// <summary>
        /// SQL Server workload type.
        /// </summary>
        public readonly string? SqlWorkloadType;

        [OutputConstructor]
        private SqlWorkloadTypeUpdateSettingsResponseResult(string? sqlWorkloadType)
        {
            SqlWorkloadType = sqlWorkloadType;
        }
    }
}
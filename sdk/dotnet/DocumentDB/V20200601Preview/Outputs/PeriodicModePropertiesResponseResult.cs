// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview.Outputs
{

    [OutputType]
    public sealed class PeriodicModePropertiesResponseResult
    {
        /// <summary>
        /// An integer representing the interval in minutes between two backups
        /// </summary>
        public readonly int? BackupIntervalInMinutes;
        /// <summary>
        /// An integer representing the time (in hours) that each backup is retained
        /// </summary>
        public readonly int? BackupRetentionIntervalInHours;

        [OutputConstructor]
        private PeriodicModePropertiesResponseResult(
            int? backupIntervalInMinutes,

            int? backupRetentionIntervalInHours)
        {
            BackupIntervalInMinutes = backupIntervalInMinutes;
            BackupRetentionIntervalInHours = backupRetentionIntervalInHours;
        }
    }
}
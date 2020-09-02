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
    public sealed class AutoPatchingSettingsResponseResult
    {
        /// <summary>
        /// Day of week to apply the patch on.
        /// </summary>
        public readonly string? DayOfWeek;
        /// <summary>
        /// Enable or disable autopatching on SQL virtual machine.
        /// </summary>
        public readonly bool? Enable;
        /// <summary>
        /// Duration of patching.
        /// </summary>
        public readonly int? MaintenanceWindowDuration;
        /// <summary>
        /// Hour of the day when patching is initiated. Local VM time.
        /// </summary>
        public readonly int? MaintenanceWindowStartingHour;

        [OutputConstructor]
        private AutoPatchingSettingsResponseResult(
            string? dayOfWeek,

            bool? enable,

            int? maintenanceWindowDuration,

            int? maintenanceWindowStartingHour)
        {
            DayOfWeek = dayOfWeek;
            Enable = enable;
            MaintenanceWindowDuration = maintenanceWindowDuration;
            MaintenanceWindowStartingHour = maintenanceWindowStartingHour;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class DataCollectionRuleResponseDataSourcesResult
    {
        /// <summary>
        /// The list of Azure VM extension data source configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExtensionDataSourceResponseResult> Extensions;
        /// <summary>
        /// The list of performance counter data source configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.PerfCounterDataSourceResponseResult> PerformanceCounters;
        /// <summary>
        /// The list of Syslog data source configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.SyslogDataSourceResponseResult> Syslog;
        /// <summary>
        /// The list of Windows Event Log data source configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.WindowsEventLogDataSourceResponseResult> WindowsEventLogs;

        [OutputConstructor]
        private DataCollectionRuleResponseDataSourcesResult(
            ImmutableArray<Outputs.ExtensionDataSourceResponseResult> extensions,

            ImmutableArray<Outputs.PerfCounterDataSourceResponseResult> performanceCounters,

            ImmutableArray<Outputs.SyslogDataSourceResponseResult> syslog,

            ImmutableArray<Outputs.WindowsEventLogDataSourceResponseResult> windowsEventLogs)
        {
            Extensions = extensions;
            PerformanceCounters = performanceCounters;
            Syslog = syslog;
            WindowsEventLogs = windowsEventLogs;
        }
    }
}
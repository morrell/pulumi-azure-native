// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate.V20180901Preview.Outputs
{

    [OutputType]
    public sealed class ProjectSummaryResponseResult
    {
        /// <summary>
        /// Gets or sets the extended summary.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ExtendedSummary;
        /// <summary>
        /// Gets the Instance type.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Gets or sets the time when summary was last refreshed.
        /// </summary>
        public readonly string? LastSummaryRefreshedTime;
        /// <summary>
        /// Gets or sets the state of refresh summary.
        /// </summary>
        public readonly string? RefreshSummaryState;

        [OutputConstructor]
        private ProjectSummaryResponseResult(
            ImmutableDictionary<string, string>? extendedSummary,

            string instanceType,

            string? lastSummaryRefreshedTime,

            string? refreshSummaryState)
        {
            ExtendedSummary = extendedSummary;
            InstanceType = instanceType;
            LastSummaryRefreshedTime = lastSummaryRefreshedTime;
            RefreshSummaryState = refreshSummaryState;
        }
    }
}
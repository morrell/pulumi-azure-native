// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.Latest.Outputs
{

    [OutputType]
    public sealed class ConnectorMappingAvailabilityResponseResult
    {
        /// <summary>
        /// The frequency to update.
        /// </summary>
        public readonly string? Frequency;
        /// <summary>
        /// The interval of the given frequency to use.
        /// </summary>
        public readonly int Interval;

        [OutputConstructor]
        private ConnectorMappingAvailabilityResponseResult(
            string? frequency,

            int interval)
        {
            Frequency = frequency;
            Interval = interval;
        }
    }
}
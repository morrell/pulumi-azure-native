// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180301.Inputs
{

    /// <summary>
    /// Min and max value of a status code range.
    /// </summary>
    public sealed class MonitorConfigExpectedStatusCodeRangesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Max status code.
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// Min status code.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        public MonitorConfigExpectedStatusCodeRangesArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.Latest.Inputs
{

    /// <summary>
    /// Self referenced tumbling window trigger dependency.
    /// </summary>
    public sealed class SelfDependencyTumblingWindowTriggerReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Timespan applied to the start time of a tumbling window when evaluating dependency.
        /// </summary>
        [Input("offset", required: true)]
        public Input<string> Offset { get; set; } = null!;

        /// <summary>
        /// The size of the window when evaluating the dependency. If undefined the frequency of the tumbling window will be used.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// The type of dependency reference.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SelfDependencyTumblingWindowTriggerReferenceArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20180301Preview.Inputs
{

    /// <summary>
    /// scale settings for BatchAI Compute
    /// </summary>
    public sealed class ScaleSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or disable auto scale
        /// </summary>
        [Input("autoScaleEnabled")]
        public Input<bool>? AutoScaleEnabled { get; set; }

        /// <summary>
        /// Max number of nodes to use
        /// </summary>
        [Input("maxNodeCount")]
        public Input<int>? MaxNodeCount { get; set; }

        /// <summary>
        /// Min number of nodes to use
        /// </summary>
        [Input("minNodeCount")]
        public Input<int>? MinNodeCount { get; set; }

        public ScaleSettingsArgs()
        {
        }
    }
}
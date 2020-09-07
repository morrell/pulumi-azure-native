// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearning.V20160501Preview.Inputs
{

    /// <summary>
    /// Holds the available configuration options for an Azure ML web service endpoint.
    /// </summary>
    public sealed class RealtimeConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the maximum concurrent calls that can be made to the web service. Minimum value: 4, Maximum value: 200.
        /// </summary>
        [Input("maxConcurrentCalls")]
        public Input<int>? MaxConcurrentCalls { get; set; }

        public RealtimeConfigurationArgs()
        {
        }
    }
}
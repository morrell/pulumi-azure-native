// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20180601Preview.Inputs
{

    /// <summary>
    /// Data source object contains configuration to collect telemetry and one or more sinks to send that telemetry data to
    /// </summary>
    public sealed class DataSourceArgs : Pulumi.ResourceArgs
    {
        [Input("configuration", required: true)]
        public Input<Inputs.DataSourceConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// Datasource kind
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        [Input("sinks", required: true)]
        private InputList<Inputs.SinkConfigurationArgs>? _sinks;
        public InputList<Inputs.SinkConfigurationArgs> Sinks
        {
            get => _sinks ?? (_sinks = new InputList<Inputs.SinkConfigurationArgs>());
            set => _sinks = value;
        }

        public DataSourceArgs()
        {
        }
    }
}
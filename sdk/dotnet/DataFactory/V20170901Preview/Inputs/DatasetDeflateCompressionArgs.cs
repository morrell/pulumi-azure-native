// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20170901Preview.Inputs
{

    /// <summary>
    /// The Deflate compression method used on a dataset.
    /// </summary>
    public sealed class DatasetDeflateCompressionArgs : Pulumi.ResourceArgs
    {
        [Input("level")]
        private InputMap<object>? _level;

        /// <summary>
        /// The Deflate compression level.
        /// </summary>
        public InputMap<object> Level
        {
            get => _level ?? (_level = new InputMap<object>());
            set => _level = value;
        }

        /// <summary>
        /// Type of dataset compression.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatasetDeflateCompressionArgs()
        {
        }
    }
}
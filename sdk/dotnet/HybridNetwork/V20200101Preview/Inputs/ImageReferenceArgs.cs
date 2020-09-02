// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridNetwork.V20200101Preview.Inputs
{

    /// <summary>
    /// The image reference properties.
    /// </summary>
    public sealed class ImageReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OS type.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// The VHD SAS URI.
        /// </summary>
        [Input("sasUri")]
        public Input<string>? SasUri { get; set; }

        /// <summary>
        /// The VHD name.
        /// </summary>
        [Input("vhdName")]
        public Input<string>? VhdName { get; set; }

        /// <summary>
        /// The VHD type.
        /// </summary>
        [Input("vhdType")]
        public Input<string>? VhdType { get; set; }

        public ImageReferenceArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190301.Inputs
{

    /// <summary>
    /// The source image from which the Image Version is going to be created.
    /// </summary>
    public sealed class UserArtifactSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The fileName of the artifact.
        /// </summary>
        [Input("fileName", required: true)]
        public Input<string> FileName { get; set; } = null!;

        /// <summary>
        /// Required. The mediaLink of the artifact, must be a readable storage blob.
        /// </summary>
        [Input("mediaLink", required: true)]
        public Input<string> MediaLink { get; set; } = null!;

        public UserArtifactSourceArgs()
        {
        }
    }
}
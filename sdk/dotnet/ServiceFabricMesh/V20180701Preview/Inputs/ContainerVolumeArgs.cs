// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180701Preview.Inputs
{

    /// <summary>
    /// Describes how a volume is attached to a container.
    /// </summary>
    public sealed class ContainerVolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path within the container at which the volume should be mounted. Only valid path characters are allowed.
        /// </summary>
        [Input("destinationPath", required: true)]
        public Input<string> DestinationPath { get; set; } = null!;

        /// <summary>
        /// Name of the volume.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The flag indicating whether the volume is read only. Default is 'false'.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        public ContainerVolumeArgs()
        {
        }
    }
}
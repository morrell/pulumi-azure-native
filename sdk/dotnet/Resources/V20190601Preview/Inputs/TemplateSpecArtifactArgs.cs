// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Resources.V20190601Preview.Inputs
{

    /// <summary>
    /// Represents a Template Spec artifact.
    /// </summary>
    public sealed class TemplateSpecArtifactArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The kind of artifact.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// A filesystem safe relative path of the artifact.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public TemplateSpecArtifactArgs()
        {
        }
    }
}
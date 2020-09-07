// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.V20180630.Inputs
{

    /// <summary>
    /// Definition of the runbook property type.
    /// </summary>
    public sealed class ContentHashArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the content hash algorithm used to hash the content.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// Gets or sets expected hash value of the content.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ContentHashArgs()
        {
        }
    }
}
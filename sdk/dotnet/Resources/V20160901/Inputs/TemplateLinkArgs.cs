// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Resources.V20160901.Inputs
{

    /// <summary>
    /// Entity representing the reference to the template.
    /// </summary>
    public sealed class TemplateLinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If included, must match the ContentVersion in the template.
        /// </summary>
        [Input("contentVersion")]
        public Input<string>? ContentVersion { get; set; }

        /// <summary>
        /// The URI of the template to deploy.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public TemplateLinkArgs()
        {
        }
    }
}
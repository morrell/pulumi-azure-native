// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20190901Preview.Inputs
{

    /// <summary>
    /// Media source
    /// </summary>
    public sealed class MediaGraphSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Source name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The discriminator for derived types.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        public MediaGraphSourceArgs()
        {
        }
    }
}
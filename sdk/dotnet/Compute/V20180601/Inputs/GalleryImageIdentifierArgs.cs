// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180601.Inputs
{

    /// <summary>
    /// This is the gallery Image Definition identifier.
    /// </summary>
    public sealed class GalleryImageIdentifierArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the gallery Image Definition offer.
        /// </summary>
        [Input("offer", required: true)]
        public Input<string> Offer { get; set; } = null!;

        /// <summary>
        /// The name of the gallery Image Definition publisher.
        /// </summary>
        [Input("publisher", required: true)]
        public Input<string> Publisher { get; set; } = null!;

        /// <summary>
        /// The name of the gallery Image Definition SKU.
        /// </summary>
        [Input("sku", required: true)]
        public Input<string> Sku { get; set; } = null!;

        public GalleryImageIdentifierArgs()
        {
        }
    }
}
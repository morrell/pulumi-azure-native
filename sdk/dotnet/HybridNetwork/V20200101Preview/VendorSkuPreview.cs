// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridNetwork.V20200101Preview
{
    /// <summary>
    /// Customer subscription which can use a sku.
    /// </summary>
    public partial class VendorSkuPreview : Pulumi.CustomResource
    {
        /// <summary>
        /// Preview subscription id
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of the resource
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a VendorSkuPreview resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VendorSkuPreview(string name, VendorSkuPreviewArgs args, CustomResourceOptions? options = null)
            : base("azurerm:hybridnetwork/v20200101preview:VendorSkuPreview", name, args ?? new VendorSkuPreviewArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VendorSkuPreview(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:hybridnetwork/v20200101preview:VendorSkuPreview", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VendorSkuPreview resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VendorSkuPreview Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VendorSkuPreview(name, id, options);
        }
    }

    public sealed class VendorSkuPreviewArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Preview subscription id.
        /// </summary>
        [Input("previewSubscription", required: true)]
        public Input<string> PreviewSubscription { get; set; } = null!;

        /// <summary>
        /// The name of the vendor sku.
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        /// <summary>
        /// The name of the vendor.
        /// </summary>
        [Input("vendorName", required: true)]
        public Input<string> VendorName { get; set; } = null!;

        public VendorSkuPreviewArgs()
        {
        }
    }
}
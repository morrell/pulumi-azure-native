// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge.Latest
{
    /// <summary>
    /// The order details.
    /// </summary>
    public partial class Order : Pulumi.CustomResource
    {
        /// <summary>
        /// The contact details.
        /// </summary>
        [Output("contactInformation")]
        public Output<Outputs.ContactDetailsResponseResult> ContactInformation { get; private set; } = null!;

        /// <summary>
        /// Current status of the order.
        /// </summary>
        [Output("currentStatus")]
        public Output<Outputs.OrderStatusResponseResult?> CurrentStatus { get; private set; } = null!;

        /// <summary>
        /// Tracking information for the package delivered to the customer whether it has an original or a replacement device.
        /// </summary>
        [Output("deliveryTrackingInfo")]
        public Output<ImmutableArray<Outputs.TrackingInfoResponseResult>> DeliveryTrackingInfo { get; private set; } = null!;

        /// <summary>
        /// The object name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of status changes in the order.
        /// </summary>
        [Output("orderHistory")]
        public Output<ImmutableArray<Outputs.OrderStatusResponseResult>> OrderHistory { get; private set; } = null!;

        /// <summary>
        /// Tracking information for the package returned from the customer whether it has an original or a replacement device.
        /// </summary>
        [Output("returnTrackingInfo")]
        public Output<ImmutableArray<Outputs.TrackingInfoResponseResult>> ReturnTrackingInfo { get; private set; } = null!;

        /// <summary>
        /// Serial number of the device.
        /// </summary>
        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// The shipping address.
        /// </summary>
        [Output("shippingAddress")]
        public Output<Outputs.AddressResponseResult> ShippingAddress { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Order resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Order(string name, OrderArgs args, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge/latest:Order", name, args ?? new OrderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Order(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge/latest:Order", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:databoxedge/v20190301:Order"},
                    new Pulumi.Alias { Type = "azurerm:databoxedge/v20190701:Order"},
                    new Pulumi.Alias { Type = "azurerm:databoxedge/v20190801:Order"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Order resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Order Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Order(name, id, options);
        }
    }

    public sealed class OrderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The contact details.
        /// </summary>
        [Input("contactInformation", required: true)]
        public Input<Inputs.ContactDetailsArgs> ContactInformation { get; set; } = null!;

        /// <summary>
        /// Current status of the order.
        /// </summary>
        [Input("currentStatus")]
        public Input<Inputs.OrderStatusArgs>? CurrentStatus { get; set; }

        /// <summary>
        /// The order details of a device.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The shipping address.
        /// </summary>
        [Input("shippingAddress", required: true)]
        public Input<Inputs.AddressArgs> ShippingAddress { get; set; } = null!;

        public OrderArgs()
        {
        }
    }
}
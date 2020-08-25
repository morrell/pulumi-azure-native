// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200601
{
    /// <summary>
    /// Route resource.
    /// </summary>
    public partial class Route : Pulumi.CustomResource
    {
        /// <summary>
        /// The destination CIDR to which the route applies.
        /// </summary>
        [Output("addressPrefix")]
        public Output<string?> AddressPrefix { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        /// </summary>
        [Output("nextHopIpAddress")]
        public Output<string?> NextHopIpAddress { get; private set; } = null!;

        /// <summary>
        /// The type of Azure hop the packet should be sent to.
        /// </summary>
        [Output("nextHopType")]
        public Output<string> NextHopType { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the route resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200601:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200601:Route", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Alias { Type = "azurerm:network/v20150615:Route"},
                    new Alias { Type = "azurerm:network/v20160330:Route"},
                    new Alias { Type = "azurerm:network/v20160601:Route"},
                    new Alias { Type = "azurerm:network/v20160901:Route"},
                    new Alias { Type = "azurerm:network/v20161201:Route"},
                    new Alias { Type = "azurerm:network/v20170301:Route"},
                    new Alias { Type = "azurerm:network/v20170601:Route"},
                    new Alias { Type = "azurerm:network/v20170801:Route"},
                    new Alias { Type = "azurerm:network/v20170901:Route"},
                    new Alias { Type = "azurerm:network/v20171001:Route"},
                    new Alias { Type = "azurerm:network/v20171101:Route"},
                    new Alias { Type = "azurerm:network/v20180101:Route"},
                    new Alias { Type = "azurerm:network/v20180201:Route"},
                    new Alias { Type = "azurerm:network/v20180401:Route"},
                    new Alias { Type = "azurerm:network/v20180601:Route"},
                    new Alias { Type = "azurerm:network/v20180701:Route"},
                    new Alias { Type = "azurerm:network/v20180801:Route"},
                    new Alias { Type = "azurerm:network/v20181001:Route"},
                    new Alias { Type = "azurerm:network/v20181101:Route"},
                    new Alias { Type = "azurerm:network/v20181201:Route"},
                    new Alias { Type = "azurerm:network/v20190201:Route"},
                    new Alias { Type = "azurerm:network/v20190401:Route"},
                    new Alias { Type = "azurerm:network/v20190601:Route"},
                    new Alias { Type = "azurerm:network/v20190701:Route"},
                    new Alias { Type = "azurerm:network/v20190801:Route"},
                    new Alias { Type = "azurerm:network/v20190901:Route"},
                    new Alias { Type = "azurerm:network/v20191101:Route"},
                    new Alias { Type = "azurerm:network/v20191201:Route"},
                    new Alias { Type = "azurerm:network/v20200301:Route"},
                    new Alias { Type = "azurerm:network/v20200401:Route"},
                    new Alias { Type = "azurerm:network/v20200501:Route"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Route(name, id, options);
        }
    }

    public sealed class RouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR to which the route applies.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the route.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        /// </summary>
        [Input("nextHopIpAddress")]
        public Input<string>? NextHopIpAddress { get; set; }

        /// <summary>
        /// The type of Azure hop the packet should be sent to.
        /// </summary>
        [Input("nextHopType", required: true)]
        public Input<string> NextHopType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the route table.
        /// </summary>
        [Input("routeTableName", required: true)]
        public Input<string> RouteTableName { get; set; } = null!;

        public RouteArgs()
        {
        }
    }
}
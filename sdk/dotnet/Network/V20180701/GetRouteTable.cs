// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180701
{
    public static class GetRouteTable
    {
        public static Task<GetRouteTableResult> InvokeAsync(GetRouteTableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteTableResult>("azurerm:network/v20180701:getRouteTable", args ?? new GetRouteTableArgs(), options.WithVersion());
    }


    public sealed class GetRouteTableArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the route table.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRouteTableArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteTableResult
    {
        /// <summary>
        /// Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
        /// </summary>
        public readonly bool? DisableBgpRoutePropagation;
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Collection of routes contained within a route table.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouteResponseResult> Routes;
        /// <summary>
        /// A collection of references to subnets.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetResponseResult> Subnets;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRouteTableResult(
            bool? disableBgpRoutePropagation,

            string? etag,

            string? location,

            string name,

            string? provisioningState,

            ImmutableArray<Outputs.RouteResponseResult> routes,

            ImmutableArray<Outputs.SubnetResponseResult> subnets,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            DisableBgpRoutePropagation = disableBgpRoutePropagation;
            Etag = etag;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Routes = routes;
            Subnets = subnets;
            Tags = tags;
            Type = type;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventHub.V20180101Preview
{
    public static class GetCluster
    {
        /// <summary>
        /// Single Event Hubs Cluster resource in List or Get operations.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azure-native:eventhub/v20180101preview:getCluster", args ?? new GetClusterArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Event Hubs Cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The UTC time when the Event Hubs Cluster was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
        /// </summary>
        public readonly string MetricId;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the cluster SKU.
        /// </summary>
        public readonly Outputs.ClusterSkuResponse? Sku;
        /// <summary>
        /// Status of the Cluster resource
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The UTC time when the Event Hubs Cluster was last updated.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetClusterResult(
            string createdAt,

            string id,

            string? location,

            string metricId,

            string name,

            Outputs.ClusterSkuResponse? sku,

            string status,

            ImmutableDictionary<string, string>? tags,

            string type,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Id = id;
            Location = location;
            MetricId = metricId;
            Name = name;
            Sku = sku;
            Status = status;
            Tags = tags;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}

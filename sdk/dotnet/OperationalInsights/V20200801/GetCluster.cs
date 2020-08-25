// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.V20200801
{
    public static class GetCluster
    {
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azurerm:operationalinsights/v20200801:getCluster", args ?? new GetClusterArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Log Analytics Cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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
        /// The ID associated with the cluster.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The identity of the resource.
        /// </summary>
        public readonly Outputs.IdentityResponseResult? Identity;
        /// <summary>
        /// The associated key properties.
        /// </summary>
        public readonly Outputs.KeyVaultPropertiesResponseResult? KeyVaultProperties;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The link used to get the next page of recommendations.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// The provisioning state of the cluster.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The sku properties.
        /// </summary>
        public readonly Outputs.ClusterSkuResponseResult? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetClusterResult(
            string clusterId,

            Outputs.IdentityResponseResult? identity,

            Outputs.KeyVaultPropertiesResponseResult? keyVaultProperties,

            string location,

            string name,

            string? nextLink,

            string provisioningState,

            Outputs.ClusterSkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            ClusterId = clusterId;
            Identity = identity;
            KeyVaultProperties = keyVaultProperties;
            Location = location;
            Name = name;
            NextLink = nextLink;
            ProvisioningState = provisioningState;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
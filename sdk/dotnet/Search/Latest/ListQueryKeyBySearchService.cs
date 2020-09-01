// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search.Latest
{
    public static class ListQueryKeyBySearchService
    {
        public static Task<ListQueryKeyBySearchServiceResult> InvokeAsync(ListQueryKeyBySearchServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListQueryKeyBySearchServiceResult>("azurerm:search/latest:listQueryKeyBySearchService", args ?? new ListQueryKeyBySearchServiceArgs(), options.WithVersion());
    }


    public sealed class ListQueryKeyBySearchServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure Cognitive Search service associated with the specified resource group.
        /// </summary>
        [Input("searchServiceName", required: true)]
        public string SearchServiceName { get; set; } = null!;

        public ListQueryKeyBySearchServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListQueryKeyBySearchServiceResult
    {
        /// <summary>
        /// Request URL that can be used to query next page of query keys. Returned when the total number of requested query keys exceed maximum page size.
        /// </summary>
        public readonly string NextLink;
        /// <summary>
        /// The query keys for the Azure Cognitive Search service.
        /// </summary>
        public readonly ImmutableArray<Outputs.QueryKeyResponseResult> Value;

        [OutputConstructor]
        private ListQueryKeyBySearchServiceResult(
            string nextLink,

            ImmutableArray<Outputs.QueryKeyResponseResult> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
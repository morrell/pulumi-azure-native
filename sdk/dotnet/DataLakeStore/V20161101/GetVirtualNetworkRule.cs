// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataLakeStore.V20161101
{
    public static class GetVirtualNetworkRule
    {
        public static Task<GetVirtualNetworkRuleResult> InvokeAsync(GetVirtualNetworkRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkRuleResult>("azurerm:datalakestore/v20161101:getVirtualNetworkRule", args ?? new GetVirtualNetworkRuleArgs(), options.WithVersion());
    }


    public sealed class GetVirtualNetworkRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Data Lake Store account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network rule to retrieve.
        /// </summary>
        [Input("virtualNetworkRuleName", required: true)]
        public string VirtualNetworkRuleName { get; set; } = null!;

        public GetVirtualNetworkRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualNetworkRuleResult
    {
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource identifier for the subnet.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVirtualNetworkRuleResult(
            string name,

            string subnetId,

            string type)
        {
            Name = name;
            SubnetId = subnetId;
            Type = type;
        }
    }
}
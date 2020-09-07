// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview
{
    public static class ListAssetContainerSas
    {
        public static Task<ListAssetContainerSasResult> InvokeAsync(ListAssetContainerSasArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListAssetContainerSasResult>("azurerm:media/v20180330preview:listAssetContainerSas", args ?? new ListAssetContainerSasArgs(), options.WithVersion());
    }


    public sealed class ListAssetContainerSasArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The Asset name.
        /// </summary>
        [Input("assetName", required: true)]
        public string AssetName { get; set; } = null!;

        /// <summary>
        /// The SAS URL expiration time.  This must be less than 24 hours from the current time.
        /// </summary>
        [Input("expiryTime")]
        public string? ExpiryTime { get; set; }

        /// <summary>
        /// The permissions to set on the SAS URL.
        /// </summary>
        [Input("permissions")]
        public string? Permissions { get; set; }

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListAssetContainerSasArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListAssetContainerSasResult
    {
        /// <summary>
        /// The list of Asset container SAS URLs.
        /// </summary>
        public readonly ImmutableArray<string> AssetContainerSasUrls;

        [OutputConstructor]
        private ListAssetContainerSasResult(ImmutableArray<string> assetContainerSasUrls)
        {
            AssetContainerSasUrls = assetContainerSasUrls;
        }
    }
}
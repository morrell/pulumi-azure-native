// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.V20210601
{
    public static class ListClusterUpgradableVersions
    {
        /// <summary>
        /// The list of intermediate cluster code versions for an upgrade or downgrade. Or minimum and maximum upgradable version if no target was given
        /// </summary>
        public static Task<ListClusterUpgradableVersionsResult> InvokeAsync(ListClusterUpgradableVersionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListClusterUpgradableVersionsResult>("azure-native:servicefabric/v20210601:listClusterUpgradableVersions", args ?? new ListClusterUpgradableVersionsArgs(), options.WithVersion());
    }


    public sealed class ListClusterUpgradableVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The target code version.
        /// </summary>
        [Input("targetVersion", required: true)]
        public string TargetVersion { get; set; } = null!;

        public ListClusterUpgradableVersionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListClusterUpgradableVersionsResult
    {
        public readonly ImmutableArray<string> SupportedPath;

        [OutputConstructor]
        private ListClusterUpgradableVersionsResult(ImmutableArray<string> supportedPath)
        {
            SupportedPath = supportedPath;
        }
    }
}

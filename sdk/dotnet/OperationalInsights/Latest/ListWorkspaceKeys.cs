// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.Latest
{
    public static class ListWorkspaceKeys
    {
        public static Task<ListWorkspaceKeysResult> InvokeAsync(ListWorkspaceKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListWorkspaceKeysResult>("azurerm:operationalinsights/latest:listWorkspaceKeys", args ?? new ListWorkspaceKeysArgs(), options.WithVersion());
    }


    public sealed class ListWorkspaceKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Resource Group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Log Analytics Workspace name.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public ListWorkspaceKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListWorkspaceKeysResult
    {
        /// <summary>
        /// The primary shared key of a workspace.
        /// </summary>
        public readonly string? PrimarySharedKey;
        /// <summary>
        /// The secondary shared key of a workspace.
        /// </summary>
        public readonly string? SecondarySharedKey;

        [OutputConstructor]
        private ListWorkspaceKeysResult(
            string? primarySharedKey,

            string? secondarySharedKey)
        {
            PrimarySharedKey = primarySharedKey;
            SecondarySharedKey = secondarySharedKey;
        }
    }
}
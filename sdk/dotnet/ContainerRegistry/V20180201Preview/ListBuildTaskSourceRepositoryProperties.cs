// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20180201Preview
{
    public static class ListBuildTaskSourceRepositoryProperties
    {
        public static Task<ListBuildTaskSourceRepositoryPropertiesResult> InvokeAsync(ListBuildTaskSourceRepositoryPropertiesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListBuildTaskSourceRepositoryPropertiesResult>("azurerm:containerregistry/v20180201preview:listBuildTaskSourceRepositoryProperties", args ?? new ListBuildTaskSourceRepositoryPropertiesArgs(), options.WithVersion());
    }


    public sealed class ListBuildTaskSourceRepositoryPropertiesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry build task.
        /// </summary>
        [Input("buildTaskName", required: true)]
        public string BuildTaskName { get; set; } = null!;

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListBuildTaskSourceRepositoryPropertiesArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListBuildTaskSourceRepositoryPropertiesResult
    {
        /// <summary>
        /// The value of this property indicates whether the source control commit trigger is enabled or not.
        /// </summary>
        public readonly bool? IsCommitTriggerEnabled;
        /// <summary>
        /// The full URL to the source code repository
        /// </summary>
        public readonly string RepositoryUrl;
        /// <summary>
        /// The authorization properties for accessing the source code repository.
        /// </summary>
        public readonly Outputs.SourceControlAuthInfoResponseResult? SourceControlAuthProperties;
        /// <summary>
        /// The type of source control service.
        /// </summary>
        public readonly string SourceControlType;

        [OutputConstructor]
        private ListBuildTaskSourceRepositoryPropertiesResult(
            bool? isCommitTriggerEnabled,

            string repositoryUrl,

            Outputs.SourceControlAuthInfoResponseResult? sourceControlAuthProperties,

            string sourceControlType)
        {
            IsCommitTriggerEnabled = isCommitTriggerEnabled;
            RepositoryUrl = repositoryUrl;
            SourceControlAuthProperties = sourceControlAuthProperties;
            SourceControlType = sourceControlType;
        }
    }
}
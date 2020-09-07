// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200301
{
    public static class GetWorkspace
    {
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("azurerm:machinelearningservices/v20200301:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithVersion());
    }


    public sealed class GetWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// The flag to indicate whether to allow public access when behind VNet.
        /// </summary>
        public readonly bool? AllowPublicAccessWhenBehindVnet;
        /// <summary>
        /// ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
        /// </summary>
        public readonly string? ApplicationInsights;
        /// <summary>
        /// ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
        /// </summary>
        public readonly string? ContainerRegistry;
        /// <summary>
        /// The creation time of the machine learning workspace in ISO8601 format.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The description of this workspace.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Url for the discovery service to identify regional endpoints for machine learning experimentation services
        /// </summary>
        public readonly string? DiscoveryUrl;
        /// <summary>
        /// The encryption settings of Azure ML workspace.
        /// </summary>
        public readonly Outputs.EncryptionPropertyResponseResult? Encryption;
        /// <summary>
        /// The friendly name for this workspace. This name in mutable
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
        /// </summary>
        public readonly bool? HbiWorkspace;
        /// <summary>
        /// The identity of the resource.
        /// </summary>
        public readonly Outputs.IdentityResponseResult? Identity;
        /// <summary>
        /// The compute name for image build
        /// </summary>
        public readonly string? ImageBuildCompute;
        /// <summary>
        /// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
        /// </summary>
        public readonly string? KeyVault;
        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Specifies the name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The notebook info of Azure ML workspace.
        /// </summary>
        public readonly Outputs.NotebookResourceInfoResponseResult NotebookInfo;
        /// <summary>
        /// The list of private endpoint connections in the workspace.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult> PrivateEndpointConnections;
        /// <summary>
        /// Count of private connections in the workspace
        /// </summary>
        public readonly int PrivateLinkCount;
        /// <summary>
        /// The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
        /// </summary>
        public readonly string ServiceProvisionedResourceGroup;
        /// <summary>
        /// The list of shared private link resources in this workspace.
        /// </summary>
        public readonly ImmutableArray<Outputs.SharedPrivateLinkResourceResponseResult> SharedPrivateLinkResources;
        /// <summary>
        /// The sku of the workspace.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
        /// </summary>
        public readonly string? StorageAccount;
        /// <summary>
        /// Contains resource tags defined as key/value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Specifies the type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The immutable id associated with this workspace.
        /// </summary>
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetWorkspaceResult(
            bool? allowPublicAccessWhenBehindVnet,

            string? applicationInsights,

            string? containerRegistry,

            string creationTime,

            string? description,

            string? discoveryUrl,

            Outputs.EncryptionPropertyResponseResult? encryption,

            string? friendlyName,

            bool? hbiWorkspace,

            Outputs.IdentityResponseResult? identity,

            string? imageBuildCompute,

            string? keyVault,

            string? location,

            string name,

            Outputs.NotebookResourceInfoResponseResult notebookInfo,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult> privateEndpointConnections,

            int privateLinkCount,

            string provisioningState,

            string serviceProvisionedResourceGroup,

            ImmutableArray<Outputs.SharedPrivateLinkResourceResponseResult> sharedPrivateLinkResources,

            Outputs.SkuResponseResult? sku,

            string? storageAccount,

            ImmutableDictionary<string, string>? tags,

            string type,

            string workspaceId)
        {
            AllowPublicAccessWhenBehindVnet = allowPublicAccessWhenBehindVnet;
            ApplicationInsights = applicationInsights;
            ContainerRegistry = containerRegistry;
            CreationTime = creationTime;
            Description = description;
            DiscoveryUrl = discoveryUrl;
            Encryption = encryption;
            FriendlyName = friendlyName;
            HbiWorkspace = hbiWorkspace;
            Identity = identity;
            ImageBuildCompute = imageBuildCompute;
            KeyVault = keyVault;
            Location = location;
            Name = name;
            NotebookInfo = notebookInfo;
            PrivateEndpointConnections = privateEndpointConnections;
            PrivateLinkCount = privateLinkCount;
            ProvisioningState = provisioningState;
            ServiceProvisionedResourceGroup = serviceProvisionedResourceGroup;
            SharedPrivateLinkResources = sharedPrivateLinkResources;
            Sku = sku;
            StorageAccount = storageAccount;
            Tags = tags;
            Type = type;
            WorkspaceId = workspaceId;
        }
    }
}
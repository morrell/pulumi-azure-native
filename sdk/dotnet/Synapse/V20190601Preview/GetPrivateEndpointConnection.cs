// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Synapse.V20190601Preview
{
    public static class GetPrivateEndpointConnection
    {
        public static Task<GetPrivateEndpointConnectionResult> InvokeAsync(GetPrivateEndpointConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateEndpointConnectionResult>("azurerm:synapse/v20190601preview:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionArgs(), options.WithVersion());
    }


    public sealed class GetPrivateEndpointConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the private endpoint connection.
        /// </summary>
        [Input("privateEndpointConnectionName", required: true)]
        public string PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetPrivateEndpointConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateEndpointConnectionResult
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The private endpoint which the connection belongs to.
        /// </summary>
        public readonly Outputs.PrivateEndpointResponseResult? PrivateEndpoint;
        /// <summary>
        /// Connection state of the private endpoint connection.
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStateResponseResult? PrivateLinkServiceConnectionState;
        /// <summary>
        /// Provisioning state of the private endpoint connection.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateEndpointConnectionResult(
            string name,

            Outputs.PrivateEndpointResponseResult? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStateResponseResult? privateLinkServiceConnectionState,

            string provisioningState,

            string type)
        {
            Name = name;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
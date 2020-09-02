// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200901Preview
{
    public static class GetLinkedService
    {
        public static Task<GetLinkedServiceResult> InvokeAsync(GetLinkedServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLinkedServiceResult>("azurerm:machinelearningservices/v20200901preview:getLinkedService", args ?? new GetLinkedServiceArgs(), options.WithVersion());
    }


    public sealed class GetLinkedServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Friendly name of the linked workspace
        /// </summary>
        [Input("linkName", required: true)]
        public string LinkName { get; set; } = null!;

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

        public GetLinkedServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLinkedServiceResult
    {
        /// <summary>
        /// Identity for the resource.
        /// </summary>
        public readonly Outputs.IdentityResponseResult? Identity;
        /// <summary>
        /// location of the linked service.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Friendly name of the linked service.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// LinkedService specific properties.
        /// </summary>
        public readonly Outputs.LinkedServicePropsResponseResult Properties;
        /// <summary>
        /// Resource type of linked service.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetLinkedServiceResult(
            Outputs.IdentityResponseResult? identity,

            string? location,

            string name,

            Outputs.LinkedServicePropsResponseResult properties,

            string type)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
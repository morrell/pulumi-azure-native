// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20210101Preview
{
    public static class GetTag
    {
        /// <summary>
        /// Tag Contract details.
        /// </summary>
        public static Task<GetTagResult> InvokeAsync(GetTagArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagResult>("azure-native:apimanagement/v20210101preview:getTag", args ?? new GetTagArgs(), options.WithVersion());
    }


    public sealed class GetTagArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// Tag identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("tagId", required: true)]
        public string TagId { get; set; } = null!;

        public GetTagArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagResult
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTagResult(
            string displayName,

            string id,

            string name,

            string type)
        {
            DisplayName = displayName;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}

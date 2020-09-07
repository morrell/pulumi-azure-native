// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20160601
{
    public static class GetConnectionGateway
    {
        public static Task<GetConnectionGatewayResult> InvokeAsync(GetConnectionGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionGatewayResult>("azurerm:web/v20160601:getConnectionGateway", args ?? new GetConnectionGatewayArgs(), options.WithVersion());
    }


    public sealed class GetConnectionGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The connection gateway name
        /// </summary>
        [Input("connectionGatewayName", required: true)]
        public string ConnectionGatewayName { get; set; } = null!;

        /// <summary>
        /// The resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConnectionGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionGatewayResult
    {
        /// <summary>
        /// Resource ETag
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        public readonly Outputs.ConnectionGatewayDefinitionResponsePropertiesResult Properties;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConnectionGatewayResult(
            string? etag,

            string? location,

            string name,

            Outputs.ConnectionGatewayDefinitionResponsePropertiesResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
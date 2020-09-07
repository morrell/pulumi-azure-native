// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20161002
{
    public static class GetCustomDomain
    {
        public static Task<GetCustomDomainResult> InvokeAsync(GetCustomDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomDomainResult>("azurerm:cdn/v20161002:getCustomDomain", args ?? new GetCustomDomainArgs(), options.WithVersion());
    }


    public sealed class GetCustomDomainArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the custom domain within an endpoint.
        /// </summary>
        [Input("customDomainName", required: true)]
        public string CustomDomainName { get; set; } = null!;

        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCustomDomainArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomDomainResult
    {
        /// <summary>
        /// Provisioning state of Custom Https of the custom domain.
        /// </summary>
        public readonly string CustomHttpsProvisioningState;
        /// <summary>
        /// The host name of the custom domain. Must be a domain name.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning status of the custom domain.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource status of the custom domain.
        /// </summary>
        public readonly string ResourceState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Special validation or data may be required when delivering CDN to some regions due to local compliance reasons. E.g. ICP license number of a custom domain is required to deliver content in China.
        /// </summary>
        public readonly string? ValidationData;

        [OutputConstructor]
        private GetCustomDomainResult(
            string customHttpsProvisioningState,

            string hostName,

            string location,

            string name,

            string provisioningState,

            string resourceState,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? validationData)
        {
            CustomHttpsProvisioningState = customHttpsProvisioningState;
            HostName = hostName;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceState = resourceState;
            Tags = tags;
            Type = type;
            ValidationData = validationData;
        }
    }
}
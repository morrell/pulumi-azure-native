// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20170301
{
    public static class GetApiManagementService
    {
        public static Task<GetApiManagementServiceResult> InvokeAsync(GetApiManagementServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiManagementServiceResult>("azurerm:apimanagement/v20170301:getApiManagementService", args ?? new GetApiManagementServiceArgs(), options.WithVersion());
    }


    public sealed class GetApiManagementServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApiManagementServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiManagementServiceResult
    {
        /// <summary>
        /// Additional datacenter locations of the API Management service.
        /// </summary>
        public readonly ImmutableArray<Outputs.AdditionalLocationResponseResult> AdditionalLocations;
        /// <summary>
        /// List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
        /// </summary>
        public readonly ImmutableArray<Outputs.CertificateConfigurationResponseResult> Certificates;
        /// <summary>
        /// Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
        /// </summary>
        public readonly string CreatedAtUtc;
        /// <summary>
        /// Custom properties of the API Management service. Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2). Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1 and setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomProperties;
        /// <summary>
        /// ETag of the resource.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Gateway URL of the API Management service in the Default Region.
        /// </summary>
        public readonly string GatewayRegionalUrl;
        /// <summary>
        /// Gateway URL of the API Management service.
        /// </summary>
        public readonly string GatewayUrl;
        /// <summary>
        /// Custom hostname configuration of the API Management service.
        /// </summary>
        public readonly ImmutableArray<Outputs.HostnameConfigurationResponseResult> HostnameConfigurations;
        /// <summary>
        /// Managed service identity of the Api Management service.
        /// </summary>
        public readonly Outputs.ApiManagementServiceIdentityResponseResult? Identity;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Management API endpoint URL of the API Management service.
        /// </summary>
        public readonly string ManagementApiUrl;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Email address from which the notification will be sent.
        /// </summary>
        public readonly string? NotificationSenderEmail;
        /// <summary>
        /// Publisher portal endpoint Url of the API Management service.
        /// </summary>
        public readonly string PortalUrl;
        /// <summary>
        /// The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Publisher email.
        /// </summary>
        public readonly string PublisherEmail;
        /// <summary>
        /// Publisher name.
        /// </summary>
        public readonly string PublisherName;
        /// <summary>
        /// SCM endpoint URL of the API Management service.
        /// </summary>
        public readonly string ScmUrl;
        /// <summary>
        /// SKU properties of the API Management service.
        /// </summary>
        public readonly Outputs.ApiManagementServiceSkuPropertiesResponseResult Sku;
        /// <summary>
        /// Static IP addresses of the API Management service virtual machines. Available only for Standard and Premium SKU.
        /// </summary>
        public readonly ImmutableArray<string> StaticIps;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
        /// </summary>
        public readonly string TargetProvisioningState;
        /// <summary>
        /// Resource type for API Management resource is set to Microsoft.ApiManagement.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Virtual network configuration of the API Management service.
        /// </summary>
        public readonly Outputs.VirtualNetworkConfigurationResponseResult? VirtualNetworkConfiguration;
        /// <summary>
        /// The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
        /// </summary>
        public readonly string? VirtualNetworkType;

        [OutputConstructor]
        private GetApiManagementServiceResult(
            ImmutableArray<Outputs.AdditionalLocationResponseResult> additionalLocations,

            ImmutableArray<Outputs.CertificateConfigurationResponseResult> certificates,

            string createdAtUtc,

            ImmutableDictionary<string, string>? customProperties,

            string etag,

            string gatewayRegionalUrl,

            string gatewayUrl,

            ImmutableArray<Outputs.HostnameConfigurationResponseResult> hostnameConfigurations,

            Outputs.ApiManagementServiceIdentityResponseResult? identity,

            string location,

            string managementApiUrl,

            string name,

            string? notificationSenderEmail,

            string portalUrl,

            string provisioningState,

            string publisherEmail,

            string publisherName,

            string scmUrl,

            Outputs.ApiManagementServiceSkuPropertiesResponseResult sku,

            ImmutableArray<string> staticIps,

            ImmutableDictionary<string, string>? tags,

            string targetProvisioningState,

            string type,

            Outputs.VirtualNetworkConfigurationResponseResult? virtualNetworkConfiguration,

            string? virtualNetworkType)
        {
            AdditionalLocations = additionalLocations;
            Certificates = certificates;
            CreatedAtUtc = createdAtUtc;
            CustomProperties = customProperties;
            Etag = etag;
            GatewayRegionalUrl = gatewayRegionalUrl;
            GatewayUrl = gatewayUrl;
            HostnameConfigurations = hostnameConfigurations;
            Identity = identity;
            Location = location;
            ManagementApiUrl = managementApiUrl;
            Name = name;
            NotificationSenderEmail = notificationSenderEmail;
            PortalUrl = portalUrl;
            ProvisioningState = provisioningState;
            PublisherEmail = publisherEmail;
            PublisherName = publisherName;
            ScmUrl = scmUrl;
            Sku = sku;
            StaticIps = staticIps;
            Tags = tags;
            TargetProvisioningState = targetProvisioningState;
            Type = type;
            VirtualNetworkConfiguration = virtualNetworkConfiguration;
            VirtualNetworkType = virtualNetworkType;
        }
    }
}
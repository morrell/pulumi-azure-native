// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330
{
    public static class GetPublicIPAddress
    {
        public static Task<GetPublicIPAddressResult> InvokeAsync(GetPublicIPAddressArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicIPAddressResult>("azurerm:network/v20160330:getPublicIPAddress", args ?? new GetPublicIPAddressArgs(), options.WithVersion());
    }


    public sealed class GetPublicIPAddressArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// expand references resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the subnet.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPublicIPAddressArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublicIPAddressResult
    {
        /// <summary>
        /// Gets or sets FQDN of the DNS record associated with the public IP address
        /// </summary>
        public readonly Outputs.PublicIPAddressDnsSettingsResponseResult? DnsSettings;
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Gets or sets the idle timeout of the public IP address
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        public readonly string? IpAddress;
        /// <summary>
        /// IPConfiguration
        /// </summary>
        public readonly Outputs.IPConfigurationResponseResult? IpConfiguration;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets PublicIP address version (IPv4/IPv6)
        /// </summary>
        public readonly string? PublicIPAddressVersion;
        /// <summary>
        /// Gets or sets PublicIP allocation method (Static/Dynamic)
        /// </summary>
        public readonly string? PublicIPAllocationMethod;
        /// <summary>
        /// Gets or sets resource GUID property of the PublicIP resource
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPublicIPAddressResult(
            Outputs.PublicIPAddressDnsSettingsResponseResult? dnsSettings,

            string? etag,

            int? idleTimeoutInMinutes,

            string? ipAddress,

            Outputs.IPConfigurationResponseResult? ipConfiguration,

            string? location,

            string name,

            string? provisioningState,

            string? publicIPAddressVersion,

            string? publicIPAllocationMethod,

            string? resourceGuid,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            DnsSettings = dnsSettings;
            Etag = etag;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            IpAddress = ipAddress;
            IpConfiguration = ipConfiguration;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            PublicIPAddressVersion = publicIPAddressVersion;
            PublicIPAllocationMethod = publicIPAllocationMethod;
            ResourceGuid = resourceGuid;
            Tags = tags;
            Type = type;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayBackendHttpSettingsResponseResult
    {
        /// <summary>
        /// Array of references to Application Gateway Authentication Certificates
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponseResult> AuthenticationCertificates;
        /// <summary>
        /// Cookie affinity
        /// </summary>
        public readonly string? CookieBasedAffinity;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Port
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Probe resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? Probe;
        /// <summary>
        /// Protocol
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Provisioning state of the backend http settings resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Request timeout
        /// </summary>
        public readonly int? RequestTimeout;

        [OutputConstructor]
        private ApplicationGatewayBackendHttpSettingsResponseResult(
            ImmutableArray<Outputs.SubResourceResponseResult> authenticationCertificates,

            string? cookieBasedAffinity,

            string? etag,

            string? id,

            string? name,

            int? port,

            Outputs.SubResourceResponseResult? probe,

            string? protocol,

            string? provisioningState,

            int? requestTimeout)
        {
            AuthenticationCertificates = authenticationCertificates;
            CookieBasedAffinity = cookieBasedAffinity;
            Etag = etag;
            Id = id;
            Name = name;
            Port = port;
            Probe = probe;
            Protocol = protocol;
            ProvisioningState = provisioningState;
            RequestTimeout = requestTimeout;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20210101.Outputs
{

    [OutputType]
    public sealed class OpenIdConnectRegistrationResponse
    {
        /// <summary>
        /// The authentication credentials of the custom Open ID Connect provider.
        /// </summary>
        public readonly Outputs.OpenIdConnectClientCredentialResponse? ClientCredential;
        /// <summary>
        /// The client id of the custom Open ID Connect provider.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The configuration settings of the endpoints used for the custom Open ID Connect provider.
        /// </summary>
        public readonly Outputs.OpenIdConnectConfigResponse? OpenIdConnectConfiguration;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private OpenIdConnectRegistrationResponse(
            Outputs.OpenIdConnectClientCredentialResponse? clientCredential,

            string? clientId,

            string id,

            string? kind,

            string name,

            Outputs.OpenIdConnectConfigResponse? openIdConnectConfiguration,

            string type)
        {
            ClientCredential = clientCredential;
            ClientId = clientId;
            Id = id;
            Kind = kind;
            Name = name;
            OpenIdConnectConfiguration = openIdConnectConfiguration;
            Type = type;
        }
    }
}

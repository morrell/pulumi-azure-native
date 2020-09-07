// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.Latest.Outputs
{

    [OutputType]
    public sealed class ConnectionGatewayDefinitionResponsePropertiesResult
    {
        /// <summary>
        /// The URI of the backend
        /// </summary>
        public readonly string? BackendUri;
        /// <summary>
        /// The gateway installation reference
        /// </summary>
        public readonly Outputs.ConnectionGatewayReferenceResponseResult? ConnectionGatewayInstallation;
        /// <summary>
        /// The gateway admin
        /// </summary>
        public readonly ImmutableArray<string> ContactInformation;
        /// <summary>
        /// The gateway description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The gateway display name
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The machine name of the gateway
        /// </summary>
        public readonly string? MachineName;
        /// <summary>
        /// The gateway status
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Status;

        [OutputConstructor]
        private ConnectionGatewayDefinitionResponsePropertiesResult(
            string? backendUri,

            Outputs.ConnectionGatewayReferenceResponseResult? connectionGatewayInstallation,

            ImmutableArray<string> contactInformation,

            string? description,

            string? displayName,

            string? machineName,

            ImmutableDictionary<string, object>? status)
        {
            BackendUri = backendUri;
            ConnectionGatewayInstallation = connectionGatewayInstallation;
            ContactInformation = contactInformation;
            Description = description;
            DisplayName = displayName;
            MachineName = machineName;
            Status = status;
        }
    }
}
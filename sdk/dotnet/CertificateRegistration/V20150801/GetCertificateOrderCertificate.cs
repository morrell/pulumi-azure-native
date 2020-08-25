// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CertificateRegistration.V20150801
{
    public static class GetCertificateOrderCertificate
    {
        public static Task<GetCertificateOrderCertificateResult> InvokeAsync(GetCertificateOrderCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateOrderCertificateResult>("azurerm:certificateregistration/v20150801:getCertificateOrderCertificate", args ?? new GetCertificateOrderCertificateArgs(), options.WithVersion());
    }


    public sealed class GetCertificateOrderCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Certificate name
        /// </summary>
        [Input("certificateOrderName", required: true)]
        public string CertificateOrderName { get; set; } = null!;

        /// <summary>
        /// Certificate name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Azure resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCertificateOrderCertificateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateOrderCertificateResult
    {
        /// <summary>
        /// Key Vault Csm resource Id
        /// </summary>
        public readonly string? KeyVaultId;
        /// <summary>
        /// Key Vault secret name
        /// </summary>
        public readonly string? KeyVaultSecretName;
        /// <summary>
        /// Kind of resource
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Status of the Key Vault secret
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetCertificateOrderCertificateResult(
            string? keyVaultId,

            string? keyVaultSecretName,

            string? kind,

            string location,

            string? name,

            string? provisioningState,

            ImmutableDictionary<string, string>? tags,

            string? type)
        {
            KeyVaultId = keyVaultId;
            KeyVaultSecretName = keyVaultSecretName;
            Kind = kind;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
        }
    }
}
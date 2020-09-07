// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190501.Outputs
{

    [OutputType]
    public sealed class CustomHttpsConfigurationResponseResult
    {
        /// <summary>
        /// Defines the source of the SSL certificate
        /// </summary>
        public readonly string CertificateSource;
        /// <summary>
        /// Defines the type of the certificate used for secure connections to a frontendEndpoint
        /// </summary>
        public readonly string? CertificateType;
        /// <summary>
        /// The minimum TLS version required from the clients to establish an SSL handshake with Front Door.
        /// </summary>
        public readonly string MinimumTlsVersion;
        /// <summary>
        /// Defines the TLS extension protocol that is used for secure delivery
        /// </summary>
        public readonly string ProtocolType;
        /// <summary>
        /// The name of the Key Vault secret representing the full certificate PFX
        /// </summary>
        public readonly string? SecretName;
        /// <summary>
        /// The version of the Key Vault secret representing the full certificate PFX
        /// </summary>
        public readonly string? SecretVersion;
        /// <summary>
        /// The Key Vault containing the SSL certificate
        /// </summary>
        public readonly Outputs.KeyVaultCertificateSourceParametersResponseVaultResult? Vault;

        [OutputConstructor]
        private CustomHttpsConfigurationResponseResult(
            string certificateSource,

            string? certificateType,

            string minimumTlsVersion,

            string protocolType,

            string? secretName,

            string? secretVersion,

            Outputs.KeyVaultCertificateSourceParametersResponseVaultResult? vault)
        {
            CertificateSource = certificateSource;
            CertificateType = certificateType;
            MinimumTlsVersion = minimumTlsVersion;
            ProtocolType = protocolType;
            SecretName = secretName;
            SecretVersion = secretVersion;
            Vault = vault;
        }
    }
}
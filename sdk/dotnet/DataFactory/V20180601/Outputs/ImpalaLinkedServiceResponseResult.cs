// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20180601.Outputs
{

    [OutputType]
    public sealed class ImpalaLinkedServiceResponseResult
    {
        /// <summary>
        /// Specifies whether to require a CA-issued SSL certificate name to match the host name of the server when connecting over SSL. The default value is false.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AllowHostNameCNMismatch;
        /// <summary>
        /// Specifies whether to allow self-signed certificates from the server. The default value is false.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AllowSelfSignedServerCert;
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// The authentication type to use.
        /// </summary>
        public readonly string AuthenticationType;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponseResult? ConnectVia;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Specifies whether the connections to the server are encrypted using SSL. The default value is false.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? EnableSsl;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? EncryptedCredential;
        /// <summary>
        /// The IP address or host name of the Impala server. (i.e. 192.168.222.160)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Host;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? Parameters;
        /// <summary>
        /// The password corresponding to the user name when using UsernameAndPassword.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? Password;
        /// <summary>
        /// The TCP port that the Impala server uses to listen for client connections. The default value is 21050.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Port;
        /// <summary>
        /// The full path of the .pem file containing trusted CA certificates for verifying the server when connecting over SSL. This property can only be set when using SSL on self-hosted IR. The default value is the cacerts.pem file installed with the IR.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? TrustedCertPath;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specifies whether to use a CA certificate from the system trust store or from a specified PEM file. The default value is false.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? UseSystemTrustStore;
        /// <summary>
        /// The user name used to access the Impala server. The default value is anonymous when using SASLUsername.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Username;

        [OutputConstructor]
        private ImpalaLinkedServiceResponseResult(
            ImmutableDictionary<string, object>? allowHostNameCNMismatch,

            ImmutableDictionary<string, object>? allowSelfSignedServerCert,

            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            string authenticationType,

            Outputs.IntegrationRuntimeReferenceResponseResult? connectVia,

            string? description,

            ImmutableDictionary<string, object>? enableSsl,

            ImmutableDictionary<string, object>? encryptedCredential,

            ImmutableDictionary<string, object> host,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? password,

            ImmutableDictionary<string, object>? port,

            ImmutableDictionary<string, object>? trustedCertPath,

            string type,

            ImmutableDictionary<string, object>? useSystemTrustStore,

            ImmutableDictionary<string, object>? username)
        {
            AllowHostNameCNMismatch = allowHostNameCNMismatch;
            AllowSelfSignedServerCert = allowSelfSignedServerCert;
            Annotations = annotations;
            AuthenticationType = authenticationType;
            ConnectVia = connectVia;
            Description = description;
            EnableSsl = enableSsl;
            EncryptedCredential = encryptedCredential;
            Host = host;
            Parameters = parameters;
            Password = password;
            Port = port;
            TrustedCertPath = trustedCertPath;
            Type = type;
            UseSystemTrustStore = useSystemTrustStore;
            Username = username;
        }
    }
}
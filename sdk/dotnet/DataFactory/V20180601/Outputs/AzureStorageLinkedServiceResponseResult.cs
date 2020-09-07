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
    public sealed class AzureStorageLinkedServiceResponseResult
    {
        /// <summary>
        /// The Azure key vault secret reference of accountKey in connection string.
        /// </summary>
        public readonly Outputs.AzureKeyVaultSecretReferenceResponseResult? AccountKey;
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponseResult? ConnectVia;
        /// <summary>
        /// The connection string. It is mutually exclusive with sasUri property. Type: string, SecureString or AzureKeyVaultSecretReference.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? ConnectionString;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly string? EncryptedCredential;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? Parameters;
        /// <summary>
        /// The Azure key vault secret reference of sasToken in sas uri.
        /// </summary>
        public readonly Outputs.AzureKeyVaultSecretReferenceResponseResult? SasToken;
        /// <summary>
        /// SAS URI of the Azure Storage resource. It is mutually exclusive with connectionString property. Type: string, SecureString or AzureKeyVaultSecretReference.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? SasUri;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AzureStorageLinkedServiceResponseResult(
            Outputs.AzureKeyVaultSecretReferenceResponseResult? accountKey,

            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            Outputs.IntegrationRuntimeReferenceResponseResult? connectVia,

            ImmutableDictionary<string, object>? connectionString,

            string? description,

            string? encryptedCredential,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            Outputs.AzureKeyVaultSecretReferenceResponseResult? sasToken,

            ImmutableDictionary<string, object>? sasUri,

            string type)
        {
            AccountKey = accountKey;
            Annotations = annotations;
            ConnectVia = connectVia;
            ConnectionString = connectionString;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Parameters = parameters;
            SasToken = sasToken;
            SasUri = sasUri;
            Type = type;
        }
    }
}
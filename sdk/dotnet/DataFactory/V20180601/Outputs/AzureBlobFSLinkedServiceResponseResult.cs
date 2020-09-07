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
    public sealed class AzureBlobFSLinkedServiceResponseResult
    {
        /// <summary>
        /// Account key for the Azure Data Lake Storage Gen2 service. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AccountKey;
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// Indicates the azure cloud type of the service principle auth. Allowed values are AzurePublic, AzureChina, AzureUsGovernment, AzureGermany. Default value is the data factory regions’ cloud type. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AzureCloudType;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponseResult? ConnectVia;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? EncryptedCredential;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? Parameters;
        /// <summary>
        /// The ID of the application used to authenticate against the Azure Data Lake Storage Gen2 account. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? ServicePrincipalId;
        /// <summary>
        /// The Key of the application used to authenticate against the Azure Data Lake Storage Gen2 account.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? ServicePrincipalKey;
        /// <summary>
        /// The name or ID of the tenant to which the service principal belongs. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tenant;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Endpoint for the Azure Data Lake Storage Gen2 service. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object> Url;

        [OutputConstructor]
        private AzureBlobFSLinkedServiceResponseResult(
            ImmutableDictionary<string, object>? accountKey,

            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            ImmutableDictionary<string, object>? azureCloudType,

            Outputs.IntegrationRuntimeReferenceResponseResult? connectVia,

            string? description,

            ImmutableDictionary<string, object>? encryptedCredential,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            ImmutableDictionary<string, object>? servicePrincipalId,

            Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? servicePrincipalKey,

            ImmutableDictionary<string, object>? tenant,

            string type,

            ImmutableDictionary<string, object> url)
        {
            AccountKey = accountKey;
            Annotations = annotations;
            AzureCloudType = azureCloudType;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Parameters = parameters;
            ServicePrincipalId = servicePrincipalId;
            ServicePrincipalKey = servicePrincipalKey;
            Tenant = tenant;
            Type = type;
            Url = url;
        }
    }
}
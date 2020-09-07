// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.Latest.Outputs
{

    [OutputType]
    public sealed class AzureDataLakeStoreLinkedServiceResponseResult
    {
        /// <summary>
        /// Data Lake Store account name. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AccountName;
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
        /// Data Lake Store service URI. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object> DataLakeStoreUri;
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
        /// Data Lake Store account resource group name (if different from Data Factory account). Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? ResourceGroupName;
        /// <summary>
        /// The ID of the application used to authenticate against the Azure Data Lake Store account. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? ServicePrincipalId;
        /// <summary>
        /// The Key of the application used to authenticate against the Azure Data Lake Store account.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? ServicePrincipalKey;
        /// <summary>
        /// Data Lake Store account subscription ID (if different from Data Factory account). Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? SubscriptionId;
        /// <summary>
        /// The name or ID of the tenant to which the service principal belongs. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tenant;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AzureDataLakeStoreLinkedServiceResponseResult(
            ImmutableDictionary<string, object>? accountName,

            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            ImmutableDictionary<string, object>? azureCloudType,

            Outputs.IntegrationRuntimeReferenceResponseResult? connectVia,

            ImmutableDictionary<string, object> dataLakeStoreUri,

            string? description,

            ImmutableDictionary<string, object>? encryptedCredential,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            ImmutableDictionary<string, object>? resourceGroupName,

            ImmutableDictionary<string, object>? servicePrincipalId,

            Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? servicePrincipalKey,

            ImmutableDictionary<string, object>? subscriptionId,

            ImmutableDictionary<string, object>? tenant,

            string type)
        {
            AccountName = accountName;
            Annotations = annotations;
            AzureCloudType = azureCloudType;
            ConnectVia = connectVia;
            DataLakeStoreUri = dataLakeStoreUri;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Parameters = parameters;
            ResourceGroupName = resourceGroupName;
            ServicePrincipalId = servicePrincipalId;
            ServicePrincipalKey = servicePrincipalKey;
            SubscriptionId = subscriptionId;
            Tenant = tenant;
            Type = type;
        }
    }
}
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
    public sealed class AzureBatchLinkedServiceResponseResult
    {
        /// <summary>
        /// The Azure Batch account access key.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? AccessKey;
        /// <summary>
        /// The Azure Batch account name. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object> AccountName;
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// The Azure Batch URI. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object> BatchUri;
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
        /// The Azure Storage linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponseResult LinkedServiceName;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? Parameters;
        /// <summary>
        /// The Azure Batch pool name. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object> PoolName;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AzureBatchLinkedServiceResponseResult(
            Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? accessKey,

            ImmutableDictionary<string, object> accountName,

            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            ImmutableDictionary<string, object> batchUri,

            Outputs.IntegrationRuntimeReferenceResponseResult? connectVia,

            string? description,

            ImmutableDictionary<string, object>? encryptedCredential,

            Outputs.LinkedServiceReferenceResponseResult linkedServiceName,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            ImmutableDictionary<string, object> poolName,

            string type)
        {
            AccessKey = accessKey;
            AccountName = accountName;
            Annotations = annotations;
            BatchUri = batchUri;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            LinkedServiceName = linkedServiceName;
            Parameters = parameters;
            PoolName = poolName;
            Type = type;
        }
    }
}
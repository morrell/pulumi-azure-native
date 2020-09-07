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
    public sealed class QuickBooksLinkedServiceResponseResult
    {
        /// <summary>
        /// The access token for OAuth 1.0 authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? AccessToken;
        /// <summary>
        /// The access token secret for OAuth 1.0 authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? AccessTokenSecret;
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// The company ID of the QuickBooks company to authorize.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? CompanyId;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponseResult? ConnectVia;
        /// <summary>
        /// Properties used to connect to QuickBooks. It is mutually exclusive with any other properties in the linked service. Type: object.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? ConnectionProperties;
        /// <summary>
        /// The consumer key for OAuth 1.0 authentication.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? ConsumerKey;
        /// <summary>
        /// The consumer secret for OAuth 1.0 authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? ConsumerSecret;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? EncryptedCredential;
        /// <summary>
        /// The endpoint of the QuickBooks server. (i.e. quickbooks.api.intuit.com)
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Endpoint;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? Parameters;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specifies whether the data source endpoints are encrypted using HTTPS. The default value is true.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? UseEncryptedEndpoints;

        [OutputConstructor]
        private QuickBooksLinkedServiceResponseResult(
            Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? accessToken,

            Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? accessTokenSecret,

            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            ImmutableDictionary<string, object>? companyId,

            Outputs.IntegrationRuntimeReferenceResponseResult? connectVia,

            ImmutableDictionary<string, object>? connectionProperties,

            ImmutableDictionary<string, object>? consumerKey,

            Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? consumerSecret,

            string? description,

            ImmutableDictionary<string, object>? encryptedCredential,

            ImmutableDictionary<string, object>? endpoint,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            string type,

            ImmutableDictionary<string, object>? useEncryptedEndpoints)
        {
            AccessToken = accessToken;
            AccessTokenSecret = accessTokenSecret;
            Annotations = annotations;
            CompanyId = companyId;
            ConnectVia = connectVia;
            ConnectionProperties = connectionProperties;
            ConsumerKey = consumerKey;
            ConsumerSecret = consumerSecret;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Endpoint = endpoint;
            Parameters = parameters;
            Type = type;
            UseEncryptedEndpoints = useEncryptedEndpoints;
        }
    }
}
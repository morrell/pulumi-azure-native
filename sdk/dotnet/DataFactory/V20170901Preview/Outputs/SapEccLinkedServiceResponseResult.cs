// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class SapEccLinkedServiceResponseResult
    {
        /// <summary>
        /// List of tags that can be used for describing the Dataset.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponseResult? ConnectVia;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Either encryptedCredential or username/password must be provided. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly string? EncryptedCredential;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? Parameters;
        /// <summary>
        /// The password for Basic authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? Password;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The URL of SAP ECC OData API. For example, '[https://hostname:port/sap/opu/odata/sap/servicename/]'. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// The username for Basic authentication. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private SapEccLinkedServiceResponseResult(
            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            Outputs.IntegrationRuntimeReferenceResponseResult? connectVia,

            string? description,

            string? encryptedCredential,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponseResult, Outputs.SecureStringResponseResult>? password,

            string type,

            string url,

            string? username)
        {
            Annotations = annotations;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Parameters = parameters;
            Password = password;
            Type = type;
            Url = url;
            Username = username;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20180601.Inputs
{

    /// <summary>
    /// Rest Service linked service.
    /// </summary>
    public sealed class RestServiceLinkedServiceArgs : Pulumi.ResourceArgs
    {
        [Input("aadResourceId")]
        private InputMap<object>? _aadResourceId;

        /// <summary>
        /// The resource you are requesting authorization to use.
        /// </summary>
        public InputMap<object> AadResourceId
        {
            get => _aadResourceId ?? (_aadResourceId = new InputMap<object>());
            set => _aadResourceId = value;
        }

        [Input("annotations")]
        private InputList<ImmutableDictionary<string, object>>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<ImmutableDictionary<string, object>>());
            set => _annotations = value;
        }

        /// <summary>
        /// Type of authentication used to connect to the REST service.
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        [Input("azureCloudType")]
        private InputMap<object>? _azureCloudType;

        /// <summary>
        /// Indicates the azure cloud type of the service principle auth. Allowed values are AzurePublic, AzureChina, AzureUsGovernment, AzureGermany. Default value is the data factory regions’ cloud type. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> AzureCloudType
        {
            get => _azureCloudType ?? (_azureCloudType = new InputMap<object>());
            set => _azureCloudType = value;
        }

        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

        /// <summary>
        /// Linked service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enableServerCertificateValidation")]
        private InputMap<object>? _enableServerCertificateValidation;

        /// <summary>
        /// Whether to validate server side SSL certificate when connecting to the endpoint.The default value is true. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public InputMap<object> EnableServerCertificateValidation
        {
            get => _enableServerCertificateValidation ?? (_enableServerCertificateValidation = new InputMap<object>());
            set => _enableServerCertificateValidation = value;
        }

        [Input("encryptedCredential")]
        private InputMap<object>? _encryptedCredential;

        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> EncryptedCredential
        {
            get => _encryptedCredential ?? (_encryptedCredential = new InputMap<object>());
            set => _encryptedCredential = value;
        }

        [Input("parameters")]
        private InputMap<Inputs.ParameterSpecificationArgs>? _parameters;

        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public InputMap<Inputs.ParameterSpecificationArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<Inputs.ParameterSpecificationArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The password used in Basic authentication type.
        /// </summary>
        [Input("password")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? Password { get; set; }

        [Input("servicePrincipalId")]
        private InputMap<object>? _servicePrincipalId;

        /// <summary>
        /// The application's client ID used in AadServicePrincipal authentication type.
        /// </summary>
        public InputMap<object> ServicePrincipalId
        {
            get => _servicePrincipalId ?? (_servicePrincipalId = new InputMap<object>());
            set => _servicePrincipalId = value;
        }

        /// <summary>
        /// The application's key used in AadServicePrincipal authentication type.
        /// </summary>
        [Input("servicePrincipalKey")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? ServicePrincipalKey { get; set; }

        [Input("tenant")]
        private InputMap<object>? _tenant;

        /// <summary>
        /// The tenant information (domain name or tenant ID) used in AadServicePrincipal authentication type under which your application resides.
        /// </summary>
        public InputMap<object> Tenant
        {
            get => _tenant ?? (_tenant = new InputMap<object>());
            set => _tenant = value;
        }

        /// <summary>
        /// Type of linked service.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("url", required: true)]
        private InputMap<object>? _url;

        /// <summary>
        /// The base URL of the REST service.
        /// </summary>
        public InputMap<object> Url
        {
            get => _url ?? (_url = new InputMap<object>());
            set => _url = value;
        }

        [Input("userName")]
        private InputMap<object>? _userName;

        /// <summary>
        /// The user name used in Basic authentication type.
        /// </summary>
        public InputMap<object> UserName
        {
            get => _userName ?? (_userName = new InputMap<object>());
            set => _userName = value;
        }

        public RestServiceLinkedServiceArgs()
        {
        }
    }
}
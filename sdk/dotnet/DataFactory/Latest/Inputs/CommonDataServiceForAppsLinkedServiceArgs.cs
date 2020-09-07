// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.Latest.Inputs
{

    /// <summary>
    /// Common Data Service for Apps linked service.
    /// </summary>
    public sealed class CommonDataServiceForAppsLinkedServiceArgs : Pulumi.ResourceArgs
    {
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
        /// The authentication type to connect to Common Data Service for Apps server. 'Office365' for online scenario, 'Ifd' for on-premises with Ifd scenario. 'AADServicePrincipal' for Server-To-Server authentication in online scenario. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

        /// <summary>
        /// The deployment type of the Common Data Service for Apps instance. 'Online' for Common Data Service for Apps Online and 'OnPremisesWithIfd' for Common Data Service for Apps on-premises with Ifd. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("deploymentType", required: true)]
        public Input<string> DeploymentType { get; set; } = null!;

        /// <summary>
        /// Linked service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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

        [Input("hostName")]
        private InputMap<object>? _hostName;

        /// <summary>
        /// The host name of the on-premises Common Data Service for Apps server. The property is required for on-prem and not allowed for online. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> HostName
        {
            get => _hostName ?? (_hostName = new InputMap<object>());
            set => _hostName = value;
        }

        [Input("organizationName")]
        private InputMap<object>? _organizationName;

        /// <summary>
        /// The organization name of the Common Data Service for Apps instance. The property is required for on-prem and required for online when there are more than one Common Data Service for Apps instances associated with the user. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> OrganizationName
        {
            get => _organizationName ?? (_organizationName = new InputMap<object>());
            set => _organizationName = value;
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
        /// Password to access the Common Data Service for Apps instance.
        /// </summary>
        [Input("password")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? Password { get; set; }

        [Input("port")]
        private InputMap<object>? _port;

        /// <summary>
        /// The port of on-premises Common Data Service for Apps server. The property is required for on-prem and not allowed for online. Default is 443. Type: integer (or Expression with resultType integer), minimum: 0.
        /// </summary>
        public InputMap<object> Port
        {
            get => _port ?? (_port = new InputMap<object>());
            set => _port = value;
        }

        /// <summary>
        /// The credential of the service principal object in Azure Active Directory. If servicePrincipalCredentialType is 'ServicePrincipalKey', servicePrincipalCredential can be SecureString or AzureKeyVaultSecretReference. If servicePrincipalCredentialType is 'ServicePrincipalCert', servicePrincipalCredential can only be AzureKeyVaultSecretReference.
        /// </summary>
        [Input("servicePrincipalCredential")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? ServicePrincipalCredential { get; set; }

        [Input("servicePrincipalCredentialType")]
        private InputMap<object>? _servicePrincipalCredentialType;

        /// <summary>
        /// The service principal credential type to use in Server-To-Server authentication. 'ServicePrincipalKey' for key/secret, 'ServicePrincipalCert' for certificate. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> ServicePrincipalCredentialType
        {
            get => _servicePrincipalCredentialType ?? (_servicePrincipalCredentialType = new InputMap<object>());
            set => _servicePrincipalCredentialType = value;
        }

        [Input("servicePrincipalId")]
        private InputMap<object>? _servicePrincipalId;

        /// <summary>
        /// The client ID of the application in Azure Active Directory used for Server-To-Server authentication. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> ServicePrincipalId
        {
            get => _servicePrincipalId ?? (_servicePrincipalId = new InputMap<object>());
            set => _servicePrincipalId = value;
        }

        [Input("serviceUri")]
        private InputMap<object>? _serviceUri;

        /// <summary>
        /// The URL to the Microsoft Common Data Service for Apps server. The property is required for on-line and not allowed for on-prem. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> ServiceUri
        {
            get => _serviceUri ?? (_serviceUri = new InputMap<object>());
            set => _serviceUri = value;
        }

        /// <summary>
        /// Type of linked service.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("username")]
        private InputMap<object>? _username;

        /// <summary>
        /// User name to access the Common Data Service for Apps instance. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> Username
        {
            get => _username ?? (_username = new InputMap<object>());
            set => _username = value;
        }

        public CommonDataServiceForAppsLinkedServiceArgs()
        {
        }
    }
}
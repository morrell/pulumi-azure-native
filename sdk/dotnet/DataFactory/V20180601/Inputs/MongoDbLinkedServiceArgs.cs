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
    /// Linked service for MongoDb data source.
    /// </summary>
    public sealed class MongoDbLinkedServiceArgs : Pulumi.ResourceArgs
    {
        [Input("allowSelfSignedServerCert")]
        private InputMap<object>? _allowSelfSignedServerCert;

        /// <summary>
        /// Specifies whether to allow self-signed certificates from the server. The default value is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public InputMap<object> AllowSelfSignedServerCert
        {
            get => _allowSelfSignedServerCert ?? (_allowSelfSignedServerCert = new InputMap<object>());
            set => _allowSelfSignedServerCert = value;
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

        [Input("authSource")]
        private InputMap<object>? _authSource;

        /// <summary>
        /// Database to verify the username and password. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> AuthSource
        {
            get => _authSource ?? (_authSource = new InputMap<object>());
            set => _authSource = value;
        }

        /// <summary>
        /// The authentication type to be used to connect to the MongoDB database.
        /// </summary>
        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

        [Input("databaseName", required: true)]
        private InputMap<object>? _databaseName;

        /// <summary>
        /// The name of the MongoDB database that you want to access. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> DatabaseName
        {
            get => _databaseName ?? (_databaseName = new InputMap<object>());
            set => _databaseName = value;
        }

        /// <summary>
        /// Linked service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enableSsl")]
        private InputMap<object>? _enableSsl;

        /// <summary>
        /// Specifies whether the connections to the server are encrypted using SSL. The default value is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public InputMap<object> EnableSsl
        {
            get => _enableSsl ?? (_enableSsl = new InputMap<object>());
            set => _enableSsl = value;
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
        /// Password for authentication.
        /// </summary>
        [Input("password")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? Password { get; set; }

        [Input("port")]
        private InputMap<object>? _port;

        /// <summary>
        /// The TCP port number that the MongoDB server uses to listen for client connections. The default value is 27017. Type: integer (or Expression with resultType integer), minimum: 0.
        /// </summary>
        public InputMap<object> Port
        {
            get => _port ?? (_port = new InputMap<object>());
            set => _port = value;
        }

        [Input("server", required: true)]
        private InputMap<object>? _server;

        /// <summary>
        /// The IP address or server name of the MongoDB server. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> Server
        {
            get => _server ?? (_server = new InputMap<object>());
            set => _server = value;
        }

        /// <summary>
        /// Type of linked service.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("username")]
        private InputMap<object>? _username;

        /// <summary>
        /// Username for authentication. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> Username
        {
            get => _username ?? (_username = new InputMap<object>());
            set => _username = value;
        }

        public MongoDbLinkedServiceArgs()
        {
        }
    }
}
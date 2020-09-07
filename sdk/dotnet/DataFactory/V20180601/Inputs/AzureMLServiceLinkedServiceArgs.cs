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
    /// Azure ML Service linked service.
    /// </summary>
    public sealed class AzureMLServiceLinkedServiceArgs : Pulumi.ResourceArgs
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
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

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

        [Input("mlWorkspaceName", required: true)]
        private InputMap<object>? _mlWorkspaceName;

        /// <summary>
        /// Azure ML Service workspace name. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> MlWorkspaceName
        {
            get => _mlWorkspaceName ?? (_mlWorkspaceName = new InputMap<object>());
            set => _mlWorkspaceName = value;
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

        [Input("resourceGroupName", required: true)]
        private InputMap<object>? _resourceGroupName;

        /// <summary>
        /// Azure ML Service workspace resource group name. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> ResourceGroupName
        {
            get => _resourceGroupName ?? (_resourceGroupName = new InputMap<object>());
            set => _resourceGroupName = value;
        }

        [Input("servicePrincipalId")]
        private InputMap<object>? _servicePrincipalId;

        /// <summary>
        /// The ID of the service principal used to authenticate against the endpoint of a published Azure ML Service pipeline. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> ServicePrincipalId
        {
            get => _servicePrincipalId ?? (_servicePrincipalId = new InputMap<object>());
            set => _servicePrincipalId = value;
        }

        /// <summary>
        /// The key of the service principal used to authenticate against the endpoint of a published Azure ML Service pipeline.
        /// </summary>
        [Input("servicePrincipalKey")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? ServicePrincipalKey { get; set; }

        [Input("subscriptionId", required: true)]
        private InputMap<object>? _subscriptionId;

        /// <summary>
        /// Azure ML Service workspace subscription ID. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> SubscriptionId
        {
            get => _subscriptionId ?? (_subscriptionId = new InputMap<object>());
            set => _subscriptionId = value;
        }

        [Input("tenant")]
        private InputMap<object>? _tenant;

        /// <summary>
        /// The name or ID of the tenant to which the service principal belongs. Type: string (or Expression with resultType string).
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

        public AzureMLServiceLinkedServiceArgs()
        {
        }
    }
}
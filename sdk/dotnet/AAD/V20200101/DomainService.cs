// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Aad.V20200101
{
    /// <summary>
    /// Domain service.
    /// </summary>
    public partial class DomainService : Pulumi.CustomResource
    {
        /// <summary>
        /// Deployment Id
        /// </summary>
        [Output("deploymentId")]
        public Output<string> DeploymentId { get; private set; } = null!;

        /// <summary>
        /// The name of the Azure domain that the user would like to deploy Domain Services to.
        /// </summary>
        [Output("domainName")]
        public Output<string?> DomainName { get; private set; } = null!;

        /// <summary>
        /// DomainSecurity Settings
        /// </summary>
        [Output("domainSecuritySettings")]
        public Output<Outputs.DomainSecuritySettingsResponseResult?> DomainSecuritySettings { get; private set; } = null!;

        /// <summary>
        /// Resource etag
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Enabled or Disabled flag to turn on Group-based filtered sync
        /// </summary>
        [Output("filteredSync")]
        public Output<string?> FilteredSync { get; private set; } = null!;

        /// <summary>
        /// Secure LDAP Settings
        /// </summary>
        [Output("ldapsSettings")]
        public Output<Outputs.LdapsSettingsResponseResult?> LdapsSettings { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Notification Settings
        /// </summary>
        [Output("notificationSettings")]
        public Output<Outputs.NotificationSettingsResponseResult?> NotificationSettings { get; private set; } = null!;

        /// <summary>
        /// the current deployment or provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// List of ReplicaSets
        /// </summary>
        [Output("replicaSets")]
        public Output<ImmutableArray<Outputs.ReplicaSetResponseResult>> ReplicaSets { get; private set; } = null!;

        /// <summary>
        /// SyncOwner ReplicaSet Id
        /// </summary>
        [Output("syncOwner")]
        public Output<string> SyncOwner { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Azure Active Directory Tenant Id
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Data Model Version
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a DomainService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainService(string name, DomainServiceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:aad/v20200101:DomainService", name, args ?? new DomainServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:aad/v20200101:DomainService", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:aad/latest:DomainService"},
                    new Pulumi.Alias { Type = "azurerm:aad/v20170101:DomainService"},
                    new Pulumi.Alias { Type = "azurerm:aad/v20170601:DomainService"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DomainService(name, id, options);
        }
    }

    public sealed class DomainServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Azure domain that the user would like to deploy Domain Services to.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// DomainSecurity Settings
        /// </summary>
        [Input("domainSecuritySettings")]
        public Input<Inputs.DomainSecuritySettingsArgs>? DomainSecuritySettings { get; set; }

        /// <summary>
        /// The name of the domain service.
        /// </summary>
        [Input("domainServiceName", required: true)]
        public Input<string> DomainServiceName { get; set; } = null!;

        /// <summary>
        /// Resource etag
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Enabled or Disabled flag to turn on Group-based filtered sync
        /// </summary>
        [Input("filteredSync")]
        public Input<string>? FilteredSync { get; set; }

        /// <summary>
        /// Secure LDAP Settings
        /// </summary>
        [Input("ldapsSettings")]
        public Input<Inputs.LdapsSettingsArgs>? LdapsSettings { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Notification Settings
        /// </summary>
        [Input("notificationSettings")]
        public Input<Inputs.NotificationSettingsArgs>? NotificationSettings { get; set; }

        [Input("replicaSets")]
        private InputList<Inputs.ReplicaSetArgs>? _replicaSets;

        /// <summary>
        /// List of ReplicaSets
        /// </summary>
        public InputList<Inputs.ReplicaSetArgs> ReplicaSets
        {
            get => _replicaSets ?? (_replicaSets = new InputList<Inputs.ReplicaSetArgs>());
            set => _replicaSets = value;
        }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DomainServiceArgs()
        {
        }
    }
}
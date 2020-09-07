// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20191212
{
    /// <summary>
    /// An Azure Cosmos DB database account.
    /// </summary>
    public partial class DatabaseAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// List of Cosmos DB capabilities for the account
        /// </summary>
        [Output("capabilities")]
        public Output<ImmutableArray<Outputs.CapabilityResponseResult>> Capabilities { get; private set; } = null!;

        /// <summary>
        /// The cassandra connector offer type for the Cosmos DB database C* account.
        /// </summary>
        [Output("connectorOffer")]
        public Output<string?> ConnectorOffer { get; private set; } = null!;

        /// <summary>
        /// The consistency policy for the Cosmos DB database account.
        /// </summary>
        [Output("consistencyPolicy")]
        public Output<Outputs.ConsistencyPolicyResponseResult?> ConsistencyPolicy { get; private set; } = null!;

        /// <summary>
        /// The offer type for the Cosmos DB database account. Default value: Standard.
        /// </summary>
        [Output("databaseAccountOfferType")]
        public Output<string> DatabaseAccountOfferType { get; private set; } = null!;

        /// <summary>
        /// Disable write operations on metadata resources (databases, containers, throughput) via account keys
        /// </summary>
        [Output("disableKeyBasedMetadataWriteAccess")]
        public Output<bool?> DisableKeyBasedMetadataWriteAccess { get; private set; } = null!;

        /// <summary>
        /// The connection endpoint for the Cosmos DB database account.
        /// </summary>
        [Output("documentEndpoint")]
        public Output<string> DocumentEndpoint { get; private set; } = null!;

        /// <summary>
        /// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
        /// </summary>
        [Output("enableAutomaticFailover")]
        public Output<bool?> EnableAutomaticFailover { get; private set; } = null!;

        /// <summary>
        /// Enables the cassandra connector on the Cosmos DB C* account
        /// </summary>
        [Output("enableCassandraConnector")]
        public Output<bool?> EnableCassandraConnector { get; private set; } = null!;

        /// <summary>
        /// Enables the account to write in multiple locations
        /// </summary>
        [Output("enableMultipleWriteLocations")]
        public Output<bool?> EnableMultipleWriteLocations { get; private set; } = null!;

        /// <summary>
        /// An array that contains the regions ordered by their failover priorities.
        /// </summary>
        [Output("failoverPolicies")]
        public Output<ImmutableArray<Outputs.FailoverPolicyResponseResult>> FailoverPolicies { get; private set; } = null!;

        /// <summary>
        /// Cosmos DB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IPs for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        /// </summary>
        [Output("ipRangeFilter")]
        public Output<string?> IpRangeFilter { get; private set; } = null!;

        /// <summary>
        /// Flag to indicate whether to enable/disable Virtual Network ACL rules.
        /// </summary>
        [Output("isVirtualNetworkFilterEnabled")]
        public Output<bool?> IsVirtualNetworkFilterEnabled { get; private set; } = null!;

        /// <summary>
        /// The URI of the key vault
        /// </summary>
        [Output("keyVaultKeyUri")]
        public Output<string?> KeyVaultKeyUri { get; private set; } = null!;

        /// <summary>
        /// Indicates the type of database account. This can only be set at database account creation.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// An array that contains all of the locations enabled for the Cosmos DB account.
        /// </summary>
        [Output("locations")]
        public Output<ImmutableArray<Outputs.LocationResponseResult>> Locations { get; private set; } = null!;

        /// <summary>
        /// The name of the ARM resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'Offline' - the Cosmos DB account is not active. 'DeletionFailed' – the Cosmos DB account deletion failed.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// An array that contains of the read locations enabled for the Cosmos DB account.
        /// </summary>
        [Output("readLocations")]
        public Output<ImmutableArray<Outputs.LocationResponseResult>> ReadLocations { get; private set; } = null!;

        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// List of Virtual Network ACL rules configured for the Cosmos DB account.
        /// </summary>
        [Output("virtualNetworkRules")]
        public Output<ImmutableArray<Outputs.VirtualNetworkRuleResponseResult>> VirtualNetworkRules { get; private set; } = null!;

        /// <summary>
        /// An array that contains the write location for the Cosmos DB account.
        /// </summary>
        [Output("writeLocations")]
        public Output<ImmutableArray<Outputs.LocationResponseResult>> WriteLocations { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseAccount(string name, DatabaseAccountArgs args, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/v20191212:DatabaseAccount", name, args ?? new DatabaseAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/v20191212:DatabaseAccount", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:documentdb/latest:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20150401:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20150408:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20151106:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20160319:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20160331:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20190801:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20200301:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20200401:DatabaseAccount"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20200601preview:DatabaseAccount"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabaseAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatabaseAccount(name, id, options);
        }
    }

    public sealed class DatabaseAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("capabilities")]
        private InputList<Inputs.CapabilityArgs>? _capabilities;

        /// <summary>
        /// List of Cosmos DB capabilities for the account
        /// </summary>
        public InputList<Inputs.CapabilityArgs> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<Inputs.CapabilityArgs>());
            set => _capabilities = value;
        }

        /// <summary>
        /// The cassandra connector offer type for the Cosmos DB database C* account.
        /// </summary>
        [Input("connectorOffer")]
        public Input<string>? ConnectorOffer { get; set; }

        /// <summary>
        /// The consistency policy for the Cosmos DB account.
        /// </summary>
        [Input("consistencyPolicy")]
        public Input<Inputs.ConsistencyPolicyArgs>? ConsistencyPolicy { get; set; }

        /// <summary>
        /// The offer type for the database
        /// </summary>
        [Input("databaseAccountOfferType", required: true)]
        public Input<string> DatabaseAccountOfferType { get; set; } = null!;

        /// <summary>
        /// Disable write operations on metadata resources (databases, containers, throughput) via account keys
        /// </summary>
        [Input("disableKeyBasedMetadataWriteAccess")]
        public Input<bool>? DisableKeyBasedMetadataWriteAccess { get; set; }

        /// <summary>
        /// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
        /// </summary>
        [Input("enableAutomaticFailover")]
        public Input<bool>? EnableAutomaticFailover { get; set; }

        /// <summary>
        /// Enables the cassandra connector on the Cosmos DB C* account
        /// </summary>
        [Input("enableCassandraConnector")]
        public Input<bool>? EnableCassandraConnector { get; set; }

        /// <summary>
        /// Enables the account to write in multiple locations
        /// </summary>
        [Input("enableMultipleWriteLocations")]
        public Input<bool>? EnableMultipleWriteLocations { get; set; }

        /// <summary>
        /// Cosmos DB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IPs for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        /// </summary>
        [Input("ipRangeFilter")]
        public Input<string>? IpRangeFilter { get; set; }

        /// <summary>
        /// Flag to indicate whether to enable/disable Virtual Network ACL rules.
        /// </summary>
        [Input("isVirtualNetworkFilterEnabled")]
        public Input<bool>? IsVirtualNetworkFilterEnabled { get; set; }

        /// <summary>
        /// The URI of the key vault
        /// </summary>
        [Input("keyVaultKeyUri")]
        public Input<string>? KeyVaultKeyUri { get; set; }

        /// <summary>
        /// Indicates the type of database account. This can only be set at database account creation.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("locations", required: true)]
        private InputList<Inputs.LocationArgs>? _locations;

        /// <summary>
        /// An array that contains the georeplication locations enabled for the Cosmos DB account.
        /// </summary>
        public InputList<Inputs.LocationArgs> Locations
        {
            get => _locations ?? (_locations = new InputList<Inputs.LocationArgs>());
            set => _locations = value;
        }

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("virtualNetworkRules")]
        private InputList<Inputs.VirtualNetworkRuleArgs>? _virtualNetworkRules;

        /// <summary>
        /// List of Virtual Network ACL rules configured for the Cosmos DB account.
        /// </summary>
        public InputList<Inputs.VirtualNetworkRuleArgs> VirtualNetworkRules
        {
            get => _virtualNetworkRules ?? (_virtualNetworkRules = new InputList<Inputs.VirtualNetworkRuleArgs>());
            set => _virtualNetworkRules = value;
        }

        public DatabaseAccountArgs()
        {
        }
    }
}
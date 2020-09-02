// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview
{
    public static class GetDatabaseAccount
    {
        public static Task<GetDatabaseAccountResult> InvokeAsync(GetDatabaseAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseAccountResult>("azurerm:documentdb/v20200601preview:getDatabaseAccount", args ?? new GetDatabaseAccountArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDatabaseAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseAccountResult
    {
        /// <summary>
        /// API specific properties.
        /// </summary>
        public readonly Outputs.ApiPropertiesResponseResult? ApiProperties;
        /// <summary>
        /// The object representing the policy for taking backups on an account.
        /// </summary>
        public readonly Outputs.BackupPolicyResponseResult? BackupPolicy;
        /// <summary>
        /// List of Cosmos DB capabilities for the account
        /// </summary>
        public readonly ImmutableArray<Outputs.CapabilityResponseResult> Capabilities;
        /// <summary>
        /// The cassandra connector offer type for the Cosmos DB database C* account.
        /// </summary>
        public readonly string? ConnectorOffer;
        /// <summary>
        /// The consistency policy for the Cosmos DB database account.
        /// </summary>
        public readonly Outputs.ConsistencyPolicyResponseResult? ConsistencyPolicy;
        /// <summary>
        /// The CORS policy for the Cosmos DB database account.
        /// </summary>
        public readonly ImmutableArray<Outputs.CorsPolicyResponseResult> Cors;
        /// <summary>
        /// Enum to indicate the mode of account creation.
        /// </summary>
        public readonly string? CreateMode;
        /// <summary>
        /// The offer type for the Cosmos DB database account. Default value: Standard.
        /// </summary>
        public readonly string DatabaseAccountOfferType;
        /// <summary>
        /// Disable write operations on metadata resources (databases, containers, throughput) via account keys
        /// </summary>
        public readonly bool? DisableKeyBasedMetadataWriteAccess;
        /// <summary>
        /// The connection endpoint for the Cosmos DB database account.
        /// </summary>
        public readonly string DocumentEndpoint;
        /// <summary>
        /// Flag to indicate whether to enable storage analytics.
        /// </summary>
        public readonly bool? EnableAnalyticalStorage;
        /// <summary>
        /// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
        /// </summary>
        public readonly bool? EnableAutomaticFailover;
        /// <summary>
        /// Enables the cassandra connector on the Cosmos DB C* account
        /// </summary>
        public readonly bool? EnableCassandraConnector;
        /// <summary>
        /// Flag to indicate whether Free Tier is enabled.
        /// </summary>
        public readonly bool? EnableFreeTier;
        /// <summary>
        /// Enables the account to write in multiple locations
        /// </summary>
        public readonly bool? EnableMultipleWriteLocations;
        /// <summary>
        /// An array that contains the regions ordered by their failover priorities.
        /// </summary>
        public readonly ImmutableArray<Outputs.FailoverPolicyResponseResult> FailoverPolicies;
        /// <summary>
        /// Identity for the resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponseResult? Identity;
        /// <summary>
        /// A unique identifier assigned to the database account
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// List of IpRules.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpAddressOrRangeResponseResult> IpRules;
        /// <summary>
        /// Flag to indicate whether to enable/disable Virtual Network ACL rules.
        /// </summary>
        public readonly bool? IsVirtualNetworkFilterEnabled;
        /// <summary>
        /// The URI of the key vault
        /// </summary>
        public readonly string? KeyVaultKeyUri;
        /// <summary>
        /// Indicates the type of database account. This can only be set at database account creation.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// An array that contains all of the locations enabled for the Cosmos DB account.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationResponseResult> Locations;
        /// <summary>
        /// The name of the ARM resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of Private Endpoint Connections configured for the Cosmos DB account.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult> PrivateEndpointConnections;
        /// <summary>
        /// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Whether requests from Public Network are allowed
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// An array that contains of the read locations enabled for the Cosmos DB account.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationResponseResult> ReadLocations;
        /// <summary>
        /// Parameters to indicate the information about the restore.
        /// </summary>
        public readonly Outputs.RestoreParametersResponseResult? RestoreParameters;
        /// <summary>
        /// The system meta data relating to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponseResult SystemData;
        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// List of Virtual Network ACL rules configured for the Cosmos DB account.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualNetworkRuleResponseResult> VirtualNetworkRules;
        /// <summary>
        /// An array that contains the write location for the Cosmos DB account.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationResponseResult> WriteLocations;

        [OutputConstructor]
        private GetDatabaseAccountResult(
            Outputs.ApiPropertiesResponseResult? apiProperties,

            Outputs.BackupPolicyResponseResult? backupPolicy,

            ImmutableArray<Outputs.CapabilityResponseResult> capabilities,

            string? connectorOffer,

            Outputs.ConsistencyPolicyResponseResult? consistencyPolicy,

            ImmutableArray<Outputs.CorsPolicyResponseResult> cors,

            string? createMode,

            string databaseAccountOfferType,

            bool? disableKeyBasedMetadataWriteAccess,

            string documentEndpoint,

            bool? enableAnalyticalStorage,

            bool? enableAutomaticFailover,

            bool? enableCassandraConnector,

            bool? enableFreeTier,

            bool? enableMultipleWriteLocations,

            ImmutableArray<Outputs.FailoverPolicyResponseResult> failoverPolicies,

            Outputs.ManagedServiceIdentityResponseResult? identity,

            string instanceId,

            ImmutableArray<Outputs.IpAddressOrRangeResponseResult> ipRules,

            bool? isVirtualNetworkFilterEnabled,

            string? keyVaultKeyUri,

            string? kind,

            string? location,

            ImmutableArray<Outputs.LocationResponseResult> locations,

            string name,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult> privateEndpointConnections,

            string? provisioningState,

            string? publicNetworkAccess,

            ImmutableArray<Outputs.LocationResponseResult> readLocations,

            Outputs.RestoreParametersResponseResult? restoreParameters,

            Outputs.SystemDataResponseResult systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.VirtualNetworkRuleResponseResult> virtualNetworkRules,

            ImmutableArray<Outputs.LocationResponseResult> writeLocations)
        {
            ApiProperties = apiProperties;
            BackupPolicy = backupPolicy;
            Capabilities = capabilities;
            ConnectorOffer = connectorOffer;
            ConsistencyPolicy = consistencyPolicy;
            Cors = cors;
            CreateMode = createMode;
            DatabaseAccountOfferType = databaseAccountOfferType;
            DisableKeyBasedMetadataWriteAccess = disableKeyBasedMetadataWriteAccess;
            DocumentEndpoint = documentEndpoint;
            EnableAnalyticalStorage = enableAnalyticalStorage;
            EnableAutomaticFailover = enableAutomaticFailover;
            EnableCassandraConnector = enableCassandraConnector;
            EnableFreeTier = enableFreeTier;
            EnableMultipleWriteLocations = enableMultipleWriteLocations;
            FailoverPolicies = failoverPolicies;
            Identity = identity;
            InstanceId = instanceId;
            IpRules = ipRules;
            IsVirtualNetworkFilterEnabled = isVirtualNetworkFilterEnabled;
            KeyVaultKeyUri = keyVaultKeyUri;
            Kind = kind;
            Location = location;
            Locations = locations;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
            ReadLocations = readLocations;
            RestoreParameters = restoreParameters;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            VirtualNetworkRules = virtualNetworkRules;
            WriteLocations = writeLocations;
        }
    }
}
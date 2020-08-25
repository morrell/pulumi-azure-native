// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache.V20150801
{
    public static class GetRedis
    {
        public static Task<GetRedisResult> InvokeAsync(GetRedisArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRedisResult>("azurerm:cache/v20150801:getRedis", args ?? new GetRedisArgs(), options.WithVersion());
    }


    public sealed class GetRedisArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRedisArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRedisResult
    {
        /// <summary>
        /// If the value is true, then the non-SLL Redis server port (6379) will be enabled.
        /// </summary>
        public readonly bool? EnableNonSslPort;
        /// <summary>
        /// Redis host name.
        /// </summary>
        public readonly string? HostName;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Redis non-SSL port.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Redis instance provisioning status.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? RedisConfiguration;
        /// <summary>
        /// RedisVersion parameter has been deprecated. As such, it is no longer necessary to provide this parameter and any value specified is ignored.
        /// </summary>
        public readonly string? RedisVersion;
        /// <summary>
        /// The number of shards to be created on a Premium Cluster Cache.
        /// </summary>
        public readonly int? ShardCount;
        /// <summary>
        /// What SKU of Redis cache to deploy.
        /// </summary>
        public readonly Outputs.SkuResponseResult Sku;
        /// <summary>
        /// Redis SSL port.
        /// </summary>
        public readonly int? SslPort;
        /// <summary>
        /// Required when deploying a Redis cache inside an existing Azure Virtual Network.
        /// </summary>
        public readonly string? StaticIP;
        /// <summary>
        /// Required when deploying a Redis cache inside an existing Azure Virtual Network.
        /// </summary>
        public readonly string? Subnet;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// tenantSettings
        /// </summary>
        public readonly ImmutableDictionary<string, string>? TenantSettings;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The exact ARM resource ID of the virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.ClassicNetwork/VirtualNetworks/vnet1
        /// </summary>
        public readonly string? VirtualNetwork;

        [OutputConstructor]
        private GetRedisResult(
            bool? enableNonSslPort,

            string? hostName,

            string location,

            string name,

            int? port,

            string? provisioningState,

            ImmutableDictionary<string, string>? redisConfiguration,

            string? redisVersion,

            int? shardCount,

            Outputs.SkuResponseResult sku,

            int? sslPort,

            string? staticIP,

            string? subnet,

            ImmutableDictionary<string, string>? tags,

            ImmutableDictionary<string, string>? tenantSettings,

            string type,

            string? virtualNetwork)
        {
            EnableNonSslPort = enableNonSslPort;
            HostName = hostName;
            Location = location;
            Name = name;
            Port = port;
            ProvisioningState = provisioningState;
            RedisConfiguration = redisConfiguration;
            RedisVersion = redisVersion;
            ShardCount = shardCount;
            Sku = sku;
            SslPort = sslPort;
            StaticIP = staticIP;
            Subnet = subnet;
            Tags = tags;
            TenantSettings = tenantSettings;
            Type = type;
            VirtualNetwork = virtualNetwork;
        }
    }
}
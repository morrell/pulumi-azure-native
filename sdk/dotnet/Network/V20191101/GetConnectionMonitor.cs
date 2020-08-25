// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101
{
    public static class GetConnectionMonitor
    {
        public static Task<GetConnectionMonitorResult> InvokeAsync(GetConnectionMonitorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionMonitorResult>("azurerm:network/v20191101:getConnectionMonitor", args ?? new GetConnectionMonitorArgs(), options.WithVersion());
    }


    public sealed class GetConnectionMonitorArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the connection monitor.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Network Watcher resource.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public string NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing Network Watcher.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConnectionMonitorArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionMonitorResult
    {
        /// <summary>
        /// Determines if the connection monitor will start automatically once created.
        /// </summary>
        public readonly bool? AutoStart;
        /// <summary>
        /// Type of connection monitor.
        /// </summary>
        public readonly string ConnectionMonitorType;
        /// <summary>
        /// Describes the destination of connection monitor.
        /// </summary>
        public readonly Outputs.ConnectionMonitorDestinationResponseResult? Destination;
        /// <summary>
        /// List of connection monitor endpoints.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionMonitorEndpointResponseResult> Endpoints;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Connection monitor location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Monitoring interval in seconds.
        /// </summary>
        public readonly int? MonitoringIntervalInSeconds;
        /// <summary>
        /// The monitoring status of the connection monitor.
        /// </summary>
        public readonly string MonitoringStatus;
        /// <summary>
        /// Name of the connection monitor.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional notes to be associated with the connection monitor.
        /// </summary>
        public readonly string? Notes;
        /// <summary>
        /// List of connection monitor outputs.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionMonitorOutputResponseResult> Outputs;
        /// <summary>
        /// The provisioning state of the connection monitor.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Describes the source of connection monitor.
        /// </summary>
        public readonly Outputs.ConnectionMonitorSourceResponseResult? Source;
        /// <summary>
        /// The date and time when the connection monitor was started.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Connection monitor tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// List of connection monitor test configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionMonitorTestConfigurationResponseResult> TestConfigurations;
        /// <summary>
        /// List of connection monitor test groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionMonitorTestGroupResponseResult> TestGroups;
        /// <summary>
        /// Connection monitor type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConnectionMonitorResult(
            bool? autoStart,

            string connectionMonitorType,

            Outputs.ConnectionMonitorDestinationResponseResult? destination,

            ImmutableArray<Outputs.ConnectionMonitorEndpointResponseResult> endpoints,

            string etag,

            string? location,

            int? monitoringIntervalInSeconds,

            string monitoringStatus,

            string name,

            string? notes,

            ImmutableArray<Outputs.ConnectionMonitorOutputResponseResult> outputs,

            string provisioningState,

            Outputs.ConnectionMonitorSourceResponseResult? source,

            string startTime,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<Outputs.ConnectionMonitorTestConfigurationResponseResult> testConfigurations,

            ImmutableArray<Outputs.ConnectionMonitorTestGroupResponseResult> testGroups,

            string type)
        {
            AutoStart = autoStart;
            ConnectionMonitorType = connectionMonitorType;
            Destination = destination;
            Endpoints = endpoints;
            Etag = etag;
            Location = location;
            MonitoringIntervalInSeconds = monitoringIntervalInSeconds;
            MonitoringStatus = monitoringStatus;
            Name = name;
            Notes = notes;
            Outputs = outputs;
            ProvisioningState = provisioningState;
            Source = source;
            StartTime = startTime;
            Tags = tags;
            TestConfigurations = testConfigurations;
            TestGroups = testGroups;
            Type = type;
        }
    }
}
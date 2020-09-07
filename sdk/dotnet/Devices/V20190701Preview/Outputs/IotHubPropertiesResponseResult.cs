// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20190701Preview.Outputs
{

    [OutputType]
    public sealed class IotHubPropertiesResponseResult
    {
        /// <summary>
        /// The shared access policies you can use to secure a connection to the IoT hub.
        /// </summary>
        public readonly ImmutableArray<Outputs.SharedAccessSignatureAuthorizationRuleResponseResult> AuthorizationPolicies;
        /// <summary>
        /// The IoT hub cloud-to-device messaging properties.
        /// </summary>
        public readonly Outputs.CloudToDevicePropertiesResponseResult? CloudToDevice;
        /// <summary>
        /// IoT hub comments.
        /// </summary>
        public readonly string? Comments;
        /// <summary>
        /// The device streams properties of iothub.
        /// </summary>
        public readonly Outputs.IotHubPropertiesResponseDeviceStreamsResult? DeviceStreams;
        /// <summary>
        /// If True, file upload notifications are enabled.
        /// </summary>
        public readonly bool? EnableFileUploadNotifications;
        /// <summary>
        /// The Event Hub-compatible endpoint properties. The only possible keys to this dictionary is events. This key has to be present in the dictionary while making create or update calls for the IoT hub.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.EventHubPropertiesResponseResult>? EventHubEndpoints;
        /// <summary>
        /// The capabilities and features enabled for the IoT hub.
        /// </summary>
        public readonly string? Features;
        /// <summary>
        /// The name of the host.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The IP filter rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpFilterRuleResponseResult> IpFilterRules;
        /// <summary>
        /// Primary and secondary location for iot hub
        /// </summary>
        public readonly ImmutableArray<Outputs.IotHubLocationDescriptionResponseResult> Locations;
        /// <summary>
        /// The messaging endpoint properties for the file upload notification queue.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.MessagingEndpointPropertiesResponseResult>? MessagingEndpoints;
        /// <summary>
        /// The provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The routing related properties of the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging
        /// </summary>
        public readonly Outputs.RoutingPropertiesResponseResult? Routing;
        /// <summary>
        /// The hub state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The list of Azure Storage endpoints where you can upload files. Currently you can configure only one Azure Storage account and that MUST have its key as $default. Specifying more than one storage account causes an error to be thrown. Not specifying a value for this property when the enableFileUploadNotifications property is set to True, causes an error to be thrown.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.StorageEndpointPropertiesResponseResult>? StorageEndpoints;

        [OutputConstructor]
        private IotHubPropertiesResponseResult(
            ImmutableArray<Outputs.SharedAccessSignatureAuthorizationRuleResponseResult> authorizationPolicies,

            Outputs.CloudToDevicePropertiesResponseResult? cloudToDevice,

            string? comments,

            Outputs.IotHubPropertiesResponseDeviceStreamsResult? deviceStreams,

            bool? enableFileUploadNotifications,

            ImmutableDictionary<string, Outputs.EventHubPropertiesResponseResult>? eventHubEndpoints,

            string? features,

            string hostName,

            ImmutableArray<Outputs.IpFilterRuleResponseResult> ipFilterRules,

            ImmutableArray<Outputs.IotHubLocationDescriptionResponseResult> locations,

            ImmutableDictionary<string, Outputs.MessagingEndpointPropertiesResponseResult>? messagingEndpoints,

            string provisioningState,

            Outputs.RoutingPropertiesResponseResult? routing,

            string state,

            ImmutableDictionary<string, Outputs.StorageEndpointPropertiesResponseResult>? storageEndpoints)
        {
            AuthorizationPolicies = authorizationPolicies;
            CloudToDevice = cloudToDevice;
            Comments = comments;
            DeviceStreams = deviceStreams;
            EnableFileUploadNotifications = enableFileUploadNotifications;
            EventHubEndpoints = eventHubEndpoints;
            Features = features;
            HostName = hostName;
            IpFilterRules = ipFilterRules;
            Locations = locations;
            MessagingEndpoints = messagingEndpoints;
            ProvisioningState = provisioningState;
            Routing = routing;
            State = state;
            StorageEndpoints = storageEndpoints;
        }
    }
}
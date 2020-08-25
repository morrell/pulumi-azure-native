// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200101.Outputs
{

    [OutputType]
    public sealed class IotDpsPropertiesDescriptionResponseResult
    {
        /// <summary>
        /// Allocation policy to be used by this provisioning service.
        /// </summary>
        public readonly string? AllocationPolicy;
        /// <summary>
        /// List of authorization keys for a provisioning service.
        /// </summary>
        public readonly ImmutableArray<Outputs.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseResult> AuthorizationPolicies;
        /// <summary>
        /// Device endpoint for this provisioning service.
        /// </summary>
        public readonly string DeviceProvisioningHostName;
        /// <summary>
        /// Unique identifier of this provisioning service.
        /// </summary>
        public readonly string IdScope;
        /// <summary>
        /// List of IoT hubs associated with this provisioning service.
        /// </summary>
        public readonly ImmutableArray<Outputs.IotHubDefinitionDescriptionResponseResult> IotHubs;
        /// <summary>
        /// The IP filter rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpFilterRuleResponseResult> IpFilterRules;
        /// <summary>
        /// The ARM provisioning state of the provisioning service.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Service endpoint for provisioning service.
        /// </summary>
        public readonly string ServiceOperationsHostName;
        /// <summary>
        /// Current state of the provisioning service.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private IotDpsPropertiesDescriptionResponseResult(
            string? allocationPolicy,

            ImmutableArray<Outputs.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseResult> authorizationPolicies,

            string deviceProvisioningHostName,

            string idScope,

            ImmutableArray<Outputs.IotHubDefinitionDescriptionResponseResult> iotHubs,

            ImmutableArray<Outputs.IpFilterRuleResponseResult> ipFilterRules,

            string? provisioningState,

            string serviceOperationsHostName,

            string? state)
        {
            AllocationPolicy = allocationPolicy;
            AuthorizationPolicies = authorizationPolicies;
            DeviceProvisioningHostName = deviceProvisioningHostName;
            IdScope = idScope;
            IotHubs = iotHubs;
            IpFilterRules = ipFilterRules;
            ProvisioningState = provisioningState;
            ServiceOperationsHostName = serviceOperationsHostName;
            State = state;
        }
    }
}
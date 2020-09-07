// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ManagedNetwork.V20190601Preview.Outputs
{

    [OutputType]
    public sealed class ManagedNetworkGroupResponseResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Responsibility role under which this Managed Network Group will be created
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The collection of management groups covered by the Managed Network
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponseResult> ManagementGroups;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the ManagedNetwork resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The collection of  subnets covered by the Managed Network
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponseResult> Subnets;
        /// <summary>
        /// The collection of subscriptions covered by the Managed Network
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponseResult> Subscriptions;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The collection of virtual nets covered by the Managed Network
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponseResult> VirtualNetworks;

        [OutputConstructor]
        private ManagedNetworkGroupResponseResult(
            string etag,

            string id,

            string? kind,

            string? location,

            ImmutableArray<Outputs.ResourceIdResponseResult> managementGroups,

            string name,

            string provisioningState,

            ImmutableArray<Outputs.ResourceIdResponseResult> subnets,

            ImmutableArray<Outputs.ResourceIdResponseResult> subscriptions,

            string type,

            ImmutableArray<Outputs.ResourceIdResponseResult> virtualNetworks)
        {
            Etag = etag;
            Id = id;
            Kind = kind;
            Location = location;
            ManagementGroups = managementGroups;
            Name = name;
            ProvisioningState = provisioningState;
            Subnets = subnets;
            Subscriptions = subscriptions;
            Type = type;
            VirtualNetworks = virtualNetworks;
        }
    }
}
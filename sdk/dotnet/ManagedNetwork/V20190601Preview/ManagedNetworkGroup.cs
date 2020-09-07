// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ManagedNetwork.V20190601Preview
{
    /// <summary>
    /// The Managed Network Group resource
    /// </summary>
    public partial class ManagedNetworkGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Responsibility role under which this Managed Network Group will be created
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The collection of management groups covered by the Managed Network
        /// </summary>
        [Output("managementGroups")]
        public Output<ImmutableArray<Outputs.ResourceIdResponseResult>> ManagementGroups { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the ManagedNetwork resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The collection of  subnets covered by the Managed Network
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<Outputs.ResourceIdResponseResult>> Subnets { get; private set; } = null!;

        /// <summary>
        /// The collection of subscriptions covered by the Managed Network
        /// </summary>
        [Output("subscriptions")]
        public Output<ImmutableArray<Outputs.ResourceIdResponseResult>> Subscriptions { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The collection of virtual nets covered by the Managed Network
        /// </summary>
        [Output("virtualNetworks")]
        public Output<ImmutableArray<Outputs.ResourceIdResponseResult>> VirtualNetworks { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedNetworkGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedNetworkGroup(string name, ManagedNetworkGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:managednetwork/v20190601preview:ManagedNetworkGroup", name, args ?? new ManagedNetworkGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedNetworkGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:managednetwork/v20190601preview:ManagedNetworkGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedNetworkGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedNetworkGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedNetworkGroup(name, id, options);
        }
    }

    public sealed class ManagedNetworkGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Responsibility role under which this Managed Network Group will be created
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Managed Network Group.
        /// </summary>
        [Input("managedNetworkGroupName", required: true)]
        public Input<string> ManagedNetworkGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Managed Network.
        /// </summary>
        [Input("managedNetworkName", required: true)]
        public Input<string> ManagedNetworkName { get; set; } = null!;

        [Input("managementGroups")]
        private InputList<Inputs.ResourceIdArgs>? _managementGroups;

        /// <summary>
        /// The collection of management groups covered by the Managed Network
        /// </summary>
        public InputList<Inputs.ResourceIdArgs> ManagementGroups
        {
            get => _managementGroups ?? (_managementGroups = new InputList<Inputs.ResourceIdArgs>());
            set => _managementGroups = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("subnets")]
        private InputList<Inputs.ResourceIdArgs>? _subnets;

        /// <summary>
        /// The collection of  subnets covered by the Managed Network
        /// </summary>
        public InputList<Inputs.ResourceIdArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.ResourceIdArgs>());
            set => _subnets = value;
        }

        [Input("subscriptions")]
        private InputList<Inputs.ResourceIdArgs>? _subscriptions;

        /// <summary>
        /// The collection of subscriptions covered by the Managed Network
        /// </summary>
        public InputList<Inputs.ResourceIdArgs> Subscriptions
        {
            get => _subscriptions ?? (_subscriptions = new InputList<Inputs.ResourceIdArgs>());
            set => _subscriptions = value;
        }

        [Input("virtualNetworks")]
        private InputList<Inputs.ResourceIdArgs>? _virtualNetworks;

        /// <summary>
        /// The collection of virtual nets covered by the Managed Network
        /// </summary>
        public InputList<Inputs.ResourceIdArgs> VirtualNetworks
        {
            get => _virtualNetworks ?? (_virtualNetworks = new InputList<Inputs.ResourceIdArgs>());
            set => _virtualNetworks = value;
        }

        public ManagedNetworkGroupArgs()
        {
        }
    }
}
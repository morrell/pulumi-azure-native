// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401
{
    /// <summary>
    /// A DDoS protection plan in a resource group.
    /// </summary>
    public partial class DdosProtectionPlan : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the DDoS protection plan resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string> ResourceGuid { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
        /// </summary>
        [Output("virtualNetworks")]
        public Output<ImmutableArray<Outputs.SubResourceResponseResult>> VirtualNetworks { get; private set; } = null!;


        /// <summary>
        /// Create a DdosProtectionPlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DdosProtectionPlan(string name, DdosProtectionPlanArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200401:DdosProtectionPlan", name, args ?? new DdosProtectionPlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DdosProtectionPlan(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200401:DdosProtectionPlan", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180201:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180401:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180601:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180701:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:DdosProtectionPlan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:DdosProtectionPlan"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DdosProtectionPlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DdosProtectionPlan Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DdosProtectionPlan(name, id, options);
        }
    }

    public sealed class DdosProtectionPlanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DDoS protection plan.
        /// </summary>
        [Input("ddosProtectionPlanName", required: true)]
        public Input<string> DdosProtectionPlanName { get; set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DdosProtectionPlanArgs()
        {
        }
    }
}
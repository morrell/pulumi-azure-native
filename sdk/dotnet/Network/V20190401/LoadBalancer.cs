// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190401
{
    /// <summary>
    /// LoadBalancer resource.
    /// </summary>
    public partial class LoadBalancer : Pulumi.CustomResource
    {
        /// <summary>
        /// Collection of backend address pools used by a load balancer.
        /// </summary>
        [Output("backendAddressPools")]
        public Output<ImmutableArray<Outputs.BackendAddressPoolResponseResult>> BackendAddressPools { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Object representing the frontend IPs to be used for the load balancer.
        /// </summary>
        [Output("frontendIPConfigurations")]
        public Output<ImmutableArray<Outputs.FrontendIPConfigurationResponseResult>> FrontendIPConfigurations { get; private set; } = null!;

        /// <summary>
        /// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
        /// </summary>
        [Output("inboundNatPools")]
        public Output<ImmutableArray<Outputs.InboundNatPoolResponseResult>> InboundNatPools { get; private set; } = null!;

        /// <summary>
        /// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
        /// </summary>
        [Output("inboundNatRules")]
        public Output<ImmutableArray<Outputs.InboundNatRuleResponseResult>> InboundNatRules { get; private set; } = null!;

        /// <summary>
        /// Object collection representing the load balancing rules Gets the provisioning.
        /// </summary>
        [Output("loadBalancingRules")]
        public Output<ImmutableArray<Outputs.LoadBalancingRuleResponseResult>> LoadBalancingRules { get; private set; } = null!;

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
        /// The outbound rules.
        /// </summary>
        [Output("outboundRules")]
        public Output<ImmutableArray<Outputs.OutboundRuleResponseResult>> OutboundRules { get; private set; } = null!;

        /// <summary>
        /// Collection of probe objects used in the load balancer.
        /// </summary>
        [Output("probes")]
        public Output<ImmutableArray<Outputs.ProbeResponseResult>> Probes { get; private set; } = null!;

        /// <summary>
        /// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource GUID property of the load balancer resource.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string?> ResourceGuid { get; private set; } = null!;

        /// <summary>
        /// The load balancer SKU.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.LoadBalancerSkuResponseResult?> Sku { get; private set; } = null!;

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
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:LoadBalancer", name, args ?? new LoadBalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:LoadBalancer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Alias { Type = "azurerm:network/v20150615:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20160330:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20160601:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20160901:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20161201:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20170301:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20170601:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20170801:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20170901:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20171001:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20171101:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20180101:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20180201:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20180401:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20180601:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20180701:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20180801:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20181001:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20181101:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20181201:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20190201:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20190601:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20190701:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20190801:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20190901:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20191101:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20191201:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20200301:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20200401:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20200501:LoadBalancer"},
                    new Alias { Type = "azurerm:network/v20200601:LoadBalancer"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, options);
        }
    }

    public sealed class LoadBalancerArgs : Pulumi.ResourceArgs
    {
        [Input("backendAddressPools")]
        private InputList<Inputs.BackendAddressPoolArgs>? _backendAddressPools;

        /// <summary>
        /// Collection of backend address pools used by a load balancer.
        /// </summary>
        public InputList<Inputs.BackendAddressPoolArgs> BackendAddressPools
        {
            get => _backendAddressPools ?? (_backendAddressPools = new InputList<Inputs.BackendAddressPoolArgs>());
            set => _backendAddressPools = value;
        }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("frontendIPConfigurations")]
        private InputList<Inputs.FrontendIPConfigurationArgs>? _frontendIPConfigurations;

        /// <summary>
        /// Object representing the frontend IPs to be used for the load balancer.
        /// </summary>
        public InputList<Inputs.FrontendIPConfigurationArgs> FrontendIPConfigurations
        {
            get => _frontendIPConfigurations ?? (_frontendIPConfigurations = new InputList<Inputs.FrontendIPConfigurationArgs>());
            set => _frontendIPConfigurations = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("inboundNatPools")]
        private InputList<Inputs.InboundNatPoolArgs>? _inboundNatPools;

        /// <summary>
        /// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
        /// </summary>
        public InputList<Inputs.InboundNatPoolArgs> InboundNatPools
        {
            get => _inboundNatPools ?? (_inboundNatPools = new InputList<Inputs.InboundNatPoolArgs>());
            set => _inboundNatPools = value;
        }

        [Input("inboundNatRules")]
        private InputList<Inputs.InboundNatRuleArgs>? _inboundNatRules;

        /// <summary>
        /// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
        /// </summary>
        public InputList<Inputs.InboundNatRuleArgs> InboundNatRules
        {
            get => _inboundNatRules ?? (_inboundNatRules = new InputList<Inputs.InboundNatRuleArgs>());
            set => _inboundNatRules = value;
        }

        [Input("loadBalancingRules")]
        private InputList<Inputs.LoadBalancingRuleArgs>? _loadBalancingRules;

        /// <summary>
        /// Object collection representing the load balancing rules Gets the provisioning.
        /// </summary>
        public InputList<Inputs.LoadBalancingRuleArgs> LoadBalancingRules
        {
            get => _loadBalancingRules ?? (_loadBalancingRules = new InputList<Inputs.LoadBalancingRuleArgs>());
            set => _loadBalancingRules = value;
        }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the load balancer.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("outboundRules")]
        private InputList<Inputs.OutboundRuleArgs>? _outboundRules;

        /// <summary>
        /// The outbound rules.
        /// </summary>
        public InputList<Inputs.OutboundRuleArgs> OutboundRules
        {
            get => _outboundRules ?? (_outboundRules = new InputList<Inputs.OutboundRuleArgs>());
            set => _outboundRules = value;
        }

        [Input("probes")]
        private InputList<Inputs.ProbeArgs>? _probes;

        /// <summary>
        /// Collection of probe objects used in the load balancer.
        /// </summary>
        public InputList<Inputs.ProbeArgs> Probes
        {
            get => _probes ?? (_probes = new InputList<Inputs.ProbeArgs>());
            set => _probes = value;
        }

        /// <summary>
        /// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource GUID property of the load balancer resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        /// <summary>
        /// The load balancer SKU.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.LoadBalancerSkuArgs>? Sku { get; set; }

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

        public LoadBalancerArgs()
        {
        }
    }
}
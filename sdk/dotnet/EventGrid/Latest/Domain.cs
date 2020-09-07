// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.Latest
{
    /// <summary>
    /// EventGrid Domain.
    /// </summary>
    public partial class Domain : Pulumi.CustomResource
    {
        /// <summary>
        /// Endpoint for the domain.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
        /// </summary>
        [Output("inboundIpRules")]
        public Output<ImmutableArray<Outputs.InboundIpRuleResponseResult>> InboundIpRules { get; private set; } = null!;

        /// <summary>
        /// This determines the format that Event Grid should expect for incoming events published to the domain.
        /// </summary>
        [Output("inputSchema")]
        public Output<string?> InputSchema { get; private set; } = null!;

        /// <summary>
        /// Information about the InputSchemaMapping which specified the info about mapping event payload.
        /// </summary>
        [Output("inputSchemaMapping")]
        public Output<Outputs.JsonInputSchemaMappingResponseResult?> InputSchemaMapping { get; private set; } = null!;

        /// <summary>
        /// Location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Metric resource id for the domain.
        /// </summary>
        [Output("metricResourceId")]
        public Output<string> MetricResourceId { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of private endpoint connections.
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the domain.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// This determines if traffic is allowed over public network. By default it is enabled. 
        /// You can further restrict to specific IPs by configuring &lt;seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" /&gt;
        /// </summary>
        [Output("publicNetworkAccess")]
        public Output<string?> PublicNetworkAccess { get; private set; } = null!;

        /// <summary>
        /// Tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("azurerm:eventgrid/latest:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:eventgrid/latest:Domain", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:eventgrid/v20180915preview:Domain"},
                    new Pulumi.Alias { Type = "azurerm:eventgrid/v20190201preview:Domain"},
                    new Pulumi.Alias { Type = "azurerm:eventgrid/v20190601:Domain"},
                    new Pulumi.Alias { Type = "azurerm:eventgrid/v20200101preview:Domain"},
                    new Pulumi.Alias { Type = "azurerm:eventgrid/v20200401preview:Domain"},
                    new Pulumi.Alias { Type = "azurerm:eventgrid/v20200601:Domain"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, options);
        }
    }

    public sealed class DomainArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        [Input("inboundIpRules")]
        private InputList<Inputs.InboundIpRuleArgs>? _inboundIpRules;

        /// <summary>
        /// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
        /// </summary>
        public InputList<Inputs.InboundIpRuleArgs> InboundIpRules
        {
            get => _inboundIpRules ?? (_inboundIpRules = new InputList<Inputs.InboundIpRuleArgs>());
            set => _inboundIpRules = value;
        }

        /// <summary>
        /// This determines the format that Event Grid should expect for incoming events published to the domain.
        /// </summary>
        [Input("inputSchema")]
        public Input<string>? InputSchema { get; set; }

        /// <summary>
        /// Information about the InputSchemaMapping which specified the info about mapping event payload.
        /// </summary>
        [Input("inputSchemaMapping")]
        public Input<Inputs.JsonInputSchemaMappingArgs>? InputSchemaMapping { get; set; }

        /// <summary>
        /// Location of the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("privateEndpointConnections")]
        private InputList<Inputs.PrivateEndpointConnectionArgs>? _privateEndpointConnections;

        /// <summary>
        /// List of private endpoint connections.
        /// </summary>
        public InputList<Inputs.PrivateEndpointConnectionArgs> PrivateEndpointConnections
        {
            get => _privateEndpointConnections ?? (_privateEndpointConnections = new InputList<Inputs.PrivateEndpointConnectionArgs>());
            set => _privateEndpointConnections = value;
        }

        /// <summary>
        /// This determines if traffic is allowed over public network. By default it is enabled. 
        /// You can further restrict to specific IPs by configuring &lt;seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" /&gt;
        /// </summary>
        [Input("publicNetworkAccess")]
        public Input<string>? PublicNetworkAccess { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DomainArgs()
        {
        }
    }
}
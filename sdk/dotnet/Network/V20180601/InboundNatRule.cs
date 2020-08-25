// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180601
{
    /// <summary>
    /// Inbound NAT rule of the load balancer.
    /// </summary>
    public partial class InboundNatRule : Pulumi.CustomResource
    {
        /// <summary>
        /// A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
        /// </summary>
        [Output("backendIPConfiguration")]
        public Output<Outputs.NetworkInterfaceIPConfigurationResponseResult> BackendIPConfiguration { get; private set; } = null!;

        /// <summary>
        /// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
        /// </summary>
        [Output("backendPort")]
        public Output<int?> BackendPort { get; private set; } = null!;

        /// <summary>
        /// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
        /// </summary>
        [Output("enableFloatingIP")]
        public Output<bool?> EnableFloatingIP { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// A reference to frontend IP addresses.
        /// </summary>
        [Output("frontendIPConfiguration")]
        public Output<Outputs.SubResourceResponseResult?> FrontendIPConfiguration { get; private set; } = null!;

        /// <summary>
        /// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
        /// </summary>
        [Output("frontendPort")]
        public Output<int?> FrontendPort { get; private set; } = null!;

        /// <summary>
        /// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
        /// </summary>
        [Output("idleTimeoutInMinutes")]
        public Output<int?> IdleTimeoutInMinutes { get; private set; } = null!;

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The transport protocol for the endpoint. Possible values are 'Udp' or 'Tcp' or 'All'.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;


        /// <summary>
        /// Create a InboundNatRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InboundNatRule(string name, InboundNatRuleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180601:InboundNatRule", name, args ?? new InboundNatRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InboundNatRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180601:InboundNatRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Alias { Type = "azurerm:network/v20170601:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20170801:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20170901:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20171001:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20171101:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20180101:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20180201:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20180401:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20180701:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20180801:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20181001:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20181101:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20181201:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20190201:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20190401:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20190601:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20190701:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20190801:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20190901:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20191101:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20191201:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20200301:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20200401:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20200501:InboundNatRule"},
                    new Alias { Type = "azurerm:network/v20200601:InboundNatRule"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InboundNatRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InboundNatRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InboundNatRule(name, id, options);
        }
    }

    public sealed class InboundNatRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
        /// </summary>
        [Input("backendPort")]
        public Input<int>? BackendPort { get; set; }

        /// <summary>
        /// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
        /// </summary>
        [Input("enableFloatingIP")]
        public Input<bool>? EnableFloatingIP { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// A reference to frontend IP addresses.
        /// </summary>
        [Input("frontendIPConfiguration")]
        public Input<Inputs.SubResourceArgs>? FrontendIPConfiguration { get; set; }

        /// <summary>
        /// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
        /// </summary>
        [Input("frontendPort")]
        public Input<int>? FrontendPort { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        /// <summary>
        /// The name of the load balancer.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public Input<string> LoadBalancerName { get; set; } = null!;

        /// <summary>
        /// The name of the inbound nat rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The transport protocol for the endpoint. Possible values are 'Udp' or 'Tcp' or 'All'.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public InboundNatRuleArgs()
        {
        }
    }
}
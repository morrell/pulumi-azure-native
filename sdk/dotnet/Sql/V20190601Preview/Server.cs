// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20190601Preview
{
    /// <summary>
    /// An Azure SQL Database server.
    /// </summary>
    public partial class Server : Pulumi.CustomResource
    {
        /// <summary>
        /// Administrator username for the server. Once created it cannot be changed.
        /// </summary>
        [Output("administratorLogin")]
        public Output<string?> AdministratorLogin { get; private set; } = null!;

        /// <summary>
        /// The administrator login password (required for server creation).
        /// </summary>
        [Output("administratorLoginPassword")]
        public Output<string?> AdministratorLoginPassword { get; private set; } = null!;

        /// <summary>
        /// The fully qualified domain name of the server.
        /// </summary>
        [Output("fullyQualifiedDomainName")]
        public Output<string> FullyQualifiedDomainName { get; private set; } = null!;

        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ResourceIdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// Kind of sql server. This is metadata used for the Azure portal experience.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
        /// </summary>
        [Output("minimalTlsVersion")]
        public Output<string?> MinimalTlsVersion { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of private endpoint connections on a server
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponseResult>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        [Output("publicNetworkAccess")]
        public Output<string?> PublicNetworkAccess { get; private set; } = null!;

        /// <summary>
        /// The state of the server.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

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
        /// The version of the server.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:sql/v20190601preview:Server", name, args ?? new ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:sql/v20190601preview:Server", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:sql/latest:Server"},
                    new Pulumi.Alias { Type = "azurerm:sql/v20140401:Server"},
                    new Pulumi.Alias { Type = "azurerm:sql/v20150501preview:Server"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Server(name, id, options);
        }
    }

    public sealed class ServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrator username for the server. Once created it cannot be changed.
        /// </summary>
        [Input("administratorLogin")]
        public Input<string>? AdministratorLogin { get; set; }

        /// <summary>
        /// The administrator login password (required for server creation).
        /// </summary>
        [Input("administratorLoginPassword")]
        public Input<string>? AdministratorLoginPassword { get; set; }

        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ResourceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
        /// </summary>
        [Input("minimalTlsVersion")]
        public Input<string>? MinimalTlsVersion { get; set; }

        /// <summary>
        /// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        [Input("publicNetworkAccess")]
        public Input<string>? PublicNetworkAccess { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

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

        /// <summary>
        /// The version of the server.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServerArgs()
        {
        }
    }
}
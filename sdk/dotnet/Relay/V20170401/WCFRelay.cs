// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Relay.V20170401
{
    /// <summary>
    /// Description of the WCF relay resource.
    /// </summary>
    public partial class WCFRelay : Pulumi.CustomResource
    {
        /// <summary>
        /// The time the WCF relay was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Returns true if the relay is dynamic; otherwise, false.
        /// </summary>
        [Output("isDynamic")]
        public Output<bool> IsDynamic { get; private set; } = null!;

        /// <summary>
        /// The number of listeners for this relay. Note that min :1 and max:25 are supported.
        /// </summary>
        [Output("listenerCount")]
        public Output<int> ListenerCount { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// WCF relay type.
        /// </summary>
        [Output("relayType")]
        public Output<string?> RelayType { get; private set; } = null!;

        /// <summary>
        /// Returns true if client authorization is needed for this relay; otherwise, false.
        /// </summary>
        [Output("requiresClientAuthorization")]
        public Output<bool?> RequiresClientAuthorization { get; private set; } = null!;

        /// <summary>
        /// Returns true if transport security is needed for this relay; otherwise, false.
        /// </summary>
        [Output("requiresTransportSecurity")]
        public Output<bool?> RequiresTransportSecurity { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The time the namespace was updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
        /// </summary>
        [Output("userMetadata")]
        public Output<string?> UserMetadata { get; private set; } = null!;


        /// <summary>
        /// Create a WCFRelay resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WCFRelay(string name, WCFRelayArgs args, CustomResourceOptions? options = null)
            : base("azurerm:relay/v20170401:WCFRelay", name, args ?? new WCFRelayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WCFRelay(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:relay/v20170401:WCFRelay", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:relay/latest:WCFRelay"},
                    new Pulumi.Alias { Type = "azurerm:relay/v20160701:WCFRelay"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WCFRelay resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WCFRelay Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WCFRelay(name, id, options);
        }
    }

    public sealed class WCFRelayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The relay name.
        /// </summary>
        [Input("relayName", required: true)]
        public Input<string> RelayName { get; set; } = null!;

        /// <summary>
        /// WCF relay type.
        /// </summary>
        [Input("relayType")]
        public Input<string>? RelayType { get; set; }

        /// <summary>
        /// Returns true if client authorization is needed for this relay; otherwise, false.
        /// </summary>
        [Input("requiresClientAuthorization")]
        public Input<bool>? RequiresClientAuthorization { get; set; }

        /// <summary>
        /// Returns true if transport security is needed for this relay; otherwise, false.
        /// </summary>
        [Input("requiresTransportSecurity")]
        public Input<bool>? RequiresTransportSecurity { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
        /// </summary>
        [Input("userMetadata")]
        public Input<string>? UserMetadata { get; set; }

        public WCFRelayArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.V20170101
{
    /// <summary>
    /// The relationship link resource format.
    /// </summary>
    public partial class RelationshipLink : Pulumi.CustomResource
    {
        /// <summary>
        /// Localized descriptions for the Relationship Link.
        /// </summary>
        [Output("description")]
        public Output<ImmutableDictionary<string, string>?> Description { get; private set; } = null!;

        /// <summary>
        /// Localized display name for the Relationship Link.
        /// </summary>
        [Output("displayName")]
        public Output<ImmutableDictionary<string, string>?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The InteractionType associated with the Relationship Link.
        /// </summary>
        [Output("interactionType")]
        public Output<string> InteractionType { get; private set; } = null!;

        /// <summary>
        /// The name of the Relationship Link.
        /// </summary>
        [Output("linkName")]
        public Output<string> LinkName { get; private set; } = null!;

        /// <summary>
        /// The mappings between Interaction and Relationship fields.
        /// </summary>
        [Output("mappings")]
        public Output<ImmutableArray<Outputs.RelationshipLinkFieldMappingResponseResult>> Mappings { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The property references for the Profile of the Relationship.
        /// </summary>
        [Output("profilePropertyReferences")]
        public Output<ImmutableArray<Outputs.ParticipantPropertyReferenceResponseResult>> ProfilePropertyReferences { get; private set; } = null!;

        /// <summary>
        /// Provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The property references for the Related Profile of the Relationship.
        /// </summary>
        [Output("relatedProfilePropertyReferences")]
        public Output<ImmutableArray<Outputs.ParticipantPropertyReferenceResponseResult>> RelatedProfilePropertyReferences { get; private set; } = null!;

        /// <summary>
        /// The relationship guid id.
        /// </summary>
        [Output("relationshipGuidId")]
        public Output<string> RelationshipGuidId { get; private set; } = null!;

        /// <summary>
        /// The Relationship associated with the Link.
        /// </summary>
        [Output("relationshipName")]
        public Output<string> RelationshipName { get; private set; } = null!;

        /// <summary>
        /// The hub name.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RelationshipLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RelationshipLink(string name, RelationshipLinkArgs args, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/v20170101:RelationshipLink", name, args ?? new RelationshipLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RelationshipLink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/v20170101:RelationshipLink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Alias { Type = "azurerm:customerinsights/v20170426:RelationshipLink"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RelationshipLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RelationshipLink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RelationshipLink(name, id, options);
        }
    }

    public sealed class RelationshipLinkArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        private InputMap<string>? _description;

        /// <summary>
        /// Localized descriptions for the Relationship Link.
        /// </summary>
        public InputMap<string> Description
        {
            get => _description ?? (_description = new InputMap<string>());
            set => _description = value;
        }

        [Input("displayName")]
        private InputMap<string>? _displayName;

        /// <summary>
        /// Localized display name for the Relationship Link.
        /// </summary>
        public InputMap<string> DisplayName
        {
            get => _displayName ?? (_displayName = new InputMap<string>());
            set => _displayName = value;
        }

        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public Input<string> HubName { get; set; } = null!;

        /// <summary>
        /// The InteractionType associated with the Relationship Link.
        /// </summary>
        [Input("interactionType", required: true)]
        public Input<string> InteractionType { get; set; } = null!;

        [Input("mappings")]
        private InputList<Inputs.RelationshipLinkFieldMappingArgs>? _mappings;

        /// <summary>
        /// The mappings between Interaction and Relationship fields.
        /// </summary>
        public InputList<Inputs.RelationshipLinkFieldMappingArgs> Mappings
        {
            get => _mappings ?? (_mappings = new InputList<Inputs.RelationshipLinkFieldMappingArgs>());
            set => _mappings = value;
        }

        /// <summary>
        /// The name of the relationship link.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("profilePropertyReferences", required: true)]
        private InputList<Inputs.ParticipantPropertyReferenceArgs>? _profilePropertyReferences;

        /// <summary>
        /// The property references for the Profile of the Relationship.
        /// </summary>
        public InputList<Inputs.ParticipantPropertyReferenceArgs> ProfilePropertyReferences
        {
            get => _profilePropertyReferences ?? (_profilePropertyReferences = new InputList<Inputs.ParticipantPropertyReferenceArgs>());
            set => _profilePropertyReferences = value;
        }

        [Input("relatedProfilePropertyReferences", required: true)]
        private InputList<Inputs.ParticipantPropertyReferenceArgs>? _relatedProfilePropertyReferences;

        /// <summary>
        /// The property references for the Related Profile of the Relationship.
        /// </summary>
        public InputList<Inputs.ParticipantPropertyReferenceArgs> RelatedProfilePropertyReferences
        {
            get => _relatedProfilePropertyReferences ?? (_relatedProfilePropertyReferences = new InputList<Inputs.ParticipantPropertyReferenceArgs>());
            set => _relatedProfilePropertyReferences = value;
        }

        /// <summary>
        /// The Relationship associated with the Link.
        /// </summary>
        [Input("relationshipName", required: true)]
        public Input<string> RelationshipName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public RelationshipLinkArgs()
        {
        }
    }
}
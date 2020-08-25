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
    /// The KPI resource format.
    /// </summary>
    public partial class Kpi : Pulumi.CustomResource
    {
        /// <summary>
        /// The aliases.
        /// </summary>
        [Output("aliases")]
        public Output<ImmutableArray<Outputs.KpiAliasResponseResult>> Aliases { get; private set; } = null!;

        /// <summary>
        /// The calculation window.
        /// </summary>
        [Output("calculationWindow")]
        public Output<string> CalculationWindow { get; private set; } = null!;

        /// <summary>
        /// Name of calculation window field.
        /// </summary>
        [Output("calculationWindowFieldName")]
        public Output<string?> CalculationWindowFieldName { get; private set; } = null!;

        /// <summary>
        /// Localized description for the KPI.
        /// </summary>
        [Output("description")]
        public Output<ImmutableDictionary<string, string>?> Description { get; private set; } = null!;

        /// <summary>
        /// Localized display name for the KPI.
        /// </summary>
        [Output("displayName")]
        public Output<ImmutableDictionary<string, string>?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The mapping entity type.
        /// </summary>
        [Output("entityType")]
        public Output<string> EntityType { get; private set; } = null!;

        /// <summary>
        /// The mapping entity name.
        /// </summary>
        [Output("entityTypeName")]
        public Output<string> EntityTypeName { get; private set; } = null!;

        /// <summary>
        /// The computation expression for the KPI.
        /// </summary>
        [Output("expression")]
        public Output<string> Expression { get; private set; } = null!;

        /// <summary>
        /// The KPI extracts.
        /// </summary>
        [Output("extracts")]
        public Output<ImmutableArray<Outputs.KpiExtractResponseResult>> Extracts { get; private set; } = null!;

        /// <summary>
        /// The filter expression for the KPI.
        /// </summary>
        [Output("filter")]
        public Output<string?> Filter { get; private set; } = null!;

        /// <summary>
        /// The computation function for the KPI.
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// the group by properties for the KPI.
        /// </summary>
        [Output("groupBy")]
        public Output<ImmutableArray<string>> GroupBy { get; private set; } = null!;

        /// <summary>
        /// The KPI GroupByMetadata.
        /// </summary>
        [Output("groupByMetadata")]
        public Output<ImmutableArray<Outputs.KpiGroupByMetadataResponseResult>> GroupByMetadata { get; private set; } = null!;

        /// <summary>
        /// The KPI name.
        /// </summary>
        [Output("kpiName")]
        public Output<string> KpiName { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The participant profiles.
        /// </summary>
        [Output("participantProfilesMetadata")]
        public Output<ImmutableArray<Outputs.KpiParticipantProfilesMetadataResponseResult>> ParticipantProfilesMetadata { get; private set; } = null!;

        /// <summary>
        /// Provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The hub name.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The KPI thresholds.
        /// </summary>
        [Output("thresHolds")]
        public Output<Outputs.KpiThresholdsResponseResult?> ThresHolds { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The unit of measurement for the KPI.
        /// </summary>
        [Output("unit")]
        public Output<string?> Unit { get; private set; } = null!;


        /// <summary>
        /// Create a Kpi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Kpi(string name, KpiArgs args, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/v20170101:Kpi", name, args ?? new KpiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Kpi(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/v20170101:Kpi", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Alias { Type = "azurerm:customerinsights/v20170426:Kpi"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Kpi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Kpi Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Kpi(name, id, options);
        }
    }

    public sealed class KpiArgs : Pulumi.ResourceArgs
    {
        [Input("aliases")]
        private InputList<Inputs.KpiAliasArgs>? _aliases;

        /// <summary>
        /// The aliases.
        /// </summary>
        public InputList<Inputs.KpiAliasArgs> Aliases
        {
            get => _aliases ?? (_aliases = new InputList<Inputs.KpiAliasArgs>());
            set => _aliases = value;
        }

        /// <summary>
        /// The calculation window.
        /// </summary>
        [Input("calculationWindow", required: true)]
        public Input<string> CalculationWindow { get; set; } = null!;

        /// <summary>
        /// Name of calculation window field.
        /// </summary>
        [Input("calculationWindowFieldName")]
        public Input<string>? CalculationWindowFieldName { get; set; }

        [Input("description")]
        private InputMap<string>? _description;

        /// <summary>
        /// Localized description for the KPI.
        /// </summary>
        public InputMap<string> Description
        {
            get => _description ?? (_description = new InputMap<string>());
            set => _description = value;
        }

        [Input("displayName")]
        private InputMap<string>? _displayName;

        /// <summary>
        /// Localized display name for the KPI.
        /// </summary>
        public InputMap<string> DisplayName
        {
            get => _displayName ?? (_displayName = new InputMap<string>());
            set => _displayName = value;
        }

        /// <summary>
        /// The mapping entity type.
        /// </summary>
        [Input("entityType", required: true)]
        public Input<string> EntityType { get; set; } = null!;

        /// <summary>
        /// The mapping entity name.
        /// </summary>
        [Input("entityTypeName", required: true)]
        public Input<string> EntityTypeName { get; set; } = null!;

        /// <summary>
        /// The computation expression for the KPI.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("extracts")]
        private InputList<Inputs.KpiExtractArgs>? _extracts;

        /// <summary>
        /// The KPI extracts.
        /// </summary>
        public InputList<Inputs.KpiExtractArgs> Extracts
        {
            get => _extracts ?? (_extracts = new InputList<Inputs.KpiExtractArgs>());
            set => _extracts = value;
        }

        /// <summary>
        /// The filter expression for the KPI.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The computation function for the KPI.
        /// </summary>
        [Input("function", required: true)]
        public Input<string> Function { get; set; } = null!;

        [Input("groupBy")]
        private InputList<string>? _groupBy;

        /// <summary>
        /// the group by properties for the KPI.
        /// </summary>
        public InputList<string> GroupBy
        {
            get => _groupBy ?? (_groupBy = new InputList<string>());
            set => _groupBy = value;
        }

        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public Input<string> HubName { get; set; } = null!;

        /// <summary>
        /// The name of the KPI.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The KPI thresholds.
        /// </summary>
        [Input("thresHolds")]
        public Input<Inputs.KpiThresholdsArgs>? ThresHolds { get; set; }

        /// <summary>
        /// The unit of measurement for the KPI.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public KpiArgs()
        {
        }
    }
}
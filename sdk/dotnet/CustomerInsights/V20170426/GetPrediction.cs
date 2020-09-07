// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.V20170426
{
    public static class GetPrediction
    {
        public static Task<GetPredictionResult> InvokeAsync(GetPredictionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPredictionResult>("azurerm:customerinsights/v20170426:getPrediction", args ?? new GetPredictionArgs(), options.WithVersion());
    }


    public sealed class GetPredictionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public string HubName { get; set; } = null!;

        /// <summary>
        /// The name of the Prediction.
        /// </summary>
        [Input("predictionName", required: true)]
        public string PredictionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPredictionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPredictionResult
    {
        /// <summary>
        /// Whether do auto analyze.
        /// </summary>
        public readonly bool AutoAnalyze;
        /// <summary>
        /// Description of the prediction.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Description;
        /// <summary>
        /// Display name of the prediction.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DisplayName;
        /// <summary>
        /// The prediction grades.
        /// </summary>
        public readonly ImmutableArray<Outputs.PredictionResponseGradesResult> Grades;
        /// <summary>
        /// Interaction types involved in the prediction.
        /// </summary>
        public readonly ImmutableArray<string> InvolvedInteractionTypes;
        /// <summary>
        /// KPI types involved in the prediction.
        /// </summary>
        public readonly ImmutableArray<string> InvolvedKpiTypes;
        /// <summary>
        /// Relationships involved in the prediction.
        /// </summary>
        public readonly ImmutableArray<string> InvolvedRelationships;
        /// <summary>
        /// Definition of the link mapping of prediction.
        /// </summary>
        public readonly Outputs.PredictionResponseMappingsResult Mappings;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Negative outcome expression.
        /// </summary>
        public readonly string NegativeOutcomeExpression;
        /// <summary>
        /// Positive outcome expression.
        /// </summary>
        public readonly string PositiveOutcomeExpression;
        /// <summary>
        /// Name of the prediction.
        /// </summary>
        public readonly string? PredictionName;
        /// <summary>
        /// Primary profile type.
        /// </summary>
        public readonly string PrimaryProfileType;
        /// <summary>
        /// Provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Scope expression.
        /// </summary>
        public readonly string ScopeExpression;
        /// <summary>
        /// Score label.
        /// </summary>
        public readonly string ScoreLabel;
        /// <summary>
        /// System generated entities.
        /// </summary>
        public readonly Outputs.PredictionResponseSystemGeneratedEntitiesResult SystemGeneratedEntities;
        /// <summary>
        /// The hub name.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPredictionResult(
            bool autoAnalyze,

            ImmutableDictionary<string, string>? description,

            ImmutableDictionary<string, string>? displayName,

            ImmutableArray<Outputs.PredictionResponseGradesResult> grades,

            ImmutableArray<string> involvedInteractionTypes,

            ImmutableArray<string> involvedKpiTypes,

            ImmutableArray<string> involvedRelationships,

            Outputs.PredictionResponseMappingsResult mappings,

            string name,

            string negativeOutcomeExpression,

            string positiveOutcomeExpression,

            string? predictionName,

            string primaryProfileType,

            string provisioningState,

            string scopeExpression,

            string scoreLabel,

            Outputs.PredictionResponseSystemGeneratedEntitiesResult systemGeneratedEntities,

            string tenantId,

            string type)
        {
            AutoAnalyze = autoAnalyze;
            Description = description;
            DisplayName = displayName;
            Grades = grades;
            InvolvedInteractionTypes = involvedInteractionTypes;
            InvolvedKpiTypes = involvedKpiTypes;
            InvolvedRelationships = involvedRelationships;
            Mappings = mappings;
            Name = name;
            NegativeOutcomeExpression = negativeOutcomeExpression;
            PositiveOutcomeExpression = positiveOutcomeExpression;
            PredictionName = predictionName;
            PrimaryProfileType = primaryProfileType;
            ProvisioningState = provisioningState;
            ScopeExpression = scopeExpression;
            ScoreLabel = scoreLabel;
            SystemGeneratedEntities = systemGeneratedEntities;
            TenantId = tenantId;
            Type = type;
        }
    }
}
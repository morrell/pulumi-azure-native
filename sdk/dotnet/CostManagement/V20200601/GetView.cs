// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CostManagement.V20200601
{
    public static class GetView
    {
        public static Task<GetViewResult> InvokeAsync(GetViewArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetViewResult>("azurerm:costmanagement/v20200601:getView", args ?? new GetViewArgs(), options.WithVersion());
    }


    public sealed class GetViewArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// View name
        /// </summary>
        [Input("viewName", required: true)]
        public string ViewName { get; set; } = null!;

        public GetViewArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetViewResult
    {
        /// <summary>
        /// Show costs accumulated over time.
        /// </summary>
        public readonly string? Accumulated;
        /// <summary>
        /// Chart type of the main view in Cost Analysis. Required.
        /// </summary>
        public readonly string? Chart;
        /// <summary>
        /// Date the user created this view.
        /// </summary>
        public readonly string CreatedOn;
        /// <summary>
        /// Has definition for data in this report config.
        /// </summary>
        public readonly Outputs.ReportConfigDatasetResponseResult? Dataset;
        /// <summary>
        /// User input name of the view. Required.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// List of KPIs to show in Cost Analysis UI.
        /// </summary>
        public readonly ImmutableArray<Outputs.KpiPropertiesResponseResult> Kpis;
        /// <summary>
        /// Metric to use when displaying costs.
        /// </summary>
        public readonly string? Metric;
        /// <summary>
        /// Date when the user last modified this view.
        /// </summary>
        public readonly string ModifiedOn;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configuration of 3 sub-views in the Cost Analysis UI.
        /// </summary>
        public readonly ImmutableArray<Outputs.PivotPropertiesResponseResult> Pivots;
        /// <summary>
        /// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// Has time period for pulling data for the report.
        /// </summary>
        public readonly Outputs.ReportConfigTimePeriodResponseResult? TimePeriod;
        /// <summary>
        /// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
        /// </summary>
        public readonly string Timeframe;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetViewResult(
            string? accumulated,

            string? chart,

            string createdOn,

            Outputs.ReportConfigDatasetResponseResult? dataset,

            string? displayName,

            string? eTag,

            ImmutableArray<Outputs.KpiPropertiesResponseResult> kpis,

            string? metric,

            string modifiedOn,

            string name,

            ImmutableArray<Outputs.PivotPropertiesResponseResult> pivots,

            string? scope,

            Outputs.ReportConfigTimePeriodResponseResult? timePeriod,

            string timeframe,

            string type)
        {
            Accumulated = accumulated;
            Chart = chart;
            CreatedOn = createdOn;
            Dataset = dataset;
            DisplayName = displayName;
            ETag = eTag;
            Kpis = kpis;
            Metric = metric;
            ModifiedOn = modifiedOn;
            Name = name;
            Pivots = pivots;
            Scope = scope;
            TimePeriod = timePeriod;
            Timeframe = timeframe;
            Type = type;
        }
    }
}
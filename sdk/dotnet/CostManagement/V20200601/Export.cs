// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CostManagement.V20200601
{
    /// <summary>
    /// An export resource.
    /// </summary>
    public partial class Export : Pulumi.CustomResource
    {
        /// <summary>
        /// Has the definition for the export.
        /// </summary>
        [Output("definition")]
        public Output<Outputs.ExportDefinitionResponseResult> Definition { get; private set; } = null!;

        /// <summary>
        /// Has delivery information for the export.
        /// </summary>
        [Output("deliveryInfo")]
        public Output<Outputs.ExportDeliveryInfoResponseResult> DeliveryInfo { get; private set; } = null!;

        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// The format of the export being delivered. Currently only 'Csv' is supported.
        /// </summary>
        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If the export has an active schedule, provides an estimate of the next execution time.
        /// </summary>
        [Output("nextRunTimeEstimate")]
        public Output<string> NextRunTimeEstimate { get; private set; } = null!;

        /// <summary>
        /// If requested, has the most recent execution history for the export.
        /// </summary>
        [Output("runHistory")]
        public Output<Outputs.ExportExecutionListResultResponseResult?> RunHistory { get; private set; } = null!;

        /// <summary>
        /// Has schedule information for the export.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.ExportScheduleResponseResult?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Export resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Export(string name, ExportArgs args, CustomResourceOptions? options = null)
            : base("azurerm:costmanagement/v20200601:Export", name, args ?? new ExportArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Export(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:costmanagement/v20200601:Export", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Alias { Type = "azurerm:costmanagement/v20190101:Export"},
                    new Alias { Type = "azurerm:costmanagement/v20190901:Export"},
                    new Alias { Type = "azurerm:costmanagement/v20191001:Export"},
                    new Alias { Type = "azurerm:costmanagement/v20191101:Export"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Export resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Export Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Export(name, id, options);
        }
    }

    public sealed class ExportArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Has the definition for the export.
        /// </summary>
        [Input("definition", required: true)]
        public Input<Inputs.ExportDefinitionArgs> Definition { get; set; } = null!;

        /// <summary>
        /// Has delivery information for the export.
        /// </summary>
        [Input("deliveryInfo", required: true)]
        public Input<Inputs.ExportDeliveryInfoArgs> DeliveryInfo { get; set; } = null!;

        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// The format of the export being delivered. Currently only 'Csv' is supported.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Export Name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Has schedule information for the export.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ExportScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// The scope associated with query and export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public ExportArgs()
        {
        }
    }
}
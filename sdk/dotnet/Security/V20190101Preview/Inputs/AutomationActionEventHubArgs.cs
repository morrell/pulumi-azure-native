// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20190101Preview.Inputs
{

    /// <summary>
    /// The target Event Hub to which event data will be exported. To learn more about Security Center continuous export capabilities, visit https://aka.ms/ASCExportLearnMore
    /// </summary>
    public sealed class AutomationActionEventHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the action that will be triggered by the Automation
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        /// <summary>
        /// The target Event Hub connection string (it will not be included in any response).
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The target Event Hub Azure Resource ID.
        /// </summary>
        [Input("eventHubResourceId")]
        public Input<string>? EventHubResourceId { get; set; }

        public AutomationActionEventHubArgs()
        {
        }
    }
}
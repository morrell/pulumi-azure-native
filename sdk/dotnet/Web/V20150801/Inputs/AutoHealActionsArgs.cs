// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801.Inputs
{

    /// <summary>
    /// AutoHealActions - Describes the actions which can be
    ///             taken by the auto-heal module when a rule is triggered.
    /// </summary>
    public sealed class AutoHealActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ActionType - predefined action to be taken
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        /// <summary>
        /// CustomAction - custom action to be taken
        /// </summary>
        [Input("customAction")]
        public Input<Inputs.AutoHealCustomActionArgs>? CustomAction { get; set; }

        /// <summary>
        /// MinProcessExecutionTime - minimum time the process must execute
        ///             before taking the action
        /// </summary>
        [Input("minProcessExecutionTime")]
        public Input<string>? MinProcessExecutionTime { get; set; }

        public AutoHealActionsArgs()
        {
        }
    }
}
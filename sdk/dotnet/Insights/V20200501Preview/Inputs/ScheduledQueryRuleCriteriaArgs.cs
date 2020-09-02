// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20200501Preview.Inputs
{

    /// <summary>
    /// The rule criteria that defines the conditions of the scheduled query rule.
    /// </summary>
    public sealed class ScheduledQueryRuleCriteriaArgs : Pulumi.ResourceArgs
    {
        [Input("allOf")]
        private InputList<Inputs.ConditionArgs>? _allOf;

        /// <summary>
        /// A list of conditions to evaluate against the specified scopes
        /// </summary>
        public InputList<Inputs.ConditionArgs> AllOf
        {
            get => _allOf ?? (_allOf = new InputList<Inputs.ConditionArgs>());
            set => _allOf = value;
        }

        public ScheduledQueryRuleCriteriaArgs()
        {
        }
    }
}
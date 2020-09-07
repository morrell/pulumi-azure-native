// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AlertsManagement.V20200804Preview.Outputs
{

    [OutputType]
    public sealed class HealthAlertCriteriaResponseResult
    {
        /// <summary>
        /// The list of metric criteria for this 'all of' operation. 
        /// </summary>
        public readonly ImmutableArray<Outputs.VmGuestHealthAlertCriterionResponseResult> AllOf;

        [OutputConstructor]
        private HealthAlertCriteriaResponseResult(ImmutableArray<Outputs.VmGuestHealthAlertCriterionResponseResult> allOf)
        {
            AllOf = allOf;
        }
    }
}
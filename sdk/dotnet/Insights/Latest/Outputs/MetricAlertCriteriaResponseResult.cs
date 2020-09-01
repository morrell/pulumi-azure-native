// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.Latest.Outputs
{

    [OutputType]
    public sealed class MetricAlertCriteriaResponseResult
    {
        /// <summary>
        /// specifies the type of the alert criteria.
        /// </summary>
        public readonly string OdataType;

        [OutputConstructor]
        private MetricAlertCriteriaResponseResult(string odataType)
        {
            OdataType = odataType;
        }
    }
}
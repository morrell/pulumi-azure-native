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
    public sealed class ActionResponseResult
    {
        /// <summary>
        /// Specifies the action. Supported values - AlertingAction, LogToMetricAction
        /// </summary>
        public readonly string OdataType;

        [OutputConstructor]
        private ActionResponseResult(string odataType)
        {
            OdataType = odataType;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20180301.Inputs
{

    /// <summary>
    /// Specifies the metric alert rule criteria for a web test resource.
    /// </summary>
    public sealed class WebtestLocationAvailabilityCriteriaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Application Insights resource Id.
        /// </summary>
        [Input("componentId", required: true)]
        public Input<string> ComponentId { get; set; } = null!;

        /// <summary>
        /// The number of failed locations.
        /// </summary>
        [Input("failedLocationCount", required: true)]
        public Input<double> FailedLocationCount { get; set; } = null!;

        /// <summary>
        /// specifies the type of the alert criteria.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        /// <summary>
        /// The Application Insights web test Id.
        /// </summary>
        [Input("webTestId", required: true)]
        public Input<string> WebTestId { get; set; } = null!;

        public WebtestLocationAvailabilityCriteriaArgs()
        {
        }
    }
}
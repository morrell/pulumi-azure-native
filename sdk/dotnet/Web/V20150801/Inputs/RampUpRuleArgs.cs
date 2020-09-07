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
    /// Routing rules for ramp up testing. This rule allows to redirect static traffic % to a slot or to gradually change routing % based on performance
    /// </summary>
    public sealed class RampUpRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname of a slot to which the traffic will be redirected if decided to. E.g. mysite-stage.azurewebsites.net
        /// </summary>
        [Input("actionHostName")]
        public Input<string>? ActionHostName { get; set; }

        /// <summary>
        /// Custom decision algorithm can be provided in TiPCallback site extension which Url can be specified. See TiPCallback site extension for the scaffold and contracts.
        ///             https://www.siteextensions.net/packages/TiPCallback/
        /// </summary>
        [Input("changeDecisionCallbackUrl")]
        public Input<string>? ChangeDecisionCallbackUrl { get; set; }

        /// <summary>
        /// [Optional] Specifies interval in minutes to reevaluate ReroutePercentage
        /// </summary>
        [Input("changeIntervalInMinutes")]
        public Input<int>? ChangeIntervalInMinutes { get; set; }

        /// <summary>
        /// [Optional] In auto ramp up scenario this is the step to add/remove from {Microsoft.Web.Hosting.Administration.RampUpRule.ReroutePercentage} until it reaches 
        ///             {Microsoft.Web.Hosting.Administration.RampUpRule.MinReroutePercentage} or {Microsoft.Web.Hosting.Administration.RampUpRule.MaxReroutePercentage}. Site metrics are checked every N minutes specified in {Microsoft.Web.Hosting.Administration.RampUpRule.ChangeIntervalInMinutes}.
        ///             Custom decision algorithm can be provided in TiPCallback site extension which Url can be specified in {Microsoft.Web.Hosting.Administration.RampUpRule.ChangeDecisionCallbackUrl}
        /// </summary>
        [Input("changeStep")]
        public Input<double>? ChangeStep { get; set; }

        /// <summary>
        /// [Optional] Specifies upper boundary below which ReroutePercentage will stay.
        /// </summary>
        [Input("maxReroutePercentage")]
        public Input<double>? MaxReroutePercentage { get; set; }

        /// <summary>
        /// [Optional] Specifies lower boundary above which ReroutePercentage will stay.
        /// </summary>
        [Input("minReroutePercentage")]
        public Input<double>? MinReroutePercentage { get; set; }

        /// <summary>
        /// Name of the routing rule. The recommended name would be to point to the slot which will receive the traffic in the experiment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Percentage of the traffic which will be redirected to {Microsoft.Web.Hosting.Administration.RampUpRule.ActionHostName}
        /// </summary>
        [Input("reroutePercentage")]
        public Input<double>? ReroutePercentage { get; set; }

        public RampUpRuleArgs()
        {
        }
    }
}
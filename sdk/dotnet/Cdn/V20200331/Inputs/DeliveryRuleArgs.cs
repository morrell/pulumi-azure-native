// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20200331.Inputs
{

    /// <summary>
    /// A rule that specifies a set of actions and conditions
    /// </summary>
    public sealed class DeliveryRuleArgs : Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Union<Inputs.DeliveryRuleCacheExpirationActionArgs, Union<Inputs.DeliveryRuleCacheKeyQueryStringActionArgs, Union<Inputs.DeliveryRuleRequestHeaderActionArgs, Union<Inputs.DeliveryRuleResponseHeaderActionArgs, Union<Inputs.UrlRedirectActionArgs, Union<Inputs.UrlRewriteActionArgs, Inputs.UrlSigningActionArgs>>>>>>>? _actions;

        /// <summary>
        /// A list of actions that are executed when all the conditions of a rule are satisfied.
        /// </summary>
        public InputList<Union<Inputs.DeliveryRuleCacheExpirationActionArgs, Union<Inputs.DeliveryRuleCacheKeyQueryStringActionArgs, Union<Inputs.DeliveryRuleRequestHeaderActionArgs, Union<Inputs.DeliveryRuleResponseHeaderActionArgs, Union<Inputs.UrlRedirectActionArgs, Union<Inputs.UrlRewriteActionArgs, Inputs.UrlSigningActionArgs>>>>>>> Actions
        {
            get => _actions ?? (_actions = new InputList<Union<Inputs.DeliveryRuleCacheExpirationActionArgs, Union<Inputs.DeliveryRuleCacheKeyQueryStringActionArgs, Union<Inputs.DeliveryRuleRequestHeaderActionArgs, Union<Inputs.DeliveryRuleResponseHeaderActionArgs, Union<Inputs.UrlRedirectActionArgs, Union<Inputs.UrlRewriteActionArgs, Inputs.UrlSigningActionArgs>>>>>>>());
            set => _actions = value;
        }

        [Input("conditions")]
        private InputList<Union<Inputs.DeliveryRuleCookiesConditionArgs, Union<Inputs.DeliveryRuleHttpVersionConditionArgs, Union<Inputs.DeliveryRuleIsDeviceConditionArgs, Union<Inputs.DeliveryRulePostArgsConditionArgs, Union<Inputs.DeliveryRuleQueryStringConditionArgs, Union<Inputs.DeliveryRuleRemoteAddressConditionArgs, Union<Inputs.DeliveryRuleRequestBodyConditionArgs, Union<Inputs.DeliveryRuleRequestHeaderConditionArgs, Union<Inputs.DeliveryRuleRequestMethodConditionArgs, Union<Inputs.DeliveryRuleRequestSchemeConditionArgs, Union<Inputs.DeliveryRuleRequestUriConditionArgs, Union<Inputs.DeliveryRuleUrlFileExtensionConditionArgs, Union<Inputs.DeliveryRuleUrlFileNameConditionArgs, Inputs.DeliveryRuleUrlPathConditionArgs>>>>>>>>>>>>>>? _conditions;

        /// <summary>
        /// A list of conditions that must be matched for the actions to be executed
        /// </summary>
        public InputList<Union<Inputs.DeliveryRuleCookiesConditionArgs, Union<Inputs.DeliveryRuleHttpVersionConditionArgs, Union<Inputs.DeliveryRuleIsDeviceConditionArgs, Union<Inputs.DeliveryRulePostArgsConditionArgs, Union<Inputs.DeliveryRuleQueryStringConditionArgs, Union<Inputs.DeliveryRuleRemoteAddressConditionArgs, Union<Inputs.DeliveryRuleRequestBodyConditionArgs, Union<Inputs.DeliveryRuleRequestHeaderConditionArgs, Union<Inputs.DeliveryRuleRequestMethodConditionArgs, Union<Inputs.DeliveryRuleRequestSchemeConditionArgs, Union<Inputs.DeliveryRuleRequestUriConditionArgs, Union<Inputs.DeliveryRuleUrlFileExtensionConditionArgs, Union<Inputs.DeliveryRuleUrlFileNameConditionArgs, Inputs.DeliveryRuleUrlPathConditionArgs>>>>>>>>>>>>>> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Union<Inputs.DeliveryRuleCookiesConditionArgs, Union<Inputs.DeliveryRuleHttpVersionConditionArgs, Union<Inputs.DeliveryRuleIsDeviceConditionArgs, Union<Inputs.DeliveryRulePostArgsConditionArgs, Union<Inputs.DeliveryRuleQueryStringConditionArgs, Union<Inputs.DeliveryRuleRemoteAddressConditionArgs, Union<Inputs.DeliveryRuleRequestBodyConditionArgs, Union<Inputs.DeliveryRuleRequestHeaderConditionArgs, Union<Inputs.DeliveryRuleRequestMethodConditionArgs, Union<Inputs.DeliveryRuleRequestSchemeConditionArgs, Union<Inputs.DeliveryRuleRequestUriConditionArgs, Union<Inputs.DeliveryRuleUrlFileExtensionConditionArgs, Union<Inputs.DeliveryRuleUrlFileNameConditionArgs, Inputs.DeliveryRuleUrlPathConditionArgs>>>>>>>>>>>>>>());
            set => _conditions = value;
        }

        /// <summary>
        /// Name of the rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        public DeliveryRuleArgs()
        {
        }
    }
}
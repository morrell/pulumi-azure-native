// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180801.Outputs
{

    [OutputType]
    public sealed class ManagedRuleSetResponseResult
    {
        /// <summary>
        /// Describes priority of the rule
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// RuleSetType - AzureManagedRuleSet or OWASP RuleSets.
        /// </summary>
        public readonly string RuleSetType;
        /// <summary>
        /// defines version of the rule set
        /// </summary>
        public readonly int? Version;

        [OutputConstructor]
        private ManagedRuleSetResponseResult(
            int? priority,

            string ruleSetType,

            int? version)
        {
            Priority = priority;
            RuleSetType = ruleSetType;
            Version = version;
        }
    }
}
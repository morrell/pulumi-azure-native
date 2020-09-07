// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200101.Outputs
{

    [OutputType]
    public sealed class IpFilterRuleResponseResult
    {
        /// <summary>
        /// The desired action for requests captured by this rule.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The name of the IP filter rule.
        /// </summary>
        public readonly string FilterName;
        /// <summary>
        /// A string that contains the IP address range in CIDR notation for the rule.
        /// </summary>
        public readonly string IpMask;
        /// <summary>
        /// Target for requests captured by this rule.
        /// </summary>
        public readonly string? Target;

        [OutputConstructor]
        private IpFilterRuleResponseResult(
            string action,

            string filterName,

            string ipMask,

            string? target)
        {
            Action = action;
            FilterName = filterName;
            IpMask = ipMask;
            Target = target;
        }
    }
}
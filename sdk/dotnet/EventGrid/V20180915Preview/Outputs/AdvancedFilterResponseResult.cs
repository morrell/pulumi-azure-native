// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20180915Preview.Outputs
{

    [OutputType]
    public sealed class AdvancedFilterResponseResult
    {
        /// <summary>
        /// The filter key. Represents an event property with up to two levels of nesting.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// Represents the filter operator
        /// </summary>
        public readonly string OperatorType;

        [OutputConstructor]
        private AdvancedFilterResponseResult(
            string? key,

            string operatorType)
        {
            Key = key;
            OperatorType = operatorType;
        }
    }
}
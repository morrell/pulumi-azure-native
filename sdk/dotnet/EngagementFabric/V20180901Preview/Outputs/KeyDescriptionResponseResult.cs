// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EngagementFabric.V20180901Preview.Outputs
{

    [OutputType]
    public sealed class KeyDescriptionResponseResult
    {
        /// <summary>
        /// The name of the key
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The rank of the key
        /// </summary>
        public readonly string Rank;
        /// <summary>
        /// The value of the key
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private KeyDescriptionResponseResult(
            string name,

            string rank,

            string value)
        {
            Name = name;
            Rank = rank;
            Value = value;
        }
    }
}
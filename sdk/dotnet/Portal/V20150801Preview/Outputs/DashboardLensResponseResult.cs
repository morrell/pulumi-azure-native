// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Portal.V20150801Preview.Outputs
{

    [OutputType]
    public sealed class DashboardLensResponseResult
    {
        /// <summary>
        /// The dashboard len's metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, ImmutableDictionary<string, object>>? Metadata;
        /// <summary>
        /// The lens order.
        /// </summary>
        public readonly int Order;
        /// <summary>
        /// The dashboard parts.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.DashboardPartsResponseResult> Parts;

        [OutputConstructor]
        private DashboardLensResponseResult(
            ImmutableDictionary<string, ImmutableDictionary<string, object>>? metadata,

            int order,

            ImmutableDictionary<string, Outputs.DashboardPartsResponseResult> parts)
        {
            Metadata = metadata;
            Order = order;
            Parts = parts;
        }
    }
}
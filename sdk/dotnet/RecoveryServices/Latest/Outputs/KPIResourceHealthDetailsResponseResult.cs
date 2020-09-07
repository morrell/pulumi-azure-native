// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.Latest.Outputs
{

    [OutputType]
    public sealed class KPIResourceHealthDetailsResponseResult
    {
        /// <summary>
        /// Resource Health Status
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceHealthDetailsResponseResult> ResourceHealthDetails;
        /// <summary>
        /// Resource Health Status
        /// </summary>
        public readonly string? ResourceHealthStatus;

        [OutputConstructor]
        private KPIResourceHealthDetailsResponseResult(
            ImmutableArray<Outputs.ResourceHealthDetailsResponseResult> resourceHealthDetails,

            string? resourceHealthStatus)
        {
            ResourceHealthDetails = resourceHealthDetails;
            ResourceHealthStatus = resourceHealthStatus;
        }
    }
}
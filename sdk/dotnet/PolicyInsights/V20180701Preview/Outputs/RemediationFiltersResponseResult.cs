// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.PolicyInsights.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class RemediationFiltersResponseResult
    {
        /// <summary>
        /// The resource locations that will be remediated.
        /// </summary>
        public readonly ImmutableArray<string> Locations;

        [OutputConstructor]
        private RemediationFiltersResponseResult(ImmutableArray<string> locations)
        {
            Locations = locations;
        }
    }
}
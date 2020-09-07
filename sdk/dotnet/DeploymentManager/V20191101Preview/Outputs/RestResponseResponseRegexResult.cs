// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DeploymentManager.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class RestResponseResponseRegexResult
    {
        /// <summary>
        /// Indicates whether any or all of the expressions should match with the response content.
        /// </summary>
        public readonly string? MatchQuantifier;
        /// <summary>
        /// The list of regular expressions.
        /// </summary>
        public readonly ImmutableArray<string> Matches;

        [OutputConstructor]
        private RestResponseResponseRegexResult(
            string? matchQuantifier,

            ImmutableArray<string> matches)
        {
            MatchQuantifier = matchQuantifier;
            Matches = matches;
        }
    }
}
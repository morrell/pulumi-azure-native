// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20200331.Outputs
{

    [OutputType]
    public sealed class UrlRewriteActionParametersResponseResult
    {
        /// <summary>
        /// Define the relative URL to which the above requests will be rewritten by.
        /// </summary>
        public readonly string Destination;
        public readonly string OdataType;
        /// <summary>
        /// Whether to preserve unmatched path. Default value is true.
        /// </summary>
        public readonly bool? PreserveUnmatchedPath;
        /// <summary>
        /// define a request URI pattern that identifies the type of requests that may be rewritten. If value is blank, all strings are matched.
        /// </summary>
        public readonly string SourcePattern;

        [OutputConstructor]
        private UrlRewriteActionParametersResponseResult(
            string destination,

            string odataType,

            bool? preserveUnmatchedPath,

            string sourcePattern)
        {
            Destination = destination;
            OdataType = odataType;
            PreserveUnmatchedPath = preserveUnmatchedPath;
            SourcePattern = sourcePattern;
        }
    }
}
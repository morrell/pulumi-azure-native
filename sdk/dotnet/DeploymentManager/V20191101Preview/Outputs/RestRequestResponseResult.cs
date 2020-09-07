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
    public sealed class RestRequestResponseResult
    {
        /// <summary>
        /// The authentication information required in the request to the health provider.
        /// </summary>
        public readonly Union<Outputs.ApiKeyAuthenticationResponseResult, Outputs.RolloutIdentityAuthenticationResponseResult> Authentication;
        /// <summary>
        /// The HTTP method to use for the request.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// The HTTP URI to use for the request.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private RestRequestResponseResult(
            Union<Outputs.ApiKeyAuthenticationResponseResult, Outputs.RolloutIdentityAuthenticationResponseResult> authentication,

            string method,

            string uri)
        {
            Authentication = authentication;
            Method = method;
            Uri = uri;
        }
    }
}
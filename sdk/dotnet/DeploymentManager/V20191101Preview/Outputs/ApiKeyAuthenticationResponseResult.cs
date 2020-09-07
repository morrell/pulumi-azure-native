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
    public sealed class ApiKeyAuthenticationResponseResult
    {
        /// <summary>
        /// The location of the authentication key/value pair in the request.
        /// </summary>
        public readonly string In;
        /// <summary>
        /// The key name of the authentication key/value pair.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The authentication type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value of the authentication key/value pair.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ApiKeyAuthenticationResponseResult(
            string @in,

            string name,

            string type,

            string value)
        {
            In = @in;
            Name = name;
            Type = type;
            Value = value;
        }
    }
}
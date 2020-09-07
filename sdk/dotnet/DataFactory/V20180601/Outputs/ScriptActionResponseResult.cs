// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20180601.Outputs
{

    [OutputType]
    public sealed class ScriptActionResponseResult
    {
        /// <summary>
        /// The user provided name of the script action.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The parameters for the script action.
        /// </summary>
        public readonly string? Parameters;
        /// <summary>
        /// The node types on which the script action should be executed.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Roles;
        /// <summary>
        /// The URI for the script action.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private ScriptActionResponseResult(
            string name,

            string? parameters,

            ImmutableDictionary<string, object> roles,

            string uri)
        {
            Name = name;
            Parameters = parameters;
            Roles = roles;
            Uri = uri;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevOps.V20190701Preview.Inputs
{

    /// <summary>
    /// Authorization info used to access a resource (like code repository).
    /// </summary>
    public sealed class AuthorizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of authorization.
        /// </summary>
        [Input("authorizationType", required: true)]
        public Input<string> AuthorizationType { get; set; } = null!;

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Authorization parameters corresponding to the authorization type.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        public AuthorizationArgs()
        {
        }
    }
}
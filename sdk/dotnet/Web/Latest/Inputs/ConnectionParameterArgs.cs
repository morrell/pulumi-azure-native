// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.Latest.Inputs
{

    /// <summary>
    /// Connection provider parameters
    /// </summary>
    public sealed class ConnectionParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// OAuth settings for the connection provider
        /// </summary>
        [Input("oAuthSettings")]
        public Input<Inputs.ApiOAuthSettingsArgs>? OAuthSettings { get; set; }

        /// <summary>
        /// Type of the parameter
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ConnectionParameterArgs()
        {
        }
    }
}
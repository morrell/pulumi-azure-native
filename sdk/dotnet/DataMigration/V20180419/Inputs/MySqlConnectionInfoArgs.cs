// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180419.Inputs
{

    /// <summary>
    /// Information for connecting to MySQL server
    /// </summary>
    public sealed class MySqlConnectionInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Password credential.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Port for Server
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Name of the server
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// Type of connection info
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// User name
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public MySqlConnectionInfoArgs()
        {
        }
    }
}
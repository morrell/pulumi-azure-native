// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180715Preview.Inputs
{

    /// <summary>
    /// Describes a connection to a MongoDB data source
    /// </summary>
    public sealed class MongoDbConnectionInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A MongoDB connection string or blob container URL. The user name and password can be specified here or in the userName and password properties
        /// </summary>
        [Input("connectionString", required: true)]
        public Input<string> ConnectionString { get; set; } = null!;

        /// <summary>
        /// Password credential.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

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

        public MongoDbConnectionInfoArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180715Preview.Outputs
{

    [OutputType]
    public sealed class MongoDbConnectionInfoResponseResult
    {
        /// <summary>
        /// A MongoDB connection string or blob container URL. The user name and password can be specified here or in the userName and password properties
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// Password credential.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Type of connection info
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// User name
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private MongoDbConnectionInfoResponseResult(
            string connectionString,

            string? password,

            string type,

            string? userName)
        {
            ConnectionString = connectionString;
            Password = password;
            Type = type;
            UserName = userName;
        }
    }
}
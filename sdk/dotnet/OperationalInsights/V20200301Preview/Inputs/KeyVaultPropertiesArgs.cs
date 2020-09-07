// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.V20200301Preview.Inputs
{

    /// <summary>
    /// The key vault properties.
    /// </summary>
    public sealed class KeyVaultPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the key associated with the Log Analytics cluster.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// The Key Vault uri which holds they key associated with the Log Analytics cluster.
        /// </summary>
        [Input("keyVaultUri")]
        public Input<string>? KeyVaultUri { get; set; }

        /// <summary>
        /// The version of the key associated with the Log Analytics cluster.
        /// </summary>
        [Input("keyVersion")]
        public Input<string>? KeyVersion { get; set; }

        public KeyVaultPropertiesArgs()
        {
        }
    }
}
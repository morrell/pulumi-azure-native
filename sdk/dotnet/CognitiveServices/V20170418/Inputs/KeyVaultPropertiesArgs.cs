// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CognitiveServices.V20170418.Inputs
{

    /// <summary>
    /// Properties to configure keyVault Properties
    /// </summary>
    public sealed class KeyVaultPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Key from KeyVault
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Uri of KeyVault
        /// </summary>
        [Input("keyVaultUri")]
        public Input<string>? KeyVaultUri { get; set; }

        /// <summary>
        /// Version of the Key from KeyVault
        /// </summary>
        [Input("keyVersion")]
        public Input<string>? KeyVersion { get; set; }

        public KeyVaultPropertiesArgs()
        {
        }
    }
}
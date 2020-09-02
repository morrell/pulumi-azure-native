// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200501Preview.Outputs
{

    [OutputType]
    public sealed class EncryptionPropertyResponseResult
    {
        /// <summary>
        /// Customer Key vault properties.
        /// </summary>
        public readonly Outputs.KeyVaultPropertiesResponseResult KeyVaultProperties;
        /// <summary>
        /// Indicates whether or not the encryption is enabled for the workspace.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private EncryptionPropertyResponseResult(
            Outputs.KeyVaultPropertiesResponseResult keyVaultProperties,

            string status)
        {
            KeyVaultProperties = keyVaultProperties;
            Status = status;
        }
    }
}
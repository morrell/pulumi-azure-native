// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20160101.Outputs
{

    [OutputType]
    public sealed class EncryptionResponseResult
    {
        /// <summary>
        /// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
        /// </summary>
        public readonly string KeySource;
        /// <summary>
        /// List of services which support encryption.
        /// </summary>
        public readonly Outputs.EncryptionServicesResponseResult? Services;

        [OutputConstructor]
        private EncryptionResponseResult(
            string keySource,

            Outputs.EncryptionServicesResponseResult? services)
        {
            KeySource = keySource;
            Services = services;
        }
    }
}
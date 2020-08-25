// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20160501.Outputs
{

    [OutputType]
    public sealed class EncryptionServicesResponseResult
    {
        /// <summary>
        /// The encryption function of the blob storage service.
        /// </summary>
        public readonly Outputs.EncryptionServiceResponseResult? Blob;

        [OutputConstructor]
        private EncryptionServicesResponseResult(Outputs.EncryptionServiceResponseResult? blob)
        {
            Blob = blob;
        }
    }
}
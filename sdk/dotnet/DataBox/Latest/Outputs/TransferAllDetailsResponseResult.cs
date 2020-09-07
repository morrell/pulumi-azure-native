// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBox.Latest.Outputs
{

    [OutputType]
    public sealed class TransferAllDetailsResponseResult
    {
        /// <summary>
        /// Type of the account of data
        /// </summary>
        public readonly string DataAccountType;
        /// <summary>
        /// To indicate if all Azure blobs have to be transferred
        /// </summary>
        public readonly bool? TransferAllBlobs;
        /// <summary>
        /// To indicate if all Azure Files have to be transferred
        /// </summary>
        public readonly bool? TransferAllFiles;

        [OutputConstructor]
        private TransferAllDetailsResponseResult(
            string dataAccountType,

            bool? transferAllBlobs,

            bool? transferAllFiles)
        {
            DataAccountType = dataAccountType;
            TransferAllBlobs = transferAllBlobs;
            TransferAllFiles = transferAllFiles;
        }
    }
}
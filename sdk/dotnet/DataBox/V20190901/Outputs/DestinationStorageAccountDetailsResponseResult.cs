// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBox.V20190901.Outputs
{

    [OutputType]
    public sealed class DestinationStorageAccountDetailsResponseResult
    {
        /// <summary>
        /// Arm Id of the destination where the data has to be moved.
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// Data Destination Type.
        /// </summary>
        public readonly string DataDestinationType;
        /// <summary>
        /// Share password to be shared by all shares in SA.
        /// </summary>
        public readonly string? SharePassword;
        /// <summary>
        /// Destination Storage Account Arm Id.
        /// </summary>
        public readonly string StorageAccountId;

        [OutputConstructor]
        private DestinationStorageAccountDetailsResponseResult(
            string? accountId,

            string dataDestinationType,

            string? sharePassword,

            string storageAccountId)
        {
            AccountId = accountId;
            DataDestinationType = dataDestinationType;
            SharePassword = sharePassword;
            StorageAccountId = storageAccountId;
        }
    }
}
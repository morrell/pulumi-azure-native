// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.Latest.Outputs
{

    [OutputType]
    public sealed class VMwareCbtProtectionContainerMappingDetailsResponseResult
    {
        /// <summary>
        /// Gets the class type. Overridden in derived classes.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The target key vault ARM Id.
        /// </summary>
        public readonly string KeyVaultId;
        /// <summary>
        /// The target key vault URI.
        /// </summary>
        public readonly string KeyVaultUri;
        /// <summary>
        /// The secret name of the service bus connection string.
        /// </summary>
        public readonly string ServiceBusConnectionStringSecretName;
        /// <summary>
        /// The storage account ARM Id.
        /// </summary>
        public readonly string StorageAccountId;
        /// <summary>
        /// The secret name of the storage account.
        /// </summary>
        public readonly string StorageAccountSasSecretName;
        /// <summary>
        /// The target location.
        /// </summary>
        public readonly string TargetLocation;

        [OutputConstructor]
        private VMwareCbtProtectionContainerMappingDetailsResponseResult(
            string instanceType,

            string keyVaultId,

            string keyVaultUri,

            string serviceBusConnectionStringSecretName,

            string storageAccountId,

            string storageAccountSasSecretName,

            string targetLocation)
        {
            InstanceType = instanceType;
            KeyVaultId = keyVaultId;
            KeyVaultUri = keyVaultUri;
            ServiceBusConnectionStringSecretName = serviceBusConnectionStringSecretName;
            StorageAccountId = storageAccountId;
            StorageAccountSasSecretName = storageAccountSasSecretName;
            TargetLocation = targetLocation;
        }
    }
}
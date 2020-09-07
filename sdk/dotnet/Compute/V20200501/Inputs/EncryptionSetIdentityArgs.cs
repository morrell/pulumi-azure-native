// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20200501.Inputs
{

    /// <summary>
    /// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
    /// </summary>
    public sealed class EncryptionSetIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of Managed Identity used by the DiskEncryptionSet. Only SystemAssigned is supported.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EncryptionSetIdentityArgs()
        {
        }
    }
}
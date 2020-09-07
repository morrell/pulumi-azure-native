// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NetApp.Latest.Inputs
{

    /// <summary>
    /// Mount target properties
    /// </summary>
    public sealed class MountTargetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID v4 used to identify the MountTarget
        /// </summary>
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        /// <summary>
        /// The SMB server's Fully Qualified Domain Name, FQDN
        /// </summary>
        [Input("smbServerFqdn")]
        public Input<string>? SmbServerFqdn { get; set; }

        public MountTargetPropertiesArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20170901Preview.Inputs
{

    /// <summary>
    /// Details of the Azure File Share to mount on the cluster.
    /// </summary>
    public sealed class AzureFileShareReferenceArgs : Pulumi.ResourceArgs
    {
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("azureFileUrl", required: true)]
        public Input<string> AzureFileUrl { get; set; } = null!;

        /// <summary>
        /// Credentials to access Azure File Share.
        /// </summary>
        [Input("credentials", required: true)]
        public Input<Inputs.AzureStorageCredentialsInfoArgs> Credentials { get; set; } = null!;

        /// <summary>
        /// Default value is 0777. Valid only if OS is linux.
        /// </summary>
        [Input("directoryMode")]
        public Input<string>? DirectoryMode { get; set; }

        /// <summary>
        /// Default value is 0777. Valid only if OS is linux.
        /// </summary>
        [Input("fileMode")]
        public Input<string>? FileMode { get; set; }

        /// <summary>
        /// Note that all file shares will be mounted under $AZ_BATCHAI_MOUNT_ROOT location.
        /// </summary>
        [Input("relativeMountPath", required: true)]
        public Input<string> RelativeMountPath { get; set; } = null!;

        public AzureFileShareReferenceArgs()
        {
        }
    }
}
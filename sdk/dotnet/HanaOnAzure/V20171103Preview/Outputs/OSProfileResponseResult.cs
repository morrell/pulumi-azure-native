// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HanaOnAzure.V20171103Preview.Outputs
{

    [OutputType]
    public sealed class OSProfileResponseResult
    {
        /// <summary>
        /// Specifies the host OS name of the HANA instance.
        /// </summary>
        public readonly string? ComputerName;
        /// <summary>
        /// This property allows you to specify the type of the OS.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// Specifies the SSH public key used to access the operating system.
        /// </summary>
        public readonly string? SshPublicKey;
        /// <summary>
        /// Specifies version of operating system.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private OSProfileResponseResult(
            string? computerName,

            string osType,

            string? sshPublicKey,

            string version)
        {
            ComputerName = computerName;
            OsType = osType;
            SshPublicKey = sshPublicKey;
            Version = version;
        }
    }
}
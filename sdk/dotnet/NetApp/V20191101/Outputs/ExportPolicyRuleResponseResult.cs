// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NetApp.V20191101.Outputs
{

    [OutputType]
    public sealed class ExportPolicyRuleResponseResult
    {
        /// <summary>
        /// Client ingress specification as comma separated string with IPv4 CIDRs, IPv4 host addresses and host names
        /// </summary>
        public readonly string? AllowedClients;
        /// <summary>
        /// Allows CIFS protocol
        /// </summary>
        public readonly bool? Cifs;
        /// <summary>
        /// Allows NFSv3 protocol
        /// </summary>
        public readonly bool? Nfsv3;
        /// <summary>
        /// Allows NFSv4.1 protocol
        /// </summary>
        public readonly bool? Nfsv41;
        /// <summary>
        /// Order index
        /// </summary>
        public readonly int? RuleIndex;
        /// <summary>
        /// Read only access
        /// </summary>
        public readonly bool? UnixReadOnly;
        /// <summary>
        /// Read and write access
        /// </summary>
        public readonly bool? UnixReadWrite;

        [OutputConstructor]
        private ExportPolicyRuleResponseResult(
            string? allowedClients,

            bool? cifs,

            bool? nfsv3,

            bool? nfsv41,

            int? ruleIndex,

            bool? unixReadOnly,

            bool? unixReadWrite)
        {
            AllowedClients = allowedClients;
            Cifs = cifs;
            Nfsv3 = nfsv3;
            Nfsv41 = nfsv41;
            RuleIndex = ruleIndex;
            UnixReadOnly = unixReadOnly;
            UnixReadWrite = unixReadWrite;
        }
    }
}
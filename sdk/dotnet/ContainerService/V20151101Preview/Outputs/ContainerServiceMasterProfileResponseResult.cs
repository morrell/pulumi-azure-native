// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20151101Preview.Outputs
{

    [OutputType]
    public sealed class ContainerServiceMasterProfileResponseResult
    {
        /// <summary>
        /// Number of masters (VMs) in the container cluster
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// DNS prefix to be used to create FQDN for master
        /// </summary>
        public readonly string DnsPrefix;
        /// <summary>
        /// FQDN for the master
        /// </summary>
        public readonly string Fqdn;

        [OutputConstructor]
        private ContainerServiceMasterProfileResponseResult(
            int? count,

            string dnsPrefix,

            string fqdn)
        {
            Count = count;
            DnsPrefix = dnsPrefix;
            Fqdn = fqdn;
        }
    }
}
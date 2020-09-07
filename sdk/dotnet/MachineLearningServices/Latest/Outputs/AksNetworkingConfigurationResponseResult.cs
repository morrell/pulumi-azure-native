// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.Latest.Outputs
{

    [OutputType]
    public sealed class AksNetworkingConfigurationResponseResult
    {
        /// <summary>
        /// An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.
        /// </summary>
        public readonly string? DnsServiceIP;
        /// <summary>
        /// A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
        /// </summary>
        public readonly string? DockerBridgeCidr;
        /// <summary>
        /// A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.
        /// </summary>
        public readonly string? ServiceCidr;
        /// <summary>
        /// Virtual network subnet resource ID the compute nodes belong to
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private AksNetworkingConfigurationResponseResult(
            string? dnsServiceIP,

            string? dockerBridgeCidr,

            string? serviceCidr,

            string? subnetId)
        {
            DnsServiceIP = dnsServiceIP;
            DockerBridgeCidr = dockerBridgeCidr;
            ServiceCidr = serviceCidr;
            SubnetId = subnetId;
        }
    }
}
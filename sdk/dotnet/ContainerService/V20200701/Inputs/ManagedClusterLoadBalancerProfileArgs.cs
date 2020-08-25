// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20200701.Inputs
{

    /// <summary>
    /// Profile of the managed cluster load balancer.
    /// </summary>
    public sealed class ManagedClusterLoadBalancerProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Desired number of allocated SNAT ports per VM. Allowed values must be in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports.
        /// </summary>
        [Input("allocatedOutboundPorts")]
        public Input<int>? AllocatedOutboundPorts { get; set; }

        [Input("effectiveOutboundIPs")]
        private InputList<Inputs.ResourceReferenceArgs>? _effectiveOutboundIPs;

        /// <summary>
        /// The effective outbound IP resources of the cluster load balancer.
        /// </summary>
        public InputList<Inputs.ResourceReferenceArgs> EffectiveOutboundIPs
        {
            get => _effectiveOutboundIPs ?? (_effectiveOutboundIPs = new InputList<Inputs.ResourceReferenceArgs>());
            set => _effectiveOutboundIPs = value;
        }

        /// <summary>
        /// Desired outbound flow idle timeout in minutes. Allowed values must be in the range of 4 to 120 (inclusive). The default value is 30 minutes.
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        /// <summary>
        /// Desired managed outbound IPs for the cluster load balancer.
        /// </summary>
        [Input("managedOutboundIPs")]
        public Input<Inputs.ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs>? ManagedOutboundIPs { get; set; }

        /// <summary>
        /// Desired outbound IP Prefix resources for the cluster load balancer.
        /// </summary>
        [Input("outboundIPPrefixes")]
        public Input<Inputs.ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs>? OutboundIPPrefixes { get; set; }

        /// <summary>
        /// Desired outbound IP resources for the cluster load balancer.
        /// </summary>
        [Input("outboundIPs")]
        public Input<Inputs.ManagedClusterLoadBalancerProfileOutboundIPsArgs>? OutboundIPs { get; set; }

        public ManagedClusterLoadBalancerProfileArgs()
        {
        }
    }
}
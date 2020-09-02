// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class LoadBalancingRuleResponseResult
    {
        /// <summary>
        /// The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.
        /// </summary>
        public readonly int BackendPort;
        /// <summary>
        /// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values are between 1 and 65534.
        /// </summary>
        public readonly int FrontendPort;
        /// <summary>
        /// the reference to the load balancer probe used by the load balancing rule.
        /// </summary>
        public readonly string ProbeProtocol;
        /// <summary>
        /// The probe request path. Only supported for HTTP/HTTPS probes.
        /// </summary>
        public readonly string? ProbeRequestPath;
        /// <summary>
        /// The reference to the transport protocol used by the load balancing rule.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private LoadBalancingRuleResponseResult(
            int backendPort,

            int frontendPort,

            string probeProtocol,

            string? probeRequestPath,

            string protocol)
        {
            BackendPort = backendPort;
            FrontendPort = frontendPort;
            ProbeProtocol = probeProtocol;
            ProbeRequestPath = probeRequestPath;
            Protocol = protocol;
        }
    }
}
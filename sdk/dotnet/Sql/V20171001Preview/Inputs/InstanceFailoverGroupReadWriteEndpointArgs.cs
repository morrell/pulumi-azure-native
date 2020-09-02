// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20171001Preview.Inputs
{

    /// <summary>
    /// Read-write endpoint of the failover group instance.
    /// </summary>
    public sealed class InstanceFailoverGroupReadWriteEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Failover policy of the read-write endpoint for the failover group. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required.
        /// </summary>
        [Input("failoverPolicy", required: true)]
        public Input<string> FailoverPolicy { get; set; } = null!;

        /// <summary>
        /// Grace period before failover with data loss is attempted for the read-write endpoint. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required.
        /// </summary>
        [Input("failoverWithDataLossGracePeriodMinutes")]
        public Input<int>? FailoverWithDataLossGracePeriodMinutes { get; set; }

        public InstanceFailoverGroupReadWriteEndpointArgs()
        {
        }
    }
}
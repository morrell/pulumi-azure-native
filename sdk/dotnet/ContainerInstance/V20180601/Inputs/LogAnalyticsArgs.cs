// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.V20180601.Inputs
{

    /// <summary>
    /// Container group log analytics information.
    /// </summary>
    public sealed class LogAnalyticsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The workspace id for log analytics
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        /// <summary>
        /// The workspace key for log analytics
        /// </summary>
        [Input("workspaceKey", required: true)]
        public Input<string> WorkspaceKey { get; set; } = null!;

        public LogAnalyticsArgs()
        {
        }
    }
}
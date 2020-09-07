// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20180901.Inputs
{

    /// <summary>
    /// The trigger based on base image dependency.
    /// </summary>
    public sealed class BaseImageTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the auto trigger for base image dependency updates.
        /// </summary>
        [Input("baseImageTriggerType", required: true)]
        public Input<string> BaseImageTriggerType { get; set; } = null!;

        /// <summary>
        /// The name of the trigger.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The current status of trigger.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BaseImageTriggerArgs()
        {
        }
    }
}
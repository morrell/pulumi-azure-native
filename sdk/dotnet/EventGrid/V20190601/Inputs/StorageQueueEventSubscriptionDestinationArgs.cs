// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20190601.Inputs
{

    /// <summary>
    /// Information about the storage queue destination for an event subscription.
    /// </summary>
    public sealed class StorageQueueEventSubscriptionDestinationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the endpoint for the event subscription destination
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// The name of the Storage queue under a storage account that is the destination of an event subscription.
        /// </summary>
        [Input("queueName")]
        public Input<string>? QueueName { get; set; }

        /// <summary>
        /// The Azure Resource ID of the storage account that contains the queue that is the destination of an event subscription.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public StorageQueueEventSubscriptionDestinationArgs()
        {
        }
    }
}
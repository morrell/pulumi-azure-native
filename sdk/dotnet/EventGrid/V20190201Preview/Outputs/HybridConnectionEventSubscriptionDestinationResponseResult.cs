// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20190201Preview.Outputs
{

    [OutputType]
    public sealed class HybridConnectionEventSubscriptionDestinationResponseResult
    {
        /// <summary>
        /// Type of the endpoint for the event subscription destination
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// The Azure Resource ID of an hybrid connection that is the destination of an event subscription.
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private HybridConnectionEventSubscriptionDestinationResponseResult(
            string endpointType,

            string? resourceId)
        {
            EndpointType = endpointType;
            ResourceId = resourceId;
        }
    }
}
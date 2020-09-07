// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions.Latest.Inputs
{

    /// <summary>
    /// Managed application notification policy.
    /// </summary>
    public sealed class ApplicationNotificationPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("notificationEndpoints", required: true)]
        private InputList<Inputs.ApplicationNotificationEndpointArgs>? _notificationEndpoints;

        /// <summary>
        /// The managed application notification endpoint.
        /// </summary>
        public InputList<Inputs.ApplicationNotificationEndpointArgs> NotificationEndpoints
        {
            get => _notificationEndpoints ?? (_notificationEndpoints = new InputList<Inputs.ApplicationNotificationEndpointArgs>());
            set => _notificationEndpoints = value;
        }

        public ApplicationNotificationPolicyArgs()
        {
        }
    }
}
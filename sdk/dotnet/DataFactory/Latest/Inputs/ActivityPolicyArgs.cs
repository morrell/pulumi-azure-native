// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.Latest.Inputs
{

    /// <summary>
    /// Execution policy for an activity.
    /// </summary>
    public sealed class ActivityPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("retry")]
        private InputMap<object>? _retry;

        /// <summary>
        /// Maximum ordinary retry attempts. Default is 0. Type: integer (or Expression with resultType integer), minimum: 0.
        /// </summary>
        public InputMap<object> Retry
        {
            get => _retry ?? (_retry = new InputMap<object>());
            set => _retry = value;
        }

        /// <summary>
        /// Interval between each retry attempt (in seconds). The default is 30 sec.
        /// </summary>
        [Input("retryIntervalInSeconds")]
        public Input<int>? RetryIntervalInSeconds { get; set; }

        /// <summary>
        /// When set to true, Input from activity is considered as secure and will not be logged to monitoring.
        /// </summary>
        [Input("secureInput")]
        public Input<bool>? SecureInput { get; set; }

        /// <summary>
        /// When set to true, Output from activity is considered as secure and will not be logged to monitoring.
        /// </summary>
        [Input("secureOutput")]
        public Input<bool>? SecureOutput { get; set; }

        [Input("timeout")]
        private InputMap<object>? _timeout;

        /// <summary>
        /// Specifies the timeout for the activity to run. The default timeout is 7 days. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public InputMap<object> Timeout
        {
            get => _timeout ?? (_timeout = new InputMap<object>());
            set => _timeout = value;
        }

        public ActivityPolicyArgs()
        {
        }
    }
}
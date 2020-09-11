// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventHub.V20180101Preview.Outputs
{

    [OutputType]
    public sealed class ConnectionStateResponseResult
    {
        /// <summary>
        /// Description of the connection state.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Status of the connection.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ConnectionStateResponseResult(
            string? description,

            string? status)
        {
            Description = description;
            Status = status;
        }
    }
}
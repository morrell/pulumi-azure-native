// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801Preview.Outputs
{

    [OutputType]
    public sealed class ConnectionStatusResponseResult
    {
        /// <summary>
        /// Error details
        /// </summary>
        public readonly Outputs.ConnectionErrorResponseResult? Error;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Kind of resource
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Status
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Target of the error
        /// </summary>
        public readonly string? Target;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ConnectionStatusResponseResult(
            Outputs.ConnectionErrorResponseResult? error,

            string? id,

            string? kind,

            string location,

            string? name,

            string? status,

            ImmutableDictionary<string, string>? tags,

            string? target,

            string? type)
        {
            Error = error;
            Id = id;
            Kind = kind;
            Location = location;
            Name = name;
            Status = status;
            Tags = tags;
            Target = target;
            Type = type;
        }
    }
}
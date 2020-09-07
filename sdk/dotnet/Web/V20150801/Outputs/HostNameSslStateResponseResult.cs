// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801.Outputs
{

    [OutputType]
    public sealed class HostNameSslStateResponseResult
    {
        /// <summary>
        /// Host name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// SSL type
        /// </summary>
        public readonly string SslState;
        /// <summary>
        /// SSL cert thumbprint
        /// </summary>
        public readonly string? Thumbprint;
        /// <summary>
        /// Set this flag to update existing host name
        /// </summary>
        public readonly bool? ToUpdate;
        /// <summary>
        /// Virtual IP address assigned to the host name if IP based SSL is enabled
        /// </summary>
        public readonly string? VirtualIP;

        [OutputConstructor]
        private HostNameSslStateResponseResult(
            string? name,

            string sslState,

            string? thumbprint,

            bool? toUpdate,

            string? virtualIP)
        {
            Name = name;
            SslState = sslState;
            Thumbprint = thumbprint;
            ToUpdate = toUpdate;
            VirtualIP = virtualIP;
        }
    }
}
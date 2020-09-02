// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20200101Preview.Inputs
{

    /// <summary>
    /// For a non-Azure machine that is not connected directly to the internet, specify a proxy server that the non-Azure machine can use.
    /// </summary>
    public sealed class ProxyServerPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Proxy server IP
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Proxy server port
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        public ProxyServerPropertiesArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HardwareSecurityModules.V20181031Preview.Inputs
{

    /// <summary>
    /// The network interface definition.
    /// </summary>
    public sealed class NetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Private Ip address of the interface
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        public NetworkInterfaceArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HanaOnAzure.V20171103Preview.Inputs
{

    /// <summary>
    /// Specifies the network settings for the HANA instance disks.
    /// </summary>
    public sealed class NetworkProfileArgs : Pulumi.ResourceArgs
    {
        [Input("networkInterfaces")]
        private InputList<Inputs.IpAddressArgs>? _networkInterfaces;

        /// <summary>
        /// Specifies the network interfaces for the HANA instance.
        /// </summary>
        public InputList<Inputs.IpAddressArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.IpAddressArgs>());
            set => _networkInterfaces = value;
        }

        public NetworkProfileArgs()
        {
        }
    }
}
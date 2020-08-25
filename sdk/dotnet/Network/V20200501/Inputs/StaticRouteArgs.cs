// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Inputs
{

    /// <summary>
    /// List of all Static Routes.
    /// </summary>
    public sealed class StaticRouteArgs : Pulumi.ResourceArgs
    {
        [Input("addressPrefixes")]
        private InputList<string>? _addressPrefixes;

        /// <summary>
        /// List of all address prefixes.
        /// </summary>
        public InputList<string> AddressPrefixes
        {
            get => _addressPrefixes ?? (_addressPrefixes = new InputList<string>());
            set => _addressPrefixes = value;
        }

        /// <summary>
        /// The name of the StaticRoute that is unique within a VnetRoute.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ip address of the next hop.
        /// </summary>
        [Input("nextHopIpAddress")]
        public Input<string>? NextHopIpAddress { get; set; }

        public StaticRouteArgs()
        {
        }
    }
}
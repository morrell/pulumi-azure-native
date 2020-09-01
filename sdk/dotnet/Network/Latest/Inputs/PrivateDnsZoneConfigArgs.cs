// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Latest.Inputs
{

    /// <summary>
    /// PrivateDnsZoneConfig resource.
    /// </summary>
    public sealed class PrivateDnsZoneConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource id of the private dns zone.
        /// </summary>
        [Input("privateDnsZoneId")]
        public Input<string>? PrivateDnsZoneId { get; set; }

        public PrivateDnsZoneConfigArgs()
        {
        }
    }
}
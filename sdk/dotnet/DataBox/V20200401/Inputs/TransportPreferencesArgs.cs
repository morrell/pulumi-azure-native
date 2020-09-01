// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBox.V20200401.Inputs
{

    /// <summary>
    /// Preferences related to the shipment logistics of the sku
    /// </summary>
    public sealed class TransportPreferencesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates Shipment Logistics type that the customer preferred.
        /// </summary>
        [Input("preferredShipmentType", required: true)]
        public Input<string> PreferredShipmentType { get; set; } = null!;

        public TransportPreferencesArgs()
        {
        }
    }
}
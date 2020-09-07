// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.Latest.Inputs
{

    /// <summary>
    /// Create network mappings input properties/behavior specific to Azure to Azure Network mapping.
    /// </summary>
    public sealed class AzureToAzureCreateNetworkMappingInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The primary azure vnet Id.
        /// </summary>
        [Input("primaryNetworkId")]
        public Input<string>? PrimaryNetworkId { get; set; }

        public AzureToAzureCreateNetworkMappingInputArgs()
        {
        }
    }
}
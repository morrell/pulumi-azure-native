// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101.Inputs
{

    /// <summary>
    /// AAD Vpn authentication type related parameters.
    /// </summary>
    public sealed class AadAuthenticationParametersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AAD Vpn authentication parameter AAD audience.
        /// </summary>
        [Input("aadAudience")]
        public Input<string>? AadAudience { get; set; }

        /// <summary>
        /// AAD Vpn authentication parameter AAD issuer.
        /// </summary>
        [Input("aadIssuer")]
        public Input<string>? AadIssuer { get; set; }

        /// <summary>
        /// AAD Vpn authentication parameter AAD tenant.
        /// </summary>
        [Input("aadTenant")]
        public Input<string>? AadTenant { get; set; }

        public AadAuthenticationParametersArgs()
        {
        }
    }
}
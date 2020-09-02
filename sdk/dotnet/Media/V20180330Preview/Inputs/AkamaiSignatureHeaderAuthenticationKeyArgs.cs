// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview.Inputs
{

    /// <summary>
    /// Akamai Signature Header authentication key.
    /// </summary>
    public sealed class AkamaiSignatureHeaderAuthenticationKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// authentication key
        /// </summary>
        [Input("base64Key")]
        public Input<string>? Base64Key { get; set; }

        /// <summary>
        /// The exact time the authentication key.
        /// </summary>
        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        /// <summary>
        /// identifier of the key
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        public AkamaiSignatureHeaderAuthenticationKeyArgs()
        {
        }
    }
}
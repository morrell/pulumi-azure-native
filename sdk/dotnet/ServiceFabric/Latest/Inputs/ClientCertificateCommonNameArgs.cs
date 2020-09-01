// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.Latest.Inputs
{

    /// <summary>
    /// Describes the client certificate details using common name.
    /// </summary>
    public sealed class ClientCertificateCommonNameArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The common name of the client certificate.
        /// </summary>
        [Input("certificateCommonName", required: true)]
        public Input<string> CertificateCommonName { get; set; } = null!;

        /// <summary>
        /// The issuer thumbprint of the client certificate.
        /// </summary>
        [Input("certificateIssuerThumbprint", required: true)]
        public Input<string> CertificateIssuerThumbprint { get; set; } = null!;

        /// <summary>
        /// Indicates if the client certificate has admin access to the cluster. Non admin clients can perform only read only operations on the cluster.
        /// </summary>
        [Input("isAdmin", required: true)]
        public Input<bool> IsAdmin { get; set; } = null!;

        public ClientCertificateCommonNameArgs()
        {
        }
    }
}
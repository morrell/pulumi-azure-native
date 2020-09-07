// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20190501Preview.Outputs
{

    [OutputType]
    public sealed class TokenCertificateResponseResult
    {
        /// <summary>
        /// Base 64 encoded string of the public certificate1 in PEM format that will be used for authenticating the token.
        /// </summary>
        public readonly string? EncodedPemCertificate;
        /// <summary>
        /// The expiry datetime of the certificate.
        /// </summary>
        public readonly string? Expiry;
        public readonly string? Name;
        /// <summary>
        /// The thumbprint of the certificate.
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private TokenCertificateResponseResult(
            string? encodedPemCertificate,

            string? expiry,

            string? name,

            string? thumbprint)
        {
            EncodedPemCertificate = encodedPemCertificate;
            Expiry = expiry;
            Name = name;
            Thumbprint = thumbprint;
        }
    }
}
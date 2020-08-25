// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190401.Outputs
{

    [OutputType]
    public sealed class VpnClientRevokedCertificateResponseResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The provisioning state of the VPN client revoked certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The revoked VPN client certificate thumbprint.
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private VpnClientRevokedCertificateResponseResult(
            string? etag,

            string? id,

            string? name,

            string provisioningState,

            string? thumbprint)
        {
            Etag = etag;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            Thumbprint = thumbprint;
        }
    }
}
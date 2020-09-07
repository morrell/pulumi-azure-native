// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kubernetes.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class ConnectedClusterIdentityResponseResult
    {
        /// <summary>
        /// The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ConnectedClusterIdentityResponseResult(
            string principalId,

            string tenantId,

            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
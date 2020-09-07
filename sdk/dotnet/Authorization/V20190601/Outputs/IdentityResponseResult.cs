// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.V20190601.Outputs
{

    [OutputType]
    public sealed class IdentityResponseResult
    {
        /// <summary>
        /// The principal ID of the resource identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant ID of the resource identity.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The identity type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private IdentityResponseResult(
            string principalId,

            string tenantId,

            string? type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
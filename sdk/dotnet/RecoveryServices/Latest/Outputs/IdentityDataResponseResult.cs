// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.Latest.Outputs
{

    [OutputType]
    public sealed class IdentityDataResponseResult
    {
        /// <summary>
        /// The principal ID of resource identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant ID of resource.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The identity type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private IdentityDataResponseResult(
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
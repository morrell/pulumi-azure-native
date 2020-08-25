// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20200215.Outputs
{

    [OutputType]
    public sealed class IdentityResponseResult
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
        /// <summary>
        /// The list of user identities associated with the Kusto cluster. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.IdentityResponseUserAssignedIdentitiesResult>? UserAssignedIdentities;

        [OutputConstructor]
        private IdentityResponseResult(
            string principalId,

            string tenantId,

            string type,

            ImmutableDictionary<string, Outputs.IdentityResponseUserAssignedIdentitiesResult>? userAssignedIdentities)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
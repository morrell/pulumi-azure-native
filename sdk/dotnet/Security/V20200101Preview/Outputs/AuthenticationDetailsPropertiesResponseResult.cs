// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class AuthenticationDetailsPropertiesResponseResult
    {
        /// <summary>
        /// State of the multi-cloud connector
        /// </summary>
        public readonly string AuthenticationProvisioningState;
        /// <summary>
        /// Connect to your cloud account, for AWS use either account credentials or role-based authentication. For GCP use account organization credentials.
        /// </summary>
        public readonly string AuthenticationType;
        /// <summary>
        /// The permissions detected in the cloud account.
        /// </summary>
        public readonly ImmutableArray<string> GrantedPermissions;

        [OutputConstructor]
        private AuthenticationDetailsPropertiesResponseResult(
            string authenticationProvisioningState,

            string authenticationType,

            ImmutableArray<string> grantedPermissions)
        {
            AuthenticationProvisioningState = authenticationProvisioningState;
            AuthenticationType = authenticationType;
            GrantedPermissions = grantedPermissions;
        }
    }
}
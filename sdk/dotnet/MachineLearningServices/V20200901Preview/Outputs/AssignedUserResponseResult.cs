// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200901Preview.Outputs
{

    [OutputType]
    public sealed class AssignedUserResponseResult
    {
        /// <summary>
        /// User’s AAD Object Id.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// User’s AAD Tenant Id.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private AssignedUserResponseResult(
            string objectId,

            string tenantId)
        {
            ObjectId = objectId;
            TenantId = tenantId;
        }
    }
}
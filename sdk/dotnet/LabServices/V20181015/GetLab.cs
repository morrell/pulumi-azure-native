// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.LabServices.V20181015
{
    public static class GetLab
    {
        public static Task<GetLabResult> InvokeAsync(GetLabArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLabResult>("azurerm:labservices/v20181015:getLab", args ?? new GetLabArgs(), options.WithVersion());
    }


    public sealed class GetLabArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the $expand query. Example: 'properties($select=maxUsersInLab)'
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the lab Account.
        /// </summary>
        [Input("labAccountName", required: true)]
        public string LabAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public string LabName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetLabArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLabResult
    {
        /// <summary>
        /// Object id of the user that created the lab.
        /// </summary>
        public readonly string CreatedByObjectId;
        /// <summary>
        /// Lab creator name
        /// </summary>
        public readonly string CreatedByUserPrincipalName;
        /// <summary>
        /// Creation date for the lab
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// Invitation code that users can use to join a lab.
        /// </summary>
        public readonly string InvitationCode;
        /// <summary>
        /// The details of the latest operation. ex: status, error
        /// </summary>
        public readonly Outputs.LatestOperationResultResponseResult LatestOperationResult;
        /// <summary>
        /// The location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Maximum number of users allowed in the lab.
        /// </summary>
        public readonly int? MaxUsersInLab;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        public readonly string? UniqueIdentifier;
        /// <summary>
        /// Maximum duration a user can use an environment for in the lab.
        /// </summary>
        public readonly string? UsageQuota;
        /// <summary>
        /// Lab user access mode (open to all vs. restricted to those listed on the lab).
        /// </summary>
        public readonly string? UserAccessMode;
        /// <summary>
        /// Maximum value MaxUsersInLab can be set to, as specified by the service
        /// </summary>
        public readonly int UserQuota;

        [OutputConstructor]
        private GetLabResult(
            string createdByObjectId,

            string createdByUserPrincipalName,

            string createdDate,

            string invitationCode,

            Outputs.LatestOperationResultResponseResult latestOperationResult,

            string? location,

            int? maxUsersInLab,

            string name,

            string? provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? uniqueIdentifier,

            string? usageQuota,

            string? userAccessMode,

            int userQuota)
        {
            CreatedByObjectId = createdByObjectId;
            CreatedByUserPrincipalName = createdByUserPrincipalName;
            CreatedDate = createdDate;
            InvitationCode = invitationCode;
            LatestOperationResult = latestOperationResult;
            Location = location;
            MaxUsersInLab = maxUsersInLab;
            Name = name;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
            UniqueIdentifier = uniqueIdentifier;
            UsageQuota = usageQuota;
            UserAccessMode = userAccessMode;
            UserQuota = userQuota;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.V20200701Preview
{
    public static class GetPolicyExemption
    {
        public static Task<GetPolicyExemptionResult> InvokeAsync(GetPolicyExemptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyExemptionResult>("azurerm:authorization/v20200701preview:getPolicyExemption", args ?? new GetPolicyExemptionArgs(), options.WithVersion());
    }


    public sealed class GetPolicyExemptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the policy exemption to delete.
        /// </summary>
        [Input("policyExemptionName", required: true)]
        public string PolicyExemptionName { get; set; } = null!;

        /// <summary>
        /// The scope of the policy exemption. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
        /// </summary>
        [Input("scope", required: true)]
        public string Scope { get; set; } = null!;

        public GetPolicyExemptionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyExemptionResult
    {
        /// <summary>
        /// The description of the policy exemption.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The display name of the policy exemption.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The policy exemption category. Possible values are Waiver and Mitigated.
        /// </summary>
        public readonly string ExemptionCategory;
        /// <summary>
        /// The expiration date and time (in UTC ISO 8601 format yyyy-MM-ddTHH:mm:ssZ) of the policy exemption.
        /// </summary>
        public readonly string? ExpiresOn;
        /// <summary>
        /// The policy exemption metadata. Metadata is an open ended object and is typically a collection of key value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Metadata;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the policy assignment that is being exempted.
        /// </summary>
        public readonly string PolicyAssignmentId;
        /// <summary>
        /// The policy definition reference ID list when the associated policy assignment is an assignment of a policy set definition.
        /// </summary>
        public readonly ImmutableArray<string> PolicyDefinitionReferenceIds;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponseResult SystemData;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyExemptionResult(
            string? description,

            string? displayName,

            string exemptionCategory,

            string? expiresOn,

            ImmutableDictionary<string, object>? metadata,

            string name,

            string policyAssignmentId,

            ImmutableArray<string> policyDefinitionReferenceIds,

            Outputs.SystemDataResponseResult systemData,

            string type)
        {
            Description = description;
            DisplayName = displayName;
            ExemptionCategory = exemptionCategory;
            ExpiresOn = expiresOn;
            Metadata = metadata;
            Name = name;
            PolicyAssignmentId = policyAssignmentId;
            PolicyDefinitionReferenceIds = policyDefinitionReferenceIds;
            SystemData = systemData;
            Type = type;
        }
    }
}
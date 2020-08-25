// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.PolicyInsights.V20190701
{
    public static class GetRemediationAtResourceGroup
    {
        public static Task<GetRemediationAtResourceGroupResult> InvokeAsync(GetRemediationAtResourceGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRemediationAtResourceGroupResult>("azurerm:policyinsights/v20190701:getRemediationAtResourceGroup", args ?? new GetRemediationAtResourceGroupArgs(), options.WithVersion());
    }


    public sealed class GetRemediationAtResourceGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the remediation.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRemediationAtResourceGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRemediationAtResourceGroupResult
    {
        /// <summary>
        /// The time at which the remediation was created.
        /// </summary>
        public readonly string CreatedOn;
        /// <summary>
        /// The deployment status summary for all deployments created by the remediation.
        /// </summary>
        public readonly Outputs.RemediationDeploymentSummaryResponseResult DeploymentStatus;
        /// <summary>
        /// The filters that will be applied to determine which resources to remediate.
        /// </summary>
        public readonly Outputs.RemediationFiltersResponseResult? Filters;
        /// <summary>
        /// The time at which the remediation was last updated.
        /// </summary>
        public readonly string LastUpdatedOn;
        /// <summary>
        /// The name of the remediation.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource ID of the policy assignment that should be remediated.
        /// </summary>
        public readonly string? PolicyAssignmentId;
        /// <summary>
        /// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        /// </summary>
        public readonly string? PolicyDefinitionReferenceId;
        /// <summary>
        /// The status of the remediation.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
        /// </summary>
        public readonly string? ResourceDiscoveryMode;
        /// <summary>
        /// The type of the remediation.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRemediationAtResourceGroupResult(
            string createdOn,

            Outputs.RemediationDeploymentSummaryResponseResult deploymentStatus,

            Outputs.RemediationFiltersResponseResult? filters,

            string lastUpdatedOn,

            string name,

            string? policyAssignmentId,

            string? policyDefinitionReferenceId,

            string provisioningState,

            string? resourceDiscoveryMode,

            string type)
        {
            CreatedOn = createdOn;
            DeploymentStatus = deploymentStatus;
            Filters = filters;
            LastUpdatedOn = lastUpdatedOn;
            Name = name;
            PolicyAssignmentId = policyAssignmentId;
            PolicyDefinitionReferenceId = policyDefinitionReferenceId;
            ProvisioningState = provisioningState;
            ResourceDiscoveryMode = resourceDiscoveryMode;
            Type = type;
        }
    }
}
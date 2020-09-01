// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.Latest
{
    public static class GetRoleAssignment
    {
        public static Task<GetRoleAssignmentResult> InvokeAsync(GetRoleAssignmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoleAssignmentResult>("azurerm:authorization/latest:getRoleAssignment", args ?? new GetRoleAssignmentArgs(), options.WithVersion());
    }


    public sealed class GetRoleAssignmentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the role assignment to get.
        /// </summary>
        [Input("roleAssignmentName", required: true)]
        public string RoleAssignmentName { get; set; } = null!;

        /// <summary>
        /// The scope of the role assignment.
        /// </summary>
        [Input("scope", required: true)]
        public string Scope { get; set; } = null!;

        public GetRoleAssignmentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRoleAssignmentResult
    {
        /// <summary>
        /// The role assignment name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Role assignment properties.
        /// </summary>
        public readonly Outputs.RoleAssignmentPropertiesWithScopeResponseResult Properties;
        /// <summary>
        /// The role assignment type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRoleAssignmentResult(
            string name,

            Outputs.RoleAssignmentPropertiesWithScopeResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Management.V20180301Preview.Outputs
{

    [OutputType]
    public sealed class ManagementGroupChildInfoResponseResult
    {
        /// <summary>
        /// The list of children.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagementGroupChildInfoResponseResult> Children;
        /// <summary>
        /// The friendly name of the child resource.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The fully qualified ID for the child resource (management group or subscription).  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the child entity.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The roles definitions associated with the management group.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// The fully qualified resource type which includes provider namespace (e.g. /providers/Microsoft.Management/managementGroups)
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ManagementGroupChildInfoResponseResult(
            ImmutableArray<Outputs.ManagementGroupChildInfoResponseResult> children,

            string? displayName,

            string? id,

            string? name,

            ImmutableArray<string> roles,

            string? type)
        {
            Children = children;
            DisplayName = displayName;
            Id = id;
            Name = name;
            Roles = roles;
            Type = type;
        }
    }
}
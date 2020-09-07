// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Management.Latest.Outputs
{

    [OutputType]
    public sealed class ManagementGroupPathElementResponseResult
    {
        /// <summary>
        /// The friendly name of the group.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The name of the group.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private ManagementGroupPathElementResponseResult(
            string? displayName,

            string? name)
        {
            DisplayName = displayName;
            Name = name;
        }
    }
}
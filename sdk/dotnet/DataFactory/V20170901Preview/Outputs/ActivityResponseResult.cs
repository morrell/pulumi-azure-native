// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class ActivityResponseResult
    {
        /// <summary>
        /// Activity depends on condition.
        /// </summary>
        public readonly ImmutableArray<Outputs.ActivityDependencyResponseResult> DependsOn;
        /// <summary>
        /// Activity description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Activity name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of activity.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ActivityResponseResult(
            ImmutableArray<Outputs.ActivityDependencyResponseResult> dependsOn,

            string? description,

            string name,

            string type)
        {
            DependsOn = dependsOn;
            Description = description;
            Name = name;
            Type = type;
        }
    }
}
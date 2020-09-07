// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20180601.Outputs
{

    [OutputType]
    public sealed class ExecutionActivityResponseResult
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
        /// Linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponseResult? LinkedServiceName;
        /// <summary>
        /// Activity name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Activity policy.
        /// </summary>
        public readonly Outputs.ActivityPolicyResponseResult? Policy;
        /// <summary>
        /// Type of activity.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Activity user properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserPropertyResponseResult> UserProperties;

        [OutputConstructor]
        private ExecutionActivityResponseResult(
            ImmutableArray<Outputs.ActivityDependencyResponseResult> dependsOn,

            string? description,

            Outputs.LinkedServiceReferenceResponseResult? linkedServiceName,

            string name,

            Outputs.ActivityPolicyResponseResult? policy,

            string type,

            ImmutableArray<Outputs.UserPropertyResponseResult> userProperties)
        {
            DependsOn = dependsOn;
            Description = description;
            LinkedServiceName = linkedServiceName;
            Name = name;
            Policy = policy;
            Type = type;
            UserProperties = userProperties;
        }
    }
}
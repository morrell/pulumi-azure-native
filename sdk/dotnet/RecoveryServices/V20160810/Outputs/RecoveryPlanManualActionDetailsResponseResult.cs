// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20160810.Outputs
{

    [OutputType]
    public sealed class RecoveryPlanManualActionDetailsResponseResult
    {
        /// <summary>
        /// The manual action description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets the type of action details (see RecoveryPlanActionDetailsTypes enum for possible values).
        /// </summary>
        public readonly string InstanceType;

        [OutputConstructor]
        private RecoveryPlanManualActionDetailsResponseResult(
            string? description,

            string instanceType)
        {
            Description = description;
            InstanceType = instanceType;
        }
    }
}
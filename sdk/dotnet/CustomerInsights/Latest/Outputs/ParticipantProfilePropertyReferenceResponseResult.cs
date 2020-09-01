// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.Latest.Outputs
{

    [OutputType]
    public sealed class ParticipantProfilePropertyReferenceResponseResult
    {
        /// <summary>
        /// The source interaction property that maps to the target profile property.
        /// </summary>
        public readonly string InteractionPropertyName;
        /// <summary>
        /// The target profile property that maps to the source interaction property.
        /// </summary>
        public readonly string ProfilePropertyName;

        [OutputConstructor]
        private ParticipantProfilePropertyReferenceResponseResult(
            string interactionPropertyName,

            string profilePropertyName)
        {
            InteractionPropertyName = interactionPropertyName;
            ProfilePropertyName = profilePropertyName;
        }
    }
}
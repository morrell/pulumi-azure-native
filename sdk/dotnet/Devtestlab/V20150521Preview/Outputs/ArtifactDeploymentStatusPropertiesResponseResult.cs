// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20150521Preview.Outputs
{

    [OutputType]
    public sealed class ArtifactDeploymentStatusPropertiesResponseResult
    {
        /// <summary>
        /// The total count of the artifacts that were successfully applied.
        /// </summary>
        public readonly int? ArtifactsApplied;
        /// <summary>
        /// The deployment status of the artifact.
        /// </summary>
        public readonly string? DeploymentStatus;
        /// <summary>
        /// The total count of the artifacts that were tentatively applied.
        /// </summary>
        public readonly int? TotalArtifacts;

        [OutputConstructor]
        private ArtifactDeploymentStatusPropertiesResponseResult(
            int? artifactsApplied,

            string? deploymentStatus,

            int? totalArtifacts)
        {
            ArtifactsApplied = artifactsApplied;
            DeploymentStatus = deploymentStatus;
            TotalArtifacts = totalArtifacts;
        }
    }
}
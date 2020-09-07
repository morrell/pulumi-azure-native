// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.VirtualMachineImages.Latest.Outputs
{

    [OutputType]
    public sealed class ImageTemplateVhdDistributorResponseResult
    {
        /// <summary>
        /// Tags that will be applied to the artifact once it has been created/updated by the distributor.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ArtifactTags;
        /// <summary>
        /// The name to be used for the associated RunOutput.
        /// </summary>
        public readonly string RunOutputName;
        /// <summary>
        /// Type of distribution.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ImageTemplateVhdDistributorResponseResult(
            ImmutableDictionary<string, string>? artifactTags,

            string runOutputName,

            string type)
        {
            ArtifactTags = artifactTags;
            RunOutputName = runOutputName;
            Type = type;
        }
    }
}
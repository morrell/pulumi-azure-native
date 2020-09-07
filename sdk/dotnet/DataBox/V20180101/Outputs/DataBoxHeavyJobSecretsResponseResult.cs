// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBox.V20180101.Outputs
{

    [OutputType]
    public sealed class DataBoxHeavyJobSecretsResponseResult
    {
        /// <summary>
        /// Contains the list of secret objects for a DataBoxHeavy job.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataBoxHeavySecretResponseResult> CabinetPodSecrets;
        /// <summary>
        /// Used to indicate what type of job secrets object.
        /// </summary>
        public readonly string JobSecretsType;

        [OutputConstructor]
        private DataBoxHeavyJobSecretsResponseResult(
            ImmutableArray<Outputs.DataBoxHeavySecretResponseResult> cabinetPodSecrets,

            string jobSecretsType)
        {
            CabinetPodSecrets = cabinetPodSecrets;
            JobSecretsType = jobSecretsType;
        }
    }
}
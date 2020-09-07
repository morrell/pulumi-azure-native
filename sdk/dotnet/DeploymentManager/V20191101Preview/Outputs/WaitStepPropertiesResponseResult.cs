// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DeploymentManager.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class WaitStepPropertiesResponseResult
    {
        /// <summary>
        /// The Wait attributes
        /// </summary>
        public readonly Outputs.WaitStepAttributesResponseResult Attributes;
        /// <summary>
        /// The type of step.
        /// </summary>
        public readonly string StepType;

        [OutputConstructor]
        private WaitStepPropertiesResponseResult(
            Outputs.WaitStepAttributesResponseResult attributes,

            string stepType)
        {
            Attributes = attributes;
            StepType = stepType;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DeploymentManager.V20191101Preview.Inputs
{

    /// <summary>
    /// The properties of a step resource.
    /// </summary>
    public sealed class StepPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of step.
        /// </summary>
        [Input("stepType", required: true)]
        public Input<string> StepType { get; set; } = null!;

        public StepPropertiesArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.Latest.Inputs
{

    /// <summary>
    /// Trigger referenced dependency.
    /// </summary>
    public sealed class TriggerDependencyReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Referenced trigger.
        /// </summary>
        [Input("referenceTrigger", required: true)]
        public Input<Inputs.TriggerReferenceArgs> ReferenceTrigger { get; set; } = null!;

        /// <summary>
        /// The type of dependency reference.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TriggerDependencyReferenceArgs()
        {
        }
    }
}
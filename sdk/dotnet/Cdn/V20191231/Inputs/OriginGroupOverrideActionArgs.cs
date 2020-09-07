// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20191231.Inputs
{

    /// <summary>
    /// Defines the Origin Group override action for the delivery rule.
    /// </summary>
    public sealed class OriginGroupOverrideActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the action for the delivery rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Defines the parameters for the action.
        /// </summary>
        [Input("parameters", required: true)]
        public Input<Inputs.OriginGroupOverrideActionParametersArgs> Parameters { get; set; } = null!;

        public OriginGroupOverrideActionArgs()
        {
        }
    }
}
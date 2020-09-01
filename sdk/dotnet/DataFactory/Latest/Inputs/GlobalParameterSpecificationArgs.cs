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
    /// Definition of a single parameter for an entity.
    /// </summary>
    public sealed class GlobalParameterSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Global Parameter type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value", required: true)]
        private InputMap<object>? _value;

        /// <summary>
        /// Value of parameter.
        /// </summary>
        public InputMap<object> Value
        {
            get => _value ?? (_value = new InputMap<object>());
            set => _value = value;
        }

        public GlobalParameterSpecificationArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HDInsight.Latest.Inputs
{

    /// <summary>
    /// Describes a script action on a running cluster.
    /// </summary>
    public sealed class RuntimeScriptActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the script action.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The parameters for the script
        /// </summary>
        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// The list of roles where script will be executed.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// The URI to the script.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public RuntimeScriptActionArgs()
        {
        }
    }
}
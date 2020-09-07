// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20190601Preview.Inputs
{

    /// <summary>
    /// The parameters for a quick task run request.
    /// </summary>
    public sealed class EncodedTaskRunRequestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The machine configuration of the run agent.
        /// </summary>
        [Input("agentConfiguration")]
        public Input<Inputs.AgentPropertiesArgs>? AgentConfiguration { get; set; }

        /// <summary>
        /// The dedicated agent pool for the run.
        /// </summary>
        [Input("agentPoolName")]
        public Input<string>? AgentPoolName { get; set; }

        /// <summary>
        /// The properties that describes a set of credentials that will be used when this run is invoked.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.CredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// Base64 encoded value of the template/definition file content.
        /// </summary>
        [Input("encodedTaskContent", required: true)]
        public Input<string> EncodedTaskContent { get; set; } = null!;

        /// <summary>
        /// Base64 encoded value of the parameters/values file content.
        /// </summary>
        [Input("encodedValuesContent")]
        public Input<string>? EncodedValuesContent { get; set; }

        /// <summary>
        /// The value that indicates whether archiving is enabled for the run or not.
        /// </summary>
        [Input("isArchiveEnabled")]
        public Input<bool>? IsArchiveEnabled { get; set; }

        /// <summary>
        /// The platform properties against which the run has to happen.
        /// </summary>
        [Input("platform", required: true)]
        public Input<Inputs.PlatformPropertiesArgs> Platform { get; set; } = null!;

        /// <summary>
        /// The URL(absolute or relative) of the source context. It can be an URL to a tar or git repository.
        /// If it is relative URL, the relative path should be obtained from calling listBuildSourceUploadUrl API.
        /// </summary>
        [Input("sourceLocation")]
        public Input<string>? SourceLocation { get; set; }

        /// <summary>
        /// Run timeout in seconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The type of the run request.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("values")]
        private InputList<Inputs.SetValueArgs>? _values;

        /// <summary>
        /// The collection of overridable values that can be passed when running a task.
        /// </summary>
        public InputList<Inputs.SetValueArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.SetValueArgs>());
            set => _values = value;
        }

        public EncodedTaskRunRequestArgs()
        {
        }
    }
}
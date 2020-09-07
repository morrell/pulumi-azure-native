// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20180601
{
    /// <summary>
    /// Pipeline resource type.
    /// </summary>
    public partial class Pipeline : Pulumi.CustomResource
    {
        /// <summary>
        /// List of activities in pipeline.
        /// </summary>
        [Output("activities")]
        public Output<ImmutableArray<Union<Outputs.ControlActivityResponseResult, Outputs.ExecutionActivityResponseResult>>> Activities { get; private set; } = null!;

        /// <summary>
        /// List of tags that can be used for describing the Pipeline.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableArray<ImmutableDictionary<string, object>>> Annotations { get; private set; } = null!;

        /// <summary>
        /// The max number of concurrent runs for the pipeline.
        /// </summary>
        [Output("concurrency")]
        public Output<int?> Concurrency { get; private set; } = null!;

        /// <summary>
        /// The description of the pipeline.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Etag identifies change in the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
        /// </summary>
        [Output("folder")]
        public Output<Outputs.PipelineResponseFolderResult?> Folder { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of parameters for pipeline.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// Dimensions emitted by Pipeline.
        /// </summary>
        [Output("runDimensions")]
        public Output<ImmutableDictionary<string, ImmutableDictionary<string, object>>?> RunDimensions { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// List of variables for pipeline.
        /// </summary>
        [Output("variables")]
        public Output<ImmutableDictionary<string, Outputs.VariableSpecificationResponseResult>?> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a Pipeline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pipeline(string name, PipelineArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datafactory/v20180601:Pipeline", name, args ?? new PipelineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pipeline(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datafactory/v20180601:Pipeline", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:datafactory/latest:Pipeline"},
                    new Pulumi.Alias { Type = "azurerm:datafactory/v20170901preview:Pipeline"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Pipeline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pipeline Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Pipeline(name, id, options);
        }
    }

    public sealed class PipelineArgs : Pulumi.ResourceArgs
    {
        [Input("activities")]
        private InputList<Union<Inputs.ControlActivityArgs, Inputs.ExecutionActivityArgs>>? _activities;

        /// <summary>
        /// List of activities in pipeline.
        /// </summary>
        public InputList<Union<Inputs.ControlActivityArgs, Inputs.ExecutionActivityArgs>> Activities
        {
            get => _activities ?? (_activities = new InputList<Union<Inputs.ControlActivityArgs, Inputs.ExecutionActivityArgs>>());
            set => _activities = value;
        }

        [Input("annotations")]
        private InputList<ImmutableDictionary<string, object>>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the Pipeline.
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<ImmutableDictionary<string, object>>());
            set => _annotations = value;
        }

        /// <summary>
        /// The max number of concurrent runs for the pipeline.
        /// </summary>
        [Input("concurrency")]
        public Input<int>? Concurrency { get; set; }

        /// <summary>
        /// The description of the pipeline.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The factory name.
        /// </summary>
        [Input("factoryName", required: true)]
        public Input<string> FactoryName { get; set; } = null!;

        /// <summary>
        /// The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
        /// </summary>
        [Input("folder")]
        public Input<Inputs.PipelineFolderArgs>? Folder { get; set; }

        [Input("parameters")]
        private InputMap<Inputs.ParameterSpecificationArgs>? _parameters;

        /// <summary>
        /// List of parameters for pipeline.
        /// </summary>
        public InputMap<Inputs.ParameterSpecificationArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<Inputs.ParameterSpecificationArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The pipeline name.
        /// </summary>
        [Input("pipelineName", required: true)]
        public Input<string> PipelineName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("runDimensions")]
        private InputMap<ImmutableDictionary<string, object>>? _runDimensions;

        /// <summary>
        /// Dimensions emitted by Pipeline.
        /// </summary>
        public InputMap<ImmutableDictionary<string, object>> RunDimensions
        {
            get => _runDimensions ?? (_runDimensions = new InputMap<ImmutableDictionary<string, object>>());
            set => _runDimensions = value;
        }

        [Input("variables")]
        private InputMap<Inputs.VariableSpecificationArgs>? _variables;

        /// <summary>
        /// List of variables for pipeline.
        /// </summary>
        public InputMap<Inputs.VariableSpecificationArgs> Variables
        {
            get => _variables ?? (_variables = new InputMap<Inputs.VariableSpecificationArgs>());
            set => _variables = value;
        }

        public PipelineArgs()
        {
        }
    }
}
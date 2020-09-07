// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200515Preview.Outputs
{

    [OutputType]
    public sealed class EnvironmentImageResponseResponseEnvironmentResult
    {
        /// <summary>
        /// The definition of a Docker container.
        /// </summary>
        public readonly Outputs.ModelEnvironmentDefinitionResponseResponseDockerResult? Docker;
        /// <summary>
        /// Definition of environment variables to be defined in the environment.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? EnvironmentVariables;
        /// <summary>
        /// The inferencing stack version added to the image. To avoid adding an inferencing stack, do not set this value. Valid values: "latest".
        /// </summary>
        public readonly string? InferencingStackVersion;
        /// <summary>
        /// The name of the environment.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Settings for a Python environment.
        /// </summary>
        public readonly Outputs.ModelEnvironmentDefinitionResponseResponsePythonResult? Python;
        /// <summary>
        /// Settings for a R environment.
        /// </summary>
        public readonly Outputs.ModelEnvironmentDefinitionResponseResponseRResult? R;
        /// <summary>
        /// The configuration for a Spark environment.
        /// </summary>
        public readonly Outputs.ModelEnvironmentDefinitionResponseResponseSparkResult? Spark;
        /// <summary>
        /// The environment version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private EnvironmentImageResponseResponseEnvironmentResult(
            Outputs.ModelEnvironmentDefinitionResponseResponseDockerResult? docker,

            ImmutableDictionary<string, string>? environmentVariables,

            string? inferencingStackVersion,

            string? name,

            Outputs.ModelEnvironmentDefinitionResponseResponsePythonResult? python,

            Outputs.ModelEnvironmentDefinitionResponseResponseRResult? r,

            Outputs.ModelEnvironmentDefinitionResponseResponseSparkResult? spark,

            string? version)
        {
            Docker = docker;
            EnvironmentVariables = environmentVariables;
            InferencingStackVersion = inferencingStackVersion;
            Name = name;
            Python = python;
            R = r;
            Spark = spark;
            Version = version;
        }
    }
}
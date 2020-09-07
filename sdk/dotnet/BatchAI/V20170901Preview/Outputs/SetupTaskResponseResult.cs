// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class SetupTaskResponseResult
    {
        public readonly string CommandLine;
        public readonly ImmutableArray<Outputs.EnvironmentSettingResponseResult> EnvironmentVariables;
        public readonly bool? RunElevated;
        /// <summary>
        /// The path where the Batch AI service will upload the stdout and stderror of setup task.
        /// </summary>
        public readonly string StdOutErrPathPrefix;

        [OutputConstructor]
        private SetupTaskResponseResult(
            string commandLine,

            ImmutableArray<Outputs.EnvironmentSettingResponseResult> environmentVariables,

            bool? runElevated,

            string stdOutErrPathPrefix)
        {
            CommandLine = commandLine;
            EnvironmentVariables = environmentVariables;
            RunElevated = runElevated;
            StdOutErrPathPrefix = stdOutErrPathPrefix;
        }
    }
}
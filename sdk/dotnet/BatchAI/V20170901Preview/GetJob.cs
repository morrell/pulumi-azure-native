// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20170901Preview
{
    public static class GetJob
    {
        public static Task<GetJobResult> InvokeAsync(GetJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("azurerm:batchai/v20170901preview:getJob", args ?? new GetJobArgs(), options.WithVersion());
    }


    public sealed class GetJobArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        /// </summary>
        [Input("jobName", required: true)]
        public string JobName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetJobArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobResult
    {
        /// <summary>
        /// Specifies the settings for Caffe job.
        /// </summary>
        public readonly Outputs.CaffeSettingsResponseResult? CaffeSettings;
        /// <summary>
        /// Specifies the settings for Chainer job.
        /// </summary>
        public readonly Outputs.ChainerSettingsResponseResult? ChainerSettings;
        /// <summary>
        /// Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
        /// </summary>
        public readonly Outputs.ResourceIdResponseResult? Cluster;
        /// <summary>
        /// Specifies the settings for CNTK (aka Microsoft Cognitive Toolkit) job.
        /// </summary>
        public readonly Outputs.CNTKsettingsResponseResult? CntkSettings;
        /// <summary>
        /// Constraints associated with the Job.
        /// </summary>
        public readonly Outputs.JobPropertiesResponseConstraintsResult? Constraints;
        /// <summary>
        /// If the container was downloaded as part of cluster setup then the same container image will be used. If not provided, the job will run on the VM.
        /// </summary>
        public readonly Outputs.ContainerSettingsResponseResult? ContainerSettings;
        /// <summary>
        /// The creation time of the job.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Specifies the settings for a custom tool kit job.
        /// </summary>
        public readonly Outputs.CustomToolkitSettingsResponseResult? CustomToolkitSettings;
        /// <summary>
        /// Batch AI services sets the following environment variables for all jobs: AZ_BATCHAI_INPUT_id, AZ_BATCHAI_OUTPUT_id, AZ_BATCHAI_NUM_GPUS_PER_NODE, For distributed TensorFlow jobs, following additional environment variables are set by the Batch AI Service: AZ_BATCHAI_PS_HOSTS, AZ_BATCHAI_WORKER_HOSTS.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnvironmentSettingResponseResult> EnvironmentVariables;
        /// <summary>
        /// Contains information about the execution of a job in the Azure Batch service.
        /// </summary>
        public readonly Outputs.JobPropertiesResponseExecutionInfoResult? ExecutionInfo;
        /// <summary>
        /// The current state of the job. Possible values are: queued - The job is queued and able to run. A job enters this state when it is created, or when it is awaiting a retry after a failed run. running - The job is running on a compute cluster. This includes job-level preparation such as downloading resource files or set up container specified on the job - it does not necessarily mean that the job command line has started executing. terminating - The job is terminated by the user, the terminate operation is in progress. succeeded - The job has completed running successfully and exited with exit code 0. failed - The job has finished unsuccessfully (failed with a non-zero exit code) and has exhausted its retry limit. A job is also marked as failed if an error occurred launching the job.
        /// </summary>
        public readonly string? ExecutionState;
        /// <summary>
        /// The time at which the job entered its current execution state.
        /// </summary>
        public readonly string ExecutionStateTransitionTime;
        /// <summary>
        /// Describe the experiment information of the job
        /// </summary>
        public readonly string? ExperimentName;
        public readonly ImmutableArray<Outputs.InputDirectoryResponseResult> InputDirectories;
        /// <summary>
        /// The specified actions will run on all the nodes that are part of the job
        /// </summary>
        public readonly Outputs.JobPreparationResponseResult? JobPreparation;
        /// <summary>
        /// The location of the resource
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The job will be gang scheduled on that many compute nodes
        /// </summary>
        public readonly int? NodeCount;
        public readonly ImmutableArray<Outputs.OutputDirectoryResponseResult> OutputDirectories;
        /// <summary>
        /// Priority associated with the job. Priority values can range from -1000 to 1000, with -1000 being the lowest priority and 1000 being the highest priority. The default value is 0.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The provisioned state of the Batch AI job
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The time at which the job entered its current provisioning state.
        /// </summary>
        public readonly string ProvisioningStateTransitionTime;
        /// <summary>
        /// The path where the Batch AI service will upload stdout and stderror of the job.
        /// </summary>
        public readonly string? StdOutErrPathPrefix;
        /// <summary>
        /// The tags of the resource
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Specifies the settings for TensorFlow job.
        /// </summary>
        public readonly Outputs.TensorFlowSettingsResponseResult? TensorFlowSettings;
        /// <summary>
        /// Possible values are: cntk, tensorflow, caffe, caffe2, chainer, custom.
        /// </summary>
        public readonly string? ToolType;
        /// <summary>
        /// The type of the resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetJobResult(
            Outputs.CaffeSettingsResponseResult? caffeSettings,

            Outputs.ChainerSettingsResponseResult? chainerSettings,

            Outputs.ResourceIdResponseResult? cluster,

            Outputs.CNTKsettingsResponseResult? cntkSettings,

            Outputs.JobPropertiesResponseConstraintsResult? constraints,

            Outputs.ContainerSettingsResponseResult? containerSettings,

            string creationTime,

            Outputs.CustomToolkitSettingsResponseResult? customToolkitSettings,

            ImmutableArray<Outputs.EnvironmentSettingResponseResult> environmentVariables,

            Outputs.JobPropertiesResponseExecutionInfoResult? executionInfo,

            string? executionState,

            string executionStateTransitionTime,

            string? experimentName,

            ImmutableArray<Outputs.InputDirectoryResponseResult> inputDirectories,

            Outputs.JobPreparationResponseResult? jobPreparation,

            string location,

            string name,

            int? nodeCount,

            ImmutableArray<Outputs.OutputDirectoryResponseResult> outputDirectories,

            int? priority,

            string provisioningState,

            string provisioningStateTransitionTime,

            string? stdOutErrPathPrefix,

            ImmutableDictionary<string, string> tags,

            Outputs.TensorFlowSettingsResponseResult? tensorFlowSettings,

            string? toolType,

            string type)
        {
            CaffeSettings = caffeSettings;
            ChainerSettings = chainerSettings;
            Cluster = cluster;
            CntkSettings = cntkSettings;
            Constraints = constraints;
            ContainerSettings = containerSettings;
            CreationTime = creationTime;
            CustomToolkitSettings = customToolkitSettings;
            EnvironmentVariables = environmentVariables;
            ExecutionInfo = executionInfo;
            ExecutionState = executionState;
            ExecutionStateTransitionTime = executionStateTransitionTime;
            ExperimentName = experimentName;
            InputDirectories = inputDirectories;
            JobPreparation = jobPreparation;
            Location = location;
            Name = name;
            NodeCount = nodeCount;
            OutputDirectories = outputDirectories;
            Priority = priority;
            ProvisioningState = provisioningState;
            ProvisioningStateTransitionTime = provisioningStateTransitionTime;
            StdOutErrPathPrefix = stdOutErrPathPrefix;
            Tags = tags;
            TensorFlowSettings = tensorFlowSettings;
            ToolType = toolType;
            Type = type;
        }
    }
}
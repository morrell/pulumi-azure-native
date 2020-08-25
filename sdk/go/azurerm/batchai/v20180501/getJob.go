// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azurerm:batchai/v20180501:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	// The name of the experiment. Experiment names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	ExperimentName string `pulumi:"experimentName"`
	// The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Information about a Job.
type LookupJobResult struct {
	// Caffe2 job settings.
	Caffe2Settings *Caffe2SettingsResponse `pulumi:"caffe2Settings"`
	// Caffe job settings.
	CaffeSettings *CaffeSettingsResponse `pulumi:"caffeSettings"`
	// Chainer job settings.
	ChainerSettings *ChainerSettingsResponse `pulumi:"chainerSettings"`
	// Resource ID of the cluster associated with the job.
	Cluster *ResourceIdResponse `pulumi:"cluster"`
	// CNTK (aka Microsoft Cognitive Toolkit) job settings.
	CntkSettings *CNTKsettingsResponse `pulumi:"cntkSettings"`
	// Constraints associated with the Job.
	Constraints *JobPropertiesResponseConstraints `pulumi:"constraints"`
	// If the container was downloaded as part of cluster setup then the same container image will be used. If not provided, the job will run on the VM.
	ContainerSettings *ContainerSettingsResponse `pulumi:"containerSettings"`
	// The creation time of the job.
	CreationTime string `pulumi:"creationTime"`
	// Custom MPI job settings.
	CustomMpiSettings *CustomMpiSettingsResponse `pulumi:"customMpiSettings"`
	// Custom tool kit job settings.
	CustomToolkitSettings *CustomToolkitSettingsResponse `pulumi:"customToolkitSettings"`
	// A collection of user defined environment variables to be setup for the job.
	EnvironmentVariables []EnvironmentVariableResponse `pulumi:"environmentVariables"`
	// Information about the execution of a job.
	ExecutionInfo *JobPropertiesResponseExecutionInfo `pulumi:"executionInfo"`
	// The current state of the job. Possible values are: queued - The job is queued and able to run. A job enters this state when it is created, or when it is awaiting a retry after a failed run. running - The job is running on a compute cluster. This includes job-level preparation such as downloading resource files or set up container specified on the job - it does not necessarily mean that the job command line has started executing. terminating - The job is terminated by the user, the terminate operation is in progress. succeeded - The job has completed running successfully and exited with exit code 0. failed - The job has finished unsuccessfully (failed with a non-zero exit code) and has exhausted its retry limit. A job is also marked as failed if an error occurred launching the job.
	ExecutionState string `pulumi:"executionState"`
	// The time at which the job entered its current execution state.
	ExecutionStateTransitionTime string `pulumi:"executionStateTransitionTime"`
	// Specifies the settings for Horovod job.
	HorovodSettings *HorovodSettingsResponse `pulumi:"horovodSettings"`
	// A list of input directories for the job.
	InputDirectories []InputDirectoryResponse `pulumi:"inputDirectories"`
	// A segment of job's output directories path created by Batch AI. Batch AI creates job's output directories under an unique path to avoid conflicts between jobs. This value contains a path segment generated by Batch AI to make the path unique and can be used to find the output directory on the node or mounted filesystem.
	JobOutputDirectoryPathSegment string `pulumi:"jobOutputDirectoryPathSegment"`
	// The specified actions will run on all the nodes that are part of the job
	JobPreparation *JobPreparationResponse `pulumi:"jobPreparation"`
	// Collection of mount volumes available to the job during execution. These volumes are mounted before the job execution and unmounted after the job completion. The volumes are mounted at location specified by $AZ_BATCHAI_JOB_MOUNT_ROOT environment variable.
	MountVolumes *MountVolumesResponse `pulumi:"mountVolumes"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The job will be gang scheduled on that many compute nodes
	NodeCount *int `pulumi:"nodeCount"`
	// A list of output directories for the job.
	OutputDirectories []OutputDirectoryResponse `pulumi:"outputDirectories"`
	// The provisioned state of the Batch AI job
	ProvisioningState string `pulumi:"provisioningState"`
	// The time at which the job entered its current provisioning state.
	ProvisioningStateTransitionTime string `pulumi:"provisioningStateTransitionTime"`
	// pyTorch job settings.
	PyTorchSettings *PyTorchSettingsResponse `pulumi:"pyTorchSettings"`
	// Scheduling priority associated with the job.
	SchedulingPriority *string `pulumi:"schedulingPriority"`
	// A collection of user defined environment variables with secret values to be setup for the job. Server will never report values of these variables back.
	Secrets []EnvironmentVariableWithSecretValueResponse `pulumi:"secrets"`
	// The path where the Batch AI service stores stdout, stderror and execution log of the job.
	StdOutErrPathPrefix *string `pulumi:"stdOutErrPathPrefix"`
	// TensorFlow job settings.
	TensorFlowSettings *TensorFlowSettingsResponse `pulumi:"tensorFlowSettings"`
	// Possible values are: cntk, tensorflow, caffe, caffe2, chainer, pytorch, custom, custommpi, horovod.
	ToolType *string `pulumi:"toolType"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
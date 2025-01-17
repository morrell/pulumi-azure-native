// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version: 2021-03-01-preview.
type BatchDeployment struct {
	pulumi.CustomResourceState

	// Service identity associated with a resource.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind     pulumi.StringPtrOutput            `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Additional attributes of the entity.
	Properties BatchDeploymentResponseOutput `pulumi:"properties"`
	// System data associated with resource provider
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBatchDeployment registers a new resource with the given unique name, arguments, and options.
func NewBatchDeployment(ctx *pulumi.Context,
	name string, args *BatchDeploymentArgs, opts ...pulumi.ResourceOption) (*BatchDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:BatchDeployment"),
		},
	})
	opts = append(opts, aliases)
	var resource BatchDeployment
	err := ctx.RegisterResource("azure-native:machinelearningservices:BatchDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBatchDeployment gets an existing BatchDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBatchDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchDeploymentState, opts ...pulumi.ResourceOption) (*BatchDeployment, error) {
	var resource BatchDeployment
	err := ctx.ReadResource("azure-native:machinelearningservices:BatchDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BatchDeployment resources.
type batchDeploymentState struct {
}

type BatchDeploymentState struct {
}

func (BatchDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*batchDeploymentState)(nil)).Elem()
}

type batchDeploymentArgs struct {
	// The identifier for the Batch inference deployment.
	DeploymentName *string `pulumi:"deploymentName"`
	// Inference endpoint name
	EndpointName string `pulumi:"endpointName"`
	// Service identity associated with a resource.
	Identity *ResourceIdentity `pulumi:"identity"`
	Kind     *string           `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Additional attributes of the entity.
	Properties BatchDeploymentType `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a BatchDeployment resource.
type BatchDeploymentArgs struct {
	// The identifier for the Batch inference deployment.
	DeploymentName pulumi.StringPtrInput
	// Inference endpoint name
	EndpointName pulumi.StringInput
	// Service identity associated with a resource.
	Identity ResourceIdentityPtrInput
	Kind     pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Additional attributes of the entity.
	Properties BatchDeploymentTypeInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (BatchDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*batchDeploymentArgs)(nil)).Elem()
}

type BatchDeploymentInput interface {
	pulumi.Input

	ToBatchDeploymentOutput() BatchDeploymentOutput
	ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput
}

func (*BatchDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchDeployment)(nil))
}

func (i *BatchDeployment) ToBatchDeploymentOutput() BatchDeploymentOutput {
	return i.ToBatchDeploymentOutputWithContext(context.Background())
}

func (i *BatchDeployment) ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchDeploymentOutput)
}

type BatchDeploymentOutput struct {
	*pulumi.OutputState
}

func (BatchDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchDeployment)(nil))
}

func (o BatchDeploymentOutput) ToBatchDeploymentOutput() BatchDeploymentOutput {
	return o
}

func (o BatchDeploymentOutput) ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BatchDeploymentOutput{})
}

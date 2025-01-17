// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Kubernetes cluster specialized for web workloads by Azure App Service
// API Version: 2021-01-01.
type KubeEnvironment struct {
	pulumi.CustomResourceState

	AksResourceID pulumi.StringPtrOutput `pulumi:"aksResourceID"`
	// Cluster configuration which enables the log daemon to export
	// app logs to a destination. Currently only "log-analytics" is
	// supported
	AppLogsConfiguration AppLogsConfigurationResponsePtrOutput `pulumi:"appLogsConfiguration"`
	// Cluster configuration which determines the ARC cluster
	// components types. Eg: Choosing between BuildService kind,
	// FrontEnd Service ArtifactsStorageType etc.
	ArcConfiguration ArcConfigurationResponsePtrOutput `pulumi:"arcConfiguration"`
	// Default Domain Name for the cluster
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// Any errors that occurred during deployment or deployment validation
	DeploymentErrors pulumi.StringOutput `pulumi:"deploymentErrors"`
	// Extended Location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Only visible within Vnet/Subnet
	InternalLoadBalancerEnabled pulumi.BoolPtrOutput `pulumi:"internalLoadBalancerEnabled"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Kubernetes Environment.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Static IP of the KubeEnvironment
	StaticIp pulumi.StringPtrOutput `pulumi:"staticIp"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKubeEnvironment registers a new resource with the given unique name, arguments, and options.
func NewKubeEnvironment(ctx *pulumi.Context,
	name string, args *KubeEnvironmentArgs, opts ...pulumi.ResourceOption) (*KubeEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:KubeEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource KubeEnvironment
	err := ctx.RegisterResource("azure-native:web:KubeEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubeEnvironment gets an existing KubeEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubeEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubeEnvironmentState, opts ...pulumi.ResourceOption) (*KubeEnvironment, error) {
	var resource KubeEnvironment
	err := ctx.ReadResource("azure-native:web:KubeEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubeEnvironment resources.
type kubeEnvironmentState struct {
}

type KubeEnvironmentState struct {
}

func (KubeEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeEnvironmentState)(nil)).Elem()
}

type kubeEnvironmentArgs struct {
	AksResourceID *string `pulumi:"aksResourceID"`
	// Cluster configuration which enables the log daemon to export
	// app logs to a destination. Currently only "log-analytics" is
	// supported
	AppLogsConfiguration *AppLogsConfiguration `pulumi:"appLogsConfiguration"`
	// Cluster configuration which determines the ARC cluster
	// components types. Eg: Choosing between BuildService kind,
	// FrontEnd Service ArtifactsStorageType etc.
	ArcConfiguration *ArcConfiguration `pulumi:"arcConfiguration"`
	// Extended Location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Only visible within Vnet/Subnet
	InternalLoadBalancerEnabled *bool `pulumi:"internalLoadBalancerEnabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Name of the Kubernetes Environment.
	Name *string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Static IP of the KubeEnvironment
	StaticIp *string `pulumi:"staticIp"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a KubeEnvironment resource.
type KubeEnvironmentArgs struct {
	AksResourceID pulumi.StringPtrInput
	// Cluster configuration which enables the log daemon to export
	// app logs to a destination. Currently only "log-analytics" is
	// supported
	AppLogsConfiguration AppLogsConfigurationPtrInput
	// Cluster configuration which determines the ARC cluster
	// components types. Eg: Choosing between BuildService kind,
	// FrontEnd Service ArtifactsStorageType etc.
	ArcConfiguration ArcConfigurationPtrInput
	// Extended Location.
	ExtendedLocation ExtendedLocationPtrInput
	// Only visible within Vnet/Subnet
	InternalLoadBalancerEnabled pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Name of the Kubernetes Environment.
	Name pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Static IP of the KubeEnvironment
	StaticIp pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (KubeEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeEnvironmentArgs)(nil)).Elem()
}

type KubeEnvironmentInput interface {
	pulumi.Input

	ToKubeEnvironmentOutput() KubeEnvironmentOutput
	ToKubeEnvironmentOutputWithContext(ctx context.Context) KubeEnvironmentOutput
}

func (*KubeEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironment)(nil))
}

func (i *KubeEnvironment) ToKubeEnvironmentOutput() KubeEnvironmentOutput {
	return i.ToKubeEnvironmentOutputWithContext(context.Background())
}

func (i *KubeEnvironment) ToKubeEnvironmentOutputWithContext(ctx context.Context) KubeEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentOutput)
}

type KubeEnvironmentOutput struct {
	*pulumi.OutputState
}

func (KubeEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironment)(nil))
}

func (o KubeEnvironmentOutput) ToKubeEnvironmentOutput() KubeEnvironmentOutput {
	return o
}

func (o KubeEnvironmentOutput) ToKubeEnvironmentOutputWithContext(ctx context.Context) KubeEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KubeEnvironmentOutput{})
}

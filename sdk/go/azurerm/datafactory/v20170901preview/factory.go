// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170901preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Factory resource type.
type Factory struct {
	pulumi.CustomResourceState

	// Time the factory was created in ISO8601 format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Managed service identity of the factory.
	Identity FactoryIdentityResponsePtrOutput `pulumi:"identity"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Factory provisioning state, example Succeeded.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the factory.
	Version pulumi.StringOutput `pulumi:"version"`
	// VSTS repo information of the factory.
	VstsConfiguration FactoryVSTSConfigurationResponsePtrOutput `pulumi:"vstsConfiguration"`
}

// NewFactory registers a new resource with the given unique name, arguments, and options.
func NewFactory(ctx *pulumi.Context,
	name string, args *FactoryArgs, opts ...pulumi.ResourceOption) (*Factory, error) {
	if args == nil || args.FactoryName == nil {
		return nil, errors.New("missing required argument 'FactoryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &FactoryArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:datafactory/latest:Factory"),
		},
		{
			Type: pulumi.String("azurerm:datafactory/v20180601:Factory"),
		},
	})
	opts = append(opts, aliases)
	var resource Factory
	err := ctx.RegisterResource("azurerm:datafactory/v20170901preview:Factory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFactory gets an existing Factory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFactory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FactoryState, opts ...pulumi.ResourceOption) (*Factory, error) {
	var resource Factory
	err := ctx.ReadResource("azurerm:datafactory/v20170901preview:Factory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Factory resources.
type factoryState struct {
	// Time the factory was created in ISO8601 format.
	CreateTime *string `pulumi:"createTime"`
	// Managed service identity of the factory.
	Identity *FactoryIdentityResponse `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Factory provisioning state, example Succeeded.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
	// Version of the factory.
	Version *string `pulumi:"version"`
	// VSTS repo information of the factory.
	VstsConfiguration *FactoryVSTSConfigurationResponse `pulumi:"vstsConfiguration"`
}

type FactoryState struct {
	// Time the factory was created in ISO8601 format.
	CreateTime pulumi.StringPtrInput
	// Managed service identity of the factory.
	Identity FactoryIdentityResponsePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Factory provisioning state, example Succeeded.
	ProvisioningState pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
	// Version of the factory.
	Version pulumi.StringPtrInput
	// VSTS repo information of the factory.
	VstsConfiguration FactoryVSTSConfigurationResponsePtrInput
}

func (FactoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*factoryState)(nil)).Elem()
}

type factoryArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// Managed service identity of the factory.
	Identity *FactoryIdentity `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// VSTS repo information of the factory.
	VstsConfiguration *FactoryVSTSConfiguration `pulumi:"vstsConfiguration"`
}

// The set of arguments for constructing a Factory resource.
type FactoryArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput
	// Managed service identity of the factory.
	Identity FactoryIdentityPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// VSTS repo information of the factory.
	VstsConfiguration FactoryVSTSConfigurationPtrInput
}

func (FactoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*factoryArgs)(nil)).Elem()
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Defines an Network Experiment Profile and lists of Experiments
type NetworkExperimentProfile struct {
	pulumi.CustomResourceState

	// The state of the Experiment
	EnabledState pulumi.StringPtrOutput `pulumi:"enabledState"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource status.
	ResourceState pulumi.StringPtrOutput `pulumi:"resourceState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkExperimentProfile registers a new resource with the given unique name, arguments, and options.
func NewNetworkExperimentProfile(ctx *pulumi.Context,
	name string, args *NetworkExperimentProfileArgs, opts ...pulumi.ResourceOption) (*NetworkExperimentProfile, error) {
	if args == nil || args.ProfileName == nil {
		return nil, errors.New("missing required argument 'ProfileName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkExperimentProfileArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/v20191101:NetworkExperimentProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkExperimentProfile
	err := ctx.RegisterResource("azurerm:network/latest:NetworkExperimentProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkExperimentProfile gets an existing NetworkExperimentProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkExperimentProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkExperimentProfileState, opts ...pulumi.ResourceOption) (*NetworkExperimentProfile, error) {
	var resource NetworkExperimentProfile
	err := ctx.ReadResource("azurerm:network/latest:NetworkExperimentProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkExperimentProfile resources.
type networkExperimentProfileState struct {
	// The state of the Experiment
	EnabledState *string `pulumi:"enabledState"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource status.
	ResourceState *string `pulumi:"resourceState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type NetworkExperimentProfileState struct {
	// The state of the Experiment
	EnabledState pulumi.StringPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource status.
	ResourceState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (NetworkExperimentProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkExperimentProfileState)(nil)).Elem()
}

type networkExperimentProfileArgs struct {
	// The state of the Experiment
	EnabledState *string `pulumi:"enabledState"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the Profile
	Name *string `pulumi:"name"`
	// The Profile identifier associated with the Tenant and Partner
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource status.
	ResourceState *string `pulumi:"resourceState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkExperimentProfile resource.
type NetworkExperimentProfileArgs struct {
	// The state of the Experiment
	EnabledState pulumi.StringPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the Profile
	Name pulumi.StringPtrInput
	// The Profile identifier associated with the Tenant and Partner
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Resource status.
	ResourceState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkExperimentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkExperimentProfileArgs)(nil)).Elem()
}
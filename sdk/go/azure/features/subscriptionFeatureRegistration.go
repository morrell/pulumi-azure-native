// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package features

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Subscription feature registration details
// API Version: 2021-07-01.
type SubscriptionFeatureRegistration struct {
	pulumi.CustomResourceState

	// Azure resource name.
	Name       pulumi.StringOutput                                     `pulumi:"name"`
	Properties SubscriptionFeatureRegistrationResponsePropertiesOutput `pulumi:"properties"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSubscriptionFeatureRegistration registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionFeatureRegistration(ctx *pulumi.Context,
	name string, args *SubscriptionFeatureRegistrationArgs, opts ...pulumi.ResourceOption) (*SubscriptionFeatureRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:features:SubscriptionFeatureRegistration"),
		},
		{
			Type: pulumi.String("azure-native:features/v20210701:SubscriptionFeatureRegistration"),
		},
		{
			Type: pulumi.String("azure-nextgen:features/v20210701:SubscriptionFeatureRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource SubscriptionFeatureRegistration
	err := ctx.RegisterResource("azure-native:features:SubscriptionFeatureRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionFeatureRegistration gets an existing SubscriptionFeatureRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionFeatureRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionFeatureRegistrationState, opts ...pulumi.ResourceOption) (*SubscriptionFeatureRegistration, error) {
	var resource SubscriptionFeatureRegistration
	err := ctx.ReadResource("azure-native:features:SubscriptionFeatureRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionFeatureRegistration resources.
type subscriptionFeatureRegistrationState struct {
}

type SubscriptionFeatureRegistrationState struct {
}

func (SubscriptionFeatureRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionFeatureRegistrationState)(nil)).Elem()
}

type subscriptionFeatureRegistrationArgs struct {
	// The feature name.
	FeatureName *string                                    `pulumi:"featureName"`
	Properties  *SubscriptionFeatureRegistrationProperties `pulumi:"properties"`
	// The provider namespace.
	ProviderNamespace string `pulumi:"providerNamespace"`
}

// The set of arguments for constructing a SubscriptionFeatureRegistration resource.
type SubscriptionFeatureRegistrationArgs struct {
	// The feature name.
	FeatureName pulumi.StringPtrInput
	Properties  SubscriptionFeatureRegistrationPropertiesPtrInput
	// The provider namespace.
	ProviderNamespace pulumi.StringInput
}

func (SubscriptionFeatureRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionFeatureRegistrationArgs)(nil)).Elem()
}

type SubscriptionFeatureRegistrationInput interface {
	pulumi.Input

	ToSubscriptionFeatureRegistrationOutput() SubscriptionFeatureRegistrationOutput
	ToSubscriptionFeatureRegistrationOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationOutput
}

func (*SubscriptionFeatureRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFeatureRegistration)(nil))
}

func (i *SubscriptionFeatureRegistration) ToSubscriptionFeatureRegistrationOutput() SubscriptionFeatureRegistrationOutput {
	return i.ToSubscriptionFeatureRegistrationOutputWithContext(context.Background())
}

func (i *SubscriptionFeatureRegistration) ToSubscriptionFeatureRegistrationOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFeatureRegistrationOutput)
}

type SubscriptionFeatureRegistrationOutput struct {
	*pulumi.OutputState
}

func (SubscriptionFeatureRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFeatureRegistration)(nil))
}

func (o SubscriptionFeatureRegistrationOutput) ToSubscriptionFeatureRegistrationOutput() SubscriptionFeatureRegistrationOutput {
	return o
}

func (o SubscriptionFeatureRegistrationOutput) ToSubscriptionFeatureRegistrationOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubscriptionFeatureRegistrationOutput{})
}

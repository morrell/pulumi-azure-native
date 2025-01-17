// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single Namespace item in List or Get Operation
type Namespace struct {
	pulumi.CustomResourceState

	// The time the Namespace was created.
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// Specifies whether this instance is enabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Identifier for Azure Insights metrics
	MetricId pulumi.StringOutput `pulumi:"metricId"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Namespace.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringPtrOutput `pulumi:"serviceBusEndpoint"`
	// SKU parameters supplied to the create Namespace operation
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// State of the Namespace.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the Namespace was updated.
	UpdatedAt pulumi.StringPtrOutput `pulumi:"updatedAt"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20180101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20210101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20210601preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:eventhub/v20150801:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:eventhub/v20150801:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
}

type NamespaceState struct {
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// The time the Namespace was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies whether this instance is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Namespace location.
	Location *string `pulumi:"location"`
	// The Namespace name
	NamespaceName *string `pulumi:"namespaceName"`
	// Provisioning state of the Namespace.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `pulumi:"serviceBusEndpoint"`
	// SKU parameters supplied to the create Namespace operation
	Sku *Sku `pulumi:"sku"`
	// State of the Namespace.
	Status *NamespaceStateEnum `pulumi:"status"`
	// Namespace tags.
	Tags map[string]string `pulumi:"tags"`
	// The time the Namespace was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// The time the Namespace was created.
	CreatedAt pulumi.StringPtrInput
	// Specifies whether this instance is enabled.
	Enabled pulumi.BoolPtrInput
	// Namespace location.
	Location pulumi.StringPtrInput
	// The Namespace name
	NamespaceName pulumi.StringPtrInput
	// Provisioning state of the Namespace.
	ProvisioningState pulumi.StringPtrInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringPtrInput
	// SKU parameters supplied to the create Namespace operation
	Sku SkuPtrInput
	// State of the Namespace.
	Status NamespaceStateEnumPtrInput
	// Namespace tags.
	Tags pulumi.StringMapInput
	// The time the Namespace was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct {
	*pulumi.OutputState
}

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}

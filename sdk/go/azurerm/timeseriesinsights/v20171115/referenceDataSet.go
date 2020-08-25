// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171115

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A reference data set provides metadata about the events in an environment. Metadata in the reference data set will be joined with events as they are read from event sources. The metadata that makes up the reference data set is uploaded or modified through the Time Series Insights data plane APIs.
type ReferenceDataSet struct {
	pulumi.CustomResourceState

	// The time the resource was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
	DataStringComparisonBehavior pulumi.StringPtrOutput `pulumi:"dataStringComparisonBehavior"`
	// The list of key properties for the reference data set.
	KeyProperties ReferenceDataSetKeyPropertyResponseArrayOutput `pulumi:"keyProperties"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReferenceDataSet registers a new resource with the given unique name, arguments, and options.
func NewReferenceDataSet(ctx *pulumi.Context,
	name string, args *ReferenceDataSetArgs, opts ...pulumi.ResourceOption) (*ReferenceDataSet, error) {
	if args == nil || args.EnvironmentName == nil {
		return nil, errors.New("missing required argument 'EnvironmentName'")
	}
	if args == nil || args.KeyProperties == nil {
		return nil, errors.New("missing required argument 'KeyProperties'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ReferenceDataSetArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:timeseriesinsights/v20200515:ReferenceDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ReferenceDataSet
	err := ctx.RegisterResource("azurerm:timeseriesinsights/v20171115:ReferenceDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReferenceDataSet gets an existing ReferenceDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReferenceDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReferenceDataSetState, opts ...pulumi.ResourceOption) (*ReferenceDataSet, error) {
	var resource ReferenceDataSet
	err := ctx.ReadResource("azurerm:timeseriesinsights/v20171115:ReferenceDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReferenceDataSet resources.
type referenceDataSetState struct {
	// The time the resource was created.
	CreationTime *string `pulumi:"creationTime"`
	// The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
	DataStringComparisonBehavior *string `pulumi:"dataStringComparisonBehavior"`
	// The list of key properties for the reference data set.
	KeyProperties []ReferenceDataSetKeyPropertyResponse `pulumi:"keyProperties"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ReferenceDataSetState struct {
	// The time the resource was created.
	CreationTime pulumi.StringPtrInput
	// The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
	DataStringComparisonBehavior pulumi.StringPtrInput
	// The list of key properties for the reference data set.
	KeyProperties ReferenceDataSetKeyPropertyResponseArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ReferenceDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceDataSetState)(nil)).Elem()
}

type referenceDataSetArgs struct {
	// The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
	DataStringComparisonBehavior *string `pulumi:"dataStringComparisonBehavior"`
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// The list of key properties for the reference data set.
	KeyProperties []ReferenceDataSetKeyProperty `pulumi:"keyProperties"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// Name of the reference data set.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Key-value pairs of additional properties for the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ReferenceDataSet resource.
type ReferenceDataSetArgs struct {
	// The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
	DataStringComparisonBehavior pulumi.StringPtrInput
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput
	// The list of key properties for the reference data set.
	KeyProperties ReferenceDataSetKeyPropertyArrayInput
	// The location of the resource.
	Location pulumi.StringInput
	// Name of the reference data set.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Key-value pairs of additional properties for the resource.
	Tags pulumi.StringMapInput
}

func (ReferenceDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceDataSetArgs)(nil)).Elem()
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type HyperVCollector struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput            `pulumi:"eTag"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties CollectorPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}

// NewHyperVCollector registers a new resource with the given unique name, arguments, and options.
func NewHyperVCollector(ctx *pulumi.Context,
	name string, args *HyperVCollectorArgs, opts ...pulumi.ResourceOption) (*HyperVCollector, error) {
	if args == nil || args.HyperVCollectorName == nil {
		return nil, errors.New("missing required argument 'HyperVCollectorName'")
	}
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &HyperVCollectorArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:migrate/v20191001:HyperVCollector"),
		},
	})
	opts = append(opts, aliases)
	var resource HyperVCollector
	err := ctx.RegisterResource("azurerm:migrate/latest:HyperVCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHyperVCollector gets an existing HyperVCollector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHyperVCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HyperVCollectorState, opts ...pulumi.ResourceOption) (*HyperVCollector, error) {
	var resource HyperVCollector
	err := ctx.ReadResource("azurerm:migrate/latest:HyperVCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HyperVCollector resources.
type hyperVCollectorState struct {
	ETag       *string                      `pulumi:"eTag"`
	Name       *string                      `pulumi:"name"`
	Properties *CollectorPropertiesResponse `pulumi:"properties"`
	Type       *string                      `pulumi:"type"`
}

type HyperVCollectorState struct {
	ETag       pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	Properties CollectorPropertiesResponsePtrInput
	Type       pulumi.StringPtrInput
}

func (HyperVCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVCollectorState)(nil)).Elem()
}

type hyperVCollectorArgs struct {
	ETag *string `pulumi:"eTag"`
	// Unique name of a Hyper-V collector within a project.
	HyperVCollectorName string `pulumi:"hyperVCollectorName"`
	// Name of the Azure Migrate project.
	ProjectName string               `pulumi:"projectName"`
	Properties  *CollectorProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a HyperVCollector resource.
type HyperVCollectorArgs struct {
	ETag pulumi.StringPtrInput
	// Unique name of a Hyper-V collector within a project.
	HyperVCollectorName pulumi.StringInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput
	Properties  CollectorPropertiesPtrInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
}

func (HyperVCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVCollectorArgs)(nil)).Elem()
}
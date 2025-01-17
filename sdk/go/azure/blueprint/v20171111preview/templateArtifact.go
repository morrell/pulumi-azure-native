// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Blueprint artifact deploys Azure resource manager template.
type TemplateArtifact struct {
	pulumi.CustomResourceState

	// Artifacts which need to be deployed before the specified artifact.
	DependsOn pulumi.StringArrayOutput `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Specifies the kind of Blueprint artifact.
	// Expected value is 'template'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Template parameter values.
	Parameters ParameterValueBaseResponseMapOutput `pulumi:"parameters"`
	// If applicable, the name of the resource group placeholder to which the template will be deployed.
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// The Azure Resource Manager template body.
	Template pulumi.AnyOutput `pulumi:"template"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTemplateArtifact registers a new resource with the given unique name, arguments, and options.
func NewTemplateArtifact(ctx *pulumi.Context,
	name string, args *TemplateArtifactArgs, opts ...pulumi.ResourceOption) (*TemplateArtifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ManagementGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupName'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	args.Kind = pulumi.String("template")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:blueprint/v20171111preview:TemplateArtifact"),
		},
	})
	opts = append(opts, aliases)
	var resource TemplateArtifact
	err := ctx.RegisterResource("azure-native:blueprint/v20171111preview:TemplateArtifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateArtifact gets an existing TemplateArtifact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateArtifactState, opts ...pulumi.ResourceOption) (*TemplateArtifact, error) {
	var resource TemplateArtifact
	err := ctx.ReadResource("azure-native:blueprint/v20171111preview:TemplateArtifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateArtifact resources.
type templateArtifactState struct {
}

type TemplateArtifactState struct {
}

func (TemplateArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArtifactState)(nil)).Elem()
}

type templateArtifactArgs struct {
	// name of the artifact.
	ArtifactName *string `pulumi:"artifactName"`
	// name of the blueprint.
	BlueprintName string `pulumi:"blueprintName"`
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []string `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the kind of Blueprint artifact.
	// Expected value is 'template'.
	Kind string `pulumi:"kind"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName string `pulumi:"managementGroupName"`
	// Template parameter values.
	Parameters map[string]ParameterValueBase `pulumi:"parameters"`
	// If applicable, the name of the resource group placeholder to which the template will be deployed.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// The Azure Resource Manager template body.
	Template interface{} `pulumi:"template"`
}

// The set of arguments for constructing a TemplateArtifact resource.
type TemplateArtifactArgs struct {
	// name of the artifact.
	ArtifactName pulumi.StringPtrInput
	// name of the blueprint.
	BlueprintName pulumi.StringInput
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn pulumi.StringArrayInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Specifies the kind of Blueprint artifact.
	// Expected value is 'template'.
	Kind pulumi.StringInput
	// ManagementGroup where blueprint stores.
	ManagementGroupName pulumi.StringInput
	// Template parameter values.
	Parameters ParameterValueBaseMapInput
	// If applicable, the name of the resource group placeholder to which the template will be deployed.
	ResourceGroup pulumi.StringPtrInput
	// The Azure Resource Manager template body.
	Template pulumi.Input
}

func (TemplateArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArtifactArgs)(nil)).Elem()
}

type TemplateArtifactInput interface {
	pulumi.Input

	ToTemplateArtifactOutput() TemplateArtifactOutput
	ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput
}

func (*TemplateArtifact) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateArtifact)(nil))
}

func (i *TemplateArtifact) ToTemplateArtifactOutput() TemplateArtifactOutput {
	return i.ToTemplateArtifactOutputWithContext(context.Background())
}

func (i *TemplateArtifact) ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArtifactOutput)
}

type TemplateArtifactOutput struct {
	*pulumi.OutputState
}

func (TemplateArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateArtifact)(nil))
}

func (o TemplateArtifactOutput) ToTemplateArtifactOutput() TemplateArtifactOutput {
	return o
}

func (o TemplateArtifactOutput) ToTemplateArtifactOutputWithContext(ctx context.Context) TemplateArtifactOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TemplateArtifactOutput{})
}

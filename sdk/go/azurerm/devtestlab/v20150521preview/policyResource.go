// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Policy.
type PolicyResource struct {
	pulumi.CustomResourceState

	// The description of the policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The evaluator type of the policy.
	EvaluatorType pulumi.StringPtrOutput `pulumi:"evaluatorType"`
	// The fact data of the policy.
	FactData pulumi.StringPtrOutput `pulumi:"factData"`
	// The fact name of the policy.
	FactName pulumi.StringPtrOutput `pulumi:"factName"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The status of the policy.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The threshold of the policy.
	Threshold pulumi.StringPtrOutput `pulumi:"threshold"`
	// The type of the resource.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewPolicyResource registers a new resource with the given unique name, arguments, and options.
func NewPolicyResource(ctx *pulumi.Context,
	name string, args *PolicyResourceArgs, opts ...pulumi.ResourceOption) (*PolicyResource, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PolicySetName == nil {
		return nil, errors.New("missing required argument 'PolicySetName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PolicyResourceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:devtestlab/latest:PolicyResource"),
		},
		{
			Type: pulumi.String("azurerm:devtestlab/v20160515:PolicyResource"),
		},
		{
			Type: pulumi.String("azurerm:devtestlab/v20180915:PolicyResource"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyResource
	err := ctx.RegisterResource("azurerm:devtestlab/v20150521preview:PolicyResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyResource gets an existing PolicyResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyResourceState, opts ...pulumi.ResourceOption) (*PolicyResource, error) {
	var resource PolicyResource
	err := ctx.ReadResource("azurerm:devtestlab/v20150521preview:PolicyResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyResource resources.
type policyResourceState struct {
	// The description of the policy.
	Description *string `pulumi:"description"`
	// The evaluator type of the policy.
	EvaluatorType *string `pulumi:"evaluatorType"`
	// The fact data of the policy.
	FactData *string `pulumi:"factData"`
	// The fact name of the policy.
	FactName *string `pulumi:"factName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The status of the policy.
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The threshold of the policy.
	Threshold *string `pulumi:"threshold"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type PolicyResourceState struct {
	// The description of the policy.
	Description pulumi.StringPtrInput
	// The evaluator type of the policy.
	EvaluatorType pulumi.StringPtrInput
	// The fact data of the policy.
	FactData pulumi.StringPtrInput
	// The fact name of the policy.
	FactName pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The status of the policy.
	Status pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The threshold of the policy.
	Threshold pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (PolicyResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyResourceState)(nil)).Elem()
}

type policyResourceArgs struct {
	// The description of the policy.
	Description *string `pulumi:"description"`
	// The evaluator type of the policy.
	EvaluatorType *string `pulumi:"evaluatorType"`
	// The fact data of the policy.
	FactData *string `pulumi:"factData"`
	// The fact name of the policy.
	FactName *string `pulumi:"factName"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The name of the policy set.
	PolicySetName string `pulumi:"policySetName"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of the policy.
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The threshold of the policy.
	Threshold *string `pulumi:"threshold"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a PolicyResource resource.
type PolicyResourceArgs struct {
	// The description of the policy.
	Description pulumi.StringPtrInput
	// The evaluator type of the policy.
	EvaluatorType pulumi.StringPtrInput
	// The fact data of the policy.
	FactData pulumi.StringPtrInput
	// The fact name of the policy.
	FactName pulumi.StringPtrInput
	// The identifier of the resource.
	Id pulumi.StringPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringInput
	// The name of the policy set.
	PolicySetName pulumi.StringInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The status of the policy.
	Status pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The threshold of the policy.
	Threshold pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (PolicyResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyResourceArgs)(nil)).Elem()
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about JIT request definition.
type JitRequest struct {
	pulumi.CustomResourceState

	// The parent application id.
	ApplicationResourceId pulumi.StringOutput `pulumi:"applicationResourceId"`
	// The client entity that created the JIT request.
	CreatedBy ApplicationClientDetailsResponseOutput `pulumi:"createdBy"`
	// The JIT authorization policies.
	JitAuthorizationPolicies JitAuthorizationPoliciesResponseArrayOutput `pulumi:"jitAuthorizationPolicies"`
	// The JIT request state.
	JitRequestState pulumi.StringOutput `pulumi:"jitRequestState"`
	// The JIT request properties.
	JitSchedulingPolicy JitSchedulingPolicyResponseOutput `pulumi:"jitSchedulingPolicy"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The JIT request provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The publisher tenant id.
	PublisherTenantId pulumi.StringOutput `pulumi:"publisherTenantId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The client entity that last updated the JIT request.
	UpdatedBy ApplicationClientDetailsResponseOutput `pulumi:"updatedBy"`
}

// NewJitRequest registers a new resource with the given unique name, arguments, and options.
func NewJitRequest(ctx *pulumi.Context,
	name string, args *JitRequestArgs, opts ...pulumi.ResourceOption) (*JitRequest, error) {
	if args == nil || args.ApplicationResourceId == nil {
		return nil, errors.New("missing required argument 'ApplicationResourceId'")
	}
	if args == nil || args.JitAuthorizationPolicies == nil {
		return nil, errors.New("missing required argument 'JitAuthorizationPolicies'")
	}
	if args == nil || args.JitRequestName == nil {
		return nil, errors.New("missing required argument 'JitRequestName'")
	}
	if args == nil || args.JitSchedulingPolicy == nil {
		return nil, errors.New("missing required argument 'JitSchedulingPolicy'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &JitRequestArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:solutions/v20190701:JitRequest"),
		},
	})
	opts = append(opts, aliases)
	var resource JitRequest
	err := ctx.RegisterResource("azurerm:solutions/latest:JitRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJitRequest gets an existing JitRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJitRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JitRequestState, opts ...pulumi.ResourceOption) (*JitRequest, error) {
	var resource JitRequest
	err := ctx.ReadResource("azurerm:solutions/latest:JitRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JitRequest resources.
type jitRequestState struct {
	// The parent application id.
	ApplicationResourceId *string `pulumi:"applicationResourceId"`
	// The client entity that created the JIT request.
	CreatedBy *ApplicationClientDetailsResponse `pulumi:"createdBy"`
	// The JIT authorization policies.
	JitAuthorizationPolicies []JitAuthorizationPoliciesResponse `pulumi:"jitAuthorizationPolicies"`
	// The JIT request state.
	JitRequestState *string `pulumi:"jitRequestState"`
	// The JIT request properties.
	JitSchedulingPolicy *JitSchedulingPolicyResponse `pulumi:"jitSchedulingPolicy"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The JIT request provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The publisher tenant id.
	PublisherTenantId *string `pulumi:"publisherTenantId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// The client entity that last updated the JIT request.
	UpdatedBy *ApplicationClientDetailsResponse `pulumi:"updatedBy"`
}

type JitRequestState struct {
	// The parent application id.
	ApplicationResourceId pulumi.StringPtrInput
	// The client entity that created the JIT request.
	CreatedBy ApplicationClientDetailsResponsePtrInput
	// The JIT authorization policies.
	JitAuthorizationPolicies JitAuthorizationPoliciesResponseArrayInput
	// The JIT request state.
	JitRequestState pulumi.StringPtrInput
	// The JIT request properties.
	JitSchedulingPolicy JitSchedulingPolicyResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The JIT request provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The publisher tenant id.
	PublisherTenantId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// The client entity that last updated the JIT request.
	UpdatedBy ApplicationClientDetailsResponsePtrInput
}

func (JitRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*jitRequestState)(nil)).Elem()
}

type jitRequestArgs struct {
	// The parent application id.
	ApplicationResourceId string `pulumi:"applicationResourceId"`
	// The JIT authorization policies.
	JitAuthorizationPolicies []JitAuthorizationPolicies `pulumi:"jitAuthorizationPolicies"`
	// The name of the JIT request.
	JitRequestName string `pulumi:"jitRequestName"`
	// The JIT request properties.
	JitSchedulingPolicy JitSchedulingPolicy `pulumi:"jitSchedulingPolicy"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a JitRequest resource.
type JitRequestArgs struct {
	// The parent application id.
	ApplicationResourceId pulumi.StringInput
	// The JIT authorization policies.
	JitAuthorizationPolicies JitAuthorizationPoliciesArrayInput
	// The name of the JIT request.
	JitRequestName pulumi.StringInput
	// The JIT request properties.
	JitSchedulingPolicy JitSchedulingPolicyInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (JitRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jitRequestArgs)(nil)).Elem()
}
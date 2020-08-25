// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Api Operation details.
type ApiOperation struct {
	pulumi.CustomResourceState

	// Description of the operation. May include HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Operation Name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
	Method pulumi.StringOutput `pulumi:"method"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Operation Policies
	Policies pulumi.StringPtrOutput `pulumi:"policies"`
	// An entity containing request details.
	Request RequestContractResponsePtrOutput `pulumi:"request"`
	// Array of Operation responses.
	Responses ResponseContractResponseArrayOutput `pulumi:"responses"`
	// Collection of URL template parameters.
	TemplateParameters ParameterContractResponseArrayOutput `pulumi:"templateParameters"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	UrlTemplate pulumi.StringOutput `pulumi:"urlTemplate"`
}

// NewApiOperation registers a new resource with the given unique name, arguments, and options.
func NewApiOperation(ctx *pulumi.Context,
	name string, args *ApiOperationArgs, opts ...pulumi.ResourceOption) (*ApiOperation, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Method == nil {
		return nil, errors.New("missing required argument 'Method'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.UrlTemplate == nil {
		return nil, errors.New("missing required argument 'UrlTemplate'")
	}
	if args == nil {
		args = &ApiOperationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/v20160707:ApiOperation"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20161010:ApiOperation"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20170301:ApiOperation"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180101:ApiOperation"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:ApiOperation"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiOperation
	err := ctx.RegisterResource("azurerm:apimanagement/v20190101:ApiOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiOperation gets an existing ApiOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationState, opts ...pulumi.ResourceOption) (*ApiOperation, error) {
	var resource ApiOperation
	err := ctx.ReadResource("azurerm:apimanagement/v20190101:ApiOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOperation resources.
type apiOperationState struct {
	// Description of the operation. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Operation Name.
	DisplayName *string `pulumi:"displayName"`
	// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
	Method *string `pulumi:"method"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Operation Policies
	Policies *string `pulumi:"policies"`
	// An entity containing request details.
	Request *RequestContractResponse `pulumi:"request"`
	// Array of Operation responses.
	Responses []ResponseContractResponse `pulumi:"responses"`
	// Collection of URL template parameters.
	TemplateParameters []ParameterContractResponse `pulumi:"templateParameters"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
	// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	UrlTemplate *string `pulumi:"urlTemplate"`
}

type ApiOperationState struct {
	// Description of the operation. May include HTML formatting tags.
	Description pulumi.StringPtrInput
	// Operation Name.
	DisplayName pulumi.StringPtrInput
	// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
	Method pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Operation Policies
	Policies pulumi.StringPtrInput
	// An entity containing request details.
	Request RequestContractResponsePtrInput
	// Array of Operation responses.
	Responses ResponseContractResponseArrayInput
	// Collection of URL template parameters.
	TemplateParameters ParameterContractResponseArrayInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
	// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	UrlTemplate pulumi.StringPtrInput
}

func (ApiOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationState)(nil)).Elem()
}

type apiOperationArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Description of the operation. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Operation Name.
	DisplayName string `pulumi:"displayName"`
	// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
	Method string `pulumi:"method"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// Operation Policies
	Policies *string `pulumi:"policies"`
	// An entity containing request details.
	Request *RequestContract `pulumi:"request"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Array of Operation responses.
	Responses []ResponseContract `pulumi:"responses"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Collection of URL template parameters.
	TemplateParameters []ParameterContract `pulumi:"templateParameters"`
	// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	UrlTemplate string `pulumi:"urlTemplate"`
}

// The set of arguments for constructing a ApiOperation resource.
type ApiOperationArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Description of the operation. May include HTML formatting tags.
	Description pulumi.StringPtrInput
	// Operation Name.
	DisplayName pulumi.StringInput
	// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
	Method pulumi.StringInput
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	Name pulumi.StringInput
	// Operation Policies
	Policies pulumi.StringPtrInput
	// An entity containing request details.
	Request RequestContractPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Array of Operation responses.
	Responses ResponseContractArrayInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Collection of URL template parameters.
	TemplateParameters ParameterContractArrayInput
	// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	UrlTemplate pulumi.StringInput
}

func (ApiOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationArgs)(nil)).Elem()
}
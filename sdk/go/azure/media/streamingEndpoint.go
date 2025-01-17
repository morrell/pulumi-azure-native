// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The streaming endpoint.
// API Version: 2020-05-01.
type StreamingEndpoint struct {
	pulumi.CustomResourceState

	// The access control definition of the streaming endpoint.
	AccessControl StreamingEndpointAccessControlResponsePtrOutput `pulumi:"accessControl"`
	// This feature is deprecated, do not set a value for this property.
	AvailabilitySetName pulumi.StringPtrOutput `pulumi:"availabilitySetName"`
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrOutput `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile pulumi.StringPtrOutput `pulumi:"cdnProfile"`
	// The CDN provider name.
	CdnProvider pulumi.StringPtrOutput `pulumi:"cdnProvider"`
	// The exact time the streaming endpoint was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// The streaming endpoint access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput `pulumi:"crossSiteAccessPolicies"`
	// The custom host names of the streaming endpoint
	CustomHostNames pulumi.StringArrayOutput `pulumi:"customHostNames"`
	// The streaming endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The free trial expiration time.
	FreeTrialEndTime pulumi.StringOutput `pulumi:"freeTrialEndTime"`
	// The streaming endpoint host name.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The exact time the streaming endpoint was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Max cache age
	MaxCacheAge pulumi.Float64PtrOutput `pulumi:"maxCacheAge"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the streaming endpoint.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource state of the streaming endpoint.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The number of scale units. Use the Scale operation to adjust this value.
	ScaleUnits pulumi.IntOutput `pulumi:"scaleUnits"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingEndpoint registers a new resource with the given unique name, arguments, and options.
func NewStreamingEndpoint(ctx *pulumi.Context,
	name string, args *StreamingEndpointArgs, opts ...pulumi.ResourceOption) (*StreamingEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScaleUnits == nil {
		return nil, errors.New("invalid value for required argument 'ScaleUnits'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:media:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180330preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180701:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20190501preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20190501preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20210601:StreamingEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingEndpoint
	err := ctx.RegisterResource("azure-native:media:StreamingEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingEndpoint gets an existing StreamingEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingEndpointState, opts ...pulumi.ResourceOption) (*StreamingEndpoint, error) {
	var resource StreamingEndpoint
	err := ctx.ReadResource("azure-native:media:StreamingEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingEndpoint resources.
type streamingEndpointState struct {
}

type StreamingEndpointState struct {
}

func (StreamingEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingEndpointState)(nil)).Elem()
}

type streamingEndpointArgs struct {
	// The access control definition of the streaming endpoint.
	AccessControl *StreamingEndpointAccessControl `pulumi:"accessControl"`
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart *bool `pulumi:"autoStart"`
	// This feature is deprecated, do not set a value for this property.
	AvailabilitySetName *string `pulumi:"availabilitySetName"`
	// The CDN enabled flag.
	CdnEnabled *bool `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile *string `pulumi:"cdnProfile"`
	// The CDN provider name.
	CdnProvider *string `pulumi:"cdnProvider"`
	// The streaming endpoint access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPolicies `pulumi:"crossSiteAccessPolicies"`
	// The custom host names of the streaming endpoint
	CustomHostNames []string `pulumi:"customHostNames"`
	// The streaming endpoint description.
	Description *string `pulumi:"description"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Max cache age
	MaxCacheAge *float64 `pulumi:"maxCacheAge"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of scale units. Use the Scale operation to adjust this value.
	ScaleUnits int `pulumi:"scaleUnits"`
	// The name of the streaming endpoint, maximum length is 24.
	StreamingEndpointName *string `pulumi:"streamingEndpointName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StreamingEndpoint resource.
type StreamingEndpointArgs struct {
	// The access control definition of the streaming endpoint.
	AccessControl StreamingEndpointAccessControlPtrInput
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart pulumi.BoolPtrInput
	// This feature is deprecated, do not set a value for this property.
	AvailabilitySetName pulumi.StringPtrInput
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrInput
	// The CDN profile name.
	CdnProfile pulumi.StringPtrInput
	// The CDN provider name.
	CdnProvider pulumi.StringPtrInput
	// The streaming endpoint access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	// The custom host names of the streaming endpoint
	CustomHostNames pulumi.StringArrayInput
	// The streaming endpoint description.
	Description pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Max cache age
	MaxCacheAge pulumi.Float64PtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The number of scale units. Use the Scale operation to adjust this value.
	ScaleUnits pulumi.IntInput
	// The name of the streaming endpoint, maximum length is 24.
	StreamingEndpointName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (StreamingEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingEndpointArgs)(nil)).Elem()
}

type StreamingEndpointInput interface {
	pulumi.Input

	ToStreamingEndpointOutput() StreamingEndpointOutput
	ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput
}

func (*StreamingEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpoint)(nil))
}

func (i *StreamingEndpoint) ToStreamingEndpointOutput() StreamingEndpointOutput {
	return i.ToStreamingEndpointOutputWithContext(context.Background())
}

func (i *StreamingEndpoint) ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointOutput)
}

type StreamingEndpointOutput struct {
	*pulumi.OutputState
}

func (StreamingEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpoint)(nil))
}

func (o StreamingEndpointOutput) ToStreamingEndpointOutput() StreamingEndpointOutput {
	return o
}

func (o StreamingEndpointOutput) ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StreamingEndpointOutput{})
}

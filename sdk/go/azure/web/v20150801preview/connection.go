// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Connection
type Connection struct {
	pulumi.CustomResourceState

	// expanded connection provider name
	Api ExpandedParentApiEntityResponsePtrOutput `pulumi:"api"`
	// Timestamp of last connection change.
	ChangedTime pulumi.StringPtrOutput `pulumi:"changedTime"`
	// Timestamp of the connection creation
	CreatedTime pulumi.StringPtrOutput `pulumi:"createdTime"`
	// Custom login setting values.
	CustomParameterValues ParameterCustomLoginSettingValuesResponseMapOutput `pulumi:"customParameterValues"`
	// display name
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Time in UTC when the first expiration of OAuth tokens
	FirstExpirationTime pulumi.StringPtrOutput `pulumi:"firstExpirationTime"`
	// List of Keywords that tag the acl
	Keywords pulumi.StringArrayOutput `pulumi:"keywords"`
	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	Metadata pulumi.AnyOutput    `pulumi:"metadata"`
	// Resource Name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Tokens/Claim
	NonSecretParameterValues pulumi.MapOutput `pulumi:"nonSecretParameterValues"`
	// Tokens/Claim
	ParameterValues pulumi.MapOutput `pulumi:"parameterValues"`
	// Status of the connection
	Statuses ConnectionStatusResponseArrayOutput `pulumi:"statuses"`
	// Resource tags
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801preview:Connection"),
		},
		{
			Type: pulumi.String("azure-native:web:Connection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:Connection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160601:Connection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160601:Connection"),
		},
	})
	opts = append(opts, aliases)
	var resource Connection
	err := ctx.RegisterResource("azure-native:web/v20150801preview:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("azure-native:web/v20150801preview:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// expanded connection provider name
	Api *ExpandedParentApiEntity `pulumi:"api"`
	// Timestamp of last connection change.
	ChangedTime *string `pulumi:"changedTime"`
	// The connection name.
	ConnectionName *string `pulumi:"connectionName"`
	// Timestamp of the connection creation
	CreatedTime *string `pulumi:"createdTime"`
	// Custom login setting values.
	CustomParameterValues map[string]ParameterCustomLoginSettingValues `pulumi:"customParameterValues"`
	// display name
	DisplayName *string `pulumi:"displayName"`
	// Time in UTC when the first expiration of OAuth tokens
	FirstExpirationTime *string `pulumi:"firstExpirationTime"`
	// Resource Id
	Id *string `pulumi:"id"`
	// List of Keywords that tag the acl
	Keywords []string `pulumi:"keywords"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string     `pulumi:"location"`
	Metadata interface{} `pulumi:"metadata"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Tokens/Claim
	NonSecretParameterValues map[string]interface{} `pulumi:"nonSecretParameterValues"`
	// Tokens/Claim
	ParameterValues map[string]interface{} `pulumi:"parameterValues"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Status of the connection
	Statuses []ConnectionStatus `pulumi:"statuses"`
	// Resource tags
	Tags     map[string]string `pulumi:"tags"`
	TenantId *string           `pulumi:"tenantId"`
	// Resource type
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// expanded connection provider name
	Api ExpandedParentApiEntityPtrInput
	// Timestamp of last connection change.
	ChangedTime pulumi.StringPtrInput
	// The connection name.
	ConnectionName pulumi.StringPtrInput
	// Timestamp of the connection creation
	CreatedTime pulumi.StringPtrInput
	// Custom login setting values.
	CustomParameterValues ParameterCustomLoginSettingValuesMapInput
	// display name
	DisplayName pulumi.StringPtrInput
	// Time in UTC when the first expiration of OAuth tokens
	FirstExpirationTime pulumi.StringPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// List of Keywords that tag the acl
	Keywords pulumi.StringArrayInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	Metadata pulumi.Input
	// Resource Name
	Name pulumi.StringPtrInput
	// Tokens/Claim
	NonSecretParameterValues pulumi.MapInput
	// Tokens/Claim
	ParameterValues pulumi.MapInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Status of the connection
	Statuses ConnectionStatusArrayInput
	// Resource tags
	Tags     pulumi.StringMapInput
	TenantId pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct {
	*pulumi.OutputState
}

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
}

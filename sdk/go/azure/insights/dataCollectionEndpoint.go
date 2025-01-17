// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of ARM tracked top level resource.
// API Version: 2021-04-01.
type DataCollectionEndpoint struct {
	pulumi.CustomResourceState

	// The endpoint used by agents to access their configuration.
	ConfigurationAccess DataCollectionEndpointResponseConfigurationAccessPtrOutput `pulumi:"configurationAccess"`
	// Description of the data collection endpoint.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource entity tag (ETag).
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The immutable ID of this data collection endpoint resource. This property is READ-ONLY.
	ImmutableId pulumi.StringPtrOutput `pulumi:"immutableId"`
	// The kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives.
	Location pulumi.StringOutput `pulumi:"location"`
	// The endpoint used by clients to ingest logs.
	LogsIngestion DataCollectionEndpointResponseLogsIngestionPtrOutput `pulumi:"logsIngestion"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network access control rules for the endpoints.
	NetworkAcls DataCollectionEndpointResponseNetworkAclsPtrOutput `pulumi:"networkAcls"`
	// The resource provisioning state. This property is READ-ONLY.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData DataCollectionEndpointResourceResponseSystemDataOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataCollectionEndpoint registers a new resource with the given unique name, arguments, and options.
func NewDataCollectionEndpoint(ctx *pulumi.Context,
	name string, args *DataCollectionEndpointArgs, opts ...pulumi.ResourceOption) (*DataCollectionEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights:DataCollectionEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210401:DataCollectionEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20210401:DataCollectionEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource DataCollectionEndpoint
	err := ctx.RegisterResource("azure-native:insights:DataCollectionEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCollectionEndpoint gets an existing DataCollectionEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCollectionEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCollectionEndpointState, opts ...pulumi.ResourceOption) (*DataCollectionEndpoint, error) {
	var resource DataCollectionEndpoint
	err := ctx.ReadResource("azure-native:insights:DataCollectionEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCollectionEndpoint resources.
type dataCollectionEndpointState struct {
}

type DataCollectionEndpointState struct {
}

func (DataCollectionEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionEndpointState)(nil)).Elem()
}

type dataCollectionEndpointArgs struct {
	// The name of the data collection endpoint. The name is case insensitive.
	DataCollectionEndpointName *string `pulumi:"dataCollectionEndpointName"`
	// Description of the data collection endpoint.
	Description *string `pulumi:"description"`
	// The immutable ID of this data collection endpoint resource. This property is READ-ONLY.
	ImmutableId *string `pulumi:"immutableId"`
	// The kind of the resource.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives.
	Location *string `pulumi:"location"`
	// Network access control rules for the endpoints.
	NetworkAcls *DataCollectionEndpointNetworkAcls `pulumi:"networkAcls"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataCollectionEndpoint resource.
type DataCollectionEndpointArgs struct {
	// The name of the data collection endpoint. The name is case insensitive.
	DataCollectionEndpointName pulumi.StringPtrInput
	// Description of the data collection endpoint.
	Description pulumi.StringPtrInput
	// The immutable ID of this data collection endpoint resource. This property is READ-ONLY.
	ImmutableId pulumi.StringPtrInput
	// The kind of the resource.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives.
	Location pulumi.StringPtrInput
	// Network access control rules for the endpoints.
	NetworkAcls DataCollectionEndpointNetworkAclsPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DataCollectionEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionEndpointArgs)(nil)).Elem()
}

type DataCollectionEndpointInput interface {
	pulumi.Input

	ToDataCollectionEndpointOutput() DataCollectionEndpointOutput
	ToDataCollectionEndpointOutputWithContext(ctx context.Context) DataCollectionEndpointOutput
}

func (*DataCollectionEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpoint)(nil))
}

func (i *DataCollectionEndpoint) ToDataCollectionEndpointOutput() DataCollectionEndpointOutput {
	return i.ToDataCollectionEndpointOutputWithContext(context.Background())
}

func (i *DataCollectionEndpoint) ToDataCollectionEndpointOutputWithContext(ctx context.Context) DataCollectionEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointOutput)
}

type DataCollectionEndpointOutput struct {
	*pulumi.OutputState
}

func (DataCollectionEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpoint)(nil))
}

func (o DataCollectionEndpointOutput) ToDataCollectionEndpointOutput() DataCollectionEndpointOutput {
	return o
}

func (o DataCollectionEndpointOutput) ToDataCollectionEndpointOutputWithContext(ctx context.Context) DataCollectionEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataCollectionEndpointOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Blob data set mapping.
type BlobDataSetMapping struct {
	pulumi.CustomResourceState

	// Container that has the file path.
	ContainerName pulumi.StringOutput `pulumi:"containerName"`
	// The id of the source data set.
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus pulumi.StringOutput `pulumi:"dataSetMappingStatus"`
	// File path within the source data set
	FilePath pulumi.StringOutput `pulumi:"filePath"`
	// Kind of data set mapping.
	// Expected value is 'Blob'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// File output type
	OutputType pulumi.StringPtrOutput `pulumi:"outputType"`
	// Provisioning state of the data set mapping.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource group of storage account.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// Storage account name of the source data set.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// Subscription id of storage account.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBlobDataSetMapping registers a new resource with the given unique name, arguments, and options.
func NewBlobDataSetMapping(ctx *pulumi.Context,
	name string, args *BlobDataSetMappingArgs, opts ...pulumi.ResourceOption) (*BlobDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.FilePath == nil {
		return nil, errors.New("invalid value for required argument 'FilePath'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("Blob")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20200901:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20201001preview:BlobDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20191101:BlobDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlobDataSetMapping gets an existing BlobDataSetMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlobDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobDataSetMappingState, opts ...pulumi.ResourceOption) (*BlobDataSetMapping, error) {
	var resource BlobDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20191101:BlobDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlobDataSetMapping resources.
type blobDataSetMappingState struct {
}

type BlobDataSetMappingState struct {
}

func (BlobDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobDataSetMappingState)(nil)).Elem()
}

type blobDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// Container that has the file path.
	ContainerName string `pulumi:"containerName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// The name of the data set mapping to be created.
	DataSetMappingName *string `pulumi:"dataSetMappingName"`
	// File path within the source data set
	FilePath string `pulumi:"filePath"`
	// Kind of data set mapping.
	// Expected value is 'Blob'.
	Kind string `pulumi:"kind"`
	// File output type
	OutputType *string `pulumi:"outputType"`
	// Resource group of storage account.
	ResourceGroup string `pulumi:"resourceGroup"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
	// Storage account name of the source data set.
	StorageAccountName string `pulumi:"storageAccountName"`
	// Subscription id of storage account.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a BlobDataSetMapping resource.
type BlobDataSetMappingArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// Container that has the file path.
	ContainerName pulumi.StringInput
	// The id of the source data set.
	DataSetId pulumi.StringInput
	// The name of the data set mapping to be created.
	DataSetMappingName pulumi.StringPtrInput
	// File path within the source data set
	FilePath pulumi.StringInput
	// Kind of data set mapping.
	// Expected value is 'Blob'.
	Kind pulumi.StringInput
	// File output type
	OutputType pulumi.StringPtrInput
	// Resource group of storage account.
	ResourceGroup pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName pulumi.StringInput
	// Storage account name of the source data set.
	StorageAccountName pulumi.StringInput
	// Subscription id of storage account.
	SubscriptionId pulumi.StringInput
}

func (BlobDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobDataSetMappingArgs)(nil)).Elem()
}

type BlobDataSetMappingInput interface {
	pulumi.Input

	ToBlobDataSetMappingOutput() BlobDataSetMappingOutput
	ToBlobDataSetMappingOutputWithContext(ctx context.Context) BlobDataSetMappingOutput
}

func (*BlobDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobDataSetMapping)(nil))
}

func (i *BlobDataSetMapping) ToBlobDataSetMappingOutput() BlobDataSetMappingOutput {
	return i.ToBlobDataSetMappingOutputWithContext(context.Background())
}

func (i *BlobDataSetMapping) ToBlobDataSetMappingOutputWithContext(ctx context.Context) BlobDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobDataSetMappingOutput)
}

type BlobDataSetMappingOutput struct {
	*pulumi.OutputState
}

func (BlobDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobDataSetMapping)(nil))
}

func (o BlobDataSetMappingOutput) ToBlobDataSetMappingOutput() BlobDataSetMappingOutput {
	return o
}

func (o BlobDataSetMappingOutput) ToBlobDataSetMappingOutputWithContext(ctx context.Context) BlobDataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobDataSetMappingOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160430preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Disk resource.
type Disk struct {
	pulumi.CustomResourceState

	// the storage account type of the disk.
	AccountType pulumi.StringPtrOutput `pulumi:"accountType"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponseOutput `pulumi:"creationData"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrOutput `pulumi:"diskSizeGB"`
	// Encryption settings for disk or snapshot
	EncryptionSettings EncryptionSettingsResponsePtrOutput `pulumi:"encryptionSettings"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The Operating System type.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// A relative URI containing the VM id that has the disk attached.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The disk provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The time when the disk was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDisk registers a new resource with the given unique name, arguments, and options.
func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreationData == nil {
		return nil, errors.New("invalid value for required argument 'CreationData'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180401:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180930:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180930:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191101:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200501:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200630:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20201201:Disk"),
		},
	})
	opts = append(opts, aliases)
	var resource Disk
	err := ctx.RegisterResource("azure-native:compute/v20160430preview:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisk gets an existing Disk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("azure-native:compute/v20160430preview:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Disk resources.
type diskState struct {
}

type DiskState struct {
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	// the storage account type of the disk.
	AccountType *StorageAccountTypes `pulumi:"accountType"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationData `pulumi:"creationData"`
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	DiskName *string `pulumi:"diskName"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption settings for disk or snapshot
	EncryptionSettings *EncryptionSettings `pulumi:"encryptionSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// The Operating System type.
	OsType *OperatingSystemTypes `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Disk resource.
type DiskArgs struct {
	// the storage account type of the disk.
	AccountType StorageAccountTypesPtrInput
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataInput
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	DiskName pulumi.StringPtrInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption settings for disk or snapshot
	EncryptionSettings EncryptionSettingsPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The Operating System type.
	OsType OperatingSystemTypesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}

type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(ctx context.Context) DiskOutput
}

func (*Disk) ElementType() reflect.Type {
	return reflect.TypeOf((*Disk)(nil))
}

func (i *Disk) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i *Disk) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}

type DiskOutput struct {
	*pulumi.OutputState
}

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Disk)(nil))
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DiskOutput{})
}

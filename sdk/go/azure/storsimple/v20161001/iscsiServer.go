// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The iSCSI server.
type IscsiServer struct {
	pulumi.CustomResourceState

	// The backup policy id.
	BackupScheduleGroupId pulumi.StringOutput `pulumi:"backupScheduleGroupId"`
	// The chap id.
	ChapId pulumi.StringPtrOutput `pulumi:"chapId"`
	// The description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The reverse chap id.
	ReverseChapId pulumi.StringPtrOutput `pulumi:"reverseChapId"`
	// The storage domain id.
	StorageDomainId pulumi.StringOutput `pulumi:"storageDomainId"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIscsiServer registers a new resource with the given unique name, arguments, and options.
func NewIscsiServer(ctx *pulumi.Context,
	name string, args *IscsiServerArgs, opts ...pulumi.ResourceOption) (*IscsiServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupScheduleGroupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupScheduleGroupId'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageDomainId == nil {
		return nil, errors.New("invalid value for required argument 'StorageDomainId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20161001:IscsiServer"),
		},
	})
	opts = append(opts, aliases)
	var resource IscsiServer
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:IscsiServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIscsiServer gets an existing IscsiServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIscsiServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IscsiServerState, opts ...pulumi.ResourceOption) (*IscsiServer, error) {
	var resource IscsiServer
	err := ctx.ReadResource("azure-native:storsimple/v20161001:IscsiServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IscsiServer resources.
type iscsiServerState struct {
}

type IscsiServerState struct {
}

func (IscsiServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiServerState)(nil)).Elem()
}

type iscsiServerArgs struct {
	// The backup policy id.
	BackupScheduleGroupId string `pulumi:"backupScheduleGroupId"`
	// The chap id.
	ChapId *string `pulumi:"chapId"`
	// The description.
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The iSCSI server name.
	IscsiServerName *string `pulumi:"iscsiServerName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reverse chap id.
	ReverseChapId *string `pulumi:"reverseChapId"`
	// The storage domain id.
	StorageDomainId string `pulumi:"storageDomainId"`
}

// The set of arguments for constructing a IscsiServer resource.
type IscsiServerArgs struct {
	// The backup policy id.
	BackupScheduleGroupId pulumi.StringInput
	// The chap id.
	ChapId pulumi.StringPtrInput
	// The description.
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// The iSCSI server name.
	IscsiServerName pulumi.StringPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The reverse chap id.
	ReverseChapId pulumi.StringPtrInput
	// The storage domain id.
	StorageDomainId pulumi.StringInput
}

func (IscsiServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiServerArgs)(nil)).Elem()
}

type IscsiServerInput interface {
	pulumi.Input

	ToIscsiServerOutput() IscsiServerOutput
	ToIscsiServerOutputWithContext(ctx context.Context) IscsiServerOutput
}

func (*IscsiServer) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiServer)(nil))
}

func (i *IscsiServer) ToIscsiServerOutput() IscsiServerOutput {
	return i.ToIscsiServerOutputWithContext(context.Background())
}

func (i *IscsiServer) ToIscsiServerOutputWithContext(ctx context.Context) IscsiServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiServerOutput)
}

type IscsiServerOutput struct {
	*pulumi.OutputState
}

func (IscsiServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiServer)(nil))
}

func (o IscsiServerOutput) ToIscsiServerOutput() IscsiServerOutput {
	return o
}

func (o IscsiServerOutput) ToIscsiServerOutputWithContext(ctx context.Context) IscsiServerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IscsiServerOutput{})
}

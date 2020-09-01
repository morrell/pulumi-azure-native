// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the gallery Image Version that you want to create or update.
type GalleryImageVersion struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The publishing profile of a gallery Image Version.
	PublishingProfile GalleryImageVersionPublishingProfileResponsePtrOutput `pulumi:"publishingProfile"`
	// This is the replication status of the gallery Image Version.
	ReplicationStatus ReplicationStatusResponseOutput `pulumi:"replicationStatus"`
	// This is the storage profile of a Gallery Image Version.
	StorageProfile GalleryImageVersionStorageProfileResponseOutput `pulumi:"storageProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGalleryImageVersion registers a new resource with the given unique name, arguments, and options.
func NewGalleryImageVersion(ctx *pulumi.Context,
	name string, args *GalleryImageVersionArgs, opts ...pulumi.ResourceOption) (*GalleryImageVersion, error) {
	if args == nil || args.GalleryImageName == nil {
		return nil, errors.New("missing required argument 'GalleryImageName'")
	}
	if args == nil || args.GalleryImageVersionName == nil {
		return nil, errors.New("missing required argument 'GalleryImageVersionName'")
	}
	if args == nil || args.GalleryName == nil {
		return nil, errors.New("missing required argument 'GalleryName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageProfile == nil {
		return nil, errors.New("missing required argument 'StorageProfile'")
	}
	if args == nil {
		args = &GalleryImageVersionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:compute/v20180601:GalleryImageVersion"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190301:GalleryImageVersion"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190701:GalleryImageVersion"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20191201:GalleryImageVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryImageVersion
	err := ctx.RegisterResource("azurerm:compute/latest:GalleryImageVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryImageVersion gets an existing GalleryImageVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryImageVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageVersionState, opts ...pulumi.ResourceOption) (*GalleryImageVersion, error) {
	var resource GalleryImageVersion
	err := ctx.ReadResource("azurerm:compute/latest:GalleryImageVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryImageVersion resources.
type galleryImageVersionState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The publishing profile of a gallery Image Version.
	PublishingProfile *GalleryImageVersionPublishingProfileResponse `pulumi:"publishingProfile"`
	// This is the replication status of the gallery Image Version.
	ReplicationStatus *ReplicationStatusResponse `pulumi:"replicationStatus"`
	// This is the storage profile of a Gallery Image Version.
	StorageProfile *GalleryImageVersionStorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type GalleryImageVersionState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// The publishing profile of a gallery Image Version.
	PublishingProfile GalleryImageVersionPublishingProfileResponsePtrInput
	// This is the replication status of the gallery Image Version.
	ReplicationStatus ReplicationStatusResponsePtrInput
	// This is the storage profile of a Gallery Image Version.
	StorageProfile GalleryImageVersionStorageProfileResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (GalleryImageVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageVersionState)(nil)).Elem()
}

type galleryImageVersionArgs struct {
	// The name of the gallery Image Definition in which the Image Version is to be created.
	GalleryImageName string `pulumi:"galleryImageName"`
	// The name of the gallery Image Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
	GalleryImageVersionName string `pulumi:"galleryImageVersionName"`
	// The name of the Shared Image Gallery in which the Image Definition resides.
	GalleryName string `pulumi:"galleryName"`
	// Resource location
	Location string `pulumi:"location"`
	// The publishing profile of a gallery Image Version.
	PublishingProfile *GalleryImageVersionPublishingProfile `pulumi:"publishingProfile"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// This is the storage profile of a Gallery Image Version.
	StorageProfile GalleryImageVersionStorageProfile `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GalleryImageVersion resource.
type GalleryImageVersionArgs struct {
	// The name of the gallery Image Definition in which the Image Version is to be created.
	GalleryImageName pulumi.StringInput
	// The name of the gallery Image Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
	GalleryImageVersionName pulumi.StringInput
	// The name of the Shared Image Gallery in which the Image Definition resides.
	GalleryName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// The publishing profile of a gallery Image Version.
	PublishingProfile GalleryImageVersionPublishingProfilePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// This is the storage profile of a Gallery Image Version.
	StorageProfile GalleryImageVersionStorageProfileInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryImageVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageVersionArgs)(nil)).Elem()
}
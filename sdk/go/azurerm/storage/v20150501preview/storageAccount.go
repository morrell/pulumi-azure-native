// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The storage account.
type StorageAccount struct {
	pulumi.CustomResourceState

	// Gets the type of the storage account.
	AccountType pulumi.StringPtrOutput `pulumi:"accountType"`
	// Gets the creation date and time of the storage account in UTC.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// Gets the user assigned custom domain assigned to this storage account.
	CustomDomain CustomDomainResponsePtrOutput `pulumi:"customDomain"`
	// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is StandardGRS or StandardRAGRS.
	LastGeoFailoverTime pulumi.StringPtrOutput `pulumi:"lastGeoFailoverTime"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object.Note that StandardZRS and PremiumLRS accounts only return the blob endpoint.
	PrimaryEndpoints EndpointsResponsePtrOutput `pulumi:"primaryEndpoints"`
	// Gets the location of the primary for the storage account.
	PrimaryLocation pulumi.StringPtrOutput `pulumi:"primaryLocation"`
	// Gets the status of the storage account at the time the operation was called.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object from the secondary location of the storage account. Only available if the accountType is StandardRAGRS.
	SecondaryEndpoints EndpointsResponsePtrOutput `pulumi:"secondaryEndpoints"`
	// Gets the location of the geo replicated secondary for the storage account. Only available if the accountType is StandardGRS or StandardRAGRS.
	SecondaryLocation pulumi.StringPtrOutput `pulumi:"secondaryLocation"`
	// Gets the status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary pulumi.StringPtrOutput `pulumi:"statusOfPrimary"`
	// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the accountType is StandardGRS or StandardRAGRS.
	StatusOfSecondary pulumi.StringPtrOutput `pulumi:"statusOfSecondary"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageAccount registers a new resource with the given unique name, arguments, and options.
func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StorageAccountArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storage/latest:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20150615:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20160101:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20160501:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20161201:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20170601:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20171001:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20180201:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20180301preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20180701:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20181101:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20190401:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20190601:StorageAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccount
	err := ctx.RegisterResource("azurerm:storage/v20150501preview:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageAccount gets an existing StorageAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azurerm:storage/v20150501preview:StorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageAccount resources.
type storageAccountState struct {
	// Gets the type of the storage account.
	AccountType *string `pulumi:"accountType"`
	// Gets the creation date and time of the storage account in UTC.
	CreationTime *string `pulumi:"creationTime"`
	// Gets the user assigned custom domain assigned to this storage account.
	CustomDomain *CustomDomainResponse `pulumi:"customDomain"`
	// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is StandardGRS or StandardRAGRS.
	LastGeoFailoverTime *string `pulumi:"lastGeoFailoverTime"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object.Note that StandardZRS and PremiumLRS accounts only return the blob endpoint.
	PrimaryEndpoints *EndpointsResponse `pulumi:"primaryEndpoints"`
	// Gets the location of the primary for the storage account.
	PrimaryLocation *string `pulumi:"primaryLocation"`
	// Gets the status of the storage account at the time the operation was called.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object from the secondary location of the storage account. Only available if the accountType is StandardRAGRS.
	SecondaryEndpoints *EndpointsResponse `pulumi:"secondaryEndpoints"`
	// Gets the location of the geo replicated secondary for the storage account. Only available if the accountType is StandardGRS or StandardRAGRS.
	SecondaryLocation *string `pulumi:"secondaryLocation"`
	// Gets the status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary *string `pulumi:"statusOfPrimary"`
	// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the accountType is StandardGRS or StandardRAGRS.
	StatusOfSecondary *string `pulumi:"statusOfSecondary"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type StorageAccountState struct {
	// Gets the type of the storage account.
	AccountType pulumi.StringPtrInput
	// Gets the creation date and time of the storage account in UTC.
	CreationTime pulumi.StringPtrInput
	// Gets the user assigned custom domain assigned to this storage account.
	CustomDomain CustomDomainResponsePtrInput
	// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is StandardGRS or StandardRAGRS.
	LastGeoFailoverTime pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object.Note that StandardZRS and PremiumLRS accounts only return the blob endpoint.
	PrimaryEndpoints EndpointsResponsePtrInput
	// Gets the location of the primary for the storage account.
	PrimaryLocation pulumi.StringPtrInput
	// Gets the status of the storage account at the time the operation was called.
	ProvisioningState pulumi.StringPtrInput
	// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object from the secondary location of the storage account. Only available if the accountType is StandardRAGRS.
	SecondaryEndpoints EndpointsResponsePtrInput
	// Gets the location of the geo replicated secondary for the storage account. Only available if the accountType is StandardGRS or StandardRAGRS.
	SecondaryLocation pulumi.StringPtrInput
	// Gets the status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary pulumi.StringPtrInput
	// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the accountType is StandardGRS or StandardRAGRS.
	StatusOfSecondary pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (StorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountState)(nil)).Elem()
}

type storageAccountArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// Gets or sets the account type.
	AccountType *string `pulumi:"accountType"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageAccount resource.
type StorageAccountArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// Gets or sets the account type.
	AccountType pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountArgs)(nil)).Elem()
}
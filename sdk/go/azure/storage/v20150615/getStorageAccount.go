// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150615

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The storage account.
func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:storage/v20150615:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The storage account.
type LookupStorageAccountResult struct {
	// The type of the storage account.
	AccountType *string `pulumi:"accountType"`
	// The creation date and time of the storage account in UTC.
	CreationTime *string `pulumi:"creationTime"`
	// The custom domain the user assigned to this storage account.
	CustomDomain *CustomDomainResponse `pulumi:"customDomain"`
	// Resource Id
	Id string `pulumi:"id"`
	// The timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime *string `pulumi:"lastGeoFailoverTime"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints *EndpointsResponse `pulumi:"primaryEndpoints"`
	// The location of the primary data center for the storage account.
	PrimaryLocation *string `pulumi:"primaryLocation"`
	// The status of the storage account at the time the operation was called.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints *EndpointsResponse `pulumi:"secondaryEndpoints"`
	// The location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation *string `pulumi:"secondaryLocation"`
	// The status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary *string `pulumi:"statusOfPrimary"`
	// The status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
	StatusOfSecondary *string `pulumi:"statusOfSecondary"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

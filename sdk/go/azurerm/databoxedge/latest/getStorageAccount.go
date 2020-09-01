// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azurerm:databoxedge/latest:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage account name.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// Represents a Storage Account on the  Data Box Edge/Gateway device.
type LookupStorageAccountResult struct {
	// BlobEndpoint of Storage Account
	BlobEndpoint string `pulumi:"blobEndpoint"`
	// The Container Count. Present only for Storage Accounts with DataPolicy set to Cloud.
	ContainerCount int `pulumi:"containerCount"`
	// Data policy of the storage Account.
	DataPolicy *string `pulumi:"dataPolicy"`
	// Description for the storage Account.
	Description *string `pulumi:"description"`
	// The object name.
	Name string `pulumi:"name"`
	// Storage Account Credential Id
	StorageAccountCredentialId *string `pulumi:"storageAccountCredentialId"`
	// Current status of the storage account
	StorageAccountStatus *string `pulumi:"storageAccountStatus"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}
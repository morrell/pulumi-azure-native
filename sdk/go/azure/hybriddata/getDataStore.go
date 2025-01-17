// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hybriddata

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data store.
// API Version: 2019-06-01.
func LookupDataStore(ctx *pulumi.Context, args *LookupDataStoreArgs, opts ...pulumi.InvokeOption) (*LookupDataStoreResult, error) {
	var rv LookupDataStoreResult
	err := ctx.Invoke("azure-native:hybriddata:getDataStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataStoreArgs struct {
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName string `pulumi:"dataManagerName"`
	// The data store/repository name queried.
	DataStoreName string `pulumi:"dataStoreName"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Data store.
type LookupDataStoreResult struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets []CustomerSecretResponse `pulumi:"customerSecrets"`
	// The arm id of the data store type.
	DataStoreTypeId string `pulumi:"dataStoreTypeId"`
	// A generic json used differently by each data source type.
	ExtendedProperties interface{} `pulumi:"extendedProperties"`
	// Id of the object.
	Id string `pulumi:"id"`
	// Name of the object.
	Name string `pulumi:"name"`
	// Arm Id for the manager resource to which the data source is associated. This is optional.
	RepositoryId *string `pulumi:"repositoryId"`
	// State of the data source.
	State string `pulumi:"state"`
	// Type of the object.
	Type string `pulumi:"type"`
}

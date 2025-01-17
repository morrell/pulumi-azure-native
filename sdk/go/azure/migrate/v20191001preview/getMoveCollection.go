// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the move collection.
func LookupMoveCollection(ctx *pulumi.Context, args *LookupMoveCollectionArgs, opts ...pulumi.InvokeOption) (*LookupMoveCollectionResult, error) {
	var rv LookupMoveCollectionResult
	err := ctx.Invoke("azure-native:migrate/v20191001preview:getMoveCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMoveCollectionArgs struct {
	// The Move Collection Name.
	MoveCollectionName string `pulumi:"moveCollectionName"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Define the move collection.
type LookupMoveCollectionResult struct {
	// The etag of the resource.
	Etag string `pulumi:"etag"`
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// Defines the MSI properties of the Move Collection.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the move collection properties.
	Properties MoveCollectionPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azurerm:dbformariadb/v20180601preview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a server.
type LookupServerResult struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `pulumi:"fullyQualifiedDomainName"`
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// The location the resource resides in.
	Location string `pulumi:"location"`
	// The master server id of a replica server.
	MasterServerId *string `pulumi:"masterServerId"`
	// Resource name.
	Name string `pulumi:"name"`
	// The maximum number of replicas that a master server can have.
	ReplicaCapacity *int `pulumi:"replicaCapacity"`
	// The replication role of the server.
	ReplicationRole *string `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku *SkuResponse `pulumi:"sku"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Storage profile of a server.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// A state of a server that is visible to user.
	UserVisibleState *string `pulumi:"userVisibleState"`
	// Server version.
	Version *string `pulumi:"version"`
}
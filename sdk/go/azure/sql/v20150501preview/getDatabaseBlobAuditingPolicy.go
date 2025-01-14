// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A database blob auditing policy.
func LookupDatabaseBlobAuditingPolicy(ctx *pulumi.Context, args *LookupDatabaseBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseBlobAuditingPolicyResult, error) {
	var rv LookupDatabaseBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql/v20150501preview:getDatabaseBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseBlobAuditingPolicyArgs struct {
	// The name of the blob auditing policy.
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	// The name of the database for which the blob audit policy is defined.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A database blob auditing policy.
type LookupDatabaseBlobAuditingPolicyResult struct {
	// Specifies the Actions and Actions-Groups to audit.
	AuditActionsAndGroups []string `pulumi:"auditActionsAndGroups"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Specifies whether storageAccountAccessKey value is the storage’s secondary key.
	IsStorageSecondaryKeyInUse *bool `pulumi:"isStorageSecondaryKeyInUse"`
	// Resource kind.
	Kind string `pulumi:"kind"`
	// Resource name.
	Name string `pulumi:"name"`
	// Specifies the number of days to keep in the audit logs.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
	State string `pulumi:"state"`
	// Specifies the blob storage subscription Id.
	StorageAccountSubscriptionId *string `pulumi:"storageAccountSubscriptionId"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint is required.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// Resource type.
	Type string `pulumi:"type"`
}

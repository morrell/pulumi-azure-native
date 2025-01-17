// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
func LookupObjectReplicationPolicy(ctx *pulumi.Context, args *LookupObjectReplicationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupObjectReplicationPolicyResult, error) {
	var rv LookupObjectReplicationPolicyResult
	err := ctx.Invoke("azure-native:storage/v20210401:getObjectReplicationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupObjectReplicationPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// For the destination account, provide the value 'default'. Configure the policy on the destination account first. For the source account, provide the value of the policy ID that is returned when you download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
	ObjectReplicationPolicyId string `pulumi:"objectReplicationPolicyId"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
type LookupObjectReplicationPolicyResult struct {
	// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
	DestinationAccount string `pulumi:"destinationAccount"`
	// Indicates when the policy is enabled on the source account.
	EnabledTime string `pulumi:"enabledTime"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// A unique id for object replication policy.
	PolicyId string `pulumi:"policyId"`
	// The storage account object replication rules.
	Rules []ObjectReplicationPolicyRuleResponse `pulumi:"rules"`
	// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
	SourceAccount string `pulumi:"sourceAccount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

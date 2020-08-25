// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDisasterRecoveryConfig(ctx *pulumi.Context, args *LookupDisasterRecoveryConfigArgs, opts ...pulumi.InvokeOption) (*LookupDisasterRecoveryConfigResult, error) {
	var rv LookupDisasterRecoveryConfigResult
	err := ctx.Invoke("azurerm:eventhub/v20170401:getDisasterRecoveryConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDisasterRecoveryConfigArgs struct {
	// The Disaster Recovery configuration name
	Name string `pulumi:"name"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in List or Get Alias(Disaster Recovery configuration) operation
type LookupDisasterRecoveryConfigResult struct {
	// Alternate name specified when alias and namespace names are same.
	AlternateName *string `pulumi:"alternateName"`
	// Resource name.
	Name string `pulumi:"name"`
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace *string `pulumi:"partnerNamespace"`
	// Number of entities pending to be replicated.
	PendingReplicationOperationsCount int `pulumi:"pendingReplicationOperationsCount"`
	// Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
	ProvisioningState string `pulumi:"provisioningState"`
	// role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
	Role string `pulumi:"role"`
	// Resource type.
	Type string `pulumi:"type"`
}
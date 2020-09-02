// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupScopeAssignment(ctx *pulumi.Context, args *LookupScopeAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupScopeAssignmentResult, error) {
	var rv LookupScopeAssignmentResult
	err := ctx.Invoke("azurerm:managednetwork/v20190601preview:getScopeAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeAssignmentArgs struct {
	// The base resource of the scope assignment.
	Scope string `pulumi:"scope"`
	// The name of the scope assignment to get.
	ScopeAssignmentName string `pulumi:"scopeAssignmentName"`
}

// The Managed Network resource
type LookupScopeAssignmentResult struct {
	// The managed network ID with scope will be assigned to.
	AssignedManagedNetwork *string `pulumi:"assignedManagedNetwork"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}
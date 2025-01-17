// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a lab account.
// API Version: 2018-10-15.
func LookupLabAccount(ctx *pulumi.Context, args *LookupLabAccountArgs, opts ...pulumi.InvokeOption) (*LookupLabAccountResult, error) {
	var rv LookupLabAccountResult
	err := ctx.Invoke("azure-native:labservices:getLabAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabAccountArgs struct {
	// Specify the $expand query. Example: 'properties($expand=sizeConfiguration)'
	Expand *string `pulumi:"expand"`
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a lab account.
type LookupLabAccountResult struct {
	// Represents if region selection is enabled
	EnabledRegionSelection *bool `pulumi:"enabledRegionSelection"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Represents the size configuration under the lab account
	SizeConfiguration SizeConfigurationPropertiesResponse `pulumi:"sizeConfiguration"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

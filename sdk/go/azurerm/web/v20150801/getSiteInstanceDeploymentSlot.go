// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSiteInstanceDeploymentSlot(ctx *pulumi.Context, args *LookupSiteInstanceDeploymentSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteInstanceDeploymentSlotResult, error) {
	var rv LookupSiteInstanceDeploymentSlotResult
	err := ctx.Invoke("azurerm:web/v20150801:getSiteInstanceDeploymentSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteInstanceDeploymentSlotArgs struct {
	// Id of web app instance
	InstanceId string `pulumi:"instanceId"`
	// Id of the deployment
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
}

// Represents user credentials used for publishing activity
type LookupSiteInstanceDeploymentSlotResult struct {
	// Active
	Active *bool `pulumi:"active"`
	// Author
	Author *string `pulumi:"author"`
	// AuthorEmail
	AuthorEmail *string `pulumi:"authorEmail"`
	// Deployer
	Deployer *string `pulumi:"deployer"`
	// Detail
	Details *string `pulumi:"details"`
	// EndTime
	EndTime *string `pulumi:"endTime"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Message
	Message *string `pulumi:"message"`
	// Resource Name
	Name *string `pulumi:"name"`
	// StartTime
	StartTime *string `pulumi:"startTime"`
	// Status
	Status *int `pulumi:"status"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}
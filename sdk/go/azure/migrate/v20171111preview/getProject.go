// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Migrate Project.
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:migrate/v20171111preview:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Migrate Project.
type LookupProjectResult struct {
	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp string `pulumi:"createdTimestamp"`
	// ARM ID of the Service Map workspace created by user.
	CustomerWorkspaceId *string `pulumi:"customerWorkspaceId"`
	// Reports whether project is under discovery.
	DiscoveryStatus string `pulumi:"discoveryStatus"`
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Path reference to this project /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}
	Id string `pulumi:"id"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Name of the project.
	Name string `pulumi:"name"`
	// Number of assessments created in the project.
	NumberOfAssessments int `pulumi:"numberOfAssessments"`
	// Number of groups created in the project.
	NumberOfGroups int `pulumi:"numberOfGroups"`
	// Number of machines in the project.
	NumberOfMachines int `pulumi:"numberOfMachines"`
	// Provisioning state of the project.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Tags provided by Azure Tagging service.
	Tags interface{} `pulumi:"tags"`
	// Type of the object = [Microsoft.Migrate/projects].
	Type string `pulumi:"type"`
	// Time when this project was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp string `pulumi:"updatedTimestamp"`
}

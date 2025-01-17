// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180202

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A group created in a Migration project.
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azure-native:migrate/v20180202:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	// Unique name of a group within a project.
	GroupName string `pulumi:"groupName"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A group created in a Migration project.
type LookupGroupResult struct {
	// List of References to Assessments created on this group.
	Assessments []string `pulumi:"assessments"`
	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp string `pulumi:"createdTimestamp"`
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Path reference to this group. /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}
	Id string `pulumi:"id"`
	// List of machine names that are part of this group.
	Machines []string `pulumi:"machines"`
	// Name of the group.
	Name string `pulumi:"name"`
	// Type of the object = [Microsoft.Migrate/projects/groups].
	Type string `pulumi:"type"`
	// Time when this project was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp string `pulumi:"updatedTimestamp"`
}

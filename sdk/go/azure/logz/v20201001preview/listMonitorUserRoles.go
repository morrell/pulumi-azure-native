// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response for list of user's role for Logz.io account.
func ListMonitorUserRoles(ctx *pulumi.Context, args *ListMonitorUserRolesArgs, opts ...pulumi.InvokeOption) (*ListMonitorUserRolesResult, error) {
	var rv ListMonitorUserRolesResult
	err := ctx.Invoke("azure-native:logz/v20201001preview:listMonitorUserRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorUserRolesArgs struct {
	// Email of the user used by Logz for contacting them if needed
	EmailAddress *string `pulumi:"emailAddress"`
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for list of user's role for Logz.io account.
type ListMonitorUserRolesResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// List of user roles for Logz.io account.
	Value []UserRoleResponseResponse `pulumi:"value"`
}

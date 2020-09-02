// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupMaintenanceWindow(ctx *pulumi.Context, args *LookupMaintenanceWindowArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceWindowResult, error) {
	var rv LookupMaintenanceWindowResult
	err := ctx.Invoke("azurerm:dbformysql/v20200701privatepreview:getMaintenanceWindow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceWindowArgs struct {
	// The name of the maintenance window.
	MaintenanceWindowName string `pulumi:"maintenanceWindowName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a server maintenance window.
type LookupMaintenanceWindowResult struct {
	// The day of week of the maintenance window to start
	DayOfWeek int `pulumi:"dayOfWeek"`
	// The duration of the maintenance window to run.
	DurationInMinutes *int `pulumi:"durationInMinutes"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The starting hour of the maintenance window.
	StartHour int `pulumi:"startHour"`
	// The starting minutes of the maintenance window.
	StartMinute int `pulumi:"startMinute"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response of a list operation.
func ListMonitorMonitoredResources(ctx *pulumi.Context, args *ListMonitorMonitoredResourcesArgs, opts ...pulumi.InvokeOption) (*ListMonitorMonitoredResourcesResult, error) {
	var rv ListMonitorMonitoredResourcesResult
	err := ctx.Invoke("azure-native:logz/v20201001preview:listMonitorMonitoredResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorMonitoredResourcesArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response of a list operation.
type ListMonitorMonitoredResourcesResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// Results of a list operation.
	Value []MonitoredResourceResponse `pulumi:"value"`
}

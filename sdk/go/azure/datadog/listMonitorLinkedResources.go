// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datadog

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response of a list operation.
// API Version: 2021-03-01.
func ListMonitorLinkedResources(ctx *pulumi.Context, args *ListMonitorLinkedResourcesArgs, opts ...pulumi.InvokeOption) (*ListMonitorLinkedResourcesResult, error) {
	var rv ListMonitorLinkedResourcesResult
	err := ctx.Invoke("azure-native:datadog:listMonitorLinkedResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorLinkedResourcesArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response of a list operation.
type ListMonitorLinkedResourcesResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// Results of a list operation.
	Value []LinkedResourceResponse `pulumi:"value"`
}

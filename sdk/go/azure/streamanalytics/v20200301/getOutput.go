// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
func LookupOutput(ctx *pulumi.Context, args *LookupOutputArgs, opts ...pulumi.InvokeOption) (*LookupOutputResult, error) {
	var rv LookupOutputResult
	err := ctx.Invoke("azure-native:streamanalytics/v20200301:getOutput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutputArgs struct {
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the output.
	OutputName string `pulumi:"outputName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
type LookupOutputResult struct {
	// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
	Datasource interface{} `pulumi:"datasource"`
	// Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
	Diagnostics DiagnosticsResponse `pulumi:"diagnostics"`
	// The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag string `pulumi:"etag"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name *string `pulumi:"name"`
	// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
	Serialization interface{} `pulumi:"serialization"`
	// The size window to constrain a Stream Analytics output to.
	SizeWindow *float64 `pulumi:"sizeWindow"`
	// The time frame for filtering Stream Analytics job outputs.
	TimeWindow *string `pulumi:"timeWindow"`
	// Resource type
	Type string `pulumi:"type"`
}

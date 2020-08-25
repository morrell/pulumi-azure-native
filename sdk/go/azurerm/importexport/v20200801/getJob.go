// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azurerm:importexport/v20200801:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	// The name of the import/export job.
	Name string `pulumi:"name"`
	// The resource group name uniquely identifies the resource group within the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Contains the job information.
type LookupJobResult struct {
	// Specifies the job identity details
	Identity *IdentityDetailsResponse `pulumi:"identity"`
	// Specifies the Azure location where the job is created.
	Location *string `pulumi:"location"`
	// Specifies the name of the job.
	Name string `pulumi:"name"`
	// Specifies the job properties
	Properties JobDetailsResponse `pulumi:"properties"`
	// Specifies the tags that are assigned to the job.
	Tags map[string]interface{} `pulumi:"tags"`
	// Specifies the type of the job resource.
	Type string `pulumi:"type"`
}
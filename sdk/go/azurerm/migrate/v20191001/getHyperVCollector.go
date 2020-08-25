// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupHyperVCollector(ctx *pulumi.Context, args *LookupHyperVCollectorArgs, opts ...pulumi.InvokeOption) (*LookupHyperVCollectorResult, error) {
	var rv LookupHyperVCollectorResult
	err := ctx.Invoke("azurerm:migrate/v20191001:getHyperVCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHyperVCollectorArgs struct {
	// Unique name of a Hyper-V collector within a project.
	Name string `pulumi:"name"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupHyperVCollectorResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Name       string                      `pulumi:"name"`
	Properties CollectorPropertiesResponse `pulumi:"properties"`
	Type       string                      `pulumi:"type"`
}
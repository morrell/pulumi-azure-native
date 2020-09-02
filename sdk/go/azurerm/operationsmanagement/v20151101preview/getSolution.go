// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSolution(ctx *pulumi.Context, args *LookupSolutionArgs, opts ...pulumi.InvokeOption) (*LookupSolutionResult, error) {
	var rv LookupSolutionResult
	err := ctx.Invoke("azurerm:operationsmanagement/v20151101preview:getSolution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSolutionArgs struct {
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// User Solution Name.
	SolutionName string `pulumi:"solutionName"`
}

// The container for solution.
type LookupSolutionResult struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Plan for solution object supported by the OperationsManagement resource provider.
	Plan *SolutionPlanResponse `pulumi:"plan"`
	// Properties for solution object supported by the OperationsManagement resource provider.
	Properties SolutionPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}
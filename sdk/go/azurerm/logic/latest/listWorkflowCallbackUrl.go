// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWorkflowCallbackUrl(ctx *pulumi.Context, args *ListWorkflowCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowCallbackUrlResult, error) {
	var rv ListWorkflowCallbackUrlResult
	err := ctx.Invoke("azurerm:logic/latest:listWorkflowCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowCallbackUrlArgs struct {
	// The key type.
	KeyType *string `pulumi:"keyType"`
	// The expiry time.
	NotAfter *string `pulumi:"notAfter"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

// The workflow trigger callback URL.
type ListWorkflowCallbackUrlResult struct {
	// Gets the workflow trigger callback URL base path.
	BasePath string `pulumi:"basePath"`
	// Gets the workflow trigger callback URL HTTP method.
	Method string `pulumi:"method"`
	// Gets the workflow trigger callback URL query parameters.
	Queries *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	// Gets the workflow trigger callback URL relative path.
	RelativePath string `pulumi:"relativePath"`
	// Gets the workflow trigger callback URL relative path parameters.
	RelativePathParameters []string `pulumi:"relativePathParameters"`
	// Gets the workflow trigger callback URL.
	Value string `pulumi:"value"`
}
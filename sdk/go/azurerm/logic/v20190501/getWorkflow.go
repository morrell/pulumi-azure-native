// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWorkflow(ctx *pulumi.Context, args *LookupWorkflowArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowResult, error) {
	var rv LookupWorkflowResult
	err := ctx.Invoke("azurerm:logic/v20190501:getWorkflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowArgs struct {
	// The workflow name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The workflow type.
type LookupWorkflowResult struct {
	// The access control configuration.
	AccessControl *FlowAccessControlConfigurationResponse `pulumi:"accessControl"`
	// Gets the access endpoint.
	AccessEndpoint string `pulumi:"accessEndpoint"`
	// Gets the changed time.
	ChangedTime string `pulumi:"changedTime"`
	// Gets the created time.
	CreatedTime string `pulumi:"createdTime"`
	// The definition.
	Definition map[string]interface{} `pulumi:"definition"`
	// The endpoints configuration.
	EndpointsConfiguration *FlowEndpointsConfigurationResponse `pulumi:"endpointsConfiguration"`
	// The integration account.
	IntegrationAccount *ResourceReferenceResponse `pulumi:"integrationAccount"`
	// The integration service environment.
	IntegrationServiceEnvironment *ResourceReferenceResponse `pulumi:"integrationServiceEnvironment"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// The parameters.
	Parameters map[string]WorkflowParameterResponse `pulumi:"parameters"`
	// Gets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sku.
	Sku SkuResponse `pulumi:"sku"`
	// The state.
	State *string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type string `pulumi:"type"`
	// Gets the version.
	Version string `pulumi:"version"`
}
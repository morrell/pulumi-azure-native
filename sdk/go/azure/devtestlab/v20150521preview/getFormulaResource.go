// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A formula.
func LookupFormulaResource(ctx *pulumi.Context, args *LookupFormulaResourceArgs, opts ...pulumi.InvokeOption) (*LookupFormulaResourceResult, error) {
	var rv LookupFormulaResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getFormulaResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFormulaResourceArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the formula.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A formula.
type LookupFormulaResourceResult struct {
	// The author of the formula.
	Author *string `pulumi:"author"`
	// The creation date of the formula.
	CreationDate *string `pulumi:"creationDate"`
	// The description of the formula.
	Description *string `pulumi:"description"`
	// The content of the formula.
	FormulaContent *LabVirtualMachineResponse `pulumi:"formulaContent"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The OS type of the formula.
	OsType *string `pulumi:"osType"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// Information about a VM from which a formula is to be created.
	Vm *FormulaPropertiesFromVmResponse `pulumi:"vm"`
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTransform(ctx *pulumi.Context, args *LookupTransformArgs, opts ...pulumi.InvokeOption) (*LookupTransformResult, error) {
	var rv LookupTransformResult
	err := ctx.Invoke("azurerm:media/v20200501:getTransform", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransformArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Transform name.
	Name string `pulumi:"name"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Transform encapsulates the rules or instructions for generating desired outputs from input media, such as by transcoding or by extracting insights. After the Transform is created, it can be applied to input media by creating Jobs.
type LookupTransformResult struct {
	// The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created string `pulumi:"created"`
	// An optional verbose description of the Transform.
	Description *string `pulumi:"description"`
	// The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified string `pulumi:"lastModified"`
	// The name of the resource
	Name string `pulumi:"name"`
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs []TransformOutputResponse `pulumi:"outputs"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this function to get an Azure authentication token for the current login context.
func GetClientToken(ctx *pulumi.Context, args *GetClientTokenArgs, opts ...pulumi.InvokeOption) (*GetClientTokenResult, error) {
	var rv GetClientTokenResult
	err := ctx.Invoke("azure-native:authorization:getClientToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetClientTokenArgs struct {
	// Optional authentication endpoint. Defaults to the endpoint of Azure Resource Manager.
	Endpoint *string `pulumi:"endpoint"`
}

// Configuration values returned by getClientToken.
type GetClientTokenResult struct {
	// OAuth token for Azure Management API and SDK authentication.
	Token string `pulumi:"token"`
}

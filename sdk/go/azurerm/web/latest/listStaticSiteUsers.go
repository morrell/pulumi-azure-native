// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListStaticSiteUsers(ctx *pulumi.Context, args *ListStaticSiteUsersArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteUsersResult, error) {
	var rv ListStaticSiteUsersResult
	err := ctx.Invoke("azurerm:web/latest:listStaticSiteUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteUsersArgs struct {
	// The auth provider for the users.
	Authprovider string `pulumi:"authprovider"`
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Collection of static site custom users.
type ListStaticSiteUsersResult struct {
	// Link to next page of resources.
	NextLink string `pulumi:"nextLink"`
	// Collection of resources.
	Value []StaticSiteUserARMResourceResponse `pulumi:"value"`
}
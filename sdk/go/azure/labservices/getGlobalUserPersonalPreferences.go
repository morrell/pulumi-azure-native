// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents the PersonalPreferences for the user
// API Version: 2018-10-15.
func GetGlobalUserPersonalPreferences(ctx *pulumi.Context, args *GetGlobalUserPersonalPreferencesArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserPersonalPreferencesResult, error) {
	var rv GetGlobalUserPersonalPreferencesResult
	err := ctx.Invoke("azure-native:labservices:getGlobalUserPersonalPreferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserPersonalPreferencesArgs struct {
	// Enum indicating if user is adding or removing a favorite lab
	AddRemove *string `pulumi:"addRemove"`
	// Resource Id of the lab account
	LabAccountResourceId *string `pulumi:"labAccountResourceId"`
	// Resource Id of the lab to add/remove from the favorites list
	LabResourceId *string `pulumi:"labResourceId"`
	// The name of the user.
	UserName string `pulumi:"userName"`
}

// Represents the PersonalPreferences for the user
type GetGlobalUserPersonalPreferencesResult struct {
	// Array of favorite lab resource ids
	FavoriteLabResourceIds []string `pulumi:"favoriteLabResourceIds"`
	// Id to be used by the cache orchestrator
	Id *string `pulumi:"id"`
}

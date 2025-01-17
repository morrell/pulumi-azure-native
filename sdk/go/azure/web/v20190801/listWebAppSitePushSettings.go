// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Push settings for the App.
func ListWebAppSitePushSettings(ctx *pulumi.Context, args *ListWebAppSitePushSettingsArgs, opts ...pulumi.InvokeOption) (*ListWebAppSitePushSettingsResult, error) {
	var rv ListWebAppSitePushSettingsResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppSitePushSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSitePushSettingsArgs struct {
	// Name of web app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Push settings for the App.
type ListWebAppSitePushSettingsResult struct {
	// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
	DynamicTagsJson *string `pulumi:"dynamicTagsJson"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Gets or sets a flag indicating whether the Push endpoint is enabled.
	IsPushEnabled bool `pulumi:"isPushEnabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
	TagWhitelistJson *string `pulumi:"tagWhitelistJson"`
	// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
	// Tags can consist of alphanumeric characters and the following:
	// '_', '@', '#', '.', ':', '-'.
	// Validation should be performed at the PushRequestHandler.
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
	// Resource type.
	Type string `pulumi:"type"`
}

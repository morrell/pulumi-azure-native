// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200630preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupConfigurationProfilePreference(ctx *pulumi.Context, args *LookupConfigurationProfilePreferenceArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfilePreferenceResult, error) {
	var rv LookupConfigurationProfilePreferenceResult
	err := ctx.Invoke("azurerm:automanage/v20200630preview:getConfigurationProfilePreference", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfilePreferenceArgs struct {
	// The configuration profile preference name.
	ConfigurationProfilePreferenceName string `pulumi:"configurationProfilePreferenceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the configuration profile preference.
type LookupConfigurationProfilePreferenceResult struct {
	// Region where the VM is located.
	Location string `pulumi:"location"`
	// Name of the Automanage assignment.
	Name string `pulumi:"name"`
	// Properties of the configuration profile preference.
	Properties ConfigurationProfilePreferencePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automanage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the configuration profile preference.
// API Version: 2020-06-30-preview.
func LookupConfigurationProfilePreference(ctx *pulumi.Context, args *LookupConfigurationProfilePreferenceArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfilePreferenceResult, error) {
	var rv LookupConfigurationProfilePreferenceResult
	err := ctx.Invoke("azure-native:automanage:getConfigurationProfilePreference", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfilePreferenceArgs struct {
	// The configuration profile preference name.
	ConfigurationProfilePreferenceName string `pulumi:"configurationProfilePreferenceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the configuration profile preference.
type LookupConfigurationProfilePreferenceResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of the configuration profile preference.
	Properties ConfigurationProfilePreferencePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

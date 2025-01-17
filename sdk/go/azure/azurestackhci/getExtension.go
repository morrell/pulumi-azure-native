// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Details of a particular extension in HCI Cluster.
// API Version: 2021-01-01-preview.
func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:azurestackhci:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	// The name of the proxy resource holding details of HCI ArcSetting information.
	ArcSettingName string `pulumi:"arcSettingName"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the machine extension.
	ExtensionName string `pulumi:"extensionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Details of a particular extension in HCI Cluster.
type LookupExtensionResult struct {
	// Aggregate state of Arc Extensions across the nodes in this HCI cluster.
	AggregateState string `pulumi:"aggregateState"`
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	// The name of the resource
	Name string `pulumi:"name"`
	// State of Arc Extension in each of the nodes.
	PerNodeExtensionDetails []PerNodeExtensionStateResponse `pulumi:"perNodeExtensionDetails"`
	// Protected settings (may contain secrets).
	ProtectedSettings interface{} `pulumi:"protectedSettings"`
	// Provisioning state of the Extension proxy resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings interface{} `pulumi:"settings"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
}

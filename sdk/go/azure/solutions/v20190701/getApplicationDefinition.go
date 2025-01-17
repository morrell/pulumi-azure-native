// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about managed application definition.
func LookupApplicationDefinition(ctx *pulumi.Context, args *LookupApplicationDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationDefinitionResult, error) {
	var rv LookupApplicationDefinitionResult
	err := ctx.Invoke("azure-native:solutions/v20190701:getApplicationDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationDefinitionArgs struct {
	// The name of the managed application definition.
	ApplicationDefinitionName string `pulumi:"applicationDefinitionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about managed application definition.
type LookupApplicationDefinitionResult struct {
	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts []ApplicationDefinitionArtifactResponse `pulumi:"artifacts"`
	// The managed application provider authorizations.
	Authorizations []ApplicationAuthorizationResponse `pulumi:"authorizations"`
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition interface{} `pulumi:"createUiDefinition"`
	// The managed application deployment policy.
	DeploymentPolicy *ApplicationDeploymentPolicyResponse `pulumi:"deploymentPolicy"`
	// The managed application definition description.
	Description *string `pulumi:"description"`
	// The managed application definition display name.
	DisplayName *string `pulumi:"displayName"`
	// Resource ID
	Id string `pulumi:"id"`
	// A value indicating whether the package is enabled or not.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Resource location
	Location *string `pulumi:"location"`
	// The managed application lock level.
	LockLevel string `pulumi:"lockLevel"`
	// The managed application locking policy.
	LockingPolicy *ApplicationPackageLockingPolicyDefinitionResponse `pulumi:"lockingPolicy"`
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate interface{} `pulumi:"mainTemplate"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The managed application management policy that determines publisher's access to the managed resource group.
	ManagementPolicy *ApplicationManagementPolicyResponse `pulumi:"managementPolicy"`
	// Resource name
	Name string `pulumi:"name"`
	// The managed application notification policy.
	NotificationPolicy *ApplicationNotificationPolicyResponse `pulumi:"notificationPolicy"`
	// The managed application definition package file Uri. Use this element
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The managed application provider policies.
	Policies []ApplicationPolicyResponse `pulumi:"policies"`
	// The SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

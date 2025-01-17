// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The link resource format.
// API Version: 2017-04-26.
func LookupLink(ctx *pulumi.Context, args *LookupLinkArgs, opts ...pulumi.InvokeOption) (*LookupLinkResult, error) {
	var rv LookupLinkResult
	err := ctx.Invoke("azure-native:customerinsights:getLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the link.
	LinkName string `pulumi:"linkName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The link resource format.
type LookupLinkResult struct {
	// Localized descriptions for the Link.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the Link.
	DisplayName map[string]string `pulumi:"displayName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The link name.
	LinkName string `pulumi:"linkName"`
	// The set of properties mappings between the source and target Types.
	Mappings []TypePropertiesMappingResponse `pulumi:"mappings"`
	// Resource name.
	Name string `pulumi:"name"`
	// Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
	OperationType *string `pulumi:"operationType"`
	// The properties that represent the participating profile.
	ParticipantPropertyReferences []ParticipantPropertyReferenceResponse `pulumi:"participantPropertyReferences"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
	ReferenceOnly *bool `pulumi:"referenceOnly"`
	// Type of source entity.
	SourceEntityType string `pulumi:"sourceEntityType"`
	// Name of the source Entity Type.
	SourceEntityTypeName string `pulumi:"sourceEntityTypeName"`
	// Type of target entity.
	TargetEntityType string `pulumi:"targetEntityType"`
	// Name of the target Entity Type.
	TargetEntityTypeName string `pulumi:"targetEntityTypeName"`
	// The hub name.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}

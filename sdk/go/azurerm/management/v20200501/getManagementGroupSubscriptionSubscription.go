// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupManagementGroupSubscriptionSubscription(ctx *pulumi.Context, args *LookupManagementGroupSubscriptionSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupSubscriptionSubscriptionResult, error) {
	var rv LookupManagementGroupSubscriptionSubscriptionResult
	err := ctx.Invoke("azurerm:management/v20200501:getManagementGroupSubscriptionSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupSubscriptionSubscriptionArgs struct {
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
}

// The details of subscription under management group.
type LookupManagementGroupSubscriptionSubscriptionResult struct {
	// The friendly name of the subscription.
	DisplayName *string `pulumi:"displayName"`
	// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
	Name string `pulumi:"name"`
	// The ID of the parent management group.
	Parent *DescendantParentGroupInfoResponse `pulumi:"parent"`
	// The state of the subscription.
	State *string `pulumi:"state"`
	// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
	Tenant *string `pulumi:"tenant"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
	Type string `pulumi:"type"`
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountPartner(ctx *pulumi.Context, args *LookupIntegrationAccountPartnerArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountPartnerResult, error) {
	var rv LookupIntegrationAccountPartnerResult
	err := ctx.Invoke("azure-native:logic/v20150801preview:getIntegrationAccountPartner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountPartnerArgs struct {
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The integration account partner name.
	PartnerName string `pulumi:"partnerName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupIntegrationAccountPartnerResult struct {
	// The changed time.
	ChangedTime string `pulumi:"changedTime"`
	// The partner content.
	Content *PartnerContentResponse `pulumi:"content"`
	// The created time.
	CreatedTime string `pulumi:"createdTime"`
	// The resource id.
	Id *string `pulumi:"id"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The partner type.
	PartnerType *string `pulumi:"partnerType"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

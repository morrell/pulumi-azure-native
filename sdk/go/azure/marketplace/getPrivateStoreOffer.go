// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package marketplace

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The privateStore offer data structure.
// API Version: 2020-01-01.
func LookupPrivateStoreOffer(ctx *pulumi.Context, args *LookupPrivateStoreOfferArgs, opts ...pulumi.InvokeOption) (*LookupPrivateStoreOfferResult, error) {
	var rv LookupPrivateStoreOfferResult
	err := ctx.Invoke("azure-native:marketplace:getPrivateStoreOffer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateStoreOfferArgs struct {
	// The offer ID to update or delete
	OfferId string `pulumi:"offerId"`
	// The store ID - must use the tenant ID
	PrivateStoreId string `pulumi:"privateStoreId"`
}

// The privateStore offer data structure.
type LookupPrivateStoreOfferResult struct {
	// Private store offer creation date
	CreatedAt string `pulumi:"createdAt"`
	// Identifier for purposes of race condition
	ETag *string `pulumi:"eTag"`
	// Icon File Uris
	IconFileUris map[string]string `pulumi:"iconFileUris"`
	// The resource ID.
	Id string `pulumi:"id"`
	// Private store offer modification date
	ModifiedAt string `pulumi:"modifiedAt"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// It will be displayed prominently in the marketplace
	OfferDisplayName string `pulumi:"offerDisplayName"`
	// Offer plans
	Plans []PlanResponse `pulumi:"plans"`
	// Private store unique id
	PrivateStoreId string `pulumi:"privateStoreId"`
	// Publisher name that will be displayed prominently in the marketplace
	PublisherDisplayName string `pulumi:"publisherDisplayName"`
	// Plan ids limitation for this offer
	SpecificPlanIdsLimitation []string `pulumi:"specificPlanIdsLimitation"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Offers unique id
	UniqueOfferId string `pulumi:"uniqueOfferId"`
	// Indicating whether the offer was not updated to db (true = not updated). If the allow list is identical to the existed one in db, the offer would not be updated.
	UpdateSuppressedDueIdempotence *bool `pulumi:"updateSuppressedDueIdempotence"`
}

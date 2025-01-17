// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180202

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An assessment created for a group in the Migration project.
func LookupAssessment(ctx *pulumi.Context, args *LookupAssessmentArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentResult, error) {
	var rv LookupAssessmentResult
	err := ctx.Invoke("azure-native:migrate/v20180202:getAssessment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentArgs struct {
	// Unique name of an assessment within a project.
	AssessmentName string `pulumi:"assessmentName"`
	// Unique name of a group within a project.
	GroupName string `pulumi:"groupName"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An assessment created for a group in the Migration project.
type LookupAssessmentResult struct {
	// AHUB discount on windows virtual machines.
	AzureHybridUseBenefit string `pulumi:"azureHybridUseBenefit"`
	// Target Azure location for which the machines should be assessed. These enums are the same as used by Compute API.
	AzureLocation string `pulumi:"azureLocation"`
	// Offer code according to which cost estimation is done.
	AzureOfferCode string `pulumi:"azureOfferCode"`
	// Pricing tier for Size evaluation.
	AzurePricingTier string `pulumi:"azurePricingTier"`
	// Storage Redundancy type offered by Azure.
	AzureStorageRedundancy string `pulumi:"azureStorageRedundancy"`
	// Confidence rating percentage for assessment. Can be in the range [0, 100].
	ConfidenceRatingInPercentage float64 `pulumi:"confidenceRatingInPercentage"`
	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp string `pulumi:"createdTimestamp"`
	// Currency to report prices in.
	Currency string `pulumi:"currency"`
	// Custom discount percentage to be applied on final costs. Can be in the range [0, 100].
	DiscountPercentage float64 `pulumi:"discountPercentage"`
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Path reference to this assessment. /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessment/{assessmentName}
	Id string `pulumi:"id"`
	// Monthly network cost estimate for the machines that are part of this assessment as a group, for a 31-day month.
	MonthlyBandwidthCost float64 `pulumi:"monthlyBandwidthCost"`
	// Monthly compute cost estimate for the machines that are part of this assessment as a group, for a 31-day month.
	MonthlyComputeCost float64 `pulumi:"monthlyComputeCost"`
	// Monthly storage cost estimate for the machines that are part of this assessment as a group, for a 31-day month.
	MonthlyStorageCost float64 `pulumi:"monthlyStorageCost"`
	// Unique name of an assessment.
	Name string `pulumi:"name"`
	// Number of assessed machines part of this assessment.
	NumberOfMachines int `pulumi:"numberOfMachines"`
	// Percentile of performance data used to recommend Azure size.
	Percentile string `pulumi:"percentile"`
	// Time when the Azure Prices were queried. Date-Time represented in ISO-8601 format.
	PricesTimestamp string `pulumi:"pricesTimestamp"`
	// Scaling factor used over utilization data to add a performance buffer for new machines to be created in Azure. Min Value = 1.0, Max value = 1.9, Default = 1.3.
	ScalingFactor float64 `pulumi:"scalingFactor"`
	// Assessment sizing criterion.
	SizingCriterion string `pulumi:"sizingCriterion"`
	// User configurable setting that describes the status of the assessment.
	Stage string `pulumi:"stage"`
	// Whether the assessment has been created and is valid.
	Status string `pulumi:"status"`
	// Time range of performance data used to recommend a size.
	TimeRange string `pulumi:"timeRange"`
	// Type of the object = [Microsoft.Migrate/projects/groups/assessments].
	Type string `pulumi:"type"`
	// Time when this project was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp string `pulumi:"updatedTimestamp"`
}

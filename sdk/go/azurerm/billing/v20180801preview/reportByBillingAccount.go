// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A report resource.
type ReportByBillingAccount struct {
	pulumi.CustomResourceState

	// Has definition for the report.
	Definition ReportDefinitionResponseOutput `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Has schedule information for the report.
	Schedule ReportScheduleResponsePtrOutput `pulumi:"schedule"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReportByBillingAccount registers a new resource with the given unique name, arguments, and options.
func NewReportByBillingAccount(ctx *pulumi.Context,
	name string, args *ReportByBillingAccountArgs, opts ...pulumi.ResourceOption) (*ReportByBillingAccount, error) {
	if args == nil || args.BillingAccountId == nil {
		return nil, errors.New("missing required argument 'BillingAccountId'")
	}
	if args == nil || args.Definition == nil {
		return nil, errors.New("missing required argument 'Definition'")
	}
	if args == nil || args.DeliveryInfo == nil {
		return nil, errors.New("missing required argument 'DeliveryInfo'")
	}
	if args == nil || args.ReportName == nil {
		return nil, errors.New("missing required argument 'ReportName'")
	}
	if args == nil {
		args = &ReportByBillingAccountArgs{}
	}
	var resource ReportByBillingAccount
	err := ctx.RegisterResource("azurerm:billing/v20180801preview:ReportByBillingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReportByBillingAccount gets an existing ReportByBillingAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReportByBillingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportByBillingAccountState, opts ...pulumi.ResourceOption) (*ReportByBillingAccount, error) {
	var resource ReportByBillingAccount
	err := ctx.ReadResource("azurerm:billing/v20180801preview:ReportByBillingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReportByBillingAccount resources.
type reportByBillingAccountState struct {
	// Has definition for the report.
	Definition *ReportDefinitionResponse `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo *ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Has schedule information for the report.
	Schedule *ReportScheduleResponse `pulumi:"schedule"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ReportByBillingAccountState struct {
	// Has definition for the report.
	Definition ReportDefinitionResponsePtrInput
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoResponsePtrInput
	// The format of the report being delivered.
	Format pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Has schedule information for the report.
	Schedule ReportScheduleResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ReportByBillingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByBillingAccountState)(nil)).Elem()
}

type reportByBillingAccountArgs struct {
	// BillingAccount ID
	BillingAccountId string `pulumi:"billingAccountId"`
	// Has definition for the report.
	Definition ReportDefinition `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfo `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Report Name.
	ReportName string `pulumi:"reportName"`
	// Has schedule information for the report.
	Schedule *ReportSchedule `pulumi:"schedule"`
}

// The set of arguments for constructing a ReportByBillingAccount resource.
type ReportByBillingAccountArgs struct {
	// BillingAccount ID
	BillingAccountId pulumi.StringInput
	// Has definition for the report.
	Definition ReportDefinitionInput
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoInput
	// The format of the report being delivered.
	Format pulumi.StringPtrInput
	// Report Name.
	ReportName pulumi.StringInput
	// Has schedule information for the report.
	Schedule ReportSchedulePtrInput
}

func (ReportByBillingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByBillingAccountArgs)(nil)).Elem()
}
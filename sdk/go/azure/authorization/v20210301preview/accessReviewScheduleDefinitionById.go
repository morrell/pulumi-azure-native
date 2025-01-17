// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Access Review Schedule Definition.
type AccessReviewScheduleDefinitionById struct {
	pulumi.CustomResourceState

	// The role assignment state eligible/active to review
	AssignmentState pulumi.StringOutput `pulumi:"assignmentState"`
	// Flag to indicate whether auto-apply capability, to automatically change the target object access resource, is enabled. If not enabled, a user must, after the review completes, apply the access review.
	AutoApplyDecisionsEnabled pulumi.BoolPtrOutput `pulumi:"autoApplyDecisionsEnabled"`
	// This is the collection of backup reviewers.
	BackupReviewers AccessReviewReviewerResponseArrayOutput `pulumi:"backupReviewers"`
	// This specifies the behavior for the autoReview feature when an access review completes.
	DefaultDecision pulumi.StringPtrOutput `pulumi:"defaultDecision"`
	// Flag to indicate whether reviewers are required to provide a justification when reviewing access.
	DefaultDecisionEnabled pulumi.BoolPtrOutput `pulumi:"defaultDecisionEnabled"`
	// The description provided by the access review creator and visible to admins.
	DescriptionForAdmins pulumi.StringPtrOutput `pulumi:"descriptionForAdmins"`
	// The description provided by the access review creator to be shown to reviewers.
	DescriptionForReviewers pulumi.StringPtrOutput `pulumi:"descriptionForReviewers"`
	// The display name for the schedule definition.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// Duration users are inactive for. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	InactiveDuration pulumi.StringPtrOutput `pulumi:"inactiveDuration"`
	// The duration in days for an instance.
	InstanceDurationInDays pulumi.IntPtrOutput `pulumi:"instanceDurationInDays"`
	// This is the collection of instances returned when one does an expand on it.
	Instances AccessReviewInstanceResponseArrayOutput `pulumi:"instances"`
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// Flag to indicate whether the reviewer is required to pass justification when recording a decision.
	JustificationRequiredOnApproval pulumi.BoolPtrOutput `pulumi:"justificationRequiredOnApproval"`
	// Flag to indicate whether sending mails to reviewers and the review creator is enabled.
	MailNotificationsEnabled pulumi.BoolPtrOutput `pulumi:"mailNotificationsEnabled"`
	// The access review schedule definition unique id.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences pulumi.IntPtrOutput `pulumi:"numberOfOccurrences"`
	// The identity id
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The identity display name
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// The identity type user/servicePrincipal to review
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// Flag to indicate whether showing recommendations to reviewers is enabled.
	RecommendationsEnabled pulumi.BoolPtrOutput `pulumi:"recommendationsEnabled"`
	// Flag to indicate whether sending reminder emails to reviewers are enabled.
	ReminderNotificationsEnabled pulumi.BoolPtrOutput `pulumi:"reminderNotificationsEnabled"`
	// ResourceId in which this review is getting created
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// This is the collection of reviewers.
	Reviewers AccessReviewReviewerResponseArrayOutput `pulumi:"reviewers"`
	// This field specifies the type of reviewers for a review. Usually for a review, reviewers are explicitly assigned. However, in some cases, the reviewers may not be assigned and instead be chosen dynamically. For example managers review or self review.
	ReviewersType pulumi.StringOutput `pulumi:"reviewersType"`
	// This is used to indicate the role being reviewed
	RoleDefinitionId pulumi.StringOutput `pulumi:"roleDefinitionId"`
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate pulumi.StringPtrOutput `pulumi:"startDate"`
	// This read-only field specifies the status of an accessReview.
	Status pulumi.StringOutput `pulumi:"status"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The user principal name(if valid)
	UserPrincipalName pulumi.StringOutput `pulumi:"userPrincipalName"`
}

// NewAccessReviewScheduleDefinitionById registers a new resource with the given unique name, arguments, and options.
func NewAccessReviewScheduleDefinitionById(ctx *pulumi.Context,
	name string, args *AccessReviewScheduleDefinitionByIdArgs, opts ...pulumi.ResourceOption) (*AccessReviewScheduleDefinitionById, error) {
	if args == nil {
		args = &AccessReviewScheduleDefinitionByIdArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:authorization/v20210301preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20180501preview:AccessReviewScheduleDefinitionById"),
		},
	})
	opts = append(opts, aliases)
	var resource AccessReviewScheduleDefinitionById
	err := ctx.RegisterResource("azure-native:authorization/v20210301preview:AccessReviewScheduleDefinitionById", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessReviewScheduleDefinitionById gets an existing AccessReviewScheduleDefinitionById resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessReviewScheduleDefinitionById(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessReviewScheduleDefinitionByIdState, opts ...pulumi.ResourceOption) (*AccessReviewScheduleDefinitionById, error) {
	var resource AccessReviewScheduleDefinitionById
	err := ctx.ReadResource("azure-native:authorization/v20210301preview:AccessReviewScheduleDefinitionById", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessReviewScheduleDefinitionById resources.
type accessReviewScheduleDefinitionByIdState struct {
}

type AccessReviewScheduleDefinitionByIdState struct {
}

func (AccessReviewScheduleDefinitionByIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewScheduleDefinitionByIdState)(nil)).Elem()
}

type accessReviewScheduleDefinitionByIdArgs struct {
	// Flag to indicate whether auto-apply capability, to automatically change the target object access resource, is enabled. If not enabled, a user must, after the review completes, apply the access review.
	AutoApplyDecisionsEnabled *bool `pulumi:"autoApplyDecisionsEnabled"`
	// This is the collection of backup reviewers.
	BackupReviewers []AccessReviewReviewer `pulumi:"backupReviewers"`
	// This specifies the behavior for the autoReview feature when an access review completes.
	DefaultDecision *string `pulumi:"defaultDecision"`
	// Flag to indicate whether reviewers are required to provide a justification when reviewing access.
	DefaultDecisionEnabled *bool `pulumi:"defaultDecisionEnabled"`
	// The description provided by the access review creator and visible to admins.
	DescriptionForAdmins *string `pulumi:"descriptionForAdmins"`
	// The description provided by the access review creator to be shown to reviewers.
	DescriptionForReviewers *string `pulumi:"descriptionForReviewers"`
	// The display name for the schedule definition.
	DisplayName *string `pulumi:"displayName"`
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate *string `pulumi:"endDate"`
	// Duration users are inactive for. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	InactiveDuration *string `pulumi:"inactiveDuration"`
	// The duration in days for an instance.
	InstanceDurationInDays *int `pulumi:"instanceDurationInDays"`
	// This is the collection of instances returned when one does an expand on it.
	Instances []AccessReviewInstance `pulumi:"instances"`
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval *int `pulumi:"interval"`
	// Flag to indicate whether the reviewer is required to pass justification when recording a decision.
	JustificationRequiredOnApproval *bool `pulumi:"justificationRequiredOnApproval"`
	// Flag to indicate whether sending mails to reviewers and the review creator is enabled.
	MailNotificationsEnabled *bool `pulumi:"mailNotificationsEnabled"`
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences *int `pulumi:"numberOfOccurrences"`
	// Flag to indicate whether showing recommendations to reviewers is enabled.
	RecommendationsEnabled *bool `pulumi:"recommendationsEnabled"`
	// Flag to indicate whether sending reminder emails to reviewers are enabled.
	ReminderNotificationsEnabled *bool `pulumi:"reminderNotificationsEnabled"`
	// This is the collection of reviewers.
	Reviewers []AccessReviewReviewer `pulumi:"reviewers"`
	// The id of the access review schedule definition.
	ScheduleDefinitionId *string `pulumi:"scheduleDefinitionId"`
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate *string `pulumi:"startDate"`
	// The recurrence range type. The possible values are: endDate, noEnd, numbered.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a AccessReviewScheduleDefinitionById resource.
type AccessReviewScheduleDefinitionByIdArgs struct {
	// Flag to indicate whether auto-apply capability, to automatically change the target object access resource, is enabled. If not enabled, a user must, after the review completes, apply the access review.
	AutoApplyDecisionsEnabled pulumi.BoolPtrInput
	// This is the collection of backup reviewers.
	BackupReviewers AccessReviewReviewerArrayInput
	// This specifies the behavior for the autoReview feature when an access review completes.
	DefaultDecision pulumi.StringPtrInput
	// Flag to indicate whether reviewers are required to provide a justification when reviewing access.
	DefaultDecisionEnabled pulumi.BoolPtrInput
	// The description provided by the access review creator and visible to admins.
	DescriptionForAdmins pulumi.StringPtrInput
	// The description provided by the access review creator to be shown to reviewers.
	DescriptionForReviewers pulumi.StringPtrInput
	// The display name for the schedule definition.
	DisplayName pulumi.StringPtrInput
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate pulumi.StringPtrInput
	// Duration users are inactive for. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	InactiveDuration pulumi.StringPtrInput
	// The duration in days for an instance.
	InstanceDurationInDays pulumi.IntPtrInput
	// This is the collection of instances returned when one does an expand on it.
	Instances AccessReviewInstanceArrayInput
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval pulumi.IntPtrInput
	// Flag to indicate whether the reviewer is required to pass justification when recording a decision.
	JustificationRequiredOnApproval pulumi.BoolPtrInput
	// Flag to indicate whether sending mails to reviewers and the review creator is enabled.
	MailNotificationsEnabled pulumi.BoolPtrInput
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences pulumi.IntPtrInput
	// Flag to indicate whether showing recommendations to reviewers is enabled.
	RecommendationsEnabled pulumi.BoolPtrInput
	// Flag to indicate whether sending reminder emails to reviewers are enabled.
	ReminderNotificationsEnabled pulumi.BoolPtrInput
	// This is the collection of reviewers.
	Reviewers AccessReviewReviewerArrayInput
	// The id of the access review schedule definition.
	ScheduleDefinitionId pulumi.StringPtrInput
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate pulumi.StringPtrInput
	// The recurrence range type. The possible values are: endDate, noEnd, numbered.
	Type pulumi.StringPtrInput
}

func (AccessReviewScheduleDefinitionByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewScheduleDefinitionByIdArgs)(nil)).Elem()
}

type AccessReviewScheduleDefinitionByIdInput interface {
	pulumi.Input

	ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput
	ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput
}

func (*AccessReviewScheduleDefinitionById) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewScheduleDefinitionById)(nil))
}

func (i *AccessReviewScheduleDefinitionById) ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput {
	return i.ToAccessReviewScheduleDefinitionByIdOutputWithContext(context.Background())
}

func (i *AccessReviewScheduleDefinitionById) ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewScheduleDefinitionByIdOutput)
}

type AccessReviewScheduleDefinitionByIdOutput struct {
	*pulumi.OutputState
}

func (AccessReviewScheduleDefinitionByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewScheduleDefinitionById)(nil))
}

func (o AccessReviewScheduleDefinitionByIdOutput) ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput {
	return o
}

func (o AccessReviewScheduleDefinitionByIdOutput) ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccessReviewScheduleDefinitionByIdOutput{})
}

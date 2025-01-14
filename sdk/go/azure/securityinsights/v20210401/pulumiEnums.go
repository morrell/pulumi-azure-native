// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The reason the incident was closed
type IncidentClassification string

const (
	// Incident classification was undetermined
	IncidentClassificationUndetermined = IncidentClassification("Undetermined")
	// Incident was true positive
	IncidentClassificationTruePositive = IncidentClassification("TruePositive")
	// Incident was benign positive
	IncidentClassificationBenignPositive = IncidentClassification("BenignPositive")
	// Incident was false positive
	IncidentClassificationFalsePositive = IncidentClassification("FalsePositive")
)

func (IncidentClassification) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentClassification)(nil)).Elem()
}

func (e IncidentClassification) ToIncidentClassificationOutput() IncidentClassificationOutput {
	return pulumi.ToOutput(e).(IncidentClassificationOutput)
}

func (e IncidentClassification) ToIncidentClassificationOutputWithContext(ctx context.Context) IncidentClassificationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentClassificationOutput)
}

func (e IncidentClassification) ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput {
	return e.ToIncidentClassificationPtrOutputWithContext(context.Background())
}

func (e IncidentClassification) ToIncidentClassificationPtrOutputWithContext(ctx context.Context) IncidentClassificationPtrOutput {
	return IncidentClassification(e).ToIncidentClassificationOutputWithContext(ctx).ToIncidentClassificationPtrOutputWithContext(ctx)
}

func (e IncidentClassification) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassification) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassification) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentClassification) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentClassificationOutput struct{ *pulumi.OutputState }

func (IncidentClassificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentClassification)(nil)).Elem()
}

func (o IncidentClassificationOutput) ToIncidentClassificationOutput() IncidentClassificationOutput {
	return o
}

func (o IncidentClassificationOutput) ToIncidentClassificationOutputWithContext(ctx context.Context) IncidentClassificationOutput {
	return o
}

func (o IncidentClassificationOutput) ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput {
	return o.ToIncidentClassificationPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationOutput) ToIncidentClassificationPtrOutputWithContext(ctx context.Context) IncidentClassificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentClassification) *IncidentClassification {
		return &v
	}).(IncidentClassificationPtrOutput)
}

func (o IncidentClassificationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentClassificationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentClassification) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentClassificationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentClassification) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentClassificationPtrOutput struct{ *pulumi.OutputState }

func (IncidentClassificationPtrOutput) ElementType() reflect.Type {
	return incidentClassificationPtrType
}

func (o IncidentClassificationPtrOutput) ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput {
	return o
}

func (o IncidentClassificationPtrOutput) ToIncidentClassificationPtrOutputWithContext(ctx context.Context) IncidentClassificationPtrOutput {
	return o
}

func (o IncidentClassificationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentClassification) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o IncidentClassificationPtrOutput) Elem() IncidentClassificationOutput {
	return o.ApplyT(func(v *IncidentClassification) IncidentClassification {
		var ret IncidentClassification
		if v != nil {
			ret = *v
		}
		return ret
	}).(IncidentClassificationOutput)
}

// IncidentClassificationInput is an input type that accepts IncidentClassificationArgs and IncidentClassificationOutput values.
// You can construct a concrete instance of `IncidentClassificationInput` via:
//
//          IncidentClassificationArgs{...}
type IncidentClassificationInput interface {
	pulumi.Input

	ToIncidentClassificationOutput() IncidentClassificationOutput
	ToIncidentClassificationOutputWithContext(context.Context) IncidentClassificationOutput
}

var incidentClassificationPtrType = reflect.TypeOf((**IncidentClassification)(nil)).Elem()

type IncidentClassificationPtrInput interface {
	pulumi.Input

	ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput
	ToIncidentClassificationPtrOutputWithContext(context.Context) IncidentClassificationPtrOutput
}

type incidentClassificationPtr string

func IncidentClassificationPtr(v string) IncidentClassificationPtrInput {
	return (*incidentClassificationPtr)(&v)
}

func (*incidentClassificationPtr) ElementType() reflect.Type {
	return incidentClassificationPtrType
}

func (in *incidentClassificationPtr) ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput {
	return pulumi.ToOutput(in).(IncidentClassificationPtrOutput)
}

func (in *incidentClassificationPtr) ToIncidentClassificationPtrOutputWithContext(ctx context.Context) IncidentClassificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentClassificationPtrOutput)
}

// The classification reason the incident was closed with
type IncidentClassificationReason string

const (
	// Classification reason was suspicious activity
	IncidentClassificationReasonSuspiciousActivity = IncidentClassificationReason("SuspiciousActivity")
	// Classification reason was suspicious but expected
	IncidentClassificationReasonSuspiciousButExpected = IncidentClassificationReason("SuspiciousButExpected")
	// Classification reason was incorrect alert logic
	IncidentClassificationReasonIncorrectAlertLogic = IncidentClassificationReason("IncorrectAlertLogic")
	// Classification reason was inaccurate data
	IncidentClassificationReasonInaccurateData = IncidentClassificationReason("InaccurateData")
)

func (IncidentClassificationReason) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentClassificationReason)(nil)).Elem()
}

func (e IncidentClassificationReason) ToIncidentClassificationReasonOutput() IncidentClassificationReasonOutput {
	return pulumi.ToOutput(e).(IncidentClassificationReasonOutput)
}

func (e IncidentClassificationReason) ToIncidentClassificationReasonOutputWithContext(ctx context.Context) IncidentClassificationReasonOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentClassificationReasonOutput)
}

func (e IncidentClassificationReason) ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput {
	return e.ToIncidentClassificationReasonPtrOutputWithContext(context.Background())
}

func (e IncidentClassificationReason) ToIncidentClassificationReasonPtrOutputWithContext(ctx context.Context) IncidentClassificationReasonPtrOutput {
	return IncidentClassificationReason(e).ToIncidentClassificationReasonOutputWithContext(ctx).ToIncidentClassificationReasonPtrOutputWithContext(ctx)
}

func (e IncidentClassificationReason) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassificationReason) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassificationReason) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentClassificationReason) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentClassificationReasonOutput struct{ *pulumi.OutputState }

func (IncidentClassificationReasonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentClassificationReason)(nil)).Elem()
}

func (o IncidentClassificationReasonOutput) ToIncidentClassificationReasonOutput() IncidentClassificationReasonOutput {
	return o
}

func (o IncidentClassificationReasonOutput) ToIncidentClassificationReasonOutputWithContext(ctx context.Context) IncidentClassificationReasonOutput {
	return o
}

func (o IncidentClassificationReasonOutput) ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput {
	return o.ToIncidentClassificationReasonPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationReasonOutput) ToIncidentClassificationReasonPtrOutputWithContext(ctx context.Context) IncidentClassificationReasonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentClassificationReason) *IncidentClassificationReason {
		return &v
	}).(IncidentClassificationReasonPtrOutput)
}

func (o IncidentClassificationReasonOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentClassificationReasonOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentClassificationReason) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentClassificationReasonOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationReasonOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentClassificationReason) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentClassificationReasonPtrOutput struct{ *pulumi.OutputState }

func (IncidentClassificationReasonPtrOutput) ElementType() reflect.Type {
	return incidentClassificationReasonPtrType
}

func (o IncidentClassificationReasonPtrOutput) ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput {
	return o
}

func (o IncidentClassificationReasonPtrOutput) ToIncidentClassificationReasonPtrOutputWithContext(ctx context.Context) IncidentClassificationReasonPtrOutput {
	return o
}

func (o IncidentClassificationReasonPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationReasonPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentClassificationReason) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o IncidentClassificationReasonPtrOutput) Elem() IncidentClassificationReasonOutput {
	return o.ApplyT(func(v *IncidentClassificationReason) IncidentClassificationReason {
		var ret IncidentClassificationReason
		if v != nil {
			ret = *v
		}
		return ret
	}).(IncidentClassificationReasonOutput)
}

// IncidentClassificationReasonInput is an input type that accepts IncidentClassificationReasonArgs and IncidentClassificationReasonOutput values.
// You can construct a concrete instance of `IncidentClassificationReasonInput` via:
//
//          IncidentClassificationReasonArgs{...}
type IncidentClassificationReasonInput interface {
	pulumi.Input

	ToIncidentClassificationReasonOutput() IncidentClassificationReasonOutput
	ToIncidentClassificationReasonOutputWithContext(context.Context) IncidentClassificationReasonOutput
}

var incidentClassificationReasonPtrType = reflect.TypeOf((**IncidentClassificationReason)(nil)).Elem()

type IncidentClassificationReasonPtrInput interface {
	pulumi.Input

	ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput
	ToIncidentClassificationReasonPtrOutputWithContext(context.Context) IncidentClassificationReasonPtrOutput
}

type incidentClassificationReasonPtr string

func IncidentClassificationReasonPtr(v string) IncidentClassificationReasonPtrInput {
	return (*incidentClassificationReasonPtr)(&v)
}

func (*incidentClassificationReasonPtr) ElementType() reflect.Type {
	return incidentClassificationReasonPtrType
}

func (in *incidentClassificationReasonPtr) ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput {
	return pulumi.ToOutput(in).(IncidentClassificationReasonPtrOutput)
}

func (in *incidentClassificationReasonPtr) ToIncidentClassificationReasonPtrOutputWithContext(ctx context.Context) IncidentClassificationReasonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentClassificationReasonPtrOutput)
}

// The severity of the incident
type IncidentSeverity string

const (
	// High severity
	IncidentSeverityHigh = IncidentSeverity("High")
	// Medium severity
	IncidentSeverityMedium = IncidentSeverity("Medium")
	// Low severity
	IncidentSeverityLow = IncidentSeverity("Low")
	// Informational severity
	IncidentSeverityInformational = IncidentSeverity("Informational")
)

func (IncidentSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentSeverity)(nil)).Elem()
}

func (e IncidentSeverity) ToIncidentSeverityOutput() IncidentSeverityOutput {
	return pulumi.ToOutput(e).(IncidentSeverityOutput)
}

func (e IncidentSeverity) ToIncidentSeverityOutputWithContext(ctx context.Context) IncidentSeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentSeverityOutput)
}

func (e IncidentSeverity) ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput {
	return e.ToIncidentSeverityPtrOutputWithContext(context.Background())
}

func (e IncidentSeverity) ToIncidentSeverityPtrOutputWithContext(ctx context.Context) IncidentSeverityPtrOutput {
	return IncidentSeverity(e).ToIncidentSeverityOutputWithContext(ctx).ToIncidentSeverityPtrOutputWithContext(ctx)
}

func (e IncidentSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentSeverityOutput struct{ *pulumi.OutputState }

func (IncidentSeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentSeverity)(nil)).Elem()
}

func (o IncidentSeverityOutput) ToIncidentSeverityOutput() IncidentSeverityOutput {
	return o
}

func (o IncidentSeverityOutput) ToIncidentSeverityOutputWithContext(ctx context.Context) IncidentSeverityOutput {
	return o
}

func (o IncidentSeverityOutput) ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput {
	return o.ToIncidentSeverityPtrOutputWithContext(context.Background())
}

func (o IncidentSeverityOutput) ToIncidentSeverityPtrOutputWithContext(ctx context.Context) IncidentSeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentSeverity) *IncidentSeverity {
		return &v
	}).(IncidentSeverityPtrOutput)
}

func (o IncidentSeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentSeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentSeverity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentSeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentSeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentSeverity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentSeverityPtrOutput struct{ *pulumi.OutputState }

func (IncidentSeverityPtrOutput) ElementType() reflect.Type {
	return incidentSeverityPtrType
}

func (o IncidentSeverityPtrOutput) ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput {
	return o
}

func (o IncidentSeverityPtrOutput) ToIncidentSeverityPtrOutputWithContext(ctx context.Context) IncidentSeverityPtrOutput {
	return o
}

func (o IncidentSeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentSeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentSeverity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o IncidentSeverityPtrOutput) Elem() IncidentSeverityOutput {
	return o.ApplyT(func(v *IncidentSeverity) IncidentSeverity {
		var ret IncidentSeverity
		if v != nil {
			ret = *v
		}
		return ret
	}).(IncidentSeverityOutput)
}

// IncidentSeverityInput is an input type that accepts IncidentSeverityArgs and IncidentSeverityOutput values.
// You can construct a concrete instance of `IncidentSeverityInput` via:
//
//          IncidentSeverityArgs{...}
type IncidentSeverityInput interface {
	pulumi.Input

	ToIncidentSeverityOutput() IncidentSeverityOutput
	ToIncidentSeverityOutputWithContext(context.Context) IncidentSeverityOutput
}

var incidentSeverityPtrType = reflect.TypeOf((**IncidentSeverity)(nil)).Elem()

type IncidentSeverityPtrInput interface {
	pulumi.Input

	ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput
	ToIncidentSeverityPtrOutputWithContext(context.Context) IncidentSeverityPtrOutput
}

type incidentSeverityPtr string

func IncidentSeverityPtr(v string) IncidentSeverityPtrInput {
	return (*incidentSeverityPtr)(&v)
}

func (*incidentSeverityPtr) ElementType() reflect.Type {
	return incidentSeverityPtrType
}

func (in *incidentSeverityPtr) ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput {
	return pulumi.ToOutput(in).(IncidentSeverityPtrOutput)
}

func (in *incidentSeverityPtr) ToIncidentSeverityPtrOutputWithContext(ctx context.Context) IncidentSeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentSeverityPtrOutput)
}

// The status of the incident
type IncidentStatus string

const (
	// An active incident which isn't being handled currently
	IncidentStatusNew = IncidentStatus("New")
	// An active incident which is being handled
	IncidentStatusActive = IncidentStatus("Active")
	// A non-active incident
	IncidentStatusClosed = IncidentStatus("Closed")
)

func (IncidentStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentStatus)(nil)).Elem()
}

func (e IncidentStatus) ToIncidentStatusOutput() IncidentStatusOutput {
	return pulumi.ToOutput(e).(IncidentStatusOutput)
}

func (e IncidentStatus) ToIncidentStatusOutputWithContext(ctx context.Context) IncidentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentStatusOutput)
}

func (e IncidentStatus) ToIncidentStatusPtrOutput() IncidentStatusPtrOutput {
	return e.ToIncidentStatusPtrOutputWithContext(context.Background())
}

func (e IncidentStatus) ToIncidentStatusPtrOutputWithContext(ctx context.Context) IncidentStatusPtrOutput {
	return IncidentStatus(e).ToIncidentStatusOutputWithContext(ctx).ToIncidentStatusPtrOutputWithContext(ctx)
}

func (e IncidentStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentStatusOutput struct{ *pulumi.OutputState }

func (IncidentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentStatus)(nil)).Elem()
}

func (o IncidentStatusOutput) ToIncidentStatusOutput() IncidentStatusOutput {
	return o
}

func (o IncidentStatusOutput) ToIncidentStatusOutputWithContext(ctx context.Context) IncidentStatusOutput {
	return o
}

func (o IncidentStatusOutput) ToIncidentStatusPtrOutput() IncidentStatusPtrOutput {
	return o.ToIncidentStatusPtrOutputWithContext(context.Background())
}

func (o IncidentStatusOutput) ToIncidentStatusPtrOutputWithContext(ctx context.Context) IncidentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentStatus) *IncidentStatus {
		return &v
	}).(IncidentStatusPtrOutput)
}

func (o IncidentStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentStatusPtrOutput struct{ *pulumi.OutputState }

func (IncidentStatusPtrOutput) ElementType() reflect.Type {
	return incidentStatusPtrType
}

func (o IncidentStatusPtrOutput) ToIncidentStatusPtrOutput() IncidentStatusPtrOutput {
	return o
}

func (o IncidentStatusPtrOutput) ToIncidentStatusPtrOutputWithContext(ctx context.Context) IncidentStatusPtrOutput {
	return o
}

func (o IncidentStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o IncidentStatusPtrOutput) Elem() IncidentStatusOutput {
	return o.ApplyT(func(v *IncidentStatus) IncidentStatus {
		var ret IncidentStatus
		if v != nil {
			ret = *v
		}
		return ret
	}).(IncidentStatusOutput)
}

// IncidentStatusInput is an input type that accepts IncidentStatusArgs and IncidentStatusOutput values.
// You can construct a concrete instance of `IncidentStatusInput` via:
//
//          IncidentStatusArgs{...}
type IncidentStatusInput interface {
	pulumi.Input

	ToIncidentStatusOutput() IncidentStatusOutput
	ToIncidentStatusOutputWithContext(context.Context) IncidentStatusOutput
}

var incidentStatusPtrType = reflect.TypeOf((**IncidentStatus)(nil)).Elem()

type IncidentStatusPtrInput interface {
	pulumi.Input

	ToIncidentStatusPtrOutput() IncidentStatusPtrOutput
	ToIncidentStatusPtrOutputWithContext(context.Context) IncidentStatusPtrOutput
}

type incidentStatusPtr string

func IncidentStatusPtr(v string) IncidentStatusPtrInput {
	return (*incidentStatusPtr)(&v)
}

func (*incidentStatusPtr) ElementType() reflect.Type {
	return incidentStatusPtrType
}

func (in *incidentStatusPtr) ToIncidentStatusPtrOutput() IncidentStatusPtrOutput {
	return pulumi.ToOutput(in).(IncidentStatusPtrOutput)
}

func (in *incidentStatusPtr) ToIncidentStatusPtrOutputWithContext(ctx context.Context) IncidentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentStatusPtrOutput)
}

// The source of the watchlist
type Source string

const (
	Source_Local_file     = Source("Local file")
	Source_Remote_storage = Source("Remote storage")
)

func (Source) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (e Source) ToSourceOutput() SourceOutput {
	return pulumi.ToOutput(e).(SourceOutput)
}

func (e Source) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceOutput)
}

func (e Source) ToSourcePtrOutput() SourcePtrOutput {
	return e.ToSourcePtrOutputWithContext(context.Background())
}

func (e Source) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return Source(e).ToSourceOutputWithContext(ctx).ToSourcePtrOutputWithContext(ctx)
}

func (e Source) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Source) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Source) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Source) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

func (o SourceOutput) ToSourcePtrOutput() SourcePtrOutput {
	return o.ToSourcePtrOutputWithContext(context.Background())
}

func (o SourceOutput) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Source) *Source {
		return &v
	}).(SourcePtrOutput)
}

func (o SourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Source) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Source) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourcePtrOutput struct{ *pulumi.OutputState }

func (SourcePtrOutput) ElementType() reflect.Type {
	return sourcePtrType
}

func (o SourcePtrOutput) ToSourcePtrOutput() SourcePtrOutput {
	return o
}

func (o SourcePtrOutput) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return o
}

func (o SourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Source) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o SourcePtrOutput) Elem() SourceOutput {
	return o.ApplyT(func(v *Source) Source {
		var ret Source
		if v != nil {
			ret = *v
		}
		return ret
	}).(SourceOutput)
}

// SourceInput is an input type that accepts SourceArgs and SourceOutput values.
// You can construct a concrete instance of `SourceInput` via:
//
//          SourceArgs{...}
type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(context.Context) SourceOutput
}

var sourcePtrType = reflect.TypeOf((**Source)(nil)).Elem()

type SourcePtrInput interface {
	pulumi.Input

	ToSourcePtrOutput() SourcePtrOutput
	ToSourcePtrOutputWithContext(context.Context) SourcePtrOutput
}

type sourcePtr string

func SourcePtr(v string) SourcePtrInput {
	return (*sourcePtr)(&v)
}

func (*sourcePtr) ElementType() reflect.Type {
	return sourcePtrType
}

func (in *sourcePtr) ToSourcePtrOutput() SourcePtrOutput {
	return pulumi.ToOutput(in).(SourcePtrOutput)
}

func (in *sourcePtr) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourcePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentClassificationOutput{})
	pulumi.RegisterOutputType(IncidentClassificationPtrOutput{})
	pulumi.RegisterOutputType(IncidentClassificationReasonOutput{})
	pulumi.RegisterOutputType(IncidentClassificationReasonPtrOutput{})
	pulumi.RegisterOutputType(IncidentSeverityOutput{})
	pulumi.RegisterOutputType(IncidentSeverityPtrOutput{})
	pulumi.RegisterOutputType(IncidentStatusOutput{})
	pulumi.RegisterOutputType(IncidentStatusPtrOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourcePtrOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of the export. Note that 'Usage' is equivalent to 'ActualCost' and is applicable to exports that do not yet provide data for charges or amortization for service reservations.
type ExportType string

const (
	ExportTypeUsage         = ExportType("Usage")
	ExportTypeActualCost    = ExportType("ActualCost")
	ExportTypeAmortizedCost = ExportType("AmortizedCost")
)

func (ExportType) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportType)(nil)).Elem()
}

func (e ExportType) ToExportTypeOutput() ExportTypeOutput {
	return pulumi.ToOutput(e).(ExportTypeOutput)
}

func (e ExportType) ToExportTypeOutputWithContext(ctx context.Context) ExportTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExportTypeOutput)
}

func (e ExportType) ToExportTypePtrOutput() ExportTypePtrOutput {
	return e.ToExportTypePtrOutputWithContext(context.Background())
}

func (e ExportType) ToExportTypePtrOutputWithContext(ctx context.Context) ExportTypePtrOutput {
	return ExportType(e).ToExportTypeOutputWithContext(ctx).ToExportTypePtrOutputWithContext(ctx)
}

func (e ExportType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExportType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExportType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExportType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExportTypeOutput struct{ *pulumi.OutputState }

func (ExportTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportType)(nil)).Elem()
}

func (o ExportTypeOutput) ToExportTypeOutput() ExportTypeOutput {
	return o
}

func (o ExportTypeOutput) ToExportTypeOutputWithContext(ctx context.Context) ExportTypeOutput {
	return o
}

func (o ExportTypeOutput) ToExportTypePtrOutput() ExportTypePtrOutput {
	return o.ToExportTypePtrOutputWithContext(context.Background())
}

func (o ExportTypeOutput) ToExportTypePtrOutputWithContext(ctx context.Context) ExportTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportType) *ExportType {
		return &v
	}).(ExportTypePtrOutput)
}

func (o ExportTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExportTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExportType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExportTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExportTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExportType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExportTypePtrOutput struct{ *pulumi.OutputState }

func (ExportTypePtrOutput) ElementType() reflect.Type {
	return exportTypePtrType
}

func (o ExportTypePtrOutput) ToExportTypePtrOutput() ExportTypePtrOutput {
	return o
}

func (o ExportTypePtrOutput) ToExportTypePtrOutputWithContext(ctx context.Context) ExportTypePtrOutput {
	return o
}

func (o ExportTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExportTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExportType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o ExportTypePtrOutput) Elem() ExportTypeOutput {
	return o.ApplyT(func(v *ExportType) ExportType {
		var ret ExportType
		if v != nil {
			ret = *v
		}
		return ret
	}).(ExportTypeOutput)
}

// ExportTypeInput is an input type that accepts ExportTypeArgs and ExportTypeOutput values.
// You can construct a concrete instance of `ExportTypeInput` via:
//
//          ExportTypeArgs{...}
type ExportTypeInput interface {
	pulumi.Input

	ToExportTypeOutput() ExportTypeOutput
	ToExportTypeOutputWithContext(context.Context) ExportTypeOutput
}

var exportTypePtrType = reflect.TypeOf((**ExportType)(nil)).Elem()

type ExportTypePtrInput interface {
	pulumi.Input

	ToExportTypePtrOutput() ExportTypePtrOutput
	ToExportTypePtrOutputWithContext(context.Context) ExportTypePtrOutput
}

type exportTypePtr string

func ExportTypePtr(v string) ExportTypePtrInput {
	return (*exportTypePtr)(&v)
}

func (*exportTypePtr) ElementType() reflect.Type {
	return exportTypePtrType
}

func (in *exportTypePtr) ToExportTypePtrOutput() ExportTypePtrOutput {
	return pulumi.ToOutput(in).(ExportTypePtrOutput)
}

func (in *exportTypePtr) ToExportTypePtrOutputWithContext(ctx context.Context) ExportTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExportTypePtrOutput)
}

// The format of the export being delivered. Currently only 'Csv' is supported.
type FormatType string

const (
	FormatTypeCsv = FormatType("Csv")
)

func (FormatType) ElementType() reflect.Type {
	return reflect.TypeOf((*FormatType)(nil)).Elem()
}

func (e FormatType) ToFormatTypeOutput() FormatTypeOutput {
	return pulumi.ToOutput(e).(FormatTypeOutput)
}

func (e FormatType) ToFormatTypeOutputWithContext(ctx context.Context) FormatTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FormatTypeOutput)
}

func (e FormatType) ToFormatTypePtrOutput() FormatTypePtrOutput {
	return e.ToFormatTypePtrOutputWithContext(context.Background())
}

func (e FormatType) ToFormatTypePtrOutputWithContext(ctx context.Context) FormatTypePtrOutput {
	return FormatType(e).ToFormatTypeOutputWithContext(ctx).ToFormatTypePtrOutputWithContext(ctx)
}

func (e FormatType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FormatType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FormatType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FormatType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FormatTypeOutput struct{ *pulumi.OutputState }

func (FormatTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FormatType)(nil)).Elem()
}

func (o FormatTypeOutput) ToFormatTypeOutput() FormatTypeOutput {
	return o
}

func (o FormatTypeOutput) ToFormatTypeOutputWithContext(ctx context.Context) FormatTypeOutput {
	return o
}

func (o FormatTypeOutput) ToFormatTypePtrOutput() FormatTypePtrOutput {
	return o.ToFormatTypePtrOutputWithContext(context.Background())
}

func (o FormatTypeOutput) ToFormatTypePtrOutputWithContext(ctx context.Context) FormatTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FormatType) *FormatType {
		return &v
	}).(FormatTypePtrOutput)
}

func (o FormatTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FormatTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FormatType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FormatTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FormatTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FormatType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FormatTypePtrOutput struct{ *pulumi.OutputState }

func (FormatTypePtrOutput) ElementType() reflect.Type {
	return formatTypePtrType
}

func (o FormatTypePtrOutput) ToFormatTypePtrOutput() FormatTypePtrOutput {
	return o
}

func (o FormatTypePtrOutput) ToFormatTypePtrOutputWithContext(ctx context.Context) FormatTypePtrOutput {
	return o
}

func (o FormatTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FormatTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FormatType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o FormatTypePtrOutput) Elem() FormatTypeOutput {
	return o.ApplyT(func(v *FormatType) FormatType {
		var ret FormatType
		if v != nil {
			ret = *v
		}
		return ret
	}).(FormatTypeOutput)
}

// FormatTypeInput is an input type that accepts FormatTypeArgs and FormatTypeOutput values.
// You can construct a concrete instance of `FormatTypeInput` via:
//
//          FormatTypeArgs{...}
type FormatTypeInput interface {
	pulumi.Input

	ToFormatTypeOutput() FormatTypeOutput
	ToFormatTypeOutputWithContext(context.Context) FormatTypeOutput
}

var formatTypePtrType = reflect.TypeOf((**FormatType)(nil)).Elem()

type FormatTypePtrInput interface {
	pulumi.Input

	ToFormatTypePtrOutput() FormatTypePtrOutput
	ToFormatTypePtrOutputWithContext(context.Context) FormatTypePtrOutput
}

type formatTypePtr string

func FormatTypePtr(v string) FormatTypePtrInput {
	return (*formatTypePtr)(&v)
}

func (*formatTypePtr) ElementType() reflect.Type {
	return formatTypePtrType
}

func (in *formatTypePtr) ToFormatTypePtrOutput() FormatTypePtrOutput {
	return pulumi.ToOutput(in).(FormatTypePtrOutput)
}

func (in *formatTypePtr) ToFormatTypePtrOutputWithContext(ctx context.Context) FormatTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FormatTypePtrOutput)
}

// The granularity of rows in the export. Currently only 'Daily' is supported.
type GranularityType string

const (
	GranularityTypeDaily = GranularityType("Daily")
)

func (GranularityType) ElementType() reflect.Type {
	return reflect.TypeOf((*GranularityType)(nil)).Elem()
}

func (e GranularityType) ToGranularityTypeOutput() GranularityTypeOutput {
	return pulumi.ToOutput(e).(GranularityTypeOutput)
}

func (e GranularityType) ToGranularityTypeOutputWithContext(ctx context.Context) GranularityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GranularityTypeOutput)
}

func (e GranularityType) ToGranularityTypePtrOutput() GranularityTypePtrOutput {
	return e.ToGranularityTypePtrOutputWithContext(context.Background())
}

func (e GranularityType) ToGranularityTypePtrOutputWithContext(ctx context.Context) GranularityTypePtrOutput {
	return GranularityType(e).ToGranularityTypeOutputWithContext(ctx).ToGranularityTypePtrOutputWithContext(ctx)
}

func (e GranularityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GranularityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GranularityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GranularityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GranularityTypeOutput struct{ *pulumi.OutputState }

func (GranularityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GranularityType)(nil)).Elem()
}

func (o GranularityTypeOutput) ToGranularityTypeOutput() GranularityTypeOutput {
	return o
}

func (o GranularityTypeOutput) ToGranularityTypeOutputWithContext(ctx context.Context) GranularityTypeOutput {
	return o
}

func (o GranularityTypeOutput) ToGranularityTypePtrOutput() GranularityTypePtrOutput {
	return o.ToGranularityTypePtrOutputWithContext(context.Background())
}

func (o GranularityTypeOutput) ToGranularityTypePtrOutputWithContext(ctx context.Context) GranularityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GranularityType) *GranularityType {
		return &v
	}).(GranularityTypePtrOutput)
}

func (o GranularityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GranularityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GranularityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GranularityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GranularityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GranularityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GranularityTypePtrOutput struct{ *pulumi.OutputState }

func (GranularityTypePtrOutput) ElementType() reflect.Type {
	return granularityTypePtrType
}

func (o GranularityTypePtrOutput) ToGranularityTypePtrOutput() GranularityTypePtrOutput {
	return o
}

func (o GranularityTypePtrOutput) ToGranularityTypePtrOutputWithContext(ctx context.Context) GranularityTypePtrOutput {
	return o
}

func (o GranularityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GranularityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GranularityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o GranularityTypePtrOutput) Elem() GranularityTypeOutput {
	return o.ApplyT(func(v *GranularityType) GranularityType {
		var ret GranularityType
		if v != nil {
			ret = *v
		}
		return ret
	}).(GranularityTypeOutput)
}

// GranularityTypeInput is an input type that accepts GranularityTypeArgs and GranularityTypeOutput values.
// You can construct a concrete instance of `GranularityTypeInput` via:
//
//          GranularityTypeArgs{...}
type GranularityTypeInput interface {
	pulumi.Input

	ToGranularityTypeOutput() GranularityTypeOutput
	ToGranularityTypeOutputWithContext(context.Context) GranularityTypeOutput
}

var granularityTypePtrType = reflect.TypeOf((**GranularityType)(nil)).Elem()

type GranularityTypePtrInput interface {
	pulumi.Input

	ToGranularityTypePtrOutput() GranularityTypePtrOutput
	ToGranularityTypePtrOutputWithContext(context.Context) GranularityTypePtrOutput
}

type granularityTypePtr string

func GranularityTypePtr(v string) GranularityTypePtrInput {
	return (*granularityTypePtr)(&v)
}

func (*granularityTypePtr) ElementType() reflect.Type {
	return granularityTypePtrType
}

func (in *granularityTypePtr) ToGranularityTypePtrOutput() GranularityTypePtrOutput {
	return pulumi.ToOutput(in).(GranularityTypePtrOutput)
}

func (in *granularityTypePtr) ToGranularityTypePtrOutputWithContext(ctx context.Context) GranularityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GranularityTypePtrOutput)
}

// The schedule recurrence.
type RecurrenceType string

const (
	RecurrenceTypeDaily    = RecurrenceType("Daily")
	RecurrenceTypeWeekly   = RecurrenceType("Weekly")
	RecurrenceTypeMonthly  = RecurrenceType("Monthly")
	RecurrenceTypeAnnually = RecurrenceType("Annually")
)

func (RecurrenceType) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceType)(nil)).Elem()
}

func (e RecurrenceType) ToRecurrenceTypeOutput() RecurrenceTypeOutput {
	return pulumi.ToOutput(e).(RecurrenceTypeOutput)
}

func (e RecurrenceType) ToRecurrenceTypeOutputWithContext(ctx context.Context) RecurrenceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecurrenceTypeOutput)
}

func (e RecurrenceType) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return e.ToRecurrenceTypePtrOutputWithContext(context.Background())
}

func (e RecurrenceType) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return RecurrenceType(e).ToRecurrenceTypeOutputWithContext(ctx).ToRecurrenceTypePtrOutputWithContext(ctx)
}

func (e RecurrenceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecurrenceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecurrenceTypeOutput struct{ *pulumi.OutputState }

func (RecurrenceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceType)(nil)).Elem()
}

func (o RecurrenceTypeOutput) ToRecurrenceTypeOutput() RecurrenceTypeOutput {
	return o
}

func (o RecurrenceTypeOutput) ToRecurrenceTypeOutputWithContext(ctx context.Context) RecurrenceTypeOutput {
	return o
}

func (o RecurrenceTypeOutput) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return o.ToRecurrenceTypePtrOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrenceType) *RecurrenceType {
		return &v
	}).(RecurrenceTypePtrOutput)
}

func (o RecurrenceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecurrenceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecurrenceTypePtrOutput struct{ *pulumi.OutputState }

func (RecurrenceTypePtrOutput) ElementType() reflect.Type {
	return recurrenceTypePtrType
}

func (o RecurrenceTypePtrOutput) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return o
}

func (o RecurrenceTypePtrOutput) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return o
}

func (o RecurrenceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecurrenceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o RecurrenceTypePtrOutput) Elem() RecurrenceTypeOutput {
	return o.ApplyT(func(v *RecurrenceType) RecurrenceType {
		var ret RecurrenceType
		if v != nil {
			ret = *v
		}
		return ret
	}).(RecurrenceTypeOutput)
}

// RecurrenceTypeInput is an input type that accepts RecurrenceTypeArgs and RecurrenceTypeOutput values.
// You can construct a concrete instance of `RecurrenceTypeInput` via:
//
//          RecurrenceTypeArgs{...}
type RecurrenceTypeInput interface {
	pulumi.Input

	ToRecurrenceTypeOutput() RecurrenceTypeOutput
	ToRecurrenceTypeOutputWithContext(context.Context) RecurrenceTypeOutput
}

var recurrenceTypePtrType = reflect.TypeOf((**RecurrenceType)(nil)).Elem()

type RecurrenceTypePtrInput interface {
	pulumi.Input

	ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput
	ToRecurrenceTypePtrOutputWithContext(context.Context) RecurrenceTypePtrOutput
}

type recurrenceTypePtr string

func RecurrenceTypePtr(v string) RecurrenceTypePtrInput {
	return (*recurrenceTypePtr)(&v)
}

func (*recurrenceTypePtr) ElementType() reflect.Type {
	return recurrenceTypePtrType
}

func (in *recurrenceTypePtr) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return pulumi.ToOutput(in).(RecurrenceTypePtrOutput)
}

func (in *recurrenceTypePtr) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecurrenceTypePtrOutput)
}

// The status of the export's schedule. If 'Inactive', the export's schedule is paused.
type StatusType string

const (
	StatusTypeActive   = StatusType("Active")
	StatusTypeInactive = StatusType("Inactive")
)

func (StatusType) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusType)(nil)).Elem()
}

func (e StatusType) ToStatusTypeOutput() StatusTypeOutput {
	return pulumi.ToOutput(e).(StatusTypeOutput)
}

func (e StatusType) ToStatusTypeOutputWithContext(ctx context.Context) StatusTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusTypeOutput)
}

func (e StatusType) ToStatusTypePtrOutput() StatusTypePtrOutput {
	return e.ToStatusTypePtrOutputWithContext(context.Background())
}

func (e StatusType) ToStatusTypePtrOutputWithContext(ctx context.Context) StatusTypePtrOutput {
	return StatusType(e).ToStatusTypeOutputWithContext(ctx).ToStatusTypePtrOutputWithContext(ctx)
}

func (e StatusType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StatusType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusTypeOutput struct{ *pulumi.OutputState }

func (StatusTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusType)(nil)).Elem()
}

func (o StatusTypeOutput) ToStatusTypeOutput() StatusTypeOutput {
	return o
}

func (o StatusTypeOutput) ToStatusTypeOutputWithContext(ctx context.Context) StatusTypeOutput {
	return o
}

func (o StatusTypeOutput) ToStatusTypePtrOutput() StatusTypePtrOutput {
	return o.ToStatusTypePtrOutputWithContext(context.Background())
}

func (o StatusTypeOutput) ToStatusTypePtrOutputWithContext(ctx context.Context) StatusTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StatusType) *StatusType {
		return &v
	}).(StatusTypePtrOutput)
}

func (o StatusTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusTypePtrOutput struct{ *pulumi.OutputState }

func (StatusTypePtrOutput) ElementType() reflect.Type {
	return statusTypePtrType
}

func (o StatusTypePtrOutput) ToStatusTypePtrOutput() StatusTypePtrOutput {
	return o
}

func (o StatusTypePtrOutput) ToStatusTypePtrOutputWithContext(ctx context.Context) StatusTypePtrOutput {
	return o
}

func (o StatusTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StatusType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o StatusTypePtrOutput) Elem() StatusTypeOutput {
	return o.ApplyT(func(v *StatusType) StatusType {
		var ret StatusType
		if v != nil {
			ret = *v
		}
		return ret
	}).(StatusTypeOutput)
}

// StatusTypeInput is an input type that accepts StatusTypeArgs and StatusTypeOutput values.
// You can construct a concrete instance of `StatusTypeInput` via:
//
//          StatusTypeArgs{...}
type StatusTypeInput interface {
	pulumi.Input

	ToStatusTypeOutput() StatusTypeOutput
	ToStatusTypeOutputWithContext(context.Context) StatusTypeOutput
}

var statusTypePtrType = reflect.TypeOf((**StatusType)(nil)).Elem()

type StatusTypePtrInput interface {
	pulumi.Input

	ToStatusTypePtrOutput() StatusTypePtrOutput
	ToStatusTypePtrOutputWithContext(context.Context) StatusTypePtrOutput
}

type statusTypePtr string

func StatusTypePtr(v string) StatusTypePtrInput {
	return (*statusTypePtr)(&v)
}

func (*statusTypePtr) ElementType() reflect.Type {
	return statusTypePtrType
}

func (in *statusTypePtr) ToStatusTypePtrOutput() StatusTypePtrOutput {
	return pulumi.ToOutput(in).(StatusTypePtrOutput)
}

func (in *statusTypePtr) ToStatusTypePtrOutputWithContext(ctx context.Context) StatusTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusTypePtrOutput)
}

// The time frame for pulling data for the export. If custom, then a specific time period must be provided.
type TimeframeType string

const (
	TimeframeTypeMonthToDate         = TimeframeType("MonthToDate")
	TimeframeTypeBillingMonthToDate  = TimeframeType("BillingMonthToDate")
	TimeframeTypeTheLastMonth        = TimeframeType("TheLastMonth")
	TimeframeTypeTheLastBillingMonth = TimeframeType("TheLastBillingMonth")
	TimeframeTypeWeekToDate          = TimeframeType("WeekToDate")
	TimeframeTypeCustom              = TimeframeType("Custom")
)

func (TimeframeType) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeframeType)(nil)).Elem()
}

func (e TimeframeType) ToTimeframeTypeOutput() TimeframeTypeOutput {
	return pulumi.ToOutput(e).(TimeframeTypeOutput)
}

func (e TimeframeType) ToTimeframeTypeOutputWithContext(ctx context.Context) TimeframeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TimeframeTypeOutput)
}

func (e TimeframeType) ToTimeframeTypePtrOutput() TimeframeTypePtrOutput {
	return e.ToTimeframeTypePtrOutputWithContext(context.Background())
}

func (e TimeframeType) ToTimeframeTypePtrOutputWithContext(ctx context.Context) TimeframeTypePtrOutput {
	return TimeframeType(e).ToTimeframeTypeOutputWithContext(ctx).ToTimeframeTypePtrOutputWithContext(ctx)
}

func (e TimeframeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeframeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeframeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TimeframeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TimeframeTypeOutput struct{ *pulumi.OutputState }

func (TimeframeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeframeType)(nil)).Elem()
}

func (o TimeframeTypeOutput) ToTimeframeTypeOutput() TimeframeTypeOutput {
	return o
}

func (o TimeframeTypeOutput) ToTimeframeTypeOutputWithContext(ctx context.Context) TimeframeTypeOutput {
	return o
}

func (o TimeframeTypeOutput) ToTimeframeTypePtrOutput() TimeframeTypePtrOutput {
	return o.ToTimeframeTypePtrOutputWithContext(context.Background())
}

func (o TimeframeTypeOutput) ToTimeframeTypePtrOutputWithContext(ctx context.Context) TimeframeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeframeType) *TimeframeType {
		return &v
	}).(TimeframeTypePtrOutput)
}

func (o TimeframeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TimeframeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeframeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TimeframeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeframeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeframeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TimeframeTypePtrOutput struct{ *pulumi.OutputState }

func (TimeframeTypePtrOutput) ElementType() reflect.Type {
	return timeframeTypePtrType
}

func (o TimeframeTypePtrOutput) ToTimeframeTypePtrOutput() TimeframeTypePtrOutput {
	return o
}

func (o TimeframeTypePtrOutput) ToTimeframeTypePtrOutputWithContext(ctx context.Context) TimeframeTypePtrOutput {
	return o
}

func (o TimeframeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeframeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TimeframeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o TimeframeTypePtrOutput) Elem() TimeframeTypeOutput {
	return o.ApplyT(func(v *TimeframeType) TimeframeType {
		var ret TimeframeType
		if v != nil {
			ret = *v
		}
		return ret
	}).(TimeframeTypeOutput)
}

// TimeframeTypeInput is an input type that accepts TimeframeTypeArgs and TimeframeTypeOutput values.
// You can construct a concrete instance of `TimeframeTypeInput` via:
//
//          TimeframeTypeArgs{...}
type TimeframeTypeInput interface {
	pulumi.Input

	ToTimeframeTypeOutput() TimeframeTypeOutput
	ToTimeframeTypeOutputWithContext(context.Context) TimeframeTypeOutput
}

var timeframeTypePtrType = reflect.TypeOf((**TimeframeType)(nil)).Elem()

type TimeframeTypePtrInput interface {
	pulumi.Input

	ToTimeframeTypePtrOutput() TimeframeTypePtrOutput
	ToTimeframeTypePtrOutputWithContext(context.Context) TimeframeTypePtrOutput
}

type timeframeTypePtr string

func TimeframeTypePtr(v string) TimeframeTypePtrInput {
	return (*timeframeTypePtr)(&v)
}

func (*timeframeTypePtr) ElementType() reflect.Type {
	return timeframeTypePtrType
}

func (in *timeframeTypePtr) ToTimeframeTypePtrOutput() TimeframeTypePtrOutput {
	return pulumi.ToOutput(in).(TimeframeTypePtrOutput)
}

func (in *timeframeTypePtr) ToTimeframeTypePtrOutputWithContext(ctx context.Context) TimeframeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TimeframeTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExportTypeOutput{})
	pulumi.RegisterOutputType(ExportTypePtrOutput{})
	pulumi.RegisterOutputType(FormatTypeOutput{})
	pulumi.RegisterOutputType(FormatTypePtrOutput{})
	pulumi.RegisterOutputType(GranularityTypeOutput{})
	pulumi.RegisterOutputType(GranularityTypePtrOutput{})
	pulumi.RegisterOutputType(RecurrenceTypeOutput{})
	pulumi.RegisterOutputType(RecurrenceTypePtrOutput{})
	pulumi.RegisterOutputType(StatusTypeOutput{})
	pulumi.RegisterOutputType(StatusTypePtrOutput{})
	pulumi.RegisterOutputType(TimeframeTypeOutput{})
	pulumi.RegisterOutputType(TimeframeTypePtrOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Protocol string

const (
	ProtocolTCP = Protocol("TCP")
	ProtocolUDP = Protocol("UDP")
	ProtocolAll = Protocol("*")
)

func (Protocol) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (e Protocol) ToProtocolOutput() ProtocolOutput {
	return pulumi.ToOutput(e).(ProtocolOutput)
}

func (e Protocol) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtocolOutput)
}

func (e Protocol) ToProtocolPtrOutput() ProtocolPtrOutput {
	return e.ToProtocolPtrOutputWithContext(context.Background())
}

func (e Protocol) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return Protocol(e).ToProtocolOutputWithContext(ctx).ToProtocolPtrOutputWithContext(ctx)
}

func (e Protocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Protocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtocolOutput struct{ *pulumi.OutputState }

func (ProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (o ProtocolOutput) ToProtocolOutput() ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o.ToProtocolPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Protocol) *Protocol {
		return &v
	}).(ProtocolPtrOutput)
}

func (o ProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProtocolPtrOutput) ElementType() reflect.Type {
	return protocolPtrType
}

func (o ProtocolPtrOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Protocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o ProtocolPtrOutput) Elem() ProtocolOutput {
	return o.ApplyT(func(v *Protocol) Protocol {
		var ret Protocol
		if v != nil {
			ret = *v
		}
		return ret
	}).(ProtocolOutput)
}

// ProtocolInput is an input type that accepts ProtocolArgs and ProtocolOutput values.
// You can construct a concrete instance of `ProtocolInput` via:
//
//          ProtocolArgs{...}
type ProtocolInput interface {
	pulumi.Input

	ToProtocolOutput() ProtocolOutput
	ToProtocolOutputWithContext(context.Context) ProtocolOutput
}

var protocolPtrType = reflect.TypeOf((**Protocol)(nil)).Elem()

type ProtocolPtrInput interface {
	pulumi.Input

	ToProtocolPtrOutput() ProtocolPtrOutput
	ToProtocolPtrOutputWithContext(context.Context) ProtocolPtrOutput
}

type protocolPtr string

func ProtocolPtr(v string) ProtocolPtrInput {
	return (*protocolPtr)(&v)
}

func (*protocolPtr) ElementType() reflect.Type {
	return protocolPtrType
}

func (in *protocolPtr) ToProtocolPtrOutput() ProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProtocolPtrOutput)
}

func (in *protocolPtr) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtocolPtrOutput)
}

// The status of the port
type Status string

const (
	StatusRevoked   = Status("Revoked")
	StatusInitiated = Status("Initiated")
)

func (Status) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (e Status) ToStatusOutput() StatusOutput {
	return pulumi.ToOutput(e).(StatusOutput)
}

func (e Status) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusOutput)
}

func (e Status) ToStatusPtrOutput() StatusPtrOutput {
	return e.ToStatusPtrOutputWithContext(context.Background())
}

func (e Status) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return Status(e).ToStatusOutputWithContext(ctx).ToStatusPtrOutputWithContext(ctx)
}

func (e Status) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Status) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusOutput struct{ *pulumi.OutputState }

func (StatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (o StatusOutput) ToStatusOutput() StatusOutput {
	return o
}

func (o StatusOutput) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return o
}

func (o StatusOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o.ToStatusPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Status) *Status {
		return &v
	}).(StatusPtrOutput)
}

func (o StatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusPtrOutput struct{ *pulumi.OutputState }

func (StatusPtrOutput) ElementType() reflect.Type {
	return statusPtrType
}

func (o StatusPtrOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Status) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o StatusPtrOutput) Elem() StatusOutput {
	return o.ApplyT(func(v *Status) Status {
		var ret Status
		if v != nil {
			ret = *v
		}
		return ret
	}).(StatusOutput)
}

// StatusInput is an input type that accepts StatusArgs and StatusOutput values.
// You can construct a concrete instance of `StatusInput` via:
//
//          StatusArgs{...}
type StatusInput interface {
	pulumi.Input

	ToStatusOutput() StatusOutput
	ToStatusOutputWithContext(context.Context) StatusOutput
}

var statusPtrType = reflect.TypeOf((**Status)(nil)).Elem()

type StatusPtrInput interface {
	pulumi.Input

	ToStatusPtrOutput() StatusPtrOutput
	ToStatusPtrOutputWithContext(context.Context) StatusPtrOutput
}

type statusPtr string

func StatusPtr(v string) StatusPtrInput {
	return (*statusPtr)(&v)
}

func (*statusPtr) ElementType() reflect.Type {
	return statusPtrType
}

func (in *statusPtr) ToStatusPtrOutput() StatusPtrOutput {
	return pulumi.ToOutput(in).(StatusPtrOutput)
}

func (in *statusPtr) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusPtrOutput)
}

// A description of why the `status` has its value
type StatusReason string

const (
	StatusReasonExpired               = StatusReason("Expired")
	StatusReasonUserRequested         = StatusReason("UserRequested")
	StatusReasonNewerRequestInitiated = StatusReason("NewerRequestInitiated")
)

func (StatusReason) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusReason)(nil)).Elem()
}

func (e StatusReason) ToStatusReasonOutput() StatusReasonOutput {
	return pulumi.ToOutput(e).(StatusReasonOutput)
}

func (e StatusReason) ToStatusReasonOutputWithContext(ctx context.Context) StatusReasonOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusReasonOutput)
}

func (e StatusReason) ToStatusReasonPtrOutput() StatusReasonPtrOutput {
	return e.ToStatusReasonPtrOutputWithContext(context.Background())
}

func (e StatusReason) ToStatusReasonPtrOutputWithContext(ctx context.Context) StatusReasonPtrOutput {
	return StatusReason(e).ToStatusReasonOutputWithContext(ctx).ToStatusReasonPtrOutputWithContext(ctx)
}

func (e StatusReason) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusReason) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusReason) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StatusReason) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusReasonOutput struct{ *pulumi.OutputState }

func (StatusReasonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusReason)(nil)).Elem()
}

func (o StatusReasonOutput) ToStatusReasonOutput() StatusReasonOutput {
	return o
}

func (o StatusReasonOutput) ToStatusReasonOutputWithContext(ctx context.Context) StatusReasonOutput {
	return o
}

func (o StatusReasonOutput) ToStatusReasonPtrOutput() StatusReasonPtrOutput {
	return o.ToStatusReasonPtrOutputWithContext(context.Background())
}

func (o StatusReasonOutput) ToStatusReasonPtrOutputWithContext(ctx context.Context) StatusReasonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StatusReason) *StatusReason {
		return &v
	}).(StatusReasonPtrOutput)
}

func (o StatusReasonOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusReasonOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusReason) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusReasonOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusReasonOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusReason) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusReasonPtrOutput struct{ *pulumi.OutputState }

func (StatusReasonPtrOutput) ElementType() reflect.Type {
	return statusReasonPtrType
}

func (o StatusReasonPtrOutput) ToStatusReasonPtrOutput() StatusReasonPtrOutput {
	return o
}

func (o StatusReasonPtrOutput) ToStatusReasonPtrOutputWithContext(ctx context.Context) StatusReasonPtrOutput {
	return o
}

func (o StatusReasonPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusReasonPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StatusReason) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o StatusReasonPtrOutput) Elem() StatusReasonOutput {
	return o.ApplyT(func(v *StatusReason) StatusReason {
		var ret StatusReason
		if v != nil {
			ret = *v
		}
		return ret
	}).(StatusReasonOutput)
}

// StatusReasonInput is an input type that accepts StatusReasonArgs and StatusReasonOutput values.
// You can construct a concrete instance of `StatusReasonInput` via:
//
//          StatusReasonArgs{...}
type StatusReasonInput interface {
	pulumi.Input

	ToStatusReasonOutput() StatusReasonOutput
	ToStatusReasonOutputWithContext(context.Context) StatusReasonOutput
}

var statusReasonPtrType = reflect.TypeOf((**StatusReason)(nil)).Elem()

type StatusReasonPtrInput interface {
	pulumi.Input

	ToStatusReasonPtrOutput() StatusReasonPtrOutput
	ToStatusReasonPtrOutputWithContext(context.Context) StatusReasonPtrOutput
}

type statusReasonPtr string

func StatusReasonPtr(v string) StatusReasonPtrInput {
	return (*statusReasonPtr)(&v)
}

func (*statusReasonPtr) ElementType() reflect.Type {
	return statusReasonPtrType
}

func (in *statusReasonPtr) ToStatusReasonPtrOutput() StatusReasonPtrOutput {
	return pulumi.ToOutput(in).(StatusReasonPtrOutput)
}

func (in *statusReasonPtr) ToStatusReasonPtrOutputWithContext(ctx context.Context) StatusReasonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusReasonPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtocolOutput{})
	pulumi.RegisterOutputType(ProtocolPtrOutput{})
	pulumi.RegisterOutputType(StatusOutput{})
	pulumi.RegisterOutputType(StatusPtrOutput{})
	pulumi.RegisterOutputType(StatusReasonOutput{})
	pulumi.RegisterOutputType(StatusReasonPtrOutput{})
}

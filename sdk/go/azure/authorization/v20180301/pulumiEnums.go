// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The policy definition mode. Possible values are NotSpecified, Indexed, and All.
type PolicyMode string

const (
	PolicyModeNotSpecified = PolicyMode("NotSpecified")
	PolicyModeIndexed      = PolicyMode("Indexed")
	PolicyModeAll          = PolicyMode("All")
)

func (PolicyMode) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyMode)(nil)).Elem()
}

func (e PolicyMode) ToPolicyModeOutput() PolicyModeOutput {
	return pulumi.ToOutput(e).(PolicyModeOutput)
}

func (e PolicyMode) ToPolicyModeOutputWithContext(ctx context.Context) PolicyModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyModeOutput)
}

func (e PolicyMode) ToPolicyModePtrOutput() PolicyModePtrOutput {
	return e.ToPolicyModePtrOutputWithContext(context.Background())
}

func (e PolicyMode) ToPolicyModePtrOutputWithContext(ctx context.Context) PolicyModePtrOutput {
	return PolicyMode(e).ToPolicyModeOutputWithContext(ctx).ToPolicyModePtrOutputWithContext(ctx)
}

func (e PolicyMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyModeOutput struct{ *pulumi.OutputState }

func (PolicyModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyMode)(nil)).Elem()
}

func (o PolicyModeOutput) ToPolicyModeOutput() PolicyModeOutput {
	return o
}

func (o PolicyModeOutput) ToPolicyModeOutputWithContext(ctx context.Context) PolicyModeOutput {
	return o
}

func (o PolicyModeOutput) ToPolicyModePtrOutput() PolicyModePtrOutput {
	return o.ToPolicyModePtrOutputWithContext(context.Background())
}

func (o PolicyModeOutput) ToPolicyModePtrOutputWithContext(ctx context.Context) PolicyModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyMode) *PolicyMode {
		return &v
	}).(PolicyModePtrOutput)
}

func (o PolicyModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyModePtrOutput struct{ *pulumi.OutputState }

func (PolicyModePtrOutput) ElementType() reflect.Type {
	return policyModePtrType
}

func (o PolicyModePtrOutput) ToPolicyModePtrOutput() PolicyModePtrOutput {
	return o
}

func (o PolicyModePtrOutput) ToPolicyModePtrOutputWithContext(ctx context.Context) PolicyModePtrOutput {
	return o
}

func (o PolicyModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o PolicyModePtrOutput) Elem() PolicyModeOutput {
	return o.ApplyT(func(v *PolicyMode) PolicyMode {
		var ret PolicyMode
		if v != nil {
			ret = *v
		}
		return ret
	}).(PolicyModeOutput)
}

// PolicyModeInput is an input type that accepts PolicyModeArgs and PolicyModeOutput values.
// You can construct a concrete instance of `PolicyModeInput` via:
//
//          PolicyModeArgs{...}
type PolicyModeInput interface {
	pulumi.Input

	ToPolicyModeOutput() PolicyModeOutput
	ToPolicyModeOutputWithContext(context.Context) PolicyModeOutput
}

var policyModePtrType = reflect.TypeOf((**PolicyMode)(nil)).Elem()

type PolicyModePtrInput interface {
	pulumi.Input

	ToPolicyModePtrOutput() PolicyModePtrOutput
	ToPolicyModePtrOutputWithContext(context.Context) PolicyModePtrOutput
}

type policyModePtr string

func PolicyModePtr(v string) PolicyModePtrInput {
	return (*policyModePtr)(&v)
}

func (*policyModePtr) ElementType() reflect.Type {
	return policyModePtrType
}

func (in *policyModePtr) ToPolicyModePtrOutput() PolicyModePtrOutput {
	return pulumi.ToOutput(in).(PolicyModePtrOutput)
}

func (in *policyModePtr) ToPolicyModePtrOutputWithContext(ctx context.Context) PolicyModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyModePtrOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
type PolicyType string

const (
	PolicyTypeNotSpecified = PolicyType("NotSpecified")
	PolicyTypeBuiltIn      = PolicyType("BuiltIn")
	PolicyTypeCustom       = PolicyType("Custom")
)

func (PolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (e PolicyType) ToPolicyTypeOutput() PolicyTypeOutput {
	return pulumi.ToOutput(e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return e.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (e PolicyType) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return PolicyType(e).ToPolicyTypeOutputWithContext(ctx).ToPolicyTypePtrOutputWithContext(ctx)
}

func (e PolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyTypeOutput struct{ *pulumi.OutputState }

func (PolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (o PolicyTypeOutput) ToPolicyTypeOutput() PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyType) *PolicyType {
		return &v
	}).(PolicyTypePtrOutput)
}

func (o PolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyTypePtrOutput struct{ *pulumi.OutputState }

func (PolicyTypePtrOutput) ElementType() reflect.Type {
	return policyTypePtrType
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o PolicyTypePtrOutput) Elem() PolicyTypeOutput {
	return o.ApplyT(func(v *PolicyType) PolicyType {
		var ret PolicyType
		if v != nil {
			ret = *v
		}
		return ret
	}).(PolicyTypeOutput)
}

// PolicyTypeInput is an input type that accepts PolicyTypeArgs and PolicyTypeOutput values.
// You can construct a concrete instance of `PolicyTypeInput` via:
//
//          PolicyTypeArgs{...}
type PolicyTypeInput interface {
	pulumi.Input

	ToPolicyTypeOutput() PolicyTypeOutput
	ToPolicyTypeOutputWithContext(context.Context) PolicyTypeOutput
}

var policyTypePtrType = reflect.TypeOf((**PolicyType)(nil)).Elem()

type PolicyTypePtrInput interface {
	pulumi.Input

	ToPolicyTypePtrOutput() PolicyTypePtrOutput
	ToPolicyTypePtrOutputWithContext(context.Context) PolicyTypePtrOutput
}

type policyTypePtr string

func PolicyTypePtr(v string) PolicyTypePtrInput {
	return (*policyTypePtr)(&v)
}

func (*policyTypePtr) ElementType() reflect.Type {
	return policyTypePtrType
}

func (in *policyTypePtr) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return pulumi.ToOutput(in).(PolicyTypePtrOutput)
}

func (in *policyTypePtr) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyModeOutput{})
	pulumi.RegisterOutputType(PolicyModePtrOutput{})
	pulumi.RegisterOutputType(PolicyTypeOutput{})
	pulumi.RegisterOutputType(PolicyTypePtrOutput{})
}

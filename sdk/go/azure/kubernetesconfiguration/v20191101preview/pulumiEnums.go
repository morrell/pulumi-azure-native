// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Option to enable Helm Operator for this git configuration.
type EnableHelmOperator string

const (
	EnableHelmOperatorTrue  = EnableHelmOperator("true")
	EnableHelmOperatorFalse = EnableHelmOperator("false")
)

func (EnableHelmOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*EnableHelmOperator)(nil)).Elem()
}

func (e EnableHelmOperator) ToEnableHelmOperatorOutput() EnableHelmOperatorOutput {
	return pulumi.ToOutput(e).(EnableHelmOperatorOutput)
}

func (e EnableHelmOperator) ToEnableHelmOperatorOutputWithContext(ctx context.Context) EnableHelmOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnableHelmOperatorOutput)
}

func (e EnableHelmOperator) ToEnableHelmOperatorPtrOutput() EnableHelmOperatorPtrOutput {
	return e.ToEnableHelmOperatorPtrOutputWithContext(context.Background())
}

func (e EnableHelmOperator) ToEnableHelmOperatorPtrOutputWithContext(ctx context.Context) EnableHelmOperatorPtrOutput {
	return EnableHelmOperator(e).ToEnableHelmOperatorOutputWithContext(ctx).ToEnableHelmOperatorPtrOutputWithContext(ctx)
}

func (e EnableHelmOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnableHelmOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnableHelmOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnableHelmOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnableHelmOperatorOutput struct{ *pulumi.OutputState }

func (EnableHelmOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnableHelmOperator)(nil)).Elem()
}

func (o EnableHelmOperatorOutput) ToEnableHelmOperatorOutput() EnableHelmOperatorOutput {
	return o
}

func (o EnableHelmOperatorOutput) ToEnableHelmOperatorOutputWithContext(ctx context.Context) EnableHelmOperatorOutput {
	return o
}

func (o EnableHelmOperatorOutput) ToEnableHelmOperatorPtrOutput() EnableHelmOperatorPtrOutput {
	return o.ToEnableHelmOperatorPtrOutputWithContext(context.Background())
}

func (o EnableHelmOperatorOutput) ToEnableHelmOperatorPtrOutputWithContext(ctx context.Context) EnableHelmOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnableHelmOperator) *EnableHelmOperator {
		return &v
	}).(EnableHelmOperatorPtrOutput)
}

func (o EnableHelmOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnableHelmOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnableHelmOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnableHelmOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnableHelmOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnableHelmOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnableHelmOperatorPtrOutput struct{ *pulumi.OutputState }

func (EnableHelmOperatorPtrOutput) ElementType() reflect.Type {
	return enableHelmOperatorPtrType
}

func (o EnableHelmOperatorPtrOutput) ToEnableHelmOperatorPtrOutput() EnableHelmOperatorPtrOutput {
	return o
}

func (o EnableHelmOperatorPtrOutput) ToEnableHelmOperatorPtrOutputWithContext(ctx context.Context) EnableHelmOperatorPtrOutput {
	return o
}

func (o EnableHelmOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnableHelmOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnableHelmOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o EnableHelmOperatorPtrOutput) Elem() EnableHelmOperatorOutput {
	return o.ApplyT(func(v *EnableHelmOperator) EnableHelmOperator {
		var ret EnableHelmOperator
		if v != nil {
			ret = *v
		}
		return ret
	}).(EnableHelmOperatorOutput)
}

// EnableHelmOperatorInput is an input type that accepts EnableHelmOperatorArgs and EnableHelmOperatorOutput values.
// You can construct a concrete instance of `EnableHelmOperatorInput` via:
//
//          EnableHelmOperatorArgs{...}
type EnableHelmOperatorInput interface {
	pulumi.Input

	ToEnableHelmOperatorOutput() EnableHelmOperatorOutput
	ToEnableHelmOperatorOutputWithContext(context.Context) EnableHelmOperatorOutput
}

var enableHelmOperatorPtrType = reflect.TypeOf((**EnableHelmOperator)(nil)).Elem()

type EnableHelmOperatorPtrInput interface {
	pulumi.Input

	ToEnableHelmOperatorPtrOutput() EnableHelmOperatorPtrOutput
	ToEnableHelmOperatorPtrOutputWithContext(context.Context) EnableHelmOperatorPtrOutput
}

type enableHelmOperatorPtr string

func EnableHelmOperatorPtr(v string) EnableHelmOperatorPtrInput {
	return (*enableHelmOperatorPtr)(&v)
}

func (*enableHelmOperatorPtr) ElementType() reflect.Type {
	return enableHelmOperatorPtrType
}

func (in *enableHelmOperatorPtr) ToEnableHelmOperatorPtrOutput() EnableHelmOperatorPtrOutput {
	return pulumi.ToOutput(in).(EnableHelmOperatorPtrOutput)
}

func (in *enableHelmOperatorPtr) ToEnableHelmOperatorPtrOutputWithContext(ctx context.Context) EnableHelmOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnableHelmOperatorPtrOutput)
}

// Scope at which the operator will be installed.
type OperatorScope string

const (
	OperatorScopeCluster   = OperatorScope("cluster")
	OperatorScopeNamespace = OperatorScope("namespace")
)

func (OperatorScope) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorScope)(nil)).Elem()
}

func (e OperatorScope) ToOperatorScopeOutput() OperatorScopeOutput {
	return pulumi.ToOutput(e).(OperatorScopeOutput)
}

func (e OperatorScope) ToOperatorScopeOutputWithContext(ctx context.Context) OperatorScopeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorScopeOutput)
}

func (e OperatorScope) ToOperatorScopePtrOutput() OperatorScopePtrOutput {
	return e.ToOperatorScopePtrOutputWithContext(context.Background())
}

func (e OperatorScope) ToOperatorScopePtrOutputWithContext(ctx context.Context) OperatorScopePtrOutput {
	return OperatorScope(e).ToOperatorScopeOutputWithContext(ctx).ToOperatorScopePtrOutputWithContext(ctx)
}

func (e OperatorScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatorScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorScopeOutput struct{ *pulumi.OutputState }

func (OperatorScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorScope)(nil)).Elem()
}

func (o OperatorScopeOutput) ToOperatorScopeOutput() OperatorScopeOutput {
	return o
}

func (o OperatorScopeOutput) ToOperatorScopeOutputWithContext(ctx context.Context) OperatorScopeOutput {
	return o
}

func (o OperatorScopeOutput) ToOperatorScopePtrOutput() OperatorScopePtrOutput {
	return o.ToOperatorScopePtrOutputWithContext(context.Background())
}

func (o OperatorScopeOutput) ToOperatorScopePtrOutputWithContext(ctx context.Context) OperatorScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatorScope) *OperatorScope {
		return &v
	}).(OperatorScopePtrOutput)
}

func (o OperatorScopeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorScopeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorScope) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorScopeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorScopeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorScope) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorScopePtrOutput struct{ *pulumi.OutputState }

func (OperatorScopePtrOutput) ElementType() reflect.Type {
	return operatorScopePtrType
}

func (o OperatorScopePtrOutput) ToOperatorScopePtrOutput() OperatorScopePtrOutput {
	return o
}

func (o OperatorScopePtrOutput) ToOperatorScopePtrOutputWithContext(ctx context.Context) OperatorScopePtrOutput {
	return o
}

func (o OperatorScopePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorScopePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatorScope) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o OperatorScopePtrOutput) Elem() OperatorScopeOutput {
	return o.ApplyT(func(v *OperatorScope) OperatorScope {
		var ret OperatorScope
		if v != nil {
			ret = *v
		}
		return ret
	}).(OperatorScopeOutput)
}

// OperatorScopeInput is an input type that accepts OperatorScopeArgs and OperatorScopeOutput values.
// You can construct a concrete instance of `OperatorScopeInput` via:
//
//          OperatorScopeArgs{...}
type OperatorScopeInput interface {
	pulumi.Input

	ToOperatorScopeOutput() OperatorScopeOutput
	ToOperatorScopeOutputWithContext(context.Context) OperatorScopeOutput
}

var operatorScopePtrType = reflect.TypeOf((**OperatorScope)(nil)).Elem()

type OperatorScopePtrInput interface {
	pulumi.Input

	ToOperatorScopePtrOutput() OperatorScopePtrOutput
	ToOperatorScopePtrOutputWithContext(context.Context) OperatorScopePtrOutput
}

type operatorScopePtr string

func OperatorScopePtr(v string) OperatorScopePtrInput {
	return (*operatorScopePtr)(&v)
}

func (*operatorScopePtr) ElementType() reflect.Type {
	return operatorScopePtrType
}

func (in *operatorScopePtr) ToOperatorScopePtrOutput() OperatorScopePtrOutput {
	return pulumi.ToOutput(in).(OperatorScopePtrOutput)
}

func (in *operatorScopePtr) ToOperatorScopePtrOutputWithContext(ctx context.Context) OperatorScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorScopePtrOutput)
}

// Type of the operator
type OperatorType string

const (
	OperatorTypeFlux = OperatorType("Flux")
)

func (OperatorType) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorType)(nil)).Elem()
}

func (e OperatorType) ToOperatorTypeOutput() OperatorTypeOutput {
	return pulumi.ToOutput(e).(OperatorTypeOutput)
}

func (e OperatorType) ToOperatorTypeOutputWithContext(ctx context.Context) OperatorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorTypeOutput)
}

func (e OperatorType) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return e.ToOperatorTypePtrOutputWithContext(context.Background())
}

func (e OperatorType) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return OperatorType(e).ToOperatorTypeOutputWithContext(ctx).ToOperatorTypePtrOutputWithContext(ctx)
}

func (e OperatorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorTypeOutput struct{ *pulumi.OutputState }

func (OperatorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorType)(nil)).Elem()
}

func (o OperatorTypeOutput) ToOperatorTypeOutput() OperatorTypeOutput {
	return o
}

func (o OperatorTypeOutput) ToOperatorTypeOutputWithContext(ctx context.Context) OperatorTypeOutput {
	return o
}

func (o OperatorTypeOutput) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return o.ToOperatorTypePtrOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatorType) *OperatorType {
		return &v
	}).(OperatorTypePtrOutput)
}

func (o OperatorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorTypePtrOutput struct{ *pulumi.OutputState }

func (OperatorTypePtrOutput) ElementType() reflect.Type {
	return operatorTypePtrType
}

func (o OperatorTypePtrOutput) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return o
}

func (o OperatorTypePtrOutput) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return o
}

func (o OperatorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o OperatorTypePtrOutput) Elem() OperatorTypeOutput {
	return o.ApplyT(func(v *OperatorType) OperatorType {
		var ret OperatorType
		if v != nil {
			ret = *v
		}
		return ret
	}).(OperatorTypeOutput)
}

// OperatorTypeInput is an input type that accepts OperatorTypeArgs and OperatorTypeOutput values.
// You can construct a concrete instance of `OperatorTypeInput` via:
//
//          OperatorTypeArgs{...}
type OperatorTypeInput interface {
	pulumi.Input

	ToOperatorTypeOutput() OperatorTypeOutput
	ToOperatorTypeOutputWithContext(context.Context) OperatorTypeOutput
}

var operatorTypePtrType = reflect.TypeOf((**OperatorType)(nil)).Elem()

type OperatorTypePtrInput interface {
	pulumi.Input

	ToOperatorTypePtrOutput() OperatorTypePtrOutput
	ToOperatorTypePtrOutputWithContext(context.Context) OperatorTypePtrOutput
}

type operatorTypePtr string

func OperatorTypePtr(v string) OperatorTypePtrInput {
	return (*operatorTypePtr)(&v)
}

func (*operatorTypePtr) ElementType() reflect.Type {
	return operatorTypePtrType
}

func (in *operatorTypePtr) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return pulumi.ToOutput(in).(OperatorTypePtrOutput)
}

func (in *operatorTypePtr) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnableHelmOperatorOutput{})
	pulumi.RegisterOutputType(EnableHelmOperatorPtrOutput{})
	pulumi.RegisterOutputType(OperatorScopeOutput{})
	pulumi.RegisterOutputType(OperatorScopePtrOutput{})
	pulumi.RegisterOutputType(OperatorTypeOutput{})
	pulumi.RegisterOutputType(OperatorTypePtrOutput{})
}

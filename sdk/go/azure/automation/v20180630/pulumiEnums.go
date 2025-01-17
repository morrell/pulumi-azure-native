// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180630

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets or sets the type of the runbook.
type RunbookTypeEnum string

const (
	RunbookTypeEnumScript                  = RunbookTypeEnum("Script")
	RunbookTypeEnumGraph                   = RunbookTypeEnum("Graph")
	RunbookTypeEnumPowerShellWorkflow      = RunbookTypeEnum("PowerShellWorkflow")
	RunbookTypeEnumPowerShell              = RunbookTypeEnum("PowerShell")
	RunbookTypeEnumGraphPowerShellWorkflow = RunbookTypeEnum("GraphPowerShellWorkflow")
	RunbookTypeEnumGraphPowerShell         = RunbookTypeEnum("GraphPowerShell")
)

func (RunbookTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookTypeEnum)(nil)).Elem()
}

func (e RunbookTypeEnum) ToRunbookTypeEnumOutput() RunbookTypeEnumOutput {
	return pulumi.ToOutput(e).(RunbookTypeEnumOutput)
}

func (e RunbookTypeEnum) ToRunbookTypeEnumOutputWithContext(ctx context.Context) RunbookTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RunbookTypeEnumOutput)
}

func (e RunbookTypeEnum) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return e.ToRunbookTypeEnumPtrOutputWithContext(context.Background())
}

func (e RunbookTypeEnum) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return RunbookTypeEnum(e).ToRunbookTypeEnumOutputWithContext(ctx).ToRunbookTypeEnumPtrOutputWithContext(ctx)
}

func (e RunbookTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RunbookTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RunbookTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RunbookTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RunbookTypeEnumOutput struct{ *pulumi.OutputState }

func (RunbookTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookTypeEnum)(nil)).Elem()
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumOutput() RunbookTypeEnumOutput {
	return o
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumOutputWithContext(ctx context.Context) RunbookTypeEnumOutput {
	return o
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return o.ToRunbookTypeEnumPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookTypeEnum) *RunbookTypeEnum {
		return &v
	}).(RunbookTypeEnumPtrOutput)
}

func (o RunbookTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RunbookTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RunbookTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RunbookTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RunbookTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (RunbookTypeEnumPtrOutput) ElementType() reflect.Type {
	return runbookTypeEnumPtrType
}

func (o RunbookTypeEnumPtrOutput) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return o
}

func (o RunbookTypeEnumPtrOutput) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return o
}

func (o RunbookTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RunbookTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o RunbookTypeEnumPtrOutput) Elem() RunbookTypeEnumOutput {
	return o.ApplyT(func(v *RunbookTypeEnum) RunbookTypeEnum {
		var ret RunbookTypeEnum
		if v != nil {
			ret = *v
		}
		return ret
	}).(RunbookTypeEnumOutput)
}

// RunbookTypeEnumInput is an input type that accepts RunbookTypeEnumArgs and RunbookTypeEnumOutput values.
// You can construct a concrete instance of `RunbookTypeEnumInput` via:
//
//          RunbookTypeEnumArgs{...}
type RunbookTypeEnumInput interface {
	pulumi.Input

	ToRunbookTypeEnumOutput() RunbookTypeEnumOutput
	ToRunbookTypeEnumOutputWithContext(context.Context) RunbookTypeEnumOutput
}

var runbookTypeEnumPtrType = reflect.TypeOf((**RunbookTypeEnum)(nil)).Elem()

type RunbookTypeEnumPtrInput interface {
	pulumi.Input

	ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput
	ToRunbookTypeEnumPtrOutputWithContext(context.Context) RunbookTypeEnumPtrOutput
}

type runbookTypeEnumPtr string

func RunbookTypeEnumPtr(v string) RunbookTypeEnumPtrInput {
	return (*runbookTypeEnumPtr)(&v)
}

func (*runbookTypeEnumPtr) ElementType() reflect.Type {
	return runbookTypeEnumPtrType
}

func (in *runbookTypeEnumPtr) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(RunbookTypeEnumPtrOutput)
}

func (in *runbookTypeEnumPtr) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RunbookTypeEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RunbookTypeEnumOutput{})
	pulumi.RegisterOutputType(RunbookTypeEnumPtrOutput{})
}

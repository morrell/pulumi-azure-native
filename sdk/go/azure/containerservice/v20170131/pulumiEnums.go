// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170131

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The orchestrator to use to manage container service cluster resources. Valid values are Swarm, DCOS, and Custom.
type ContainerServiceOrchestratorTypes string

const (
	ContainerServiceOrchestratorTypesSwarm      = ContainerServiceOrchestratorTypes("Swarm")
	ContainerServiceOrchestratorTypesDCOS       = ContainerServiceOrchestratorTypes("DCOS")
	ContainerServiceOrchestratorTypesCustom     = ContainerServiceOrchestratorTypes("Custom")
	ContainerServiceOrchestratorTypesKubernetes = ContainerServiceOrchestratorTypes("Kubernetes")
)

func (ContainerServiceOrchestratorTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceOrchestratorTypes)(nil)).Elem()
}

func (e ContainerServiceOrchestratorTypes) ToContainerServiceOrchestratorTypesOutput() ContainerServiceOrchestratorTypesOutput {
	return pulumi.ToOutput(e).(ContainerServiceOrchestratorTypesOutput)
}

func (e ContainerServiceOrchestratorTypes) ToContainerServiceOrchestratorTypesOutputWithContext(ctx context.Context) ContainerServiceOrchestratorTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerServiceOrchestratorTypesOutput)
}

func (e ContainerServiceOrchestratorTypes) ToContainerServiceOrchestratorTypesPtrOutput() ContainerServiceOrchestratorTypesPtrOutput {
	return e.ToContainerServiceOrchestratorTypesPtrOutputWithContext(context.Background())
}

func (e ContainerServiceOrchestratorTypes) ToContainerServiceOrchestratorTypesPtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorTypesPtrOutput {
	return ContainerServiceOrchestratorTypes(e).ToContainerServiceOrchestratorTypesOutputWithContext(ctx).ToContainerServiceOrchestratorTypesPtrOutputWithContext(ctx)
}

func (e ContainerServiceOrchestratorTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceOrchestratorTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceOrchestratorTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerServiceOrchestratorTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerServiceOrchestratorTypesOutput struct{ *pulumi.OutputState }

func (ContainerServiceOrchestratorTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceOrchestratorTypes)(nil)).Elem()
}

func (o ContainerServiceOrchestratorTypesOutput) ToContainerServiceOrchestratorTypesOutput() ContainerServiceOrchestratorTypesOutput {
	return o
}

func (o ContainerServiceOrchestratorTypesOutput) ToContainerServiceOrchestratorTypesOutputWithContext(ctx context.Context) ContainerServiceOrchestratorTypesOutput {
	return o
}

func (o ContainerServiceOrchestratorTypesOutput) ToContainerServiceOrchestratorTypesPtrOutput() ContainerServiceOrchestratorTypesPtrOutput {
	return o.ToContainerServiceOrchestratorTypesPtrOutputWithContext(context.Background())
}

func (o ContainerServiceOrchestratorTypesOutput) ToContainerServiceOrchestratorTypesPtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceOrchestratorTypes) *ContainerServiceOrchestratorTypes {
		return &v
	}).(ContainerServiceOrchestratorTypesPtrOutput)
}

func (o ContainerServiceOrchestratorTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerServiceOrchestratorTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceOrchestratorTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerServiceOrchestratorTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceOrchestratorTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceOrchestratorTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceOrchestratorTypesPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceOrchestratorTypesPtrOutput) ElementType() reflect.Type {
	return containerServiceOrchestratorTypesPtrType
}

func (o ContainerServiceOrchestratorTypesPtrOutput) ToContainerServiceOrchestratorTypesPtrOutput() ContainerServiceOrchestratorTypesPtrOutput {
	return o
}

func (o ContainerServiceOrchestratorTypesPtrOutput) ToContainerServiceOrchestratorTypesPtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorTypesPtrOutput {
	return o
}

func (o ContainerServiceOrchestratorTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceOrchestratorTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerServiceOrchestratorTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceOrchestratorTypesPtrOutput) Elem() ContainerServiceOrchestratorTypesOutput {
	return o.ApplyT(func(v *ContainerServiceOrchestratorTypes) ContainerServiceOrchestratorTypes {
		var ret ContainerServiceOrchestratorTypes
		if v != nil {
			ret = *v
		}
		return ret
	}).(ContainerServiceOrchestratorTypesOutput)
}

// ContainerServiceOrchestratorTypesInput is an input type that accepts ContainerServiceOrchestratorTypesArgs and ContainerServiceOrchestratorTypesOutput values.
// You can construct a concrete instance of `ContainerServiceOrchestratorTypesInput` via:
//
//          ContainerServiceOrchestratorTypesArgs{...}
type ContainerServiceOrchestratorTypesInput interface {
	pulumi.Input

	ToContainerServiceOrchestratorTypesOutput() ContainerServiceOrchestratorTypesOutput
	ToContainerServiceOrchestratorTypesOutputWithContext(context.Context) ContainerServiceOrchestratorTypesOutput
}

var containerServiceOrchestratorTypesPtrType = reflect.TypeOf((**ContainerServiceOrchestratorTypes)(nil)).Elem()

type ContainerServiceOrchestratorTypesPtrInput interface {
	pulumi.Input

	ToContainerServiceOrchestratorTypesPtrOutput() ContainerServiceOrchestratorTypesPtrOutput
	ToContainerServiceOrchestratorTypesPtrOutputWithContext(context.Context) ContainerServiceOrchestratorTypesPtrOutput
}

type containerServiceOrchestratorTypesPtr string

func ContainerServiceOrchestratorTypesPtr(v string) ContainerServiceOrchestratorTypesPtrInput {
	return (*containerServiceOrchestratorTypesPtr)(&v)
}

func (*containerServiceOrchestratorTypesPtr) ElementType() reflect.Type {
	return containerServiceOrchestratorTypesPtrType
}

func (in *containerServiceOrchestratorTypesPtr) ToContainerServiceOrchestratorTypesPtrOutput() ContainerServiceOrchestratorTypesPtrOutput {
	return pulumi.ToOutput(in).(ContainerServiceOrchestratorTypesPtrOutput)
}

func (in *containerServiceOrchestratorTypesPtr) ToContainerServiceOrchestratorTypesPtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerServiceOrchestratorTypesPtrOutput)
}

// Size of agent VMs.
type ContainerServiceVMSizeTypes string

const (
	ContainerServiceVMSizeTypes_Standard_A0     = ContainerServiceVMSizeTypes("Standard_A0")
	ContainerServiceVMSizeTypes_Standard_A1     = ContainerServiceVMSizeTypes("Standard_A1")
	ContainerServiceVMSizeTypes_Standard_A2     = ContainerServiceVMSizeTypes("Standard_A2")
	ContainerServiceVMSizeTypes_Standard_A3     = ContainerServiceVMSizeTypes("Standard_A3")
	ContainerServiceVMSizeTypes_Standard_A4     = ContainerServiceVMSizeTypes("Standard_A4")
	ContainerServiceVMSizeTypes_Standard_A5     = ContainerServiceVMSizeTypes("Standard_A5")
	ContainerServiceVMSizeTypes_Standard_A6     = ContainerServiceVMSizeTypes("Standard_A6")
	ContainerServiceVMSizeTypes_Standard_A7     = ContainerServiceVMSizeTypes("Standard_A7")
	ContainerServiceVMSizeTypes_Standard_A8     = ContainerServiceVMSizeTypes("Standard_A8")
	ContainerServiceVMSizeTypes_Standard_A9     = ContainerServiceVMSizeTypes("Standard_A9")
	ContainerServiceVMSizeTypes_Standard_A10    = ContainerServiceVMSizeTypes("Standard_A10")
	ContainerServiceVMSizeTypes_Standard_A11    = ContainerServiceVMSizeTypes("Standard_A11")
	ContainerServiceVMSizeTypes_Standard_D1     = ContainerServiceVMSizeTypes("Standard_D1")
	ContainerServiceVMSizeTypes_Standard_D2     = ContainerServiceVMSizeTypes("Standard_D2")
	ContainerServiceVMSizeTypes_Standard_D3     = ContainerServiceVMSizeTypes("Standard_D3")
	ContainerServiceVMSizeTypes_Standard_D4     = ContainerServiceVMSizeTypes("Standard_D4")
	ContainerServiceVMSizeTypes_Standard_D11    = ContainerServiceVMSizeTypes("Standard_D11")
	ContainerServiceVMSizeTypes_Standard_D12    = ContainerServiceVMSizeTypes("Standard_D12")
	ContainerServiceVMSizeTypes_Standard_D13    = ContainerServiceVMSizeTypes("Standard_D13")
	ContainerServiceVMSizeTypes_Standard_D14    = ContainerServiceVMSizeTypes("Standard_D14")
	ContainerServiceVMSizeTypes_Standard_D1_v2  = ContainerServiceVMSizeTypes("Standard_D1_v2")
	ContainerServiceVMSizeTypes_Standard_D2_v2  = ContainerServiceVMSizeTypes("Standard_D2_v2")
	ContainerServiceVMSizeTypes_Standard_D3_v2  = ContainerServiceVMSizeTypes("Standard_D3_v2")
	ContainerServiceVMSizeTypes_Standard_D4_v2  = ContainerServiceVMSizeTypes("Standard_D4_v2")
	ContainerServiceVMSizeTypes_Standard_D5_v2  = ContainerServiceVMSizeTypes("Standard_D5_v2")
	ContainerServiceVMSizeTypes_Standard_D11_v2 = ContainerServiceVMSizeTypes("Standard_D11_v2")
	ContainerServiceVMSizeTypes_Standard_D12_v2 = ContainerServiceVMSizeTypes("Standard_D12_v2")
	ContainerServiceVMSizeTypes_Standard_D13_v2 = ContainerServiceVMSizeTypes("Standard_D13_v2")
	ContainerServiceVMSizeTypes_Standard_D14_v2 = ContainerServiceVMSizeTypes("Standard_D14_v2")
	ContainerServiceVMSizeTypes_Standard_G1     = ContainerServiceVMSizeTypes("Standard_G1")
	ContainerServiceVMSizeTypes_Standard_G2     = ContainerServiceVMSizeTypes("Standard_G2")
	ContainerServiceVMSizeTypes_Standard_G3     = ContainerServiceVMSizeTypes("Standard_G3")
	ContainerServiceVMSizeTypes_Standard_G4     = ContainerServiceVMSizeTypes("Standard_G4")
	ContainerServiceVMSizeTypes_Standard_G5     = ContainerServiceVMSizeTypes("Standard_G5")
	ContainerServiceVMSizeTypes_Standard_DS1    = ContainerServiceVMSizeTypes("Standard_DS1")
	ContainerServiceVMSizeTypes_Standard_DS2    = ContainerServiceVMSizeTypes("Standard_DS2")
	ContainerServiceVMSizeTypes_Standard_DS3    = ContainerServiceVMSizeTypes("Standard_DS3")
	ContainerServiceVMSizeTypes_Standard_DS4    = ContainerServiceVMSizeTypes("Standard_DS4")
	ContainerServiceVMSizeTypes_Standard_DS11   = ContainerServiceVMSizeTypes("Standard_DS11")
	ContainerServiceVMSizeTypes_Standard_DS12   = ContainerServiceVMSizeTypes("Standard_DS12")
	ContainerServiceVMSizeTypes_Standard_DS13   = ContainerServiceVMSizeTypes("Standard_DS13")
	ContainerServiceVMSizeTypes_Standard_DS14   = ContainerServiceVMSizeTypes("Standard_DS14")
	ContainerServiceVMSizeTypes_Standard_GS1    = ContainerServiceVMSizeTypes("Standard_GS1")
	ContainerServiceVMSizeTypes_Standard_GS2    = ContainerServiceVMSizeTypes("Standard_GS2")
	ContainerServiceVMSizeTypes_Standard_GS3    = ContainerServiceVMSizeTypes("Standard_GS3")
	ContainerServiceVMSizeTypes_Standard_GS4    = ContainerServiceVMSizeTypes("Standard_GS4")
	ContainerServiceVMSizeTypes_Standard_GS5    = ContainerServiceVMSizeTypes("Standard_GS5")
)

func (ContainerServiceVMSizeTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceVMSizeTypes)(nil)).Elem()
}

func (e ContainerServiceVMSizeTypes) ToContainerServiceVMSizeTypesOutput() ContainerServiceVMSizeTypesOutput {
	return pulumi.ToOutput(e).(ContainerServiceVMSizeTypesOutput)
}

func (e ContainerServiceVMSizeTypes) ToContainerServiceVMSizeTypesOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerServiceVMSizeTypesOutput)
}

func (e ContainerServiceVMSizeTypes) ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput {
	return e.ToContainerServiceVMSizeTypesPtrOutputWithContext(context.Background())
}

func (e ContainerServiceVMSizeTypes) ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesPtrOutput {
	return ContainerServiceVMSizeTypes(e).ToContainerServiceVMSizeTypesOutputWithContext(ctx).ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx)
}

func (e ContainerServiceVMSizeTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceVMSizeTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceVMSizeTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerServiceVMSizeTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerServiceVMSizeTypesOutput struct{ *pulumi.OutputState }

func (ContainerServiceVMSizeTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceVMSizeTypes)(nil)).Elem()
}

func (o ContainerServiceVMSizeTypesOutput) ToContainerServiceVMSizeTypesOutput() ContainerServiceVMSizeTypesOutput {
	return o
}

func (o ContainerServiceVMSizeTypesOutput) ToContainerServiceVMSizeTypesOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesOutput {
	return o
}

func (o ContainerServiceVMSizeTypesOutput) ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput {
	return o.ToContainerServiceVMSizeTypesPtrOutputWithContext(context.Background())
}

func (o ContainerServiceVMSizeTypesOutput) ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceVMSizeTypes) *ContainerServiceVMSizeTypes {
		return &v
	}).(ContainerServiceVMSizeTypesPtrOutput)
}

func (o ContainerServiceVMSizeTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerServiceVMSizeTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceVMSizeTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerServiceVMSizeTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceVMSizeTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceVMSizeTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceVMSizeTypesPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceVMSizeTypesPtrOutput) ElementType() reflect.Type {
	return containerServiceVMSizeTypesPtrType
}

func (o ContainerServiceVMSizeTypesPtrOutput) ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput {
	return o
}

func (o ContainerServiceVMSizeTypesPtrOutput) ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesPtrOutput {
	return o
}

func (o ContainerServiceVMSizeTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceVMSizeTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerServiceVMSizeTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceVMSizeTypesPtrOutput) Elem() ContainerServiceVMSizeTypesOutput {
	return o.ApplyT(func(v *ContainerServiceVMSizeTypes) ContainerServiceVMSizeTypes {
		var ret ContainerServiceVMSizeTypes
		if v != nil {
			ret = *v
		}
		return ret
	}).(ContainerServiceVMSizeTypesOutput)
}

// ContainerServiceVMSizeTypesInput is an input type that accepts ContainerServiceVMSizeTypesArgs and ContainerServiceVMSizeTypesOutput values.
// You can construct a concrete instance of `ContainerServiceVMSizeTypesInput` via:
//
//          ContainerServiceVMSizeTypesArgs{...}
type ContainerServiceVMSizeTypesInput interface {
	pulumi.Input

	ToContainerServiceVMSizeTypesOutput() ContainerServiceVMSizeTypesOutput
	ToContainerServiceVMSizeTypesOutputWithContext(context.Context) ContainerServiceVMSizeTypesOutput
}

var containerServiceVMSizeTypesPtrType = reflect.TypeOf((**ContainerServiceVMSizeTypes)(nil)).Elem()

type ContainerServiceVMSizeTypesPtrInput interface {
	pulumi.Input

	ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput
	ToContainerServiceVMSizeTypesPtrOutputWithContext(context.Context) ContainerServiceVMSizeTypesPtrOutput
}

type containerServiceVMSizeTypesPtr string

func ContainerServiceVMSizeTypesPtr(v string) ContainerServiceVMSizeTypesPtrInput {
	return (*containerServiceVMSizeTypesPtr)(&v)
}

func (*containerServiceVMSizeTypesPtr) ElementType() reflect.Type {
	return containerServiceVMSizeTypesPtrType
}

func (in *containerServiceVMSizeTypesPtr) ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput {
	return pulumi.ToOutput(in).(ContainerServiceVMSizeTypesPtrOutput)
}

func (in *containerServiceVMSizeTypesPtr) ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerServiceVMSizeTypesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerServiceOrchestratorTypesOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorTypesPtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMSizeTypesOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMSizeTypesPtrOutput{})
}

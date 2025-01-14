// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The multi-factor authorization provider to be used for just-in-time access requests.
type MultiFactorAuthProvider string

const (
	MultiFactorAuthProviderAzure = MultiFactorAuthProvider("Azure")
	MultiFactorAuthProviderNone  = MultiFactorAuthProvider("None")
)

func (MultiFactorAuthProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiFactorAuthProvider)(nil)).Elem()
}

func (e MultiFactorAuthProvider) ToMultiFactorAuthProviderOutput() MultiFactorAuthProviderOutput {
	return pulumi.ToOutput(e).(MultiFactorAuthProviderOutput)
}

func (e MultiFactorAuthProvider) ToMultiFactorAuthProviderOutputWithContext(ctx context.Context) MultiFactorAuthProviderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MultiFactorAuthProviderOutput)
}

func (e MultiFactorAuthProvider) ToMultiFactorAuthProviderPtrOutput() MultiFactorAuthProviderPtrOutput {
	return e.ToMultiFactorAuthProviderPtrOutputWithContext(context.Background())
}

func (e MultiFactorAuthProvider) ToMultiFactorAuthProviderPtrOutputWithContext(ctx context.Context) MultiFactorAuthProviderPtrOutput {
	return MultiFactorAuthProvider(e).ToMultiFactorAuthProviderOutputWithContext(ctx).ToMultiFactorAuthProviderPtrOutputWithContext(ctx)
}

func (e MultiFactorAuthProvider) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MultiFactorAuthProvider) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MultiFactorAuthProvider) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MultiFactorAuthProvider) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MultiFactorAuthProviderOutput struct{ *pulumi.OutputState }

func (MultiFactorAuthProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiFactorAuthProvider)(nil)).Elem()
}

func (o MultiFactorAuthProviderOutput) ToMultiFactorAuthProviderOutput() MultiFactorAuthProviderOutput {
	return o
}

func (o MultiFactorAuthProviderOutput) ToMultiFactorAuthProviderOutputWithContext(ctx context.Context) MultiFactorAuthProviderOutput {
	return o
}

func (o MultiFactorAuthProviderOutput) ToMultiFactorAuthProviderPtrOutput() MultiFactorAuthProviderPtrOutput {
	return o.ToMultiFactorAuthProviderPtrOutputWithContext(context.Background())
}

func (o MultiFactorAuthProviderOutput) ToMultiFactorAuthProviderPtrOutputWithContext(ctx context.Context) MultiFactorAuthProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MultiFactorAuthProvider) *MultiFactorAuthProvider {
		return &v
	}).(MultiFactorAuthProviderPtrOutput)
}

func (o MultiFactorAuthProviderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MultiFactorAuthProviderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MultiFactorAuthProvider) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MultiFactorAuthProviderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MultiFactorAuthProviderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MultiFactorAuthProvider) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MultiFactorAuthProviderPtrOutput struct{ *pulumi.OutputState }

func (MultiFactorAuthProviderPtrOutput) ElementType() reflect.Type {
	return multiFactorAuthProviderPtrType
}

func (o MultiFactorAuthProviderPtrOutput) ToMultiFactorAuthProviderPtrOutput() MultiFactorAuthProviderPtrOutput {
	return o
}

func (o MultiFactorAuthProviderPtrOutput) ToMultiFactorAuthProviderPtrOutputWithContext(ctx context.Context) MultiFactorAuthProviderPtrOutput {
	return o
}

func (o MultiFactorAuthProviderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MultiFactorAuthProviderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MultiFactorAuthProvider) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o MultiFactorAuthProviderPtrOutput) Elem() MultiFactorAuthProviderOutput {
	return o.ApplyT(func(v *MultiFactorAuthProvider) MultiFactorAuthProvider {
		var ret MultiFactorAuthProvider
		if v != nil {
			ret = *v
		}
		return ret
	}).(MultiFactorAuthProviderOutput)
}

// MultiFactorAuthProviderInput is an input type that accepts MultiFactorAuthProviderArgs and MultiFactorAuthProviderOutput values.
// You can construct a concrete instance of `MultiFactorAuthProviderInput` via:
//
//          MultiFactorAuthProviderArgs{...}
type MultiFactorAuthProviderInput interface {
	pulumi.Input

	ToMultiFactorAuthProviderOutput() MultiFactorAuthProviderOutput
	ToMultiFactorAuthProviderOutputWithContext(context.Context) MultiFactorAuthProviderOutput
}

var multiFactorAuthProviderPtrType = reflect.TypeOf((**MultiFactorAuthProvider)(nil)).Elem()

type MultiFactorAuthProviderPtrInput interface {
	pulumi.Input

	ToMultiFactorAuthProviderPtrOutput() MultiFactorAuthProviderPtrOutput
	ToMultiFactorAuthProviderPtrOutputWithContext(context.Context) MultiFactorAuthProviderPtrOutput
}

type multiFactorAuthProviderPtr string

func MultiFactorAuthProviderPtr(v string) MultiFactorAuthProviderPtrInput {
	return (*multiFactorAuthProviderPtr)(&v)
}

func (*multiFactorAuthProviderPtr) ElementType() reflect.Type {
	return multiFactorAuthProviderPtrType
}

func (in *multiFactorAuthProviderPtr) ToMultiFactorAuthProviderPtrOutput() MultiFactorAuthProviderPtrOutput {
	return pulumi.ToOutput(in).(MultiFactorAuthProviderPtrOutput)
}

func (in *multiFactorAuthProviderPtr) ToMultiFactorAuthProviderPtrOutputWithContext(ctx context.Context) MultiFactorAuthProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MultiFactorAuthProviderPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MultiFactorAuthProviderOutput{})
	pulumi.RegisterOutputType(MultiFactorAuthProviderPtrOutput{})
}

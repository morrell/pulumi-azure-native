// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// FeatureFlags is the supported features of Azure SignalR service.
// - ServiceMode: Flag for backend server for SignalR service. Values allowed: "Default": have your own backend server; "Serverless": your application doesn't have a backend server; "Classic": for backward compatibility. Support both Default and Serverless mode but not recommended; "PredefinedOnly": for future use.
// - EnableConnectivityLogs: "true"/"false", to enable/disable the connectivity log category respectively.
type FeatureFlags string

const (
	FeatureFlagsServiceMode            = FeatureFlags("ServiceMode")
	FeatureFlagsEnableConnectivityLogs = FeatureFlags("EnableConnectivityLogs")
)

func (FeatureFlags) ElementType() reflect.Type {
	return reflect.TypeOf((*FeatureFlags)(nil)).Elem()
}

func (e FeatureFlags) ToFeatureFlagsOutput() FeatureFlagsOutput {
	return pulumi.ToOutput(e).(FeatureFlagsOutput)
}

func (e FeatureFlags) ToFeatureFlagsOutputWithContext(ctx context.Context) FeatureFlagsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FeatureFlagsOutput)
}

func (e FeatureFlags) ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput {
	return e.ToFeatureFlagsPtrOutputWithContext(context.Background())
}

func (e FeatureFlags) ToFeatureFlagsPtrOutputWithContext(ctx context.Context) FeatureFlagsPtrOutput {
	return FeatureFlags(e).ToFeatureFlagsOutputWithContext(ctx).ToFeatureFlagsPtrOutputWithContext(ctx)
}

func (e FeatureFlags) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FeatureFlags) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FeatureFlags) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FeatureFlags) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FeatureFlagsOutput struct{ *pulumi.OutputState }

func (FeatureFlagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FeatureFlags)(nil)).Elem()
}

func (o FeatureFlagsOutput) ToFeatureFlagsOutput() FeatureFlagsOutput {
	return o
}

func (o FeatureFlagsOutput) ToFeatureFlagsOutputWithContext(ctx context.Context) FeatureFlagsOutput {
	return o
}

func (o FeatureFlagsOutput) ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput {
	return o.ToFeatureFlagsPtrOutputWithContext(context.Background())
}

func (o FeatureFlagsOutput) ToFeatureFlagsPtrOutputWithContext(ctx context.Context) FeatureFlagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FeatureFlags) *FeatureFlags {
		return &v
	}).(FeatureFlagsPtrOutput)
}

func (o FeatureFlagsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FeatureFlagsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FeatureFlags) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FeatureFlagsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FeatureFlagsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FeatureFlags) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FeatureFlagsPtrOutput struct{ *pulumi.OutputState }

func (FeatureFlagsPtrOutput) ElementType() reflect.Type {
	return featureFlagsPtrType
}

func (o FeatureFlagsPtrOutput) ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput {
	return o
}

func (o FeatureFlagsPtrOutput) ToFeatureFlagsPtrOutputWithContext(ctx context.Context) FeatureFlagsPtrOutput {
	return o
}

func (o FeatureFlagsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FeatureFlagsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FeatureFlags) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o FeatureFlagsPtrOutput) Elem() FeatureFlagsOutput {
	return o.ApplyT(func(v *FeatureFlags) FeatureFlags {
		var ret FeatureFlags
		if v != nil {
			ret = *v
		}
		return ret
	}).(FeatureFlagsOutput)
}

// FeatureFlagsInput is an input type that accepts FeatureFlagsArgs and FeatureFlagsOutput values.
// You can construct a concrete instance of `FeatureFlagsInput` via:
//
//          FeatureFlagsArgs{...}
type FeatureFlagsInput interface {
	pulumi.Input

	ToFeatureFlagsOutput() FeatureFlagsOutput
	ToFeatureFlagsOutputWithContext(context.Context) FeatureFlagsOutput
}

var featureFlagsPtrType = reflect.TypeOf((**FeatureFlags)(nil)).Elem()

type FeatureFlagsPtrInput interface {
	pulumi.Input

	ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput
	ToFeatureFlagsPtrOutputWithContext(context.Context) FeatureFlagsPtrOutput
}

type featureFlagsPtr string

func FeatureFlagsPtr(v string) FeatureFlagsPtrInput {
	return (*featureFlagsPtr)(&v)
}

func (*featureFlagsPtr) ElementType() reflect.Type {
	return featureFlagsPtrType
}

func (in *featureFlagsPtr) ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput {
	return pulumi.ToOutput(in).(FeatureFlagsPtrOutput)
}

func (in *featureFlagsPtr) ToFeatureFlagsPtrOutputWithContext(ctx context.Context) FeatureFlagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FeatureFlagsPtrOutput)
}

// Optional tier of this particular SKU. 'Standard' or 'Free'.
//
// `Basic` is deprecated, use `Standard` instead.
type SignalRSkuTier string

const (
	SignalRSkuTierFree     = SignalRSkuTier("Free")
	SignalRSkuTierBasic    = SignalRSkuTier("Basic")
	SignalRSkuTierStandard = SignalRSkuTier("Standard")
	SignalRSkuTierPremium  = SignalRSkuTier("Premium")
)

func (SignalRSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRSkuTier)(nil)).Elem()
}

func (e SignalRSkuTier) ToSignalRSkuTierOutput() SignalRSkuTierOutput {
	return pulumi.ToOutput(e).(SignalRSkuTierOutput)
}

func (e SignalRSkuTier) ToSignalRSkuTierOutputWithContext(ctx context.Context) SignalRSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SignalRSkuTierOutput)
}

func (e SignalRSkuTier) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return e.ToSignalRSkuTierPtrOutputWithContext(context.Background())
}

func (e SignalRSkuTier) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return SignalRSkuTier(e).ToSignalRSkuTierOutputWithContext(ctx).ToSignalRSkuTierPtrOutputWithContext(ctx)
}

func (e SignalRSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignalRSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignalRSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SignalRSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SignalRSkuTierOutput struct{ *pulumi.OutputState }

func (SignalRSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRSkuTier)(nil)).Elem()
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierOutput() SignalRSkuTierOutput {
	return o
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierOutputWithContext(ctx context.Context) SignalRSkuTierOutput {
	return o
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return o.ToSignalRSkuTierPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRSkuTier) *SignalRSkuTier {
		return &v
	}).(SignalRSkuTierPtrOutput)
}

func (o SignalRSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignalRSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SignalRSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignalRSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SignalRSkuTierPtrOutput struct{ *pulumi.OutputState }

func (SignalRSkuTierPtrOutput) ElementType() reflect.Type {
	return signalRSkuTierPtrType
}

func (o SignalRSkuTierPtrOutput) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return o
}

func (o SignalRSkuTierPtrOutput) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return o
}

func (o SignalRSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SignalRSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o SignalRSkuTierPtrOutput) Elem() SignalRSkuTierOutput {
	return o.ApplyT(func(v *SignalRSkuTier) SignalRSkuTier {
		var ret SignalRSkuTier
		if v != nil {
			ret = *v
		}
		return ret
	}).(SignalRSkuTierOutput)
}

// SignalRSkuTierInput is an input type that accepts SignalRSkuTierArgs and SignalRSkuTierOutput values.
// You can construct a concrete instance of `SignalRSkuTierInput` via:
//
//          SignalRSkuTierArgs{...}
type SignalRSkuTierInput interface {
	pulumi.Input

	ToSignalRSkuTierOutput() SignalRSkuTierOutput
	ToSignalRSkuTierOutputWithContext(context.Context) SignalRSkuTierOutput
}

var signalRSkuTierPtrType = reflect.TypeOf((**SignalRSkuTier)(nil)).Elem()

type SignalRSkuTierPtrInput interface {
	pulumi.Input

	ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput
	ToSignalRSkuTierPtrOutputWithContext(context.Context) SignalRSkuTierPtrOutput
}

type signalRSkuTierPtr string

func SignalRSkuTierPtr(v string) SignalRSkuTierPtrInput {
	return (*signalRSkuTierPtr)(&v)
}

func (*signalRSkuTierPtr) ElementType() reflect.Type {
	return signalRSkuTierPtrType
}

func (in *signalRSkuTierPtr) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return pulumi.ToOutput(in).(SignalRSkuTierPtrOutput)
}

func (in *signalRSkuTierPtr) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SignalRSkuTierPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FeatureFlagsOutput{})
	pulumi.RegisterOutputType(FeatureFlagsPtrOutput{})
	pulumi.RegisterOutputType(SignalRSkuTierOutput{})
	pulumi.RegisterOutputType(SignalRSkuTierPtrOutput{})
}

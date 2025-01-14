// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The mode of client authentication.
type AuthenticationMethod string

const (
	AuthenticationMethodToken = AuthenticationMethod("Token")
	AuthenticationMethodAAD   = AuthenticationMethod("AAD")
)

func (AuthenticationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationMethod)(nil)).Elem()
}

func (e AuthenticationMethod) ToAuthenticationMethodOutput() AuthenticationMethodOutput {
	return pulumi.ToOutput(e).(AuthenticationMethodOutput)
}

func (e AuthenticationMethod) ToAuthenticationMethodOutputWithContext(ctx context.Context) AuthenticationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthenticationMethodOutput)
}

func (e AuthenticationMethod) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return e.ToAuthenticationMethodPtrOutputWithContext(context.Background())
}

func (e AuthenticationMethod) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return AuthenticationMethod(e).ToAuthenticationMethodOutputWithContext(ctx).ToAuthenticationMethodPtrOutputWithContext(ctx)
}

func (e AuthenticationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthenticationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthenticationMethodOutput struct{ *pulumi.OutputState }

func (AuthenticationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationMethod)(nil)).Elem()
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodOutput() AuthenticationMethodOutput {
	return o
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodOutputWithContext(ctx context.Context) AuthenticationMethodOutput {
	return o
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return o.ToAuthenticationMethodPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationMethod) *AuthenticationMethod {
		return &v
	}).(AuthenticationMethodPtrOutput)
}

func (o AuthenticationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthenticationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthenticationMethodPtrOutput struct{ *pulumi.OutputState }

func (AuthenticationMethodPtrOutput) ElementType() reflect.Type {
	return authenticationMethodPtrType
}

func (o AuthenticationMethodPtrOutput) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return o
}

func (o AuthenticationMethodPtrOutput) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return o
}

func (o AuthenticationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthenticationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o AuthenticationMethodPtrOutput) Elem() AuthenticationMethodOutput {
	return o.ApplyT(func(v *AuthenticationMethod) AuthenticationMethod {
		var ret AuthenticationMethod
		if v != nil {
			ret = *v
		}
		return ret
	}).(AuthenticationMethodOutput)
}

// AuthenticationMethodInput is an input type that accepts AuthenticationMethodArgs and AuthenticationMethodOutput values.
// You can construct a concrete instance of `AuthenticationMethodInput` via:
//
//          AuthenticationMethodArgs{...}
type AuthenticationMethodInput interface {
	pulumi.Input

	ToAuthenticationMethodOutput() AuthenticationMethodOutput
	ToAuthenticationMethodOutputWithContext(context.Context) AuthenticationMethodOutput
}

var authenticationMethodPtrType = reflect.TypeOf((**AuthenticationMethod)(nil)).Elem()

type AuthenticationMethodPtrInput interface {
	pulumi.Input

	ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput
	ToAuthenticationMethodPtrOutputWithContext(context.Context) AuthenticationMethodPtrOutput
}

type authenticationMethodPtr string

func AuthenticationMethodPtr(v string) AuthenticationMethodPtrInput {
	return (*authenticationMethodPtr)(&v)
}

func (*authenticationMethodPtr) ElementType() reflect.Type {
	return authenticationMethodPtrType
}

func (in *authenticationMethodPtr) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return pulumi.ToOutput(in).(AuthenticationMethodPtrOutput)
}

func (in *authenticationMethodPtr) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthenticationMethodPtrOutput)
}

// Property which describes the state of private link on a connected cluster resource.
type PrivateLinkState string

const (
	PrivateLinkStateEnabled  = PrivateLinkState("Enabled")
	PrivateLinkStateDisabled = PrivateLinkState("Disabled")
)

func (PrivateLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkState)(nil)).Elem()
}

func (e PrivateLinkState) ToPrivateLinkStateOutput() PrivateLinkStateOutput {
	return pulumi.ToOutput(e).(PrivateLinkStateOutput)
}

func (e PrivateLinkState) ToPrivateLinkStateOutputWithContext(ctx context.Context) PrivateLinkStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateLinkStateOutput)
}

func (e PrivateLinkState) ToPrivateLinkStatePtrOutput() PrivateLinkStatePtrOutput {
	return e.ToPrivateLinkStatePtrOutputWithContext(context.Background())
}

func (e PrivateLinkState) ToPrivateLinkStatePtrOutputWithContext(ctx context.Context) PrivateLinkStatePtrOutput {
	return PrivateLinkState(e).ToPrivateLinkStateOutputWithContext(ctx).ToPrivateLinkStatePtrOutputWithContext(ctx)
}

func (e PrivateLinkState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateLinkStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkState)(nil)).Elem()
}

func (o PrivateLinkStateOutput) ToPrivateLinkStateOutput() PrivateLinkStateOutput {
	return o
}

func (o PrivateLinkStateOutput) ToPrivateLinkStateOutputWithContext(ctx context.Context) PrivateLinkStateOutput {
	return o
}

func (o PrivateLinkStateOutput) ToPrivateLinkStatePtrOutput() PrivateLinkStatePtrOutput {
	return o.ToPrivateLinkStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkStateOutput) ToPrivateLinkStatePtrOutputWithContext(ctx context.Context) PrivateLinkStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkState) *PrivateLinkState {
		return &v
	}).(PrivateLinkStatePtrOutput)
}

func (o PrivateLinkStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateLinkStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateLinkStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkStatePtrOutput) ElementType() reflect.Type {
	return privateLinkStatePtrType
}

func (o PrivateLinkStatePtrOutput) ToPrivateLinkStatePtrOutput() PrivateLinkStatePtrOutput {
	return o
}

func (o PrivateLinkStatePtrOutput) ToPrivateLinkStatePtrOutputWithContext(ctx context.Context) PrivateLinkStatePtrOutput {
	return o
}

func (o PrivateLinkStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateLinkState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkStatePtrOutput) Elem() PrivateLinkStateOutput {
	return o.ApplyT(func(v *PrivateLinkState) PrivateLinkState {
		var ret PrivateLinkState
		if v != nil {
			ret = *v
		}
		return ret
	}).(PrivateLinkStateOutput)
}

// PrivateLinkStateInput is an input type that accepts PrivateLinkStateArgs and PrivateLinkStateOutput values.
// You can construct a concrete instance of `PrivateLinkStateInput` via:
//
//          PrivateLinkStateArgs{...}
type PrivateLinkStateInput interface {
	pulumi.Input

	ToPrivateLinkStateOutput() PrivateLinkStateOutput
	ToPrivateLinkStateOutputWithContext(context.Context) PrivateLinkStateOutput
}

var privateLinkStatePtrType = reflect.TypeOf((**PrivateLinkState)(nil)).Elem()

type PrivateLinkStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkStatePtrOutput() PrivateLinkStatePtrOutput
	ToPrivateLinkStatePtrOutputWithContext(context.Context) PrivateLinkStatePtrOutput
}

type privateLinkStatePtr string

func PrivateLinkStatePtr(v string) PrivateLinkStatePtrInput {
	return (*privateLinkStatePtr)(&v)
}

func (*privateLinkStatePtr) ElementType() reflect.Type {
	return privateLinkStatePtrType
}

func (in *privateLinkStatePtr) ToPrivateLinkStatePtrOutput() PrivateLinkStatePtrOutput {
	return pulumi.ToOutput(in).(PrivateLinkStatePtrOutput)
}

func (in *privateLinkStatePtr) ToPrivateLinkStatePtrOutputWithContext(ctx context.Context) PrivateLinkStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateLinkStatePtrOutput)
}

// Provisioning state of the connected cluster resource.
type ProvisioningState string

const (
	ProvisioningStateSucceeded    = ProvisioningState("Succeeded")
	ProvisioningStateFailed       = ProvisioningState("Failed")
	ProvisioningStateCanceled     = ProvisioningState("Canceled")
	ProvisioningStateProvisioning = ProvisioningState("Provisioning")
	ProvisioningStateUpdating     = ProvisioningState("Updating")
	ProvisioningStateDeleting     = ProvisioningState("Deleting")
	ProvisioningStateAccepted     = ProvisioningState("Accepted")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		var ret ProvisioningState
		if v != nil {
			ret = *v
		}
		return ret
	}).(ProvisioningStateOutput)
}

// ProvisioningStateInput is an input type that accepts ProvisioningStateArgs and ProvisioningStateOutput values.
// You can construct a concrete instance of `ProvisioningStateInput` via:
//
//          ProvisioningStateArgs{...}
type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		var ret ResourceIdentityType
		if v != nil {
			ret = *v
		}
		return ret
	}).(ResourceIdentityTypeOutput)
}

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//          ResourceIdentityTypeArgs{...}
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthenticationMethodOutput{})
	pulumi.RegisterOutputType(AuthenticationMethodPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkStatePtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}

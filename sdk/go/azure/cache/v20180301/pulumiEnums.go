// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Day of the week when a cache can be patched.
type DayOfWeek string

const (
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekEveryday  = DayOfWeek("Everyday")
	DayOfWeekWeekend   = DayOfWeek("Weekend")
)

func (DayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (e DayOfWeek) ToDayOfWeekOutput() DayOfWeekOutput {
	return pulumi.ToOutput(e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return e.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return DayOfWeek(e).ToDayOfWeekOutputWithContext(ctx).ToDayOfWeekPtrOutputWithContext(ctx)
}

func (e DayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DayOfWeekOutput struct{ *pulumi.OutputState }

func (DayOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekOutput) ToDayOfWeekOutput() DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DayOfWeek) *DayOfWeek {
		return &v
	}).(DayOfWeekPtrOutput)
}

func (o DayOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DayOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DayOfWeekPtrOutput struct{ *pulumi.OutputState }

func (DayOfWeekPtrOutput) ElementType() reflect.Type {
	return dayOfWeekPtrType
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DayOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o DayOfWeekPtrOutput) Elem() DayOfWeekOutput {
	return o.ApplyT(func(v *DayOfWeek) DayOfWeek {
		var ret DayOfWeek
		if v != nil {
			ret = *v
		}
		return ret
	}).(DayOfWeekOutput)
}

// DayOfWeekInput is an input type that accepts DayOfWeekArgs and DayOfWeekOutput values.
// You can construct a concrete instance of `DayOfWeekInput` via:
//
//          DayOfWeekArgs{...}
type DayOfWeekInput interface {
	pulumi.Input

	ToDayOfWeekOutput() DayOfWeekOutput
	ToDayOfWeekOutputWithContext(context.Context) DayOfWeekOutput
}

var dayOfWeekPtrType = reflect.TypeOf((**DayOfWeek)(nil)).Elem()

type DayOfWeekPtrInput interface {
	pulumi.Input

	ToDayOfWeekPtrOutput() DayOfWeekPtrOutput
	ToDayOfWeekPtrOutputWithContext(context.Context) DayOfWeekPtrOutput
}

type dayOfWeekPtr string

func DayOfWeekPtr(v string) DayOfWeekPtrInput {
	return (*dayOfWeekPtr)(&v)
}

func (*dayOfWeekPtr) ElementType() reflect.Type {
	return dayOfWeekPtrType
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return pulumi.ToOutput(in).(DayOfWeekPtrOutput)
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DayOfWeekPtrOutput)
}

// Role of the linked server.
type ReplicationRole string

const (
	ReplicationRolePrimary   = ReplicationRole("Primary")
	ReplicationRoleSecondary = ReplicationRole("Secondary")
)

func (ReplicationRole) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationRole)(nil)).Elem()
}

func (e ReplicationRole) ToReplicationRoleOutput() ReplicationRoleOutput {
	return pulumi.ToOutput(e).(ReplicationRoleOutput)
}

func (e ReplicationRole) ToReplicationRoleOutputWithContext(ctx context.Context) ReplicationRoleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReplicationRoleOutput)
}

func (e ReplicationRole) ToReplicationRolePtrOutput() ReplicationRolePtrOutput {
	return e.ToReplicationRolePtrOutputWithContext(context.Background())
}

func (e ReplicationRole) ToReplicationRolePtrOutputWithContext(ctx context.Context) ReplicationRolePtrOutput {
	return ReplicationRole(e).ToReplicationRoleOutputWithContext(ctx).ToReplicationRolePtrOutputWithContext(ctx)
}

func (e ReplicationRole) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationRole) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationRole) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReplicationRole) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReplicationRoleOutput struct{ *pulumi.OutputState }

func (ReplicationRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationRole)(nil)).Elem()
}

func (o ReplicationRoleOutput) ToReplicationRoleOutput() ReplicationRoleOutput {
	return o
}

func (o ReplicationRoleOutput) ToReplicationRoleOutputWithContext(ctx context.Context) ReplicationRoleOutput {
	return o
}

func (o ReplicationRoleOutput) ToReplicationRolePtrOutput() ReplicationRolePtrOutput {
	return o.ToReplicationRolePtrOutputWithContext(context.Background())
}

func (o ReplicationRoleOutput) ToReplicationRolePtrOutputWithContext(ctx context.Context) ReplicationRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationRole) *ReplicationRole {
		return &v
	}).(ReplicationRolePtrOutput)
}

func (o ReplicationRoleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReplicationRoleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationRole) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReplicationRoleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationRoleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationRole) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReplicationRolePtrOutput struct{ *pulumi.OutputState }

func (ReplicationRolePtrOutput) ElementType() reflect.Type {
	return replicationRolePtrType
}

func (o ReplicationRolePtrOutput) ToReplicationRolePtrOutput() ReplicationRolePtrOutput {
	return o
}

func (o ReplicationRolePtrOutput) ToReplicationRolePtrOutputWithContext(ctx context.Context) ReplicationRolePtrOutput {
	return o
}

func (o ReplicationRolePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationRolePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReplicationRole) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationRolePtrOutput) Elem() ReplicationRoleOutput {
	return o.ApplyT(func(v *ReplicationRole) ReplicationRole {
		var ret ReplicationRole
		if v != nil {
			ret = *v
		}
		return ret
	}).(ReplicationRoleOutput)
}

// ReplicationRoleInput is an input type that accepts ReplicationRoleArgs and ReplicationRoleOutput values.
// You can construct a concrete instance of `ReplicationRoleInput` via:
//
//          ReplicationRoleArgs{...}
type ReplicationRoleInput interface {
	pulumi.Input

	ToReplicationRoleOutput() ReplicationRoleOutput
	ToReplicationRoleOutputWithContext(context.Context) ReplicationRoleOutput
}

var replicationRolePtrType = reflect.TypeOf((**ReplicationRole)(nil)).Elem()

type ReplicationRolePtrInput interface {
	pulumi.Input

	ToReplicationRolePtrOutput() ReplicationRolePtrOutput
	ToReplicationRolePtrOutputWithContext(context.Context) ReplicationRolePtrOutput
}

type replicationRolePtr string

func ReplicationRolePtr(v string) ReplicationRolePtrInput {
	return (*replicationRolePtr)(&v)
}

func (*replicationRolePtr) ElementType() reflect.Type {
	return replicationRolePtrType
}

func (in *replicationRolePtr) ToReplicationRolePtrOutput() ReplicationRolePtrOutput {
	return pulumi.ToOutput(in).(ReplicationRolePtrOutput)
}

func (in *replicationRolePtr) ToReplicationRolePtrOutputWithContext(ctx context.Context) ReplicationRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReplicationRolePtrOutput)
}

// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
type SkuFamily string

const (
	SkuFamilyC = SkuFamily("C")
	SkuFamilyP = SkuFamily("P")
)

func (SkuFamily) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuFamily)(nil)).Elem()
}

func (e SkuFamily) ToSkuFamilyOutput() SkuFamilyOutput {
	return pulumi.ToOutput(e).(SkuFamilyOutput)
}

func (e SkuFamily) ToSkuFamilyOutputWithContext(ctx context.Context) SkuFamilyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuFamilyOutput)
}

func (e SkuFamily) ToSkuFamilyPtrOutput() SkuFamilyPtrOutput {
	return e.ToSkuFamilyPtrOutputWithContext(context.Background())
}

func (e SkuFamily) ToSkuFamilyPtrOutputWithContext(ctx context.Context) SkuFamilyPtrOutput {
	return SkuFamily(e).ToSkuFamilyOutputWithContext(ctx).ToSkuFamilyPtrOutputWithContext(ctx)
}

func (e SkuFamily) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuFamily) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuFamily) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuFamily) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuFamilyOutput struct{ *pulumi.OutputState }

func (SkuFamilyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuFamily)(nil)).Elem()
}

func (o SkuFamilyOutput) ToSkuFamilyOutput() SkuFamilyOutput {
	return o
}

func (o SkuFamilyOutput) ToSkuFamilyOutputWithContext(ctx context.Context) SkuFamilyOutput {
	return o
}

func (o SkuFamilyOutput) ToSkuFamilyPtrOutput() SkuFamilyPtrOutput {
	return o.ToSkuFamilyPtrOutputWithContext(context.Background())
}

func (o SkuFamilyOutput) ToSkuFamilyPtrOutputWithContext(ctx context.Context) SkuFamilyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuFamily) *SkuFamily {
		return &v
	}).(SkuFamilyPtrOutput)
}

func (o SkuFamilyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuFamilyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuFamily) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuFamilyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuFamilyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuFamily) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuFamilyPtrOutput struct{ *pulumi.OutputState }

func (SkuFamilyPtrOutput) ElementType() reflect.Type {
	return skuFamilyPtrType
}

func (o SkuFamilyPtrOutput) ToSkuFamilyPtrOutput() SkuFamilyPtrOutput {
	return o
}

func (o SkuFamilyPtrOutput) ToSkuFamilyPtrOutputWithContext(ctx context.Context) SkuFamilyPtrOutput {
	return o
}

func (o SkuFamilyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuFamilyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuFamily) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o SkuFamilyPtrOutput) Elem() SkuFamilyOutput {
	return o.ApplyT(func(v *SkuFamily) SkuFamily {
		var ret SkuFamily
		if v != nil {
			ret = *v
		}
		return ret
	}).(SkuFamilyOutput)
}

// SkuFamilyInput is an input type that accepts SkuFamilyArgs and SkuFamilyOutput values.
// You can construct a concrete instance of `SkuFamilyInput` via:
//
//          SkuFamilyArgs{...}
type SkuFamilyInput interface {
	pulumi.Input

	ToSkuFamilyOutput() SkuFamilyOutput
	ToSkuFamilyOutputWithContext(context.Context) SkuFamilyOutput
}

var skuFamilyPtrType = reflect.TypeOf((**SkuFamily)(nil)).Elem()

type SkuFamilyPtrInput interface {
	pulumi.Input

	ToSkuFamilyPtrOutput() SkuFamilyPtrOutput
	ToSkuFamilyPtrOutputWithContext(context.Context) SkuFamilyPtrOutput
}

type skuFamilyPtr string

func SkuFamilyPtr(v string) SkuFamilyPtrInput {
	return (*skuFamilyPtr)(&v)
}

func (*skuFamilyPtr) ElementType() reflect.Type {
	return skuFamilyPtrType
}

func (in *skuFamilyPtr) ToSkuFamilyPtrOutput() SkuFamilyPtrOutput {
	return pulumi.ToOutput(in).(SkuFamilyPtrOutput)
}

func (in *skuFamilyPtr) ToSkuFamilyPtrOutputWithContext(ctx context.Context) SkuFamilyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuFamilyPtrOutput)
}

// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return skuNamePtrType
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		var ret SkuName
		if v != nil {
			ret = *v
		}
		return ret
	}).(SkuNameOutput)
}

// SkuNameInput is an input type that accepts SkuNameArgs and SkuNameOutput values.
// You can construct a concrete instance of `SkuNameInput` via:
//
//          SkuNameArgs{...}
type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
type TlsVersion string

const (
	TlsVersion_1_0 = TlsVersion("1.0")
	TlsVersion_1_1 = TlsVersion("1.1")
	TlsVersion_1_2 = TlsVersion("1.2")
)

func (TlsVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsVersion)(nil)).Elem()
}

func (e TlsVersion) ToTlsVersionOutput() TlsVersionOutput {
	return pulumi.ToOutput(e).(TlsVersionOutput)
}

func (e TlsVersion) ToTlsVersionOutputWithContext(ctx context.Context) TlsVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TlsVersionOutput)
}

func (e TlsVersion) ToTlsVersionPtrOutput() TlsVersionPtrOutput {
	return e.ToTlsVersionPtrOutputWithContext(context.Background())
}

func (e TlsVersion) ToTlsVersionPtrOutputWithContext(ctx context.Context) TlsVersionPtrOutput {
	return TlsVersion(e).ToTlsVersionOutputWithContext(ctx).ToTlsVersionPtrOutputWithContext(ctx)
}

func (e TlsVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TlsVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TlsVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TlsVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TlsVersionOutput struct{ *pulumi.OutputState }

func (TlsVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsVersion)(nil)).Elem()
}

func (o TlsVersionOutput) ToTlsVersionOutput() TlsVersionOutput {
	return o
}

func (o TlsVersionOutput) ToTlsVersionOutputWithContext(ctx context.Context) TlsVersionOutput {
	return o
}

func (o TlsVersionOutput) ToTlsVersionPtrOutput() TlsVersionPtrOutput {
	return o.ToTlsVersionPtrOutputWithContext(context.Background())
}

func (o TlsVersionOutput) ToTlsVersionPtrOutputWithContext(ctx context.Context) TlsVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TlsVersion) *TlsVersion {
		return &v
	}).(TlsVersionPtrOutput)
}

func (o TlsVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TlsVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TlsVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TlsVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TlsVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TlsVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TlsVersionPtrOutput struct{ *pulumi.OutputState }

func (TlsVersionPtrOutput) ElementType() reflect.Type {
	return tlsVersionPtrType
}

func (o TlsVersionPtrOutput) ToTlsVersionPtrOutput() TlsVersionPtrOutput {
	return o
}

func (o TlsVersionPtrOutput) ToTlsVersionPtrOutputWithContext(ctx context.Context) TlsVersionPtrOutput {
	return o
}

func (o TlsVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TlsVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TlsVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o TlsVersionPtrOutput) Elem() TlsVersionOutput {
	return o.ApplyT(func(v *TlsVersion) TlsVersion {
		var ret TlsVersion
		if v != nil {
			ret = *v
		}
		return ret
	}).(TlsVersionOutput)
}

// TlsVersionInput is an input type that accepts TlsVersionArgs and TlsVersionOutput values.
// You can construct a concrete instance of `TlsVersionInput` via:
//
//          TlsVersionArgs{...}
type TlsVersionInput interface {
	pulumi.Input

	ToTlsVersionOutput() TlsVersionOutput
	ToTlsVersionOutputWithContext(context.Context) TlsVersionOutput
}

var tlsVersionPtrType = reflect.TypeOf((**TlsVersion)(nil)).Elem()

type TlsVersionPtrInput interface {
	pulumi.Input

	ToTlsVersionPtrOutput() TlsVersionPtrOutput
	ToTlsVersionPtrOutputWithContext(context.Context) TlsVersionPtrOutput
}

type tlsVersionPtr string

func TlsVersionPtr(v string) TlsVersionPtrInput {
	return (*tlsVersionPtr)(&v)
}

func (*tlsVersionPtr) ElementType() reflect.Type {
	return tlsVersionPtrType
}

func (in *tlsVersionPtr) ToTlsVersionPtrOutput() TlsVersionPtrOutput {
	return pulumi.ToOutput(in).(TlsVersionPtrOutput)
}

func (in *tlsVersionPtr) ToTlsVersionPtrOutputWithContext(ctx context.Context) TlsVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TlsVersionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DayOfWeekOutput{})
	pulumi.RegisterOutputType(DayOfWeekPtrOutput{})
	pulumi.RegisterOutputType(ReplicationRoleOutput{})
	pulumi.RegisterOutputType(ReplicationRolePtrOutput{})
	pulumi.RegisterOutputType(SkuFamilyOutput{})
	pulumi.RegisterOutputType(SkuFamilyPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(TlsVersionOutput{})
	pulumi.RegisterOutputType(TlsVersionPtrOutput{})
}

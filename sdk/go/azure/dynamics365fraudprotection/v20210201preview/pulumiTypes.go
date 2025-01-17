// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An array of administrator user identities
type DFPInstanceAdministrators struct {
	// An array of administrator user identities.
	Members []string `pulumi:"members"`
}

// DFPInstanceAdministratorsInput is an input type that accepts DFPInstanceAdministratorsArgs and DFPInstanceAdministratorsOutput values.
// You can construct a concrete instance of `DFPInstanceAdministratorsInput` via:
//
//          DFPInstanceAdministratorsArgs{...}
type DFPInstanceAdministratorsInput interface {
	pulumi.Input

	ToDFPInstanceAdministratorsOutput() DFPInstanceAdministratorsOutput
	ToDFPInstanceAdministratorsOutputWithContext(context.Context) DFPInstanceAdministratorsOutput
}

// An array of administrator user identities
type DFPInstanceAdministratorsArgs struct {
	// An array of administrator user identities.
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (DFPInstanceAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DFPInstanceAdministrators)(nil)).Elem()
}

func (i DFPInstanceAdministratorsArgs) ToDFPInstanceAdministratorsOutput() DFPInstanceAdministratorsOutput {
	return i.ToDFPInstanceAdministratorsOutputWithContext(context.Background())
}

func (i DFPInstanceAdministratorsArgs) ToDFPInstanceAdministratorsOutputWithContext(ctx context.Context) DFPInstanceAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsOutput)
}

func (i DFPInstanceAdministratorsArgs) ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput {
	return i.ToDFPInstanceAdministratorsPtrOutputWithContext(context.Background())
}

func (i DFPInstanceAdministratorsArgs) ToDFPInstanceAdministratorsPtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsOutput).ToDFPInstanceAdministratorsPtrOutputWithContext(ctx)
}

// DFPInstanceAdministratorsPtrInput is an input type that accepts DFPInstanceAdministratorsArgs, DFPInstanceAdministratorsPtr and DFPInstanceAdministratorsPtrOutput values.
// You can construct a concrete instance of `DFPInstanceAdministratorsPtrInput` via:
//
//          DFPInstanceAdministratorsArgs{...}
//
//  or:
//
//          nil
type DFPInstanceAdministratorsPtrInput interface {
	pulumi.Input

	ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput
	ToDFPInstanceAdministratorsPtrOutputWithContext(context.Context) DFPInstanceAdministratorsPtrOutput
}

type dfpinstanceAdministratorsPtrType DFPInstanceAdministratorsArgs

func DFPInstanceAdministratorsPtr(v *DFPInstanceAdministratorsArgs) DFPInstanceAdministratorsPtrInput {
	return (*dfpinstanceAdministratorsPtrType)(v)
}

func (*dfpinstanceAdministratorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DFPInstanceAdministrators)(nil)).Elem()
}

func (i *dfpinstanceAdministratorsPtrType) ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput {
	return i.ToDFPInstanceAdministratorsPtrOutputWithContext(context.Background())
}

func (i *dfpinstanceAdministratorsPtrType) ToDFPInstanceAdministratorsPtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsPtrOutput)
}

// An array of administrator user identities
type DFPInstanceAdministratorsOutput struct{ *pulumi.OutputState }

func (DFPInstanceAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DFPInstanceAdministrators)(nil)).Elem()
}

func (o DFPInstanceAdministratorsOutput) ToDFPInstanceAdministratorsOutput() DFPInstanceAdministratorsOutput {
	return o
}

func (o DFPInstanceAdministratorsOutput) ToDFPInstanceAdministratorsOutputWithContext(ctx context.Context) DFPInstanceAdministratorsOutput {
	return o
}

func (o DFPInstanceAdministratorsOutput) ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput {
	return o.ToDFPInstanceAdministratorsPtrOutputWithContext(context.Background())
}

func (o DFPInstanceAdministratorsOutput) ToDFPInstanceAdministratorsPtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsPtrOutput {
	return o.ApplyT(func(v DFPInstanceAdministrators) *DFPInstanceAdministrators {
		return &v
	}).(DFPInstanceAdministratorsPtrOutput)
}

// An array of administrator user identities.
func (o DFPInstanceAdministratorsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DFPInstanceAdministrators) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DFPInstanceAdministratorsPtrOutput struct{ *pulumi.OutputState }

func (DFPInstanceAdministratorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DFPInstanceAdministrators)(nil)).Elem()
}

func (o DFPInstanceAdministratorsPtrOutput) ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput {
	return o
}

func (o DFPInstanceAdministratorsPtrOutput) ToDFPInstanceAdministratorsPtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsPtrOutput {
	return o
}

func (o DFPInstanceAdministratorsPtrOutput) Elem() DFPInstanceAdministratorsOutput {
	return o.ApplyT(func(v *DFPInstanceAdministrators) DFPInstanceAdministrators { return *v }).(DFPInstanceAdministratorsOutput)
}

// An array of administrator user identities.
func (o DFPInstanceAdministratorsPtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DFPInstanceAdministrators) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

// An array of administrator user identities
type DFPInstanceAdministratorsResponse struct {
	// An array of administrator user identities.
	Members []string `pulumi:"members"`
}

// DFPInstanceAdministratorsResponseInput is an input type that accepts DFPInstanceAdministratorsResponseArgs and DFPInstanceAdministratorsResponseOutput values.
// You can construct a concrete instance of `DFPInstanceAdministratorsResponseInput` via:
//
//          DFPInstanceAdministratorsResponseArgs{...}
type DFPInstanceAdministratorsResponseInput interface {
	pulumi.Input

	ToDFPInstanceAdministratorsResponseOutput() DFPInstanceAdministratorsResponseOutput
	ToDFPInstanceAdministratorsResponseOutputWithContext(context.Context) DFPInstanceAdministratorsResponseOutput
}

// An array of administrator user identities
type DFPInstanceAdministratorsResponseArgs struct {
	// An array of administrator user identities.
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (DFPInstanceAdministratorsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DFPInstanceAdministratorsResponse)(nil)).Elem()
}

func (i DFPInstanceAdministratorsResponseArgs) ToDFPInstanceAdministratorsResponseOutput() DFPInstanceAdministratorsResponseOutput {
	return i.ToDFPInstanceAdministratorsResponseOutputWithContext(context.Background())
}

func (i DFPInstanceAdministratorsResponseArgs) ToDFPInstanceAdministratorsResponseOutputWithContext(ctx context.Context) DFPInstanceAdministratorsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsResponseOutput)
}

func (i DFPInstanceAdministratorsResponseArgs) ToDFPInstanceAdministratorsResponsePtrOutput() DFPInstanceAdministratorsResponsePtrOutput {
	return i.ToDFPInstanceAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i DFPInstanceAdministratorsResponseArgs) ToDFPInstanceAdministratorsResponsePtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsResponseOutput).ToDFPInstanceAdministratorsResponsePtrOutputWithContext(ctx)
}

// DFPInstanceAdministratorsResponsePtrInput is an input type that accepts DFPInstanceAdministratorsResponseArgs, DFPInstanceAdministratorsResponsePtr and DFPInstanceAdministratorsResponsePtrOutput values.
// You can construct a concrete instance of `DFPInstanceAdministratorsResponsePtrInput` via:
//
//          DFPInstanceAdministratorsResponseArgs{...}
//
//  or:
//
//          nil
type DFPInstanceAdministratorsResponsePtrInput interface {
	pulumi.Input

	ToDFPInstanceAdministratorsResponsePtrOutput() DFPInstanceAdministratorsResponsePtrOutput
	ToDFPInstanceAdministratorsResponsePtrOutputWithContext(context.Context) DFPInstanceAdministratorsResponsePtrOutput
}

type dfpinstanceAdministratorsResponsePtrType DFPInstanceAdministratorsResponseArgs

func DFPInstanceAdministratorsResponsePtr(v *DFPInstanceAdministratorsResponseArgs) DFPInstanceAdministratorsResponsePtrInput {
	return (*dfpinstanceAdministratorsResponsePtrType)(v)
}

func (*dfpinstanceAdministratorsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DFPInstanceAdministratorsResponse)(nil)).Elem()
}

func (i *dfpinstanceAdministratorsResponsePtrType) ToDFPInstanceAdministratorsResponsePtrOutput() DFPInstanceAdministratorsResponsePtrOutput {
	return i.ToDFPInstanceAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i *dfpinstanceAdministratorsResponsePtrType) ToDFPInstanceAdministratorsResponsePtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsResponsePtrOutput)
}

// An array of administrator user identities
type DFPInstanceAdministratorsResponseOutput struct{ *pulumi.OutputState }

func (DFPInstanceAdministratorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DFPInstanceAdministratorsResponse)(nil)).Elem()
}

func (o DFPInstanceAdministratorsResponseOutput) ToDFPInstanceAdministratorsResponseOutput() DFPInstanceAdministratorsResponseOutput {
	return o
}

func (o DFPInstanceAdministratorsResponseOutput) ToDFPInstanceAdministratorsResponseOutputWithContext(ctx context.Context) DFPInstanceAdministratorsResponseOutput {
	return o
}

func (o DFPInstanceAdministratorsResponseOutput) ToDFPInstanceAdministratorsResponsePtrOutput() DFPInstanceAdministratorsResponsePtrOutput {
	return o.ToDFPInstanceAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (o DFPInstanceAdministratorsResponseOutput) ToDFPInstanceAdministratorsResponsePtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v DFPInstanceAdministratorsResponse) *DFPInstanceAdministratorsResponse {
		return &v
	}).(DFPInstanceAdministratorsResponsePtrOutput)
}

// An array of administrator user identities.
func (o DFPInstanceAdministratorsResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DFPInstanceAdministratorsResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DFPInstanceAdministratorsResponsePtrOutput struct{ *pulumi.OutputState }

func (DFPInstanceAdministratorsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DFPInstanceAdministratorsResponse)(nil)).Elem()
}

func (o DFPInstanceAdministratorsResponsePtrOutput) ToDFPInstanceAdministratorsResponsePtrOutput() DFPInstanceAdministratorsResponsePtrOutput {
	return o
}

func (o DFPInstanceAdministratorsResponsePtrOutput) ToDFPInstanceAdministratorsResponsePtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsResponsePtrOutput {
	return o
}

func (o DFPInstanceAdministratorsResponsePtrOutput) Elem() DFPInstanceAdministratorsResponseOutput {
	return o.ApplyT(func(v *DFPInstanceAdministratorsResponse) DFPInstanceAdministratorsResponse { return *v }).(DFPInstanceAdministratorsResponseOutput)
}

// An array of administrator user identities.
func (o DFPInstanceAdministratorsResponsePtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DFPInstanceAdministratorsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// SystemDataResponseInput is an input type that accepts SystemDataResponseArgs and SystemDataResponseOutput values.
// You can construct a concrete instance of `SystemDataResponseInput` via:
//
//          SystemDataResponseArgs{...}
type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseArgs struct {
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrInput `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrInput `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}

// SystemDataResponsePtrInput is an input type that accepts SystemDataResponseArgs, SystemDataResponsePtr and SystemDataResponsePtrOutput values.
// You can construct a concrete instance of `SystemDataResponsePtrInput` via:
//
//          SystemDataResponseArgs{...}
//
//  or:
//
//          nil
type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse { return *v }).(SystemDataResponseOutput)
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DFPInstanceAdministratorsOutput{})
	pulumi.RegisterOutputType(DFPInstanceAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(DFPInstanceAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(DFPInstanceAdministratorsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a Template Spec artifact.
type TemplateSpecArtifact struct {
	// The kind of artifact.
	Kind string `pulumi:"kind"`
	// A filesystem safe relative path of the artifact.
	Path string `pulumi:"path"`
}

// TemplateSpecArtifactInput is an input type that accepts TemplateSpecArtifactArgs and TemplateSpecArtifactOutput values.
// You can construct a concrete instance of `TemplateSpecArtifactInput` via:
//
//          TemplateSpecArtifactArgs{...}
type TemplateSpecArtifactInput interface {
	pulumi.Input

	ToTemplateSpecArtifactOutput() TemplateSpecArtifactOutput
	ToTemplateSpecArtifactOutputWithContext(context.Context) TemplateSpecArtifactOutput
}

// Represents a Template Spec artifact.
type TemplateSpecArtifactArgs struct {
	// The kind of artifact.
	Kind pulumi.StringInput `pulumi:"kind"`
	// A filesystem safe relative path of the artifact.
	Path pulumi.StringInput `pulumi:"path"`
}

func (TemplateSpecArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecArtifact)(nil)).Elem()
}

func (i TemplateSpecArtifactArgs) ToTemplateSpecArtifactOutput() TemplateSpecArtifactOutput {
	return i.ToTemplateSpecArtifactOutputWithContext(context.Background())
}

func (i TemplateSpecArtifactArgs) ToTemplateSpecArtifactOutputWithContext(ctx context.Context) TemplateSpecArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecArtifactOutput)
}

// TemplateSpecArtifactArrayInput is an input type that accepts TemplateSpecArtifactArray and TemplateSpecArtifactArrayOutput values.
// You can construct a concrete instance of `TemplateSpecArtifactArrayInput` via:
//
//          TemplateSpecArtifactArray{ TemplateSpecArtifactArgs{...} }
type TemplateSpecArtifactArrayInput interface {
	pulumi.Input

	ToTemplateSpecArtifactArrayOutput() TemplateSpecArtifactArrayOutput
	ToTemplateSpecArtifactArrayOutputWithContext(context.Context) TemplateSpecArtifactArrayOutput
}

type TemplateSpecArtifactArray []TemplateSpecArtifactInput

func (TemplateSpecArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateSpecArtifact)(nil)).Elem()
}

func (i TemplateSpecArtifactArray) ToTemplateSpecArtifactArrayOutput() TemplateSpecArtifactArrayOutput {
	return i.ToTemplateSpecArtifactArrayOutputWithContext(context.Background())
}

func (i TemplateSpecArtifactArray) ToTemplateSpecArtifactArrayOutputWithContext(ctx context.Context) TemplateSpecArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecArtifactArrayOutput)
}

// Represents a Template Spec artifact.
type TemplateSpecArtifactOutput struct{ *pulumi.OutputState }

func (TemplateSpecArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecArtifact)(nil)).Elem()
}

func (o TemplateSpecArtifactOutput) ToTemplateSpecArtifactOutput() TemplateSpecArtifactOutput {
	return o
}

func (o TemplateSpecArtifactOutput) ToTemplateSpecArtifactOutputWithContext(ctx context.Context) TemplateSpecArtifactOutput {
	return o
}

// The kind of artifact.
func (o TemplateSpecArtifactOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecArtifact) string { return v.Kind }).(pulumi.StringOutput)
}

// A filesystem safe relative path of the artifact.
func (o TemplateSpecArtifactOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecArtifact) string { return v.Path }).(pulumi.StringOutput)
}

type TemplateSpecArtifactArrayOutput struct{ *pulumi.OutputState }

func (TemplateSpecArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateSpecArtifact)(nil)).Elem()
}

func (o TemplateSpecArtifactArrayOutput) ToTemplateSpecArtifactArrayOutput() TemplateSpecArtifactArrayOutput {
	return o
}

func (o TemplateSpecArtifactArrayOutput) ToTemplateSpecArtifactArrayOutputWithContext(ctx context.Context) TemplateSpecArtifactArrayOutput {
	return o
}

func (o TemplateSpecArtifactArrayOutput) Index(i pulumi.IntInput) TemplateSpecArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TemplateSpecArtifact {
		return vs[0].([]TemplateSpecArtifact)[vs[1].(int)]
	}).(TemplateSpecArtifactOutput)
}

// Represents a Template Spec artifact.
type TemplateSpecArtifactResponse struct {
	// The kind of artifact.
	Kind string `pulumi:"kind"`
	// A filesystem safe relative path of the artifact.
	Path string `pulumi:"path"`
}

// TemplateSpecArtifactResponseInput is an input type that accepts TemplateSpecArtifactResponseArgs and TemplateSpecArtifactResponseOutput values.
// You can construct a concrete instance of `TemplateSpecArtifactResponseInput` via:
//
//          TemplateSpecArtifactResponseArgs{...}
type TemplateSpecArtifactResponseInput interface {
	pulumi.Input

	ToTemplateSpecArtifactResponseOutput() TemplateSpecArtifactResponseOutput
	ToTemplateSpecArtifactResponseOutputWithContext(context.Context) TemplateSpecArtifactResponseOutput
}

// Represents a Template Spec artifact.
type TemplateSpecArtifactResponseArgs struct {
	// The kind of artifact.
	Kind pulumi.StringInput `pulumi:"kind"`
	// A filesystem safe relative path of the artifact.
	Path pulumi.StringInput `pulumi:"path"`
}

func (TemplateSpecArtifactResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecArtifactResponse)(nil)).Elem()
}

func (i TemplateSpecArtifactResponseArgs) ToTemplateSpecArtifactResponseOutput() TemplateSpecArtifactResponseOutput {
	return i.ToTemplateSpecArtifactResponseOutputWithContext(context.Background())
}

func (i TemplateSpecArtifactResponseArgs) ToTemplateSpecArtifactResponseOutputWithContext(ctx context.Context) TemplateSpecArtifactResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecArtifactResponseOutput)
}

// TemplateSpecArtifactResponseArrayInput is an input type that accepts TemplateSpecArtifactResponseArray and TemplateSpecArtifactResponseArrayOutput values.
// You can construct a concrete instance of `TemplateSpecArtifactResponseArrayInput` via:
//
//          TemplateSpecArtifactResponseArray{ TemplateSpecArtifactResponseArgs{...} }
type TemplateSpecArtifactResponseArrayInput interface {
	pulumi.Input

	ToTemplateSpecArtifactResponseArrayOutput() TemplateSpecArtifactResponseArrayOutput
	ToTemplateSpecArtifactResponseArrayOutputWithContext(context.Context) TemplateSpecArtifactResponseArrayOutput
}

type TemplateSpecArtifactResponseArray []TemplateSpecArtifactResponseInput

func (TemplateSpecArtifactResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateSpecArtifactResponse)(nil)).Elem()
}

func (i TemplateSpecArtifactResponseArray) ToTemplateSpecArtifactResponseArrayOutput() TemplateSpecArtifactResponseArrayOutput {
	return i.ToTemplateSpecArtifactResponseArrayOutputWithContext(context.Background())
}

func (i TemplateSpecArtifactResponseArray) ToTemplateSpecArtifactResponseArrayOutputWithContext(ctx context.Context) TemplateSpecArtifactResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecArtifactResponseArrayOutput)
}

// Represents a Template Spec artifact.
type TemplateSpecArtifactResponseOutput struct{ *pulumi.OutputState }

func (TemplateSpecArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecArtifactResponse)(nil)).Elem()
}

func (o TemplateSpecArtifactResponseOutput) ToTemplateSpecArtifactResponseOutput() TemplateSpecArtifactResponseOutput {
	return o
}

func (o TemplateSpecArtifactResponseOutput) ToTemplateSpecArtifactResponseOutputWithContext(ctx context.Context) TemplateSpecArtifactResponseOutput {
	return o
}

// The kind of artifact.
func (o TemplateSpecArtifactResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecArtifactResponse) string { return v.Kind }).(pulumi.StringOutput)
}

// A filesystem safe relative path of the artifact.
func (o TemplateSpecArtifactResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecArtifactResponse) string { return v.Path }).(pulumi.StringOutput)
}

type TemplateSpecArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (TemplateSpecArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateSpecArtifactResponse)(nil)).Elem()
}

func (o TemplateSpecArtifactResponseArrayOutput) ToTemplateSpecArtifactResponseArrayOutput() TemplateSpecArtifactResponseArrayOutput {
	return o
}

func (o TemplateSpecArtifactResponseArrayOutput) ToTemplateSpecArtifactResponseArrayOutputWithContext(ctx context.Context) TemplateSpecArtifactResponseArrayOutput {
	return o
}

func (o TemplateSpecArtifactResponseArrayOutput) Index(i pulumi.IntInput) TemplateSpecArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TemplateSpecArtifactResponse {
		return vs[0].([]TemplateSpecArtifactResponse)[vs[1].(int)]
	}).(TemplateSpecArtifactResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The type of identity that last modified the resource.
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
	// The type of identity that last modified the resource.
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

// The type of identity that last modified the resource.
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

// The type of identity that last modified the resource.
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
	pulumi.RegisterOutputType(TemplateSpecArtifactOutput{})
	pulumi.RegisterOutputType(TemplateSpecArtifactArrayOutput{})
	pulumi.RegisterOutputType(TemplateSpecArtifactResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
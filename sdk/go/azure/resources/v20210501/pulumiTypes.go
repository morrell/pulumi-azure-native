// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifact struct {
	// A filesystem safe relative path of the artifact.
	Path string `pulumi:"path"`
	// The Azure Resource Manager template.
	Template interface{} `pulumi:"template"`
}

// LinkedTemplateArtifactInput is an input type that accepts LinkedTemplateArtifactArgs and LinkedTemplateArtifactOutput values.
// You can construct a concrete instance of `LinkedTemplateArtifactInput` via:
//
//          LinkedTemplateArtifactArgs{...}
type LinkedTemplateArtifactInput interface {
	pulumi.Input

	ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput
	ToLinkedTemplateArtifactOutputWithContext(context.Context) LinkedTemplateArtifactOutput
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactArgs struct {
	// A filesystem safe relative path of the artifact.
	Path pulumi.StringInput `pulumi:"path"`
	// The Azure Resource Manager template.
	Template pulumi.Input `pulumi:"template"`
}

func (LinkedTemplateArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifact)(nil)).Elem()
}

func (i LinkedTemplateArtifactArgs) ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput {
	return i.ToLinkedTemplateArtifactOutputWithContext(context.Background())
}

func (i LinkedTemplateArtifactArgs) ToLinkedTemplateArtifactOutputWithContext(ctx context.Context) LinkedTemplateArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedTemplateArtifactOutput)
}

// LinkedTemplateArtifactArrayInput is an input type that accepts LinkedTemplateArtifactArray and LinkedTemplateArtifactArrayOutput values.
// You can construct a concrete instance of `LinkedTemplateArtifactArrayInput` via:
//
//          LinkedTemplateArtifactArray{ LinkedTemplateArtifactArgs{...} }
type LinkedTemplateArtifactArrayInput interface {
	pulumi.Input

	ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput
	ToLinkedTemplateArtifactArrayOutputWithContext(context.Context) LinkedTemplateArtifactArrayOutput
}

type LinkedTemplateArtifactArray []LinkedTemplateArtifactInput

func (LinkedTemplateArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifact)(nil)).Elem()
}

func (i LinkedTemplateArtifactArray) ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput {
	return i.ToLinkedTemplateArtifactArrayOutputWithContext(context.Background())
}

func (i LinkedTemplateArtifactArray) ToLinkedTemplateArtifactArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedTemplateArtifactArrayOutput)
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifact)(nil)).Elem()
}

func (o LinkedTemplateArtifactOutput) ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput {
	return o
}

func (o LinkedTemplateArtifactOutput) ToLinkedTemplateArtifactOutputWithContext(ctx context.Context) LinkedTemplateArtifactOutput {
	return o
}

// A filesystem safe relative path of the artifact.
func (o LinkedTemplateArtifactOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedTemplateArtifact) string { return v.Path }).(pulumi.StringOutput)
}

// The Azure Resource Manager template.
func (o LinkedTemplateArtifactOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LinkedTemplateArtifact) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type LinkedTemplateArtifactArrayOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifact)(nil)).Elem()
}

func (o LinkedTemplateArtifactArrayOutput) ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput {
	return o
}

func (o LinkedTemplateArtifactArrayOutput) ToLinkedTemplateArtifactArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactArrayOutput {
	return o
}

func (o LinkedTemplateArtifactArrayOutput) Index(i pulumi.IntInput) LinkedTemplateArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedTemplateArtifact {
		return vs[0].([]LinkedTemplateArtifact)[vs[1].(int)]
	}).(LinkedTemplateArtifactOutput)
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactResponse struct {
	// A filesystem safe relative path of the artifact.
	Path string `pulumi:"path"`
	// The Azure Resource Manager template.
	Template interface{} `pulumi:"template"`
}

// LinkedTemplateArtifactResponseInput is an input type that accepts LinkedTemplateArtifactResponseArgs and LinkedTemplateArtifactResponseOutput values.
// You can construct a concrete instance of `LinkedTemplateArtifactResponseInput` via:
//
//          LinkedTemplateArtifactResponseArgs{...}
type LinkedTemplateArtifactResponseInput interface {
	pulumi.Input

	ToLinkedTemplateArtifactResponseOutput() LinkedTemplateArtifactResponseOutput
	ToLinkedTemplateArtifactResponseOutputWithContext(context.Context) LinkedTemplateArtifactResponseOutput
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactResponseArgs struct {
	// A filesystem safe relative path of the artifact.
	Path pulumi.StringInput `pulumi:"path"`
	// The Azure Resource Manager template.
	Template pulumi.Input `pulumi:"template"`
}

func (LinkedTemplateArtifactResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifactResponse)(nil)).Elem()
}

func (i LinkedTemplateArtifactResponseArgs) ToLinkedTemplateArtifactResponseOutput() LinkedTemplateArtifactResponseOutput {
	return i.ToLinkedTemplateArtifactResponseOutputWithContext(context.Background())
}

func (i LinkedTemplateArtifactResponseArgs) ToLinkedTemplateArtifactResponseOutputWithContext(ctx context.Context) LinkedTemplateArtifactResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedTemplateArtifactResponseOutput)
}

// LinkedTemplateArtifactResponseArrayInput is an input type that accepts LinkedTemplateArtifactResponseArray and LinkedTemplateArtifactResponseArrayOutput values.
// You can construct a concrete instance of `LinkedTemplateArtifactResponseArrayInput` via:
//
//          LinkedTemplateArtifactResponseArray{ LinkedTemplateArtifactResponseArgs{...} }
type LinkedTemplateArtifactResponseArrayInput interface {
	pulumi.Input

	ToLinkedTemplateArtifactResponseArrayOutput() LinkedTemplateArtifactResponseArrayOutput
	ToLinkedTemplateArtifactResponseArrayOutputWithContext(context.Context) LinkedTemplateArtifactResponseArrayOutput
}

type LinkedTemplateArtifactResponseArray []LinkedTemplateArtifactResponseInput

func (LinkedTemplateArtifactResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifactResponse)(nil)).Elem()
}

func (i LinkedTemplateArtifactResponseArray) ToLinkedTemplateArtifactResponseArrayOutput() LinkedTemplateArtifactResponseArrayOutput {
	return i.ToLinkedTemplateArtifactResponseArrayOutputWithContext(context.Background())
}

func (i LinkedTemplateArtifactResponseArray) ToLinkedTemplateArtifactResponseArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedTemplateArtifactResponseArrayOutput)
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactResponseOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifactResponse)(nil)).Elem()
}

func (o LinkedTemplateArtifactResponseOutput) ToLinkedTemplateArtifactResponseOutput() LinkedTemplateArtifactResponseOutput {
	return o
}

func (o LinkedTemplateArtifactResponseOutput) ToLinkedTemplateArtifactResponseOutputWithContext(ctx context.Context) LinkedTemplateArtifactResponseOutput {
	return o
}

// A filesystem safe relative path of the artifact.
func (o LinkedTemplateArtifactResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedTemplateArtifactResponse) string { return v.Path }).(pulumi.StringOutput)
}

// The Azure Resource Manager template.
func (o LinkedTemplateArtifactResponseOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LinkedTemplateArtifactResponse) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type LinkedTemplateArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifactResponse)(nil)).Elem()
}

func (o LinkedTemplateArtifactResponseArrayOutput) ToLinkedTemplateArtifactResponseArrayOutput() LinkedTemplateArtifactResponseArrayOutput {
	return o
}

func (o LinkedTemplateArtifactResponseArrayOutput) ToLinkedTemplateArtifactResponseArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactResponseArrayOutput {
	return o
}

func (o LinkedTemplateArtifactResponseArrayOutput) Index(i pulumi.IntInput) LinkedTemplateArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedTemplateArtifactResponse {
		return vs[0].([]LinkedTemplateArtifactResponse)[vs[1].(int)]
	}).(LinkedTemplateArtifactResponseOutput)
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

// High-level information about a Template Spec version.
type TemplateSpecVersionInfoResponse struct {
	// Template Spec version description.
	Description string `pulumi:"description"`
	// The timestamp of when the version was created.
	TimeCreated string `pulumi:"timeCreated"`
	// The timestamp of when the version was last modified.
	TimeModified string `pulumi:"timeModified"`
}

// TemplateSpecVersionInfoResponseInput is an input type that accepts TemplateSpecVersionInfoResponseArgs and TemplateSpecVersionInfoResponseOutput values.
// You can construct a concrete instance of `TemplateSpecVersionInfoResponseInput` via:
//
//          TemplateSpecVersionInfoResponseArgs{...}
type TemplateSpecVersionInfoResponseInput interface {
	pulumi.Input

	ToTemplateSpecVersionInfoResponseOutput() TemplateSpecVersionInfoResponseOutput
	ToTemplateSpecVersionInfoResponseOutputWithContext(context.Context) TemplateSpecVersionInfoResponseOutput
}

// High-level information about a Template Spec version.
type TemplateSpecVersionInfoResponseArgs struct {
	// Template Spec version description.
	Description pulumi.StringInput `pulumi:"description"`
	// The timestamp of when the version was created.
	TimeCreated pulumi.StringInput `pulumi:"timeCreated"`
	// The timestamp of when the version was last modified.
	TimeModified pulumi.StringInput `pulumi:"timeModified"`
}

func (TemplateSpecVersionInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (i TemplateSpecVersionInfoResponseArgs) ToTemplateSpecVersionInfoResponseOutput() TemplateSpecVersionInfoResponseOutput {
	return i.ToTemplateSpecVersionInfoResponseOutputWithContext(context.Background())
}

func (i TemplateSpecVersionInfoResponseArgs) ToTemplateSpecVersionInfoResponseOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecVersionInfoResponseOutput)
}

// TemplateSpecVersionInfoResponseMapInput is an input type that accepts TemplateSpecVersionInfoResponseMap and TemplateSpecVersionInfoResponseMapOutput values.
// You can construct a concrete instance of `TemplateSpecVersionInfoResponseMapInput` via:
//
//          TemplateSpecVersionInfoResponseMap{ "key": TemplateSpecVersionInfoResponseArgs{...} }
type TemplateSpecVersionInfoResponseMapInput interface {
	pulumi.Input

	ToTemplateSpecVersionInfoResponseMapOutput() TemplateSpecVersionInfoResponseMapOutput
	ToTemplateSpecVersionInfoResponseMapOutputWithContext(context.Context) TemplateSpecVersionInfoResponseMapOutput
}

type TemplateSpecVersionInfoResponseMap map[string]TemplateSpecVersionInfoResponseInput

func (TemplateSpecVersionInfoResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (i TemplateSpecVersionInfoResponseMap) ToTemplateSpecVersionInfoResponseMapOutput() TemplateSpecVersionInfoResponseMapOutput {
	return i.ToTemplateSpecVersionInfoResponseMapOutputWithContext(context.Background())
}

func (i TemplateSpecVersionInfoResponseMap) ToTemplateSpecVersionInfoResponseMapOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecVersionInfoResponseMapOutput)
}

// High-level information about a Template Spec version.
type TemplateSpecVersionInfoResponseOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (o TemplateSpecVersionInfoResponseOutput) ToTemplateSpecVersionInfoResponseOutput() TemplateSpecVersionInfoResponseOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseOutput) ToTemplateSpecVersionInfoResponseOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseOutput {
	return o
}

// Template Spec version description.
func (o TemplateSpecVersionInfoResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.Description }).(pulumi.StringOutput)
}

// The timestamp of when the version was created.
func (o TemplateSpecVersionInfoResponseOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// The timestamp of when the version was last modified.
func (o TemplateSpecVersionInfoResponseOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.TimeModified }).(pulumi.StringOutput)
}

type TemplateSpecVersionInfoResponseMapOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionInfoResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToTemplateSpecVersionInfoResponseMapOutput() TemplateSpecVersionInfoResponseMapOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToTemplateSpecVersionInfoResponseMapOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseMapOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseMapOutput) MapIndex(k pulumi.StringInput) TemplateSpecVersionInfoResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TemplateSpecVersionInfoResponse {
		return vs[0].(map[string]TemplateSpecVersionInfoResponse)[vs[1].(string)]
	}).(TemplateSpecVersionInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedTemplateArtifactOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactArrayOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactResponseOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TemplateSpecVersionInfoResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecVersionInfoResponseMapOutput{})
}

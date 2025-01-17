// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180907preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SKU name.
type AzureSkuName string

const (
	AzureSkuNameKC8     = AzureSkuName("KC8")
	AzureSkuNameKC16    = AzureSkuName("KC16")
	AzureSkuNameKS8     = AzureSkuName("KS8")
	AzureSkuNameKS16    = AzureSkuName("KS16")
	AzureSkuName_D13_v2 = AzureSkuName("D13_v2")
	AzureSkuName_D14_v2 = AzureSkuName("D14_v2")
	AzureSkuNameL8      = AzureSkuName("L8")
	AzureSkuNameL16     = AzureSkuName("L16")
)

func (AzureSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuName)(nil)).Elem()
}

func (e AzureSkuName) ToAzureSkuNameOutput() AzureSkuNameOutput {
	return pulumi.ToOutput(e).(AzureSkuNameOutput)
}

func (e AzureSkuName) ToAzureSkuNameOutputWithContext(ctx context.Context) AzureSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSkuNameOutput)
}

func (e AzureSkuName) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return e.ToAzureSkuNamePtrOutputWithContext(context.Background())
}

func (e AzureSkuName) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return AzureSkuName(e).ToAzureSkuNameOutputWithContext(ctx).ToAzureSkuNamePtrOutputWithContext(ctx)
}

func (e AzureSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSkuNameOutput struct{ *pulumi.OutputState }

func (AzureSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuName)(nil)).Elem()
}

func (o AzureSkuNameOutput) ToAzureSkuNameOutput() AzureSkuNameOutput {
	return o
}

func (o AzureSkuNameOutput) ToAzureSkuNameOutputWithContext(ctx context.Context) AzureSkuNameOutput {
	return o
}

func (o AzureSkuNameOutput) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return o.ToAzureSkuNamePtrOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuName) *AzureSkuName {
		return &v
	}).(AzureSkuNamePtrOutput)
}

func (o AzureSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSkuNamePtrOutput struct{ *pulumi.OutputState }

func (AzureSkuNamePtrOutput) ElementType() reflect.Type {
	return azureSkuNamePtrType
}

func (o AzureSkuNamePtrOutput) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return o
}

func (o AzureSkuNamePtrOutput) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return o
}

func (o AzureSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o AzureSkuNamePtrOutput) Elem() AzureSkuNameOutput {
	return o.ApplyT(func(v *AzureSkuName) AzureSkuName {
		var ret AzureSkuName
		if v != nil {
			ret = *v
		}
		return ret
	}).(AzureSkuNameOutput)
}

// AzureSkuNameInput is an input type that accepts AzureSkuNameArgs and AzureSkuNameOutput values.
// You can construct a concrete instance of `AzureSkuNameInput` via:
//
//          AzureSkuNameArgs{...}
type AzureSkuNameInput interface {
	pulumi.Input

	ToAzureSkuNameOutput() AzureSkuNameOutput
	ToAzureSkuNameOutputWithContext(context.Context) AzureSkuNameOutput
}

var azureSkuNamePtrType = reflect.TypeOf((**AzureSkuName)(nil)).Elem()

type AzureSkuNamePtrInput interface {
	pulumi.Input

	ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput
	ToAzureSkuNamePtrOutputWithContext(context.Context) AzureSkuNamePtrOutput
}

type azureSkuNamePtr string

func AzureSkuNamePtr(v string) AzureSkuNamePtrInput {
	return (*azureSkuNamePtr)(&v)
}

func (*azureSkuNamePtr) ElementType() reflect.Type {
	return azureSkuNamePtrType
}

func (in *azureSkuNamePtr) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return pulumi.ToOutput(in).(AzureSkuNamePtrOutput)
}

func (in *azureSkuNamePtr) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSkuNamePtrOutput)
}

// SKU tier.
type AzureSkuTier string

const (
	AzureSkuTierStandard = AzureSkuTier("Standard")
)

func (AzureSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuTier)(nil)).Elem()
}

func (e AzureSkuTier) ToAzureSkuTierOutput() AzureSkuTierOutput {
	return pulumi.ToOutput(e).(AzureSkuTierOutput)
}

func (e AzureSkuTier) ToAzureSkuTierOutputWithContext(ctx context.Context) AzureSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSkuTierOutput)
}

func (e AzureSkuTier) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return e.ToAzureSkuTierPtrOutputWithContext(context.Background())
}

func (e AzureSkuTier) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return AzureSkuTier(e).ToAzureSkuTierOutputWithContext(ctx).ToAzureSkuTierPtrOutputWithContext(ctx)
}

func (e AzureSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSkuTierOutput struct{ *pulumi.OutputState }

func (AzureSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuTier)(nil)).Elem()
}

func (o AzureSkuTierOutput) ToAzureSkuTierOutput() AzureSkuTierOutput {
	return o
}

func (o AzureSkuTierOutput) ToAzureSkuTierOutputWithContext(ctx context.Context) AzureSkuTierOutput {
	return o
}

func (o AzureSkuTierOutput) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return o.ToAzureSkuTierPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuTier) *AzureSkuTier {
		return &v
	}).(AzureSkuTierPtrOutput)
}

func (o AzureSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSkuTierPtrOutput struct{ *pulumi.OutputState }

func (AzureSkuTierPtrOutput) ElementType() reflect.Type {
	return azureSkuTierPtrType
}

func (o AzureSkuTierPtrOutput) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return o
}

func (o AzureSkuTierPtrOutput) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return o
}

func (o AzureSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o AzureSkuTierPtrOutput) Elem() AzureSkuTierOutput {
	return o.ApplyT(func(v *AzureSkuTier) AzureSkuTier {
		var ret AzureSkuTier
		if v != nil {
			ret = *v
		}
		return ret
	}).(AzureSkuTierOutput)
}

// AzureSkuTierInput is an input type that accepts AzureSkuTierArgs and AzureSkuTierOutput values.
// You can construct a concrete instance of `AzureSkuTierInput` via:
//
//          AzureSkuTierArgs{...}
type AzureSkuTierInput interface {
	pulumi.Input

	ToAzureSkuTierOutput() AzureSkuTierOutput
	ToAzureSkuTierOutputWithContext(context.Context) AzureSkuTierOutput
}

var azureSkuTierPtrType = reflect.TypeOf((**AzureSkuTier)(nil)).Elem()

type AzureSkuTierPtrInput interface {
	pulumi.Input

	ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput
	ToAzureSkuTierPtrOutputWithContext(context.Context) AzureSkuTierPtrOutput
}

type azureSkuTierPtr string

func AzureSkuTierPtr(v string) AzureSkuTierPtrInput {
	return (*azureSkuTierPtr)(&v)
}

func (*azureSkuTierPtr) ElementType() reflect.Type {
	return azureSkuTierPtrType
}

func (in *azureSkuTierPtr) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return pulumi.ToOutput(in).(AzureSkuTierPtrOutput)
}

func (in *azureSkuTierPtr) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSkuTierPtrOutput)
}

// The data format of the message. Optionally the data format can be added to each message.
type DataFormat string

const (
	DataFormatMULTIJSON = DataFormat("MULTIJSON")
	DataFormatJSON      = DataFormat("JSON")
	DataFormatCSV       = DataFormat("CSV")
)

func (DataFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFormat)(nil)).Elem()
}

func (e DataFormat) ToDataFormatOutput() DataFormatOutput {
	return pulumi.ToOutput(e).(DataFormatOutput)
}

func (e DataFormat) ToDataFormatOutputWithContext(ctx context.Context) DataFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataFormatOutput)
}

func (e DataFormat) ToDataFormatPtrOutput() DataFormatPtrOutput {
	return e.ToDataFormatPtrOutputWithContext(context.Background())
}

func (e DataFormat) ToDataFormatPtrOutputWithContext(ctx context.Context) DataFormatPtrOutput {
	return DataFormat(e).ToDataFormatOutputWithContext(ctx).ToDataFormatPtrOutputWithContext(ctx)
}

func (e DataFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataFormatOutput struct{ *pulumi.OutputState }

func (DataFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFormat)(nil)).Elem()
}

func (o DataFormatOutput) ToDataFormatOutput() DataFormatOutput {
	return o
}

func (o DataFormatOutput) ToDataFormatOutputWithContext(ctx context.Context) DataFormatOutput {
	return o
}

func (o DataFormatOutput) ToDataFormatPtrOutput() DataFormatPtrOutput {
	return o.ToDataFormatPtrOutputWithContext(context.Background())
}

func (o DataFormatOutput) ToDataFormatPtrOutputWithContext(ctx context.Context) DataFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataFormat) *DataFormat {
		return &v
	}).(DataFormatPtrOutput)
}

func (o DataFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataFormatPtrOutput struct{ *pulumi.OutputState }

func (DataFormatPtrOutput) ElementType() reflect.Type {
	return dataFormatPtrType
}

func (o DataFormatPtrOutput) ToDataFormatPtrOutput() DataFormatPtrOutput {
	return o
}

func (o DataFormatPtrOutput) ToDataFormatPtrOutputWithContext(ctx context.Context) DataFormatPtrOutput {
	return o
}

func (o DataFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func (o DataFormatPtrOutput) Elem() DataFormatOutput {
	return o.ApplyT(func(v *DataFormat) DataFormat {
		var ret DataFormat
		if v != nil {
			ret = *v
		}
		return ret
	}).(DataFormatOutput)
}

// DataFormatInput is an input type that accepts DataFormatArgs and DataFormatOutput values.
// You can construct a concrete instance of `DataFormatInput` via:
//
//          DataFormatArgs{...}
type DataFormatInput interface {
	pulumi.Input

	ToDataFormatOutput() DataFormatOutput
	ToDataFormatOutputWithContext(context.Context) DataFormatOutput
}

var dataFormatPtrType = reflect.TypeOf((**DataFormat)(nil)).Elem()

type DataFormatPtrInput interface {
	pulumi.Input

	ToDataFormatPtrOutput() DataFormatPtrOutput
	ToDataFormatPtrOutputWithContext(context.Context) DataFormatPtrOutput
}

type dataFormatPtr string

func DataFormatPtr(v string) DataFormatPtrInput {
	return (*dataFormatPtr)(&v)
}

func (*dataFormatPtr) ElementType() reflect.Type {
	return dataFormatPtrType
}

func (in *dataFormatPtr) ToDataFormatPtrOutput() DataFormatPtrOutput {
	return pulumi.ToOutput(in).(DataFormatPtrOutput)
}

func (in *dataFormatPtr) ToDataFormatPtrOutputWithContext(ctx context.Context) DataFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataFormatPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureSkuNameOutput{})
	pulumi.RegisterOutputType(AzureSkuNamePtrOutput{})
	pulumi.RegisterOutputType(AzureSkuTierOutput{})
	pulumi.RegisterOutputType(AzureSkuTierPtrOutput{})
	pulumi.RegisterOutputType(DataFormatOutput{})
	pulumi.RegisterOutputType(DataFormatPtrOutput{})
}

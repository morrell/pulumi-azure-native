// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Properties to configure Encryption
type Encryption struct {
	// Enumerates the possible value of keySource for Encryption
	KeySource *string `pulumi:"keySource"`
	// Properties of KeyVault
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
}

// EncryptionInput is an input type that accepts EncryptionArgs and EncryptionOutput values.
// You can construct a concrete instance of `EncryptionInput` via:
//
//          EncryptionArgs{...}
type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

// Properties to configure Encryption
type EncryptionArgs struct {
	// Enumerates the possible value of keySource for Encryption
	KeySource pulumi.StringPtrInput `pulumi:"keySource"`
	// Properties of KeyVault
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}

// EncryptionPtrInput is an input type that accepts EncryptionArgs, EncryptionPtr and EncryptionPtrOutput values.
// You can construct a concrete instance of `EncryptionPtrInput` via:
//
//          EncryptionArgs{...}
//
//  or:
//
//          nil
type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

// Properties to configure Encryption
type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyT(func(v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

// Enumerates the possible value of keySource for Encryption
func (o EncryptionOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

// Properties of KeyVault
func (o EncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v Encryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption { return *v }).(EncryptionOutput)
}

// Enumerates the possible value of keySource for Encryption
func (o EncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

// Properties of KeyVault
func (o EncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *Encryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

// Properties to configure Encryption
type EncryptionResponse struct {
	// Enumerates the possible value of keySource for Encryption
	KeySource *string `pulumi:"keySource"`
	// Properties of KeyVault
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
}

// EncryptionResponseInput is an input type that accepts EncryptionResponseArgs and EncryptionResponseOutput values.
// You can construct a concrete instance of `EncryptionResponseInput` via:
//
//          EncryptionResponseArgs{...}
type EncryptionResponseInput interface {
	pulumi.Input

	ToEncryptionResponseOutput() EncryptionResponseOutput
	ToEncryptionResponseOutputWithContext(context.Context) EncryptionResponseOutput
}

// Properties to configure Encryption
type EncryptionResponseArgs struct {
	// Enumerates the possible value of keySource for Encryption
	KeySource pulumi.StringPtrInput `pulumi:"keySource"`
	// Properties of KeyVault
	KeyVaultProperties KeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
}

func (EncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return i.ToEncryptionResponseOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput)
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput).ToEncryptionResponsePtrOutputWithContext(ctx)
}

// EncryptionResponsePtrInput is an input type that accepts EncryptionResponseArgs, EncryptionResponsePtr and EncryptionResponsePtrOutput values.
// You can construct a concrete instance of `EncryptionResponsePtrInput` via:
//
//          EncryptionResponseArgs{...}
//
//  or:
//
//          nil
type EncryptionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput
	ToEncryptionResponsePtrOutputWithContext(context.Context) EncryptionResponsePtrOutput
}

type encryptionResponsePtrType EncryptionResponseArgs

func EncryptionResponsePtr(v *EncryptionResponseArgs) EncryptionResponsePtrInput {
	return (*encryptionResponsePtrType)(v)
}

func (*encryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponsePtrOutput)
}

// Properties to configure Encryption
type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
}

// Enumerates the possible value of keySource for Encryption
func (o EncryptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

// Properties of KeyVault
func (o EncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse { return *v }).(EncryptionResponseOutput)
}

// Enumerates the possible value of keySource for Encryption
func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

// Properties of KeyVault
func (o EncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

// Properties to configure keyVault Properties
type KeyVaultProperties struct {
	// Name of the Key from KeyVault
	KeyName *string `pulumi:"keyName"`
	// Uri of KeyVault
	KeyVaultUri *string `pulumi:"keyVaultUri"`
}

// KeyVaultPropertiesInput is an input type that accepts KeyVaultPropertiesArgs and KeyVaultPropertiesOutput values.
// You can construct a concrete instance of `KeyVaultPropertiesInput` via:
//
//          KeyVaultPropertiesArgs{...}
type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

// Properties to configure keyVault Properties
type KeyVaultPropertiesArgs struct {
	// Name of the Key from KeyVault
	KeyName pulumi.StringPtrInput `pulumi:"keyName"`
	// Uri of KeyVault
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}

// KeyVaultPropertiesPtrInput is an input type that accepts KeyVaultPropertiesArgs, KeyVaultPropertiesPtr and KeyVaultPropertiesPtrOutput values.
// You can construct a concrete instance of `KeyVaultPropertiesPtrInput` via:
//
//          KeyVaultPropertiesArgs{...}
//
//  or:
//
//          nil
type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

// Properties to configure keyVault Properties
type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

// Name of the Key from KeyVault
func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

// Uri of KeyVault
func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties { return *v }).(KeyVaultPropertiesOutput)
}

// Name of the Key from KeyVault
func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

// Uri of KeyVault
func (o KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

// Properties to configure keyVault Properties
type KeyVaultPropertiesResponse struct {
	// Name of the Key from KeyVault
	KeyName *string `pulumi:"keyName"`
	// Uri of KeyVault
	KeyVaultUri *string `pulumi:"keyVaultUri"`
}

// KeyVaultPropertiesResponseInput is an input type that accepts KeyVaultPropertiesResponseArgs and KeyVaultPropertiesResponseOutput values.
// You can construct a concrete instance of `KeyVaultPropertiesResponseInput` via:
//
//          KeyVaultPropertiesResponseArgs{...}
type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

// Properties to configure keyVault Properties
type KeyVaultPropertiesResponseArgs struct {
	// Name of the Key from KeyVault
	KeyName pulumi.StringPtrInput `pulumi:"keyName"`
	// Uri of KeyVault
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
}

func (KeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return i.ToKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput)
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput).ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}

// KeyVaultPropertiesResponsePtrInput is an input type that accepts KeyVaultPropertiesResponseArgs, KeyVaultPropertiesResponsePtr and KeyVaultPropertiesResponsePtrOutput values.
// You can construct a concrete instance of `KeyVaultPropertiesResponsePtrInput` via:
//
//          KeyVaultPropertiesResponseArgs{...}
//
//  or:
//
//          nil
type KeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput
	ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) KeyVaultPropertiesResponsePtrOutput
}

type keyVaultPropertiesResponsePtrType KeyVaultPropertiesResponseArgs

func KeyVaultPropertiesResponsePtr(v *KeyVaultPropertiesResponseArgs) KeyVaultPropertiesResponsePtrInput {
	return (*keyVaultPropertiesResponsePtrType)(v)
}

func (*keyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponsePtrOutput)
}

// Properties to configure keyVault Properties
type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *KeyVaultPropertiesResponse {
		return &v
	}).(KeyVaultPropertiesResponsePtrOutput)
}

// Name of the Key from KeyVault
func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

// Uri of KeyVault
func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse { return *v }).(KeyVaultPropertiesResponseOutput)
}

// Name of the Key from KeyVault
func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

// Uri of KeyVault
func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

// SKU of the namespace.
type SBSku struct {
	// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity *int `pulumi:"capacity"`
	// Name of this SKU.
	Name string `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier *string `pulumi:"tier"`
}

// SBSkuInput is an input type that accepts SBSkuArgs and SBSkuOutput values.
// You can construct a concrete instance of `SBSkuInput` via:
//
//          SBSkuArgs{...}
type SBSkuInput interface {
	pulumi.Input

	ToSBSkuOutput() SBSkuOutput
	ToSBSkuOutputWithContext(context.Context) SBSkuOutput
}

// SKU of the namespace.
type SBSkuArgs struct {
	// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// Name of this SKU.
	Name pulumi.StringInput `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SBSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSku)(nil)).Elem()
}

func (i SBSkuArgs) ToSBSkuOutput() SBSkuOutput {
	return i.ToSBSkuOutputWithContext(context.Background())
}

func (i SBSkuArgs) ToSBSkuOutputWithContext(ctx context.Context) SBSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuOutput)
}

func (i SBSkuArgs) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return i.ToSBSkuPtrOutputWithContext(context.Background())
}

func (i SBSkuArgs) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuOutput).ToSBSkuPtrOutputWithContext(ctx)
}

// SBSkuPtrInput is an input type that accepts SBSkuArgs, SBSkuPtr and SBSkuPtrOutput values.
// You can construct a concrete instance of `SBSkuPtrInput` via:
//
//          SBSkuArgs{...}
//
//  or:
//
//          nil
type SBSkuPtrInput interface {
	pulumi.Input

	ToSBSkuPtrOutput() SBSkuPtrOutput
	ToSBSkuPtrOutputWithContext(context.Context) SBSkuPtrOutput
}

type sbskuPtrType SBSkuArgs

func SBSkuPtr(v *SBSkuArgs) SBSkuPtrInput {
	return (*sbskuPtrType)(v)
}

func (*sbskuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSku)(nil)).Elem()
}

func (i *sbskuPtrType) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return i.ToSBSkuPtrOutputWithContext(context.Background())
}

func (i *sbskuPtrType) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuPtrOutput)
}

// SKU of the namespace.
type SBSkuOutput struct{ *pulumi.OutputState }

func (SBSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSku)(nil)).Elem()
}

func (o SBSkuOutput) ToSBSkuOutput() SBSkuOutput {
	return o
}

func (o SBSkuOutput) ToSBSkuOutputWithContext(ctx context.Context) SBSkuOutput {
	return o
}

func (o SBSkuOutput) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return o.ToSBSkuPtrOutputWithContext(context.Background())
}

func (o SBSkuOutput) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return o.ApplyT(func(v SBSku) *SBSku {
		return &v
	}).(SBSkuPtrOutput)
}

// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
func (o SBSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SBSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SBSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SBSku) string { return v.Name }).(pulumi.StringOutput)
}

// The billing tier of this particular SKU.
func (o SBSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SBSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SBSkuPtrOutput struct{ *pulumi.OutputState }

func (SBSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSku)(nil)).Elem()
}

func (o SBSkuPtrOutput) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return o
}

func (o SBSkuPtrOutput) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return o
}

func (o SBSkuPtrOutput) Elem() SBSkuOutput {
	return o.ApplyT(func(v *SBSku) SBSku { return *v }).(SBSkuOutput)
}

// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
func (o SBSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SBSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SBSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The billing tier of this particular SKU.
func (o SBSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

// SKU of the namespace.
type SBSkuResponse struct {
	// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity *int `pulumi:"capacity"`
	// Name of this SKU.
	Name string `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier *string `pulumi:"tier"`
}

// SBSkuResponseInput is an input type that accepts SBSkuResponseArgs and SBSkuResponseOutput values.
// You can construct a concrete instance of `SBSkuResponseInput` via:
//
//          SBSkuResponseArgs{...}
type SBSkuResponseInput interface {
	pulumi.Input

	ToSBSkuResponseOutput() SBSkuResponseOutput
	ToSBSkuResponseOutputWithContext(context.Context) SBSkuResponseOutput
}

// SKU of the namespace.
type SBSkuResponseArgs struct {
	// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// Name of this SKU.
	Name pulumi.StringInput `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SBSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSkuResponse)(nil)).Elem()
}

func (i SBSkuResponseArgs) ToSBSkuResponseOutput() SBSkuResponseOutput {
	return i.ToSBSkuResponseOutputWithContext(context.Background())
}

func (i SBSkuResponseArgs) ToSBSkuResponseOutputWithContext(ctx context.Context) SBSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuResponseOutput)
}

func (i SBSkuResponseArgs) ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput {
	return i.ToSBSkuResponsePtrOutputWithContext(context.Background())
}

func (i SBSkuResponseArgs) ToSBSkuResponsePtrOutputWithContext(ctx context.Context) SBSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuResponseOutput).ToSBSkuResponsePtrOutputWithContext(ctx)
}

// SBSkuResponsePtrInput is an input type that accepts SBSkuResponseArgs, SBSkuResponsePtr and SBSkuResponsePtrOutput values.
// You can construct a concrete instance of `SBSkuResponsePtrInput` via:
//
//          SBSkuResponseArgs{...}
//
//  or:
//
//          nil
type SBSkuResponsePtrInput interface {
	pulumi.Input

	ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput
	ToSBSkuResponsePtrOutputWithContext(context.Context) SBSkuResponsePtrOutput
}

type sbskuResponsePtrType SBSkuResponseArgs

func SBSkuResponsePtr(v *SBSkuResponseArgs) SBSkuResponsePtrInput {
	return (*sbskuResponsePtrType)(v)
}

func (*sbskuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSkuResponse)(nil)).Elem()
}

func (i *sbskuResponsePtrType) ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput {
	return i.ToSBSkuResponsePtrOutputWithContext(context.Background())
}

func (i *sbskuResponsePtrType) ToSBSkuResponsePtrOutputWithContext(ctx context.Context) SBSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuResponsePtrOutput)
}

// SKU of the namespace.
type SBSkuResponseOutput struct{ *pulumi.OutputState }

func (SBSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSkuResponse)(nil)).Elem()
}

func (o SBSkuResponseOutput) ToSBSkuResponseOutput() SBSkuResponseOutput {
	return o
}

func (o SBSkuResponseOutput) ToSBSkuResponseOutputWithContext(ctx context.Context) SBSkuResponseOutput {
	return o
}

func (o SBSkuResponseOutput) ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput {
	return o.ToSBSkuResponsePtrOutputWithContext(context.Background())
}

func (o SBSkuResponseOutput) ToSBSkuResponsePtrOutputWithContext(ctx context.Context) SBSkuResponsePtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *SBSkuResponse {
		return &v
	}).(SBSkuResponsePtrOutput)
}

// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
func (o SBSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SBSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SBSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The billing tier of this particular SKU.
func (o SBSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SBSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SBSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSkuResponse)(nil)).Elem()
}

func (o SBSkuResponsePtrOutput) ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput {
	return o
}

func (o SBSkuResponsePtrOutput) ToSBSkuResponsePtrOutputWithContext(ctx context.Context) SBSkuResponsePtrOutput {
	return o
}

func (o SBSkuResponsePtrOutput) Elem() SBSkuResponseOutput {
	return o.ApplyT(func(v *SBSkuResponse) SBSkuResponse { return *v }).(SBSkuResponseOutput)
}

// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
func (o SBSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SBSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The billing tier of this particular SKU.
func (o SBSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SBSkuOutput{})
	pulumi.RegisterOutputType(SBSkuPtrOutput{})
	pulumi.RegisterOutputType(SBSkuResponseOutput{})
	pulumi.RegisterOutputType(SBSkuResponsePtrOutput{})
}
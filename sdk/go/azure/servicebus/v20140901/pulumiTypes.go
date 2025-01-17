// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Message Count Details.
type MessageCountDetailsResponse struct {
	// Number of active messages in the queue, topic, or subscription.
	ActiveMessageCount float64 `pulumi:"activeMessageCount"`
	// Number of messages that are dead lettered.
	DeadLetterMessageCount float64 `pulumi:"deadLetterMessageCount"`
	// Number of scheduled messages.
	ScheduledMessageCount float64 `pulumi:"scheduledMessageCount"`
	// Number of messages transferred into dead letters.
	TransferDeadLetterMessageCount float64 `pulumi:"transferDeadLetterMessageCount"`
	// Number of messages transferred to another queue, topic, or subscription.
	TransferMessageCount float64 `pulumi:"transferMessageCount"`
}

// MessageCountDetailsResponseInput is an input type that accepts MessageCountDetailsResponseArgs and MessageCountDetailsResponseOutput values.
// You can construct a concrete instance of `MessageCountDetailsResponseInput` via:
//
//          MessageCountDetailsResponseArgs{...}
type MessageCountDetailsResponseInput interface {
	pulumi.Input

	ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput
	ToMessageCountDetailsResponseOutputWithContext(context.Context) MessageCountDetailsResponseOutput
}

// Message Count Details.
type MessageCountDetailsResponseArgs struct {
	// Number of active messages in the queue, topic, or subscription.
	ActiveMessageCount pulumi.Float64Input `pulumi:"activeMessageCount"`
	// Number of messages that are dead lettered.
	DeadLetterMessageCount pulumi.Float64Input `pulumi:"deadLetterMessageCount"`
	// Number of scheduled messages.
	ScheduledMessageCount pulumi.Float64Input `pulumi:"scheduledMessageCount"`
	// Number of messages transferred into dead letters.
	TransferDeadLetterMessageCount pulumi.Float64Input `pulumi:"transferDeadLetterMessageCount"`
	// Number of messages transferred to another queue, topic, or subscription.
	TransferMessageCount pulumi.Float64Input `pulumi:"transferMessageCount"`
}

func (MessageCountDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageCountDetailsResponse)(nil)).Elem()
}

func (i MessageCountDetailsResponseArgs) ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput {
	return i.ToMessageCountDetailsResponseOutputWithContext(context.Background())
}

func (i MessageCountDetailsResponseArgs) ToMessageCountDetailsResponseOutputWithContext(ctx context.Context) MessageCountDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageCountDetailsResponseOutput)
}

func (i MessageCountDetailsResponseArgs) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return i.ToMessageCountDetailsResponsePtrOutputWithContext(context.Background())
}

func (i MessageCountDetailsResponseArgs) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageCountDetailsResponseOutput).ToMessageCountDetailsResponsePtrOutputWithContext(ctx)
}

// MessageCountDetailsResponsePtrInput is an input type that accepts MessageCountDetailsResponseArgs, MessageCountDetailsResponsePtr and MessageCountDetailsResponsePtrOutput values.
// You can construct a concrete instance of `MessageCountDetailsResponsePtrInput` via:
//
//          MessageCountDetailsResponseArgs{...}
//
//  or:
//
//          nil
type MessageCountDetailsResponsePtrInput interface {
	pulumi.Input

	ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput
	ToMessageCountDetailsResponsePtrOutputWithContext(context.Context) MessageCountDetailsResponsePtrOutput
}

type messageCountDetailsResponsePtrType MessageCountDetailsResponseArgs

func MessageCountDetailsResponsePtr(v *MessageCountDetailsResponseArgs) MessageCountDetailsResponsePtrInput {
	return (*messageCountDetailsResponsePtrType)(v)
}

func (*messageCountDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageCountDetailsResponse)(nil)).Elem()
}

func (i *messageCountDetailsResponsePtrType) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return i.ToMessageCountDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *messageCountDetailsResponsePtrType) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageCountDetailsResponsePtrOutput)
}

// Message Count Details.
type MessageCountDetailsResponseOutput struct{ *pulumi.OutputState }

func (MessageCountDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageCountDetailsResponse)(nil)).Elem()
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutputWithContext(ctx context.Context) MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return o.ToMessageCountDetailsResponsePtrOutputWithContext(context.Background())
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return o.ApplyT(func(v MessageCountDetailsResponse) *MessageCountDetailsResponse {
		return &v
	}).(MessageCountDetailsResponsePtrOutput)
}

// Number of active messages in the queue, topic, or subscription.
func (o MessageCountDetailsResponseOutput) ActiveMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.ActiveMessageCount }).(pulumi.Float64Output)
}

// Number of messages that are dead lettered.
func (o MessageCountDetailsResponseOutput) DeadLetterMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.DeadLetterMessageCount }).(pulumi.Float64Output)
}

// Number of scheduled messages.
func (o MessageCountDetailsResponseOutput) ScheduledMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.ScheduledMessageCount }).(pulumi.Float64Output)
}

// Number of messages transferred into dead letters.
func (o MessageCountDetailsResponseOutput) TransferDeadLetterMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.TransferDeadLetterMessageCount }).(pulumi.Float64Output)
}

// Number of messages transferred to another queue, topic, or subscription.
func (o MessageCountDetailsResponseOutput) TransferMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.TransferMessageCount }).(pulumi.Float64Output)
}

type MessageCountDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (MessageCountDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageCountDetailsResponse)(nil)).Elem()
}

func (o MessageCountDetailsResponsePtrOutput) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return o
}

func (o MessageCountDetailsResponsePtrOutput) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return o
}

func (o MessageCountDetailsResponsePtrOutput) Elem() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) MessageCountDetailsResponse { return *v }).(MessageCountDetailsResponseOutput)
}

// Number of active messages in the queue, topic, or subscription.
func (o MessageCountDetailsResponsePtrOutput) ActiveMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ActiveMessageCount
	}).(pulumi.Float64PtrOutput)
}

// Number of messages that are dead lettered.
func (o MessageCountDetailsResponsePtrOutput) DeadLetterMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.DeadLetterMessageCount
	}).(pulumi.Float64PtrOutput)
}

// Number of scheduled messages.
func (o MessageCountDetailsResponsePtrOutput) ScheduledMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ScheduledMessageCount
	}).(pulumi.Float64PtrOutput)
}

// Number of messages transferred into dead letters.
func (o MessageCountDetailsResponsePtrOutput) TransferDeadLetterMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransferDeadLetterMessageCount
	}).(pulumi.Float64PtrOutput)
}

// Number of messages transferred to another queue, topic, or subscription.
func (o MessageCountDetailsResponsePtrOutput) TransferMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransferMessageCount
	}).(pulumi.Float64PtrOutput)
}

// SKU of the namespace.
type Sku struct {
	// The specified messaging units for the tier.
	Capacity *int `pulumi:"capacity"`
	// Name of this SKU.
	Name *string `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier string `pulumi:"tier"`
}

// SkuInput is an input type that accepts SkuArgs and SkuOutput values.
// You can construct a concrete instance of `SkuInput` via:
//
//          SkuArgs{...}
type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

// SKU of the namespace.
type SkuArgs struct {
	// The specified messaging units for the tier.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// Name of this SKU.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}

// SkuPtrInput is an input type that accepts SkuArgs, SkuPtr and SkuPtrOutput values.
// You can construct a concrete instance of `SkuPtrInput` via:
//
//          SkuArgs{...}
//
//  or:
//
//          nil
type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

// SKU of the namespace.
type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyT(func(v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

// The specified messaging units for the tier.
func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The billing tier of this particular SKU.
func (o SkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku { return *v }).(SkuOutput)
}

// The specified messaging units for the tier.
func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The billing tier of this particular SKU.
func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

// SKU of the namespace.
type SkuResponse struct {
	// The specified messaging units for the tier.
	Capacity *int `pulumi:"capacity"`
	// Name of this SKU.
	Name *string `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier string `pulumi:"tier"`
}

// SkuResponseInput is an input type that accepts SkuResponseArgs and SkuResponseOutput values.
// You can construct a concrete instance of `SkuResponseInput` via:
//
//          SkuResponseArgs{...}
type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

// SKU of the namespace.
type SkuResponseArgs struct {
	// The specified messaging units for the tier.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// Name of this SKU.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}

// SkuResponsePtrInput is an input type that accepts SkuResponseArgs, SkuResponsePtr and SkuResponsePtrOutput values.
// You can construct a concrete instance of `SkuResponsePtrInput` via:
//
//          SkuResponseArgs{...}
//
//  or:
//
//          nil
type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

// SKU of the namespace.
type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyT(func(v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

// The specified messaging units for the tier.
func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The billing tier of this particular SKU.
func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse { return *v }).(SkuResponseOutput)
}

// The specified messaging units for the tier.
func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The billing tier of this particular SKU.
func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MessageCountDetailsResponseOutput{})
	pulumi.RegisterOutputType(MessageCountDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}

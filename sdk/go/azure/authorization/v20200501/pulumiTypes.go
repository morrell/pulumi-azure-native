// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceManagementPrivateLinkEndpointConnectionsResponse struct {
	// The private endpoint connections.
	PrivateEndpointConnections []string `pulumi:"privateEndpointConnections"`
}

// ResourceManagementPrivateLinkEndpointConnectionsResponseInput is an input type that accepts ResourceManagementPrivateLinkEndpointConnectionsResponseArgs and ResourceManagementPrivateLinkEndpointConnectionsResponseOutput values.
// You can construct a concrete instance of `ResourceManagementPrivateLinkEndpointConnectionsResponseInput` via:
//
//          ResourceManagementPrivateLinkEndpointConnectionsResponseArgs{...}
type ResourceManagementPrivateLinkEndpointConnectionsResponseInput interface {
	pulumi.Input

	ToResourceManagementPrivateLinkEndpointConnectionsResponseOutput() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput
	ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponseOutput
}

type ResourceManagementPrivateLinkEndpointConnectionsResponseArgs struct {
	// The private endpoint connections.
	PrivateEndpointConnections pulumi.StringArrayInput `pulumi:"privateEndpointConnections"`
}

func (ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (i ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutput() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return i.ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(context.Background())
}

func (i ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput)
}

func (i ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return i.ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(context.Background())
}

func (i ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput).ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx)
}

// ResourceManagementPrivateLinkEndpointConnectionsResponsePtrInput is an input type that accepts ResourceManagementPrivateLinkEndpointConnectionsResponseArgs, ResourceManagementPrivateLinkEndpointConnectionsResponsePtr and ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput values.
// You can construct a concrete instance of `ResourceManagementPrivateLinkEndpointConnectionsResponsePtrInput` via:
//
//          ResourceManagementPrivateLinkEndpointConnectionsResponseArgs{...}
//
//  or:
//
//          nil
type ResourceManagementPrivateLinkEndpointConnectionsResponsePtrInput interface {
	pulumi.Input

	ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput
	ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput
}

type resourceManagementPrivateLinkEndpointConnectionsResponsePtrType ResourceManagementPrivateLinkEndpointConnectionsResponseArgs

func ResourceManagementPrivateLinkEndpointConnectionsResponsePtr(v *ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrInput {
	return (*resourceManagementPrivateLinkEndpointConnectionsResponsePtrType)(v)
}

func (*resourceManagementPrivateLinkEndpointConnectionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (i *resourceManagementPrivateLinkEndpointConnectionsResponsePtrType) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return i.ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(context.Background())
}

func (i *resourceManagementPrivateLinkEndpointConnectionsResponsePtrType) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput)
}

type ResourceManagementPrivateLinkEndpointConnectionsResponseOutput struct{ *pulumi.OutputState }

func (ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutput() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return o.ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(context.Background())
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return o.ApplyT(func(v ResourceManagementPrivateLinkEndpointConnectionsResponse) *ResourceManagementPrivateLinkEndpointConnectionsResponse {
		return &v
	}).(ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput)
}

// The private endpoint connections.
func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) PrivateEndpointConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceManagementPrivateLinkEndpointConnectionsResponse) []string {
		return v.PrivateEndpointConnections
	}).(pulumi.StringArrayOutput)
}

type ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) Elem() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o.ApplyT(func(v *ResourceManagementPrivateLinkEndpointConnectionsResponse) ResourceManagementPrivateLinkEndpointConnectionsResponse {
		return *v
	}).(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput)
}

// The private endpoint connections.
func (o ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) PrivateEndpointConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceManagementPrivateLinkEndpointConnectionsResponse) []string {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput{})
	pulumi.RegisterOutputType(ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Cluster node details.
type ClusterNodeResponse struct {
	// Number of physical cores on the cluster node.
	CoreCount float64 `pulumi:"coreCount"`
	// Id of the node in the cluster.
	Id float64 `pulumi:"id"`
	// Manufacturer of the cluster node hardware.
	Manufacturer string `pulumi:"manufacturer"`
	// Total available memory on the cluster node (in GiB).
	MemoryInGiB float64 `pulumi:"memoryInGiB"`
	// Model name of the cluster node hardware.
	Model string `pulumi:"model"`
	// Name of the cluster node.
	Name string `pulumi:"name"`
	// Operating system running on the cluster node.
	OsName string `pulumi:"osName"`
	// Version of the operating system running on the cluster node.
	OsVersion string `pulumi:"osVersion"`
	// Immutable id of the cluster node.
	SerialNumber string `pulumi:"serialNumber"`
}

// ClusterNodeResponseInput is an input type that accepts ClusterNodeResponseArgs and ClusterNodeResponseOutput values.
// You can construct a concrete instance of `ClusterNodeResponseInput` via:
//
//          ClusterNodeResponseArgs{...}
type ClusterNodeResponseInput interface {
	pulumi.Input

	ToClusterNodeResponseOutput() ClusterNodeResponseOutput
	ToClusterNodeResponseOutputWithContext(context.Context) ClusterNodeResponseOutput
}

// Cluster node details.
type ClusterNodeResponseArgs struct {
	// Number of physical cores on the cluster node.
	CoreCount pulumi.Float64Input `pulumi:"coreCount"`
	// Id of the node in the cluster.
	Id pulumi.Float64Input `pulumi:"id"`
	// Manufacturer of the cluster node hardware.
	Manufacturer pulumi.StringInput `pulumi:"manufacturer"`
	// Total available memory on the cluster node (in GiB).
	MemoryInGiB pulumi.Float64Input `pulumi:"memoryInGiB"`
	// Model name of the cluster node hardware.
	Model pulumi.StringInput `pulumi:"model"`
	// Name of the cluster node.
	Name pulumi.StringInput `pulumi:"name"`
	// Operating system running on the cluster node.
	OsName pulumi.StringInput `pulumi:"osName"`
	// Version of the operating system running on the cluster node.
	OsVersion pulumi.StringInput `pulumi:"osVersion"`
	// Immutable id of the cluster node.
	SerialNumber pulumi.StringInput `pulumi:"serialNumber"`
}

func (ClusterNodeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNodeResponse)(nil)).Elem()
}

func (i ClusterNodeResponseArgs) ToClusterNodeResponseOutput() ClusterNodeResponseOutput {
	return i.ToClusterNodeResponseOutputWithContext(context.Background())
}

func (i ClusterNodeResponseArgs) ToClusterNodeResponseOutputWithContext(ctx context.Context) ClusterNodeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterNodeResponseOutput)
}

// ClusterNodeResponseArrayInput is an input type that accepts ClusterNodeResponseArray and ClusterNodeResponseArrayOutput values.
// You can construct a concrete instance of `ClusterNodeResponseArrayInput` via:
//
//          ClusterNodeResponseArray{ ClusterNodeResponseArgs{...} }
type ClusterNodeResponseArrayInput interface {
	pulumi.Input

	ToClusterNodeResponseArrayOutput() ClusterNodeResponseArrayOutput
	ToClusterNodeResponseArrayOutputWithContext(context.Context) ClusterNodeResponseArrayOutput
}

type ClusterNodeResponseArray []ClusterNodeResponseInput

func (ClusterNodeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterNodeResponse)(nil)).Elem()
}

func (i ClusterNodeResponseArray) ToClusterNodeResponseArrayOutput() ClusterNodeResponseArrayOutput {
	return i.ToClusterNodeResponseArrayOutputWithContext(context.Background())
}

func (i ClusterNodeResponseArray) ToClusterNodeResponseArrayOutputWithContext(ctx context.Context) ClusterNodeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterNodeResponseArrayOutput)
}

// Cluster node details.
type ClusterNodeResponseOutput struct{ *pulumi.OutputState }

func (ClusterNodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNodeResponse)(nil)).Elem()
}

func (o ClusterNodeResponseOutput) ToClusterNodeResponseOutput() ClusterNodeResponseOutput {
	return o
}

func (o ClusterNodeResponseOutput) ToClusterNodeResponseOutputWithContext(ctx context.Context) ClusterNodeResponseOutput {
	return o
}

// Number of physical cores on the cluster node.
func (o ClusterNodeResponseOutput) CoreCount() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.CoreCount }).(pulumi.Float64Output)
}

// Id of the node in the cluster.
func (o ClusterNodeResponseOutput) Id() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.Id }).(pulumi.Float64Output)
}

// Manufacturer of the cluster node hardware.
func (o ClusterNodeResponseOutput) Manufacturer() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Manufacturer }).(pulumi.StringOutput)
}

// Total available memory on the cluster node (in GiB).
func (o ClusterNodeResponseOutput) MemoryInGiB() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.MemoryInGiB }).(pulumi.Float64Output)
}

// Model name of the cluster node hardware.
func (o ClusterNodeResponseOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Model }).(pulumi.StringOutput)
}

// Name of the cluster node.
func (o ClusterNodeResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Operating system running on the cluster node.
func (o ClusterNodeResponseOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.OsName }).(pulumi.StringOutput)
}

// Version of the operating system running on the cluster node.
func (o ClusterNodeResponseOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.OsVersion }).(pulumi.StringOutput)
}

// Immutable id of the cluster node.
func (o ClusterNodeResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

type ClusterNodeResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterNodeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterNodeResponse)(nil)).Elem()
}

func (o ClusterNodeResponseArrayOutput) ToClusterNodeResponseArrayOutput() ClusterNodeResponseArrayOutput {
	return o
}

func (o ClusterNodeResponseArrayOutput) ToClusterNodeResponseArrayOutputWithContext(ctx context.Context) ClusterNodeResponseArrayOutput {
	return o
}

func (o ClusterNodeResponseArrayOutput) Index(i pulumi.IntInput) ClusterNodeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterNodeResponse {
		return vs[0].([]ClusterNodeResponse)[vs[1].(int)]
	}).(ClusterNodeResponseOutput)
}

// Properties reported by cluster agent.
type ClusterReportedPropertiesResponse struct {
	// Unique id generated by the on-prem cluster.
	ClusterId string `pulumi:"clusterId"`
	// Name of the on-prem cluster connected to this resource.
	ClusterName string `pulumi:"clusterName"`
	// Version of the cluster software.
	ClusterVersion string `pulumi:"clusterVersion"`
	// Last time the cluster reported the data.
	LastUpdated string `pulumi:"lastUpdated"`
	// List of nodes reported by the cluster.
	Nodes []ClusterNodeResponse `pulumi:"nodes"`
}

// ClusterReportedPropertiesResponseInput is an input type that accepts ClusterReportedPropertiesResponseArgs and ClusterReportedPropertiesResponseOutput values.
// You can construct a concrete instance of `ClusterReportedPropertiesResponseInput` via:
//
//          ClusterReportedPropertiesResponseArgs{...}
type ClusterReportedPropertiesResponseInput interface {
	pulumi.Input

	ToClusterReportedPropertiesResponseOutput() ClusterReportedPropertiesResponseOutput
	ToClusterReportedPropertiesResponseOutputWithContext(context.Context) ClusterReportedPropertiesResponseOutput
}

// Properties reported by cluster agent.
type ClusterReportedPropertiesResponseArgs struct {
	// Unique id generated by the on-prem cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Name of the on-prem cluster connected to this resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Version of the cluster software.
	ClusterVersion pulumi.StringInput `pulumi:"clusterVersion"`
	// Last time the cluster reported the data.
	LastUpdated pulumi.StringInput `pulumi:"lastUpdated"`
	// List of nodes reported by the cluster.
	Nodes ClusterNodeResponseArrayInput `pulumi:"nodes"`
}

func (ClusterReportedPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (i ClusterReportedPropertiesResponseArgs) ToClusterReportedPropertiesResponseOutput() ClusterReportedPropertiesResponseOutput {
	return i.ToClusterReportedPropertiesResponseOutputWithContext(context.Background())
}

func (i ClusterReportedPropertiesResponseArgs) ToClusterReportedPropertiesResponseOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterReportedPropertiesResponseOutput)
}

func (i ClusterReportedPropertiesResponseArgs) ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput {
	return i.ToClusterReportedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ClusterReportedPropertiesResponseArgs) ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterReportedPropertiesResponseOutput).ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx)
}

// ClusterReportedPropertiesResponsePtrInput is an input type that accepts ClusterReportedPropertiesResponseArgs, ClusterReportedPropertiesResponsePtr and ClusterReportedPropertiesResponsePtrOutput values.
// You can construct a concrete instance of `ClusterReportedPropertiesResponsePtrInput` via:
//
//          ClusterReportedPropertiesResponseArgs{...}
//
//  or:
//
//          nil
type ClusterReportedPropertiesResponsePtrInput interface {
	pulumi.Input

	ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput
	ToClusterReportedPropertiesResponsePtrOutputWithContext(context.Context) ClusterReportedPropertiesResponsePtrOutput
}

type clusterReportedPropertiesResponsePtrType ClusterReportedPropertiesResponseArgs

func ClusterReportedPropertiesResponsePtr(v *ClusterReportedPropertiesResponseArgs) ClusterReportedPropertiesResponsePtrInput {
	return (*clusterReportedPropertiesResponsePtrType)(v)
}

func (*clusterReportedPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (i *clusterReportedPropertiesResponsePtrType) ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput {
	return i.ToClusterReportedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *clusterReportedPropertiesResponsePtrType) ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterReportedPropertiesResponsePtrOutput)
}

// Properties reported by cluster agent.
type ClusterReportedPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterReportedPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponseOutput() ClusterReportedPropertiesResponseOutput {
	return o
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponseOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponseOutput {
	return o
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput {
	return o.ToClusterReportedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) *ClusterReportedPropertiesResponse {
		return &v
	}).(ClusterReportedPropertiesResponsePtrOutput)
}

// Unique id generated by the on-prem cluster.
func (o ClusterReportedPropertiesResponseOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Name of the on-prem cluster connected to this resource.
func (o ClusterReportedPropertiesResponseOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterName }).(pulumi.StringOutput)
}

// Version of the cluster software.
func (o ClusterReportedPropertiesResponseOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterVersion }).(pulumi.StringOutput)
}

// Last time the cluster reported the data.
func (o ClusterReportedPropertiesResponseOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.LastUpdated }).(pulumi.StringOutput)
}

// List of nodes reported by the cluster.
func (o ClusterReportedPropertiesResponseOutput) Nodes() ClusterNodeResponseArrayOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) []ClusterNodeResponse { return v.Nodes }).(ClusterNodeResponseArrayOutput)
}

type ClusterReportedPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterReportedPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (o ClusterReportedPropertiesResponsePtrOutput) ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput {
	return o
}

func (o ClusterReportedPropertiesResponsePtrOutput) ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponsePtrOutput {
	return o
}

func (o ClusterReportedPropertiesResponsePtrOutput) Elem() ClusterReportedPropertiesResponseOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) ClusterReportedPropertiesResponse { return *v }).(ClusterReportedPropertiesResponseOutput)
}

// Unique id generated by the on-prem cluster.
func (o ClusterReportedPropertiesResponsePtrOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterId
	}).(pulumi.StringPtrOutput)
}

// Name of the on-prem cluster connected to this resource.
func (o ClusterReportedPropertiesResponsePtrOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterName
	}).(pulumi.StringPtrOutput)
}

// Version of the cluster software.
func (o ClusterReportedPropertiesResponsePtrOutput) ClusterVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterVersion
	}).(pulumi.StringPtrOutput)
}

// Last time the cluster reported the data.
func (o ClusterReportedPropertiesResponsePtrOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdated
	}).(pulumi.StringPtrOutput)
}

// List of nodes reported by the cluster.
func (o ClusterReportedPropertiesResponsePtrOutput) Nodes() ClusterNodeResponseArrayOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) []ClusterNodeResponse {
		if v == nil {
			return nil
		}
		return v.Nodes
	}).(ClusterNodeResponseArrayOutput)
}

// Status of Arc Extension for a particular node in HCI Cluster.
type PerNodeExtensionStateResponse struct {
	// Fully qualified resource ID for the particular Arc Extension on this node.
	Extension string `pulumi:"extension"`
	// Name of the node in HCI Cluster.
	Name string `pulumi:"name"`
	// State of Arc Extension in this node.
	State string `pulumi:"state"`
}

// PerNodeExtensionStateResponseInput is an input type that accepts PerNodeExtensionStateResponseArgs and PerNodeExtensionStateResponseOutput values.
// You can construct a concrete instance of `PerNodeExtensionStateResponseInput` via:
//
//          PerNodeExtensionStateResponseArgs{...}
type PerNodeExtensionStateResponseInput interface {
	pulumi.Input

	ToPerNodeExtensionStateResponseOutput() PerNodeExtensionStateResponseOutput
	ToPerNodeExtensionStateResponseOutputWithContext(context.Context) PerNodeExtensionStateResponseOutput
}

// Status of Arc Extension for a particular node in HCI Cluster.
type PerNodeExtensionStateResponseArgs struct {
	// Fully qualified resource ID for the particular Arc Extension on this node.
	Extension pulumi.StringInput `pulumi:"extension"`
	// Name of the node in HCI Cluster.
	Name pulumi.StringInput `pulumi:"name"`
	// State of Arc Extension in this node.
	State pulumi.StringInput `pulumi:"state"`
}

func (PerNodeExtensionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeExtensionStateResponse)(nil)).Elem()
}

func (i PerNodeExtensionStateResponseArgs) ToPerNodeExtensionStateResponseOutput() PerNodeExtensionStateResponseOutput {
	return i.ToPerNodeExtensionStateResponseOutputWithContext(context.Background())
}

func (i PerNodeExtensionStateResponseArgs) ToPerNodeExtensionStateResponseOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerNodeExtensionStateResponseOutput)
}

// PerNodeExtensionStateResponseArrayInput is an input type that accepts PerNodeExtensionStateResponseArray and PerNodeExtensionStateResponseArrayOutput values.
// You can construct a concrete instance of `PerNodeExtensionStateResponseArrayInput` via:
//
//          PerNodeExtensionStateResponseArray{ PerNodeExtensionStateResponseArgs{...} }
type PerNodeExtensionStateResponseArrayInput interface {
	pulumi.Input

	ToPerNodeExtensionStateResponseArrayOutput() PerNodeExtensionStateResponseArrayOutput
	ToPerNodeExtensionStateResponseArrayOutputWithContext(context.Context) PerNodeExtensionStateResponseArrayOutput
}

type PerNodeExtensionStateResponseArray []PerNodeExtensionStateResponseInput

func (PerNodeExtensionStateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeExtensionStateResponse)(nil)).Elem()
}

func (i PerNodeExtensionStateResponseArray) ToPerNodeExtensionStateResponseArrayOutput() PerNodeExtensionStateResponseArrayOutput {
	return i.ToPerNodeExtensionStateResponseArrayOutputWithContext(context.Background())
}

func (i PerNodeExtensionStateResponseArray) ToPerNodeExtensionStateResponseArrayOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerNodeExtensionStateResponseArrayOutput)
}

// Status of Arc Extension for a particular node in HCI Cluster.
type PerNodeExtensionStateResponseOutput struct{ *pulumi.OutputState }

func (PerNodeExtensionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeExtensionStateResponse)(nil)).Elem()
}

func (o PerNodeExtensionStateResponseOutput) ToPerNodeExtensionStateResponseOutput() PerNodeExtensionStateResponseOutput {
	return o
}

func (o PerNodeExtensionStateResponseOutput) ToPerNodeExtensionStateResponseOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseOutput {
	return o
}

// Fully qualified resource ID for the particular Arc Extension on this node.
func (o PerNodeExtensionStateResponseOutput) Extension() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.Extension }).(pulumi.StringOutput)
}

// Name of the node in HCI Cluster.
func (o PerNodeExtensionStateResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.Name }).(pulumi.StringOutput)
}

// State of Arc Extension in this node.
func (o PerNodeExtensionStateResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.State }).(pulumi.StringOutput)
}

type PerNodeExtensionStateResponseArrayOutput struct{ *pulumi.OutputState }

func (PerNodeExtensionStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeExtensionStateResponse)(nil)).Elem()
}

func (o PerNodeExtensionStateResponseArrayOutput) ToPerNodeExtensionStateResponseArrayOutput() PerNodeExtensionStateResponseArrayOutput {
	return o
}

func (o PerNodeExtensionStateResponseArrayOutput) ToPerNodeExtensionStateResponseArrayOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseArrayOutput {
	return o
}

func (o PerNodeExtensionStateResponseArrayOutput) Index(i pulumi.IntInput) PerNodeExtensionStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerNodeExtensionStateResponse {
		return vs[0].([]PerNodeExtensionStateResponse)[vs[1].(int)]
	}).(PerNodeExtensionStateResponseOutput)
}

// Status of Arc agent for a particular node in HCI Cluster.
type PerNodeStateResponse struct {
	// Fully qualified resource ID for the Arc agent of this node.
	ArcInstance string `pulumi:"arcInstance"`
	// Name of the Node in HCI Cluster
	Name string `pulumi:"name"`
	// State of Arc agent in this node.
	State string `pulumi:"state"`
}

// PerNodeStateResponseInput is an input type that accepts PerNodeStateResponseArgs and PerNodeStateResponseOutput values.
// You can construct a concrete instance of `PerNodeStateResponseInput` via:
//
//          PerNodeStateResponseArgs{...}
type PerNodeStateResponseInput interface {
	pulumi.Input

	ToPerNodeStateResponseOutput() PerNodeStateResponseOutput
	ToPerNodeStateResponseOutputWithContext(context.Context) PerNodeStateResponseOutput
}

// Status of Arc agent for a particular node in HCI Cluster.
type PerNodeStateResponseArgs struct {
	// Fully qualified resource ID for the Arc agent of this node.
	ArcInstance pulumi.StringInput `pulumi:"arcInstance"`
	// Name of the Node in HCI Cluster
	Name pulumi.StringInput `pulumi:"name"`
	// State of Arc agent in this node.
	State pulumi.StringInput `pulumi:"state"`
}

func (PerNodeStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeStateResponse)(nil)).Elem()
}

func (i PerNodeStateResponseArgs) ToPerNodeStateResponseOutput() PerNodeStateResponseOutput {
	return i.ToPerNodeStateResponseOutputWithContext(context.Background())
}

func (i PerNodeStateResponseArgs) ToPerNodeStateResponseOutputWithContext(ctx context.Context) PerNodeStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerNodeStateResponseOutput)
}

// PerNodeStateResponseArrayInput is an input type that accepts PerNodeStateResponseArray and PerNodeStateResponseArrayOutput values.
// You can construct a concrete instance of `PerNodeStateResponseArrayInput` via:
//
//          PerNodeStateResponseArray{ PerNodeStateResponseArgs{...} }
type PerNodeStateResponseArrayInput interface {
	pulumi.Input

	ToPerNodeStateResponseArrayOutput() PerNodeStateResponseArrayOutput
	ToPerNodeStateResponseArrayOutputWithContext(context.Context) PerNodeStateResponseArrayOutput
}

type PerNodeStateResponseArray []PerNodeStateResponseInput

func (PerNodeStateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeStateResponse)(nil)).Elem()
}

func (i PerNodeStateResponseArray) ToPerNodeStateResponseArrayOutput() PerNodeStateResponseArrayOutput {
	return i.ToPerNodeStateResponseArrayOutputWithContext(context.Background())
}

func (i PerNodeStateResponseArray) ToPerNodeStateResponseArrayOutputWithContext(ctx context.Context) PerNodeStateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerNodeStateResponseArrayOutput)
}

// Status of Arc agent for a particular node in HCI Cluster.
type PerNodeStateResponseOutput struct{ *pulumi.OutputState }

func (PerNodeStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeStateResponse)(nil)).Elem()
}

func (o PerNodeStateResponseOutput) ToPerNodeStateResponseOutput() PerNodeStateResponseOutput {
	return o
}

func (o PerNodeStateResponseOutput) ToPerNodeStateResponseOutputWithContext(ctx context.Context) PerNodeStateResponseOutput {
	return o
}

// Fully qualified resource ID for the Arc agent of this node.
func (o PerNodeStateResponseOutput) ArcInstance() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.ArcInstance }).(pulumi.StringOutput)
}

// Name of the Node in HCI Cluster
func (o PerNodeStateResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.Name }).(pulumi.StringOutput)
}

// State of Arc agent in this node.
func (o PerNodeStateResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.State }).(pulumi.StringOutput)
}

type PerNodeStateResponseArrayOutput struct{ *pulumi.OutputState }

func (PerNodeStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeStateResponse)(nil)).Elem()
}

func (o PerNodeStateResponseArrayOutput) ToPerNodeStateResponseArrayOutput() PerNodeStateResponseArrayOutput {
	return o
}

func (o PerNodeStateResponseArrayOutput) ToPerNodeStateResponseArrayOutputWithContext(ctx context.Context) PerNodeStateResponseArrayOutput {
	return o
}

func (o PerNodeStateResponseArrayOutput) Index(i pulumi.IntInput) PerNodeStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerNodeStateResponse {
		return vs[0].([]PerNodeStateResponse)[vs[1].(int)]
	}).(PerNodeStateResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterNodeResponseOutput{})
	pulumi.RegisterOutputType(ClusterNodeResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterReportedPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterReportedPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PerNodeExtensionStateResponseOutput{})
	pulumi.RegisterOutputType(PerNodeExtensionStateResponseArrayOutput{})
	pulumi.RegisterOutputType(PerNodeStateResponseOutput{})
	pulumi.RegisterOutputType(PerNodeStateResponseArrayOutput{})
}

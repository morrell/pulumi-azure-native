// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// DSCP Configuration in a resource group.
type DscpConfiguration struct {
	pulumi.CustomResourceState

	// Associated Network Interfaces to the DSCP Configuration.
	AssociatedNetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"associatedNetworkInterfaces"`
	// Destination IP ranges.
	DestinationIpRanges QosIpRangeResponseArrayOutput `pulumi:"destinationIpRanges"`
	// Destination port ranges.
	DestinationPortRanges QosPortRangeResponseArrayOutput `pulumi:"destinationPortRanges"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// List of markings to be used in the configuration.
	Markings pulumi.IntArrayOutput `pulumi:"markings"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// RNM supported protocol types.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The provisioning state of the DSCP Configuration resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Qos Collection ID generated by RNM.
	QosCollectionId pulumi.StringOutput `pulumi:"qosCollectionId"`
	// The resource GUID property of the DSCP Configuration resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Source IP ranges.
	SourceIpRanges QosIpRangeResponseArrayOutput `pulumi:"sourceIpRanges"`
	// Sources port ranges.
	SourcePortRanges QosPortRangeResponseArrayOutput `pulumi:"sourcePortRanges"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDscpConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDscpConfiguration(ctx *pulumi.Context,
	name string, args *DscpConfigurationArgs, opts ...pulumi.ResourceOption) (*DscpConfiguration, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DscpConfigurationArgs{}
	}
	var resource DscpConfiguration
	err := ctx.RegisterResource("azurerm:network/v20200601:DscpConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDscpConfiguration gets an existing DscpConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDscpConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscpConfigurationState, opts ...pulumi.ResourceOption) (*DscpConfiguration, error) {
	var resource DscpConfiguration
	err := ctx.ReadResource("azurerm:network/v20200601:DscpConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DscpConfiguration resources.
type dscpConfigurationState struct {
	// Associated Network Interfaces to the DSCP Configuration.
	AssociatedNetworkInterfaces []NetworkInterfaceResponse `pulumi:"associatedNetworkInterfaces"`
	// Destination IP ranges.
	DestinationIpRanges []QosIpRangeResponse `pulumi:"destinationIpRanges"`
	// Destination port ranges.
	DestinationPortRanges []QosPortRangeResponse `pulumi:"destinationPortRanges"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// List of markings to be used in the configuration.
	Markings []int `pulumi:"markings"`
	// Resource name.
	Name *string `pulumi:"name"`
	// RNM supported protocol types.
	Protocol *string `pulumi:"protocol"`
	// The provisioning state of the DSCP Configuration resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Qos Collection ID generated by RNM.
	QosCollectionId *string `pulumi:"qosCollectionId"`
	// The resource GUID property of the DSCP Configuration resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Source IP ranges.
	SourceIpRanges []QosIpRangeResponse `pulumi:"sourceIpRanges"`
	// Sources port ranges.
	SourcePortRanges []QosPortRangeResponse `pulumi:"sourcePortRanges"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type DscpConfigurationState struct {
	// Associated Network Interfaces to the DSCP Configuration.
	AssociatedNetworkInterfaces NetworkInterfaceResponseArrayInput
	// Destination IP ranges.
	DestinationIpRanges QosIpRangeResponseArrayInput
	// Destination port ranges.
	DestinationPortRanges QosPortRangeResponseArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// List of markings to be used in the configuration.
	Markings pulumi.IntArrayInput
	// Resource name.
	Name pulumi.StringPtrInput
	// RNM supported protocol types.
	Protocol pulumi.StringPtrInput
	// The provisioning state of the DSCP Configuration resource.
	ProvisioningState pulumi.StringPtrInput
	// Qos Collection ID generated by RNM.
	QosCollectionId pulumi.StringPtrInput
	// The resource GUID property of the DSCP Configuration resource.
	ResourceGuid pulumi.StringPtrInput
	// Source IP ranges.
	SourceIpRanges QosIpRangeResponseArrayInput
	// Sources port ranges.
	SourcePortRanges QosPortRangeResponseArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (DscpConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscpConfigurationState)(nil)).Elem()
}

type dscpConfigurationArgs struct {
	// Destination IP ranges.
	DestinationIpRanges []QosIpRange `pulumi:"destinationIpRanges"`
	// Destination port ranges.
	DestinationPortRanges []QosPortRange `pulumi:"destinationPortRanges"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// List of markings to be used in the configuration.
	Markings []int `pulumi:"markings"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// RNM supported protocol types.
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Source IP ranges.
	SourceIpRanges []QosIpRange `pulumi:"sourceIpRanges"`
	// Sources port ranges.
	SourcePortRanges []QosPortRange `pulumi:"sourcePortRanges"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DscpConfiguration resource.
type DscpConfigurationArgs struct {
	// Destination IP ranges.
	DestinationIpRanges QosIpRangeArrayInput
	// Destination port ranges.
	DestinationPortRanges QosPortRangeArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// List of markings to be used in the configuration.
	Markings pulumi.IntArrayInput
	// The name of the resource.
	Name pulumi.StringInput
	// RNM supported protocol types.
	Protocol pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Source IP ranges.
	SourceIpRanges QosIpRangeArrayInput
	// Sources port ranges.
	SourcePortRanges QosPortRangeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DscpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscpConfigurationArgs)(nil)).Elem()
}
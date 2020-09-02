// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a node type in the cluster, each node type represents sub set of nodes in the cluster.
type NodeType struct {
	pulumi.CustomResourceState

	// The range of ports from which cluster assigned port to Service Fabric applications.
	ApplicationPorts EndpointRangeDescriptionResponsePtrOutput `pulumi:"applicationPorts"`
	// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
	Capacities pulumi.StringMapOutput `pulumi:"capacities"`
	// Disk size for each vm in the node type in GBs.
	DataDiskSizeGB pulumi.IntOutput `pulumi:"dataDiskSizeGB"`
	// The range of ephemeral ports that nodes in this node type should be configured with.
	EphemeralPorts EndpointRangeDescriptionResponsePtrOutput `pulumi:"ephemeralPorts"`
	// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
	IsPrimary pulumi.BoolOutput `pulumi:"isPrimary"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
	PlacementProperties pulumi.StringMapOutput `pulumi:"placementProperties"`
	// The provisioning state of the managed cluster resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Set of extensions that should be installed onto the virtual machines.
	VmExtensions VMSSExtensionResponseArrayOutput `pulumi:"vmExtensions"`
	// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
	VmImageOffer pulumi.StringPtrOutput `pulumi:"vmImageOffer"`
	// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
	VmImagePublisher pulumi.StringPtrOutput `pulumi:"vmImagePublisher"`
	// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
	VmImageSku pulumi.StringPtrOutput `pulumi:"vmImageSku"`
	// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
	VmImageVersion pulumi.StringPtrOutput `pulumi:"vmImageVersion"`
	// The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
	VmInstanceCount pulumi.IntOutput `pulumi:"vmInstanceCount"`
	// The secrets to install in the virtual machines.
	VmSecrets VaultSecretGroupResponseArrayOutput `pulumi:"vmSecrets"`
	// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
	VmSize pulumi.StringPtrOutput `pulumi:"vmSize"`
}

// NewNodeType registers a new resource with the given unique name, arguments, and options.
func NewNodeType(ctx *pulumi.Context,
	name string, args *NodeTypeArgs, opts ...pulumi.ResourceOption) (*NodeType, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.DataDiskSizeGB == nil {
		return nil, errors.New("missing required argument 'DataDiskSizeGB'")
	}
	if args == nil || args.IsPrimary == nil {
		return nil, errors.New("missing required argument 'IsPrimary'")
	}
	if args == nil || args.NodeTypeName == nil {
		return nil, errors.New("missing required argument 'NodeTypeName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VmInstanceCount == nil {
		return nil, errors.New("missing required argument 'VmInstanceCount'")
	}
	if args == nil {
		args = &NodeTypeArgs{}
	}
	var resource NodeType
	err := ctx.RegisterResource("azurerm:servicefabric/v20200101preview:NodeType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeType gets an existing NodeType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTypeState, opts ...pulumi.ResourceOption) (*NodeType, error) {
	var resource NodeType
	err := ctx.ReadResource("azurerm:servicefabric/v20200101preview:NodeType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeType resources.
type nodeTypeState struct {
	// The range of ports from which cluster assigned port to Service Fabric applications.
	ApplicationPorts *EndpointRangeDescriptionResponse `pulumi:"applicationPorts"`
	// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
	Capacities map[string]string `pulumi:"capacities"`
	// Disk size for each vm in the node type in GBs.
	DataDiskSizeGB *int `pulumi:"dataDiskSizeGB"`
	// The range of ephemeral ports that nodes in this node type should be configured with.
	EphemeralPorts *EndpointRangeDescriptionResponse `pulumi:"ephemeralPorts"`
	// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
	IsPrimary *bool `pulumi:"isPrimary"`
	// Azure resource name.
	Name *string `pulumi:"name"`
	// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
	PlacementProperties map[string]string `pulumi:"placementProperties"`
	// The provisioning state of the managed cluster resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type *string `pulumi:"type"`
	// Set of extensions that should be installed onto the virtual machines.
	VmExtensions []VMSSExtensionResponse `pulumi:"vmExtensions"`
	// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
	VmImageOffer *string `pulumi:"vmImageOffer"`
	// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
	VmImagePublisher *string `pulumi:"vmImagePublisher"`
	// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
	VmImageSku *string `pulumi:"vmImageSku"`
	// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
	VmImageVersion *string `pulumi:"vmImageVersion"`
	// The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
	VmInstanceCount *int `pulumi:"vmInstanceCount"`
	// The secrets to install in the virtual machines.
	VmSecrets []VaultSecretGroupResponse `pulumi:"vmSecrets"`
	// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
	VmSize *string `pulumi:"vmSize"`
}

type NodeTypeState struct {
	// The range of ports from which cluster assigned port to Service Fabric applications.
	ApplicationPorts EndpointRangeDescriptionResponsePtrInput
	// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
	Capacities pulumi.StringMapInput
	// Disk size for each vm in the node type in GBs.
	DataDiskSizeGB pulumi.IntPtrInput
	// The range of ephemeral ports that nodes in this node type should be configured with.
	EphemeralPorts EndpointRangeDescriptionResponsePtrInput
	// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
	IsPrimary pulumi.BoolPtrInput
	// Azure resource name.
	Name pulumi.StringPtrInput
	// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
	PlacementProperties pulumi.StringMapInput
	// The provisioning state of the managed cluster resource.
	ProvisioningState pulumi.StringPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// Azure resource type.
	Type pulumi.StringPtrInput
	// Set of extensions that should be installed onto the virtual machines.
	VmExtensions VMSSExtensionResponseArrayInput
	// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
	VmImageOffer pulumi.StringPtrInput
	// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
	VmImagePublisher pulumi.StringPtrInput
	// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
	VmImageSku pulumi.StringPtrInput
	// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
	VmImageVersion pulumi.StringPtrInput
	// The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
	VmInstanceCount pulumi.IntPtrInput
	// The secrets to install in the virtual machines.
	VmSecrets VaultSecretGroupResponseArrayInput
	// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
	VmSize pulumi.StringPtrInput
}

func (NodeTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTypeState)(nil)).Elem()
}

type nodeTypeArgs struct {
	// The range of ports from which cluster assigned port to Service Fabric applications.
	ApplicationPorts *EndpointRangeDescription `pulumi:"applicationPorts"`
	// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
	Capacities map[string]string `pulumi:"capacities"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Disk size for each vm in the node type in GBs.
	DataDiskSizeGB int `pulumi:"dataDiskSizeGB"`
	// The range of ephemeral ports that nodes in this node type should be configured with.
	EphemeralPorts *EndpointRangeDescription `pulumi:"ephemeralPorts"`
	// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
	IsPrimary bool `pulumi:"isPrimary"`
	// The name of the node type.
	NodeTypeName string `pulumi:"nodeTypeName"`
	// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
	PlacementProperties map[string]string `pulumi:"placementProperties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Set of extensions that should be installed onto the virtual machines.
	VmExtensions []VMSSExtension `pulumi:"vmExtensions"`
	// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
	VmImageOffer *string `pulumi:"vmImageOffer"`
	// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
	VmImagePublisher *string `pulumi:"vmImagePublisher"`
	// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
	VmImageSku *string `pulumi:"vmImageSku"`
	// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
	VmImageVersion *string `pulumi:"vmImageVersion"`
	// The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
	VmInstanceCount int `pulumi:"vmInstanceCount"`
	// The secrets to install in the virtual machines.
	VmSecrets []VaultSecretGroup `pulumi:"vmSecrets"`
	// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
	VmSize *string `pulumi:"vmSize"`
}

// The set of arguments for constructing a NodeType resource.
type NodeTypeArgs struct {
	// The range of ports from which cluster assigned port to Service Fabric applications.
	ApplicationPorts EndpointRangeDescriptionPtrInput
	// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
	Capacities pulumi.StringMapInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Disk size for each vm in the node type in GBs.
	DataDiskSizeGB pulumi.IntInput
	// The range of ephemeral ports that nodes in this node type should be configured with.
	EphemeralPorts EndpointRangeDescriptionPtrInput
	// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
	IsPrimary pulumi.BoolInput
	// The name of the node type.
	NodeTypeName pulumi.StringInput
	// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
	PlacementProperties pulumi.StringMapInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// Set of extensions that should be installed onto the virtual machines.
	VmExtensions VMSSExtensionArrayInput
	// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
	VmImageOffer pulumi.StringPtrInput
	// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
	VmImagePublisher pulumi.StringPtrInput
	// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
	VmImageSku pulumi.StringPtrInput
	// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
	VmImageVersion pulumi.StringPtrInput
	// The number of nodes in the node type. This count should match the capacity property in the corresponding VirtualMachineScaleSet resource.
	VmInstanceCount pulumi.IntInput
	// The secrets to install in the virtual machines.
	VmSecrets VaultSecretGroupArrayInput
	// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
	VmSize pulumi.StringPtrInput
}

func (NodeTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTypeArgs)(nil)).Elem()
}
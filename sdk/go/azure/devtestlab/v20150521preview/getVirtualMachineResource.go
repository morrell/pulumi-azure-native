// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A virtual machine.
func LookupVirtualMachineResource(ctx *pulumi.Context, args *LookupVirtualMachineResourceArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResourceResult, error) {
	var rv LookupVirtualMachineResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getVirtualMachineResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineResourceArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the virtual Machine.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A virtual machine.
type LookupVirtualMachineResourceResult struct {
	// The artifact deployment status for the virtual machine.
	ArtifactDeploymentStatus *ArtifactDeploymentStatusPropertiesResponse `pulumi:"artifactDeploymentStatus"`
	// The artifacts to be installed on the virtual machine.
	Artifacts []ArtifactInstallPropertiesResponse `pulumi:"artifacts"`
	// The resource identifier (Microsoft.Compute) of the virtual machine.
	ComputeId *string `pulumi:"computeId"`
	// The email address of creator of the virtual machine.
	CreatedByUser *string `pulumi:"createdByUser"`
	// The object identifier of the creator of the virtual machine.
	CreatedByUserId *string `pulumi:"createdByUserId"`
	// The custom image identifier of the virtual machine.
	CustomImageId *string `pulumi:"customImageId"`
	// Indicates whether the virtual machine is to be created without a public IP address.
	DisallowPublicIpAddress *bool `pulumi:"disallowPublicIpAddress"`
	// The fully-qualified domain name of the virtual machine.
	Fqdn *string `pulumi:"fqdn"`
	// The Microsoft Azure Marketplace image reference of the virtual machine.
	GalleryImageReference *GalleryImageReferenceResponse `pulumi:"galleryImageReference"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// A value indicating whether this virtual machine uses an SSH key for authentication.
	IsAuthenticationWithSshKey *bool `pulumi:"isAuthenticationWithSshKey"`
	// The lab subnet name of the virtual machine.
	LabSubnetName *string `pulumi:"labSubnetName"`
	// The lab virtual network identifier of the virtual machine.
	LabVirtualNetworkId *string `pulumi:"labVirtualNetworkId"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The notes of the virtual machine.
	Notes *string `pulumi:"notes"`
	// The OS type of the virtual machine.
	OsType *string `pulumi:"osType"`
	// The object identifier of the owner of the virtual machine.
	OwnerObjectId *string `pulumi:"ownerObjectId"`
	// The password of the virtual machine administrator.
	Password *string `pulumi:"password"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The size of the virtual machine.
	Size *string `pulumi:"size"`
	// The SSH key of the virtual machine administrator.
	SshKey *string `pulumi:"sshKey"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The user name of the virtual machine.
	UserName *string `pulumi:"userName"`
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupFileServer(ctx *pulumi.Context, args *LookupFileServerArgs, opts ...pulumi.InvokeOption) (*LookupFileServerResult, error) {
	var rv LookupFileServerResult
	err := ctx.Invoke("azurerm:batchai/v20180301:getFileServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileServerArgs struct {
	// The name of the file server within the specified resource group. File server names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Contains information about the File Server.
type LookupFileServerResult struct {
	CreationTime string `pulumi:"creationTime"`
	// Settings for the data disk which would be created for the File Server.
	DataDisks *DataDisksResponse `pulumi:"dataDisks"`
	// The location of the resource
	Location string `pulumi:"location"`
	// Details of the File Server.
	MountSettings MountSettingsResponse `pulumi:"mountSettings"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Possible values: creating - The File Server is getting created. updating - The File Server creation has been accepted and it is getting updated. deleting - The user has requested that the File Server be deleted, and it is in the process of being deleted. failed - The File Server creation has failed with the specified errorCode. Details about the error code are specified in the message field. succeeded - The File Server creation has succeeded.
	ProvisioningState               string `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime string `pulumi:"provisioningStateTransitionTime"`
	// SSH configuration settings for the VM
	SshConfiguration *SshConfigurationResponse `pulumi:"sshConfiguration"`
	// Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
	Subnet *ResourceIdResponse `pulumi:"subnet"`
	// The tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource
	Type string `pulumi:"type"`
	// For information about available VM sizes for File Server from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
	VmSize *string `pulumi:"vmSize"`
}
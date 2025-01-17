// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A lab.
func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabArgs struct {
	// Specify the $expand query. Example: 'properties($select=defaultStorageAccount)'
	Expand *string `pulumi:"expand"`
	// The name of the lab.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A lab.
type LookupLabResult struct {
	// The properties of any lab announcement associated with this lab
	Announcement *LabAnnouncementPropertiesResponse `pulumi:"announcement"`
	// The lab's artifact storage account.
	ArtifactsStorageAccount string `pulumi:"artifactsStorageAccount"`
	// The creation date of the lab.
	CreatedDate string `pulumi:"createdDate"`
	// The lab's default premium storage account.
	DefaultPremiumStorageAccount string `pulumi:"defaultPremiumStorageAccount"`
	// The lab's default storage account.
	DefaultStorageAccount string `pulumi:"defaultStorageAccount"`
	// The access rights to be granted to the user when provisioning an environment
	EnvironmentPermission *string `pulumi:"environmentPermission"`
	// Extended properties of the lab used for experimental features
	ExtendedProperties map[string]string `pulumi:"extendedProperties"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// Type of storage used by the lab. It can be either Premium or Standard. Default is Premium.
	LabStorageType *string `pulumi:"labStorageType"`
	// The load balancer used to for lab VMs that use shared IP address.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The ordered list of artifact resource IDs that should be applied on all Linux VM creations by default, prior to the artifacts specified by the user.
	MandatoryArtifactsResourceIdsLinux []string `pulumi:"mandatoryArtifactsResourceIdsLinux"`
	// The ordered list of artifact resource IDs that should be applied on all Windows VM creations by default, prior to the artifacts specified by the user.
	MandatoryArtifactsResourceIdsWindows []string `pulumi:"mandatoryArtifactsResourceIdsWindows"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The Network Security Group attached to the lab VMs Network interfaces to restrict open ports.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
	// The lab's premium data disk storage account.
	PremiumDataDiskStorageAccount string `pulumi:"premiumDataDiskStorageAccount"`
	// The setting to enable usage of premium data disks.
	// When its value is 'Enabled', creation of standard or premium data disks is allowed.
	// When its value is 'Disabled', only creation of standard data disks is allowed.
	PremiumDataDisks *string `pulumi:"premiumDataDisks"`
	// The provisioning status of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The public IP address for the lab's load balancer.
	PublicIpId string `pulumi:"publicIpId"`
	// The properties of any lab support message associated with this lab
	Support *LabSupportPropertiesResponse `pulumi:"support"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier string `pulumi:"uniqueIdentifier"`
	// The lab's Key vault.
	VaultName string `pulumi:"vaultName"`
	// The resource group in which all new lab virtual machines will be created. To let DevTest Labs manage resource group creation, set this value to null.
	VmCreationResourceGroup string `pulumi:"vmCreationResourceGroup"`
}

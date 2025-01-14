// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200207preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SAP monitor info on Azure (ARM properties and SAP monitor properties)
func LookupSapMonitor(ctx *pulumi.Context, args *LookupSapMonitorArgs, opts ...pulumi.InvokeOption) (*LookupSapMonitorResult, error) {
	var rv LookupSapMonitorResult
	err := ctx.Invoke("azure-native:hanaonazure/v20200207preview:getSapMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSapMonitorArgs struct {
	// Name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SAP monitor resource.
	SapMonitorName string `pulumi:"sapMonitorName"`
}

// SAP monitor info on Azure (ARM properties and SAP monitor properties)
type LookupSapMonitorResult struct {
	// The value indicating whether to send analytics to Microsoft
	EnableCustomerAnalytics *bool `pulumi:"enableCustomerAnalytics"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The ARM ID of the Log Analytics Workspace that is used for monitoring
	LogAnalyticsWorkspaceArmId *string `pulumi:"logAnalyticsWorkspaceArmId"`
	// The workspace ID of the log analytics workspace to be used for monitoring
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The shared key of the log analytics workspace that is used for monitoring
	LogAnalyticsWorkspaceSharedKey *string `pulumi:"logAnalyticsWorkspaceSharedKey"`
	// The name of the resource group the SAP Monitor resources get deployed into.
	ManagedResourceGroupName string `pulumi:"managedResourceGroupName"`
	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet *string `pulumi:"monitorSubnet"`
	// The name of the resource
	Name string `pulumi:"name"`
	// State of provisioning of the HanaInstance
	ProvisioningState string `pulumi:"provisioningState"`
	// The version of the payload running in the Collector VM
	SapMonitorCollectorVersion string `pulumi:"sapMonitorCollectorVersion"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

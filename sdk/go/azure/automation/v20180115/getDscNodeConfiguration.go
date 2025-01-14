// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180115

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the dsc node configuration.
func LookupDscNodeConfiguration(ctx *pulumi.Context, args *LookupDscNodeConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDscNodeConfigurationResult, error) {
	var rv LookupDscNodeConfigurationResult
	err := ctx.Invoke("azure-native:automation/v20180115:getDscNodeConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDscNodeConfigurationArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The Dsc node configuration name.
	NodeConfigurationName string `pulumi:"nodeConfigurationName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the dsc node configuration.
type LookupDscNodeConfigurationResult struct {
	// Gets or sets the configuration of the node.
	Configuration *DscConfigurationAssociationPropertyResponse `pulumi:"configuration"`
	// Gets or sets creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// If a new build version of NodeConfiguration is required.
	IncrementNodeConfigurationBuild *bool `pulumi:"incrementNodeConfigurationBuild"`
	// Gets or sets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Number of nodes with this node configuration assigned
	NodeCount *float64 `pulumi:"nodeCount"`
	// Source of node configuration.
	Source *string `pulumi:"source"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

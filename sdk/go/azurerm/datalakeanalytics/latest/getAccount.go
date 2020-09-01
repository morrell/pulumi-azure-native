// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azurerm:datalakeanalytics/latest:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName string `pulumi:"accountName"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics account.
type LookupAccountResult struct {
	// The unique identifier associated with this Data Lake Analytics account.
	AccountId string `pulumi:"accountId"`
	// The list of compute policies associated with this account.
	ComputePolicies []ComputePolicyResponse `pulumi:"computePolicies"`
	// The account creation time.
	CreationTime string `pulumi:"creationTime"`
	// The commitment tier in use for the current month.
	CurrentTier string `pulumi:"currentTier"`
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts []DataLakeStoreAccountInformationResponse `pulumi:"dataLakeStoreAccounts"`
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount string `pulumi:"defaultDataLakeStoreAccount"`
	// The full CName endpoint for this account.
	Endpoint string `pulumi:"endpoint"`
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps string `pulumi:"firewallAllowAzureIps"`
	// The list of firewall rules associated with this account.
	FirewallRules []FirewallRuleResponse `pulumi:"firewallRules"`
	// The current state of the IP address firewall for this account.
	FirewallState string `pulumi:"firewallState"`
	// The account last modified time.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// The resource location.
	Location string `pulumi:"location"`
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism int `pulumi:"maxDegreeOfParallelism"`
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob int `pulumi:"maxDegreeOfParallelismPerJob"`
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount int `pulumi:"maxJobCount"`
	// The minimum supported priority per job for this account.
	MinPriorityPerJob int `pulumi:"minPriorityPerJob"`
	// The resource name.
	Name string `pulumi:"name"`
	// The commitment tier for the next month.
	NewTier string `pulumi:"newTier"`
	// The provisioning status of the Data Lake Analytics account.
	ProvisioningState string `pulumi:"provisioningState"`
	// The number of days that job metadata is retained.
	QueryStoreRetention int `pulumi:"queryStoreRetention"`
	// The state of the Data Lake Analytics account.
	State string `pulumi:"state"`
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts []StorageAccountInformationResponse `pulumi:"storageAccounts"`
	// The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account.
	SystemMaxDegreeOfParallelism int `pulumi:"systemMaxDegreeOfParallelism"`
	// The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account.
	SystemMaxJobCount int `pulumi:"systemMaxJobCount"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}
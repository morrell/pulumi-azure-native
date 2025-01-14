// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A server security alert policy.
func LookupServerSecurityAlertPolicy(ctx *pulumi.Context, args *LookupServerSecurityAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServerSecurityAlertPolicyResult, error) {
	var rv LookupServerSecurityAlertPolicyResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20171201preview:getServerSecurityAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerSecurityAlertPolicyArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security alert policy.
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A server security alert policy.
type LookupServerSecurityAlertPolicyResult struct {
	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly
	DisabledAlerts []string `pulumi:"disabledAlerts"`
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *bool `pulumi:"emailAccountAdmins"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the state of the policy, whether it is enabled or disabled.
	State string `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

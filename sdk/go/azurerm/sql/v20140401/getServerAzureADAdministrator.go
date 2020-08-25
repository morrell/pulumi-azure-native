// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupServerAzureADAdministrator(ctx *pulumi.Context, args *LookupServerAzureADAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupServerAzureADAdministratorResult, error) {
	var rv LookupServerAzureADAdministratorResult
	err := ctx.Invoke("azurerm:sql/v20140401:getServerAzureADAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerAzureADAdministratorArgs struct {
	// Name of the server administrator resource.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// An server Active Directory Administrator.
type LookupServerAzureADAdministratorResult struct {
	// The type of administrator.
	AdministratorType string `pulumi:"administratorType"`
	// The server administrator login value.
	Login string `pulumi:"login"`
	// Resource name.
	Name string `pulumi:"name"`
	// The server administrator Sid (Secure ID).
	Sid string `pulumi:"sid"`
	// The server Active Directory Administrator tenant id.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}
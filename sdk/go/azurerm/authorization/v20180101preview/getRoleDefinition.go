// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRoleDefinition(ctx *pulumi.Context, args *LookupRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupRoleDefinitionResult, error) {
	var rv LookupRoleDefinitionResult
	err := ctx.Invoke("azurerm:authorization/v20180101preview:getRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleDefinitionArgs struct {
	// The ID of the role definition.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	// The scope of the role definition.
	Scope string `pulumi:"scope"`
}

// Role definition.
type LookupRoleDefinitionResult struct {
	// Role definition assignable scopes.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// The role definition description.
	Description *string `pulumi:"description"`
	// The role definition name.
	Name string `pulumi:"name"`
	// Role definition permissions.
	Permissions []PermissionResponse `pulumi:"permissions"`
	// The role name.
	RoleName *string `pulumi:"roleName"`
	// The role type.
	RoleType *string `pulumi:"roleType"`
	// The role definition type.
	Type string `pulumi:"type"`
}
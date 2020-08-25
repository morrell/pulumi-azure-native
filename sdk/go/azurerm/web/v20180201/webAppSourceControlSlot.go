// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Source control configuration for an app.
type WebAppSourceControlSlot struct {
	pulumi.CustomResourceState

	// Name of branch to use for deployment.
	Branch pulumi.StringPtrOutput `pulumi:"branch"`
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled pulumi.BoolPtrOutput `pulumi:"deploymentRollbackEnabled"`
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration pulumi.BoolPtrOutput `pulumi:"isManualIntegration"`
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial pulumi.BoolPtrOutput `pulumi:"isMercurial"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Repository or source control URL.
	RepoUrl pulumi.StringPtrOutput `pulumi:"repoUrl"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppSourceControlSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppSourceControlSlot(ctx *pulumi.Context,
	name string, args *WebAppSourceControlSlotArgs, opts ...pulumi.ResourceOption) (*WebAppSourceControlSlot, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &WebAppSourceControlSlotArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:web/v20150801:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azurerm:web/v20160801:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azurerm:web/v20181101:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azurerm:web/v20190801:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azurerm:web/v20200601:WebAppSourceControlSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSourceControlSlot
	err := ctx.RegisterResource("azurerm:web/v20180201:WebAppSourceControlSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSourceControlSlot gets an existing WebAppSourceControlSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSourceControlSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSourceControlSlotState, opts ...pulumi.ResourceOption) (*WebAppSourceControlSlot, error) {
	var resource WebAppSourceControlSlot
	err := ctx.ReadResource("azurerm:web/v20180201:WebAppSourceControlSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSourceControlSlot resources.
type webAppSourceControlSlotState struct {
	// Name of branch to use for deployment.
	Branch *string `pulumi:"branch"`
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled *bool `pulumi:"deploymentRollbackEnabled"`
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration *bool `pulumi:"isManualIntegration"`
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial *bool `pulumi:"isMercurial"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Repository or source control URL.
	RepoUrl *string `pulumi:"repoUrl"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppSourceControlSlotState struct {
	// Name of branch to use for deployment.
	Branch pulumi.StringPtrInput
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled pulumi.BoolPtrInput
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration pulumi.BoolPtrInput
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Repository or source control URL.
	RepoUrl pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppSourceControlSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSourceControlSlotState)(nil)).Elem()
}

type webAppSourceControlSlotArgs struct {
	// Name of branch to use for deployment.
	Branch *string `pulumi:"branch"`
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled *bool `pulumi:"deploymentRollbackEnabled"`
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration *bool `pulumi:"isManualIntegration"`
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial *bool `pulumi:"isMercurial"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the deployment slot. If a slot is not specified, the API will update the source control configuration for the production slot.
	Name string `pulumi:"name"`
	// Repository or source control URL.
	RepoUrl *string `pulumi:"repoUrl"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppSourceControlSlot resource.
type WebAppSourceControlSlotArgs struct {
	// Name of branch to use for deployment.
	Branch pulumi.StringPtrInput
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled pulumi.BoolPtrInput
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration pulumi.BoolPtrInput
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the deployment slot. If a slot is not specified, the API will update the source control configuration for the production slot.
	Name pulumi.StringInput
	// Repository or source control URL.
	RepoUrl pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppSourceControlSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSourceControlSlotArgs)(nil)).Elem()
}
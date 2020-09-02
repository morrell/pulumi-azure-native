// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170515preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the source control.
type SourceControl struct {
	pulumi.CustomResourceState

	// The auto sync of the source control. Default is false.
	AutoSync pulumi.BoolPtrOutput `pulumi:"autoSync"`
	// The repo branch of the source control. Include branch as empty string for VsoTfvc.
	Branch pulumi.StringPtrOutput `pulumi:"branch"`
	// The creation time.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// The description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder path of the source control.
	FolderPath pulumi.StringPtrOutput `pulumi:"folderPath"`
	// The last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The auto publish of the source control. Default is true.
	PublishRunbook pulumi.BoolPtrOutput `pulumi:"publishRunbook"`
	// The repo url of the source control.
	RepoUrl pulumi.StringPtrOutput `pulumi:"repoUrl"`
	// The source type. Must be one of VsoGit, VsoTfvc, GitHub.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSourceControl registers a new resource with the given unique name, arguments, and options.
func NewSourceControl(ctx *pulumi.Context,
	name string, args *SourceControlArgs, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SourceControlName == nil {
		return nil, errors.New("missing required argument 'SourceControlName'")
	}
	if args == nil {
		args = &SourceControlArgs{}
	}
	var resource SourceControl
	err := ctx.RegisterResource("azurerm:automation/v20170515preview:SourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceControl gets an existing SourceControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceControlState, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	var resource SourceControl
	err := ctx.ReadResource("azurerm:automation/v20170515preview:SourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceControl resources.
type sourceControlState struct {
	// The auto sync of the source control. Default is false.
	AutoSync *bool `pulumi:"autoSync"`
	// The repo branch of the source control. Include branch as empty string for VsoTfvc.
	Branch *string `pulumi:"branch"`
	// The creation time.
	CreationTime *string `pulumi:"creationTime"`
	// The description.
	Description *string `pulumi:"description"`
	// The folder path of the source control.
	FolderPath *string `pulumi:"folderPath"`
	// The last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The auto publish of the source control. Default is true.
	PublishRunbook *bool `pulumi:"publishRunbook"`
	// The repo url of the source control.
	RepoUrl *string `pulumi:"repoUrl"`
	// The source type. Must be one of VsoGit, VsoTfvc, GitHub.
	SourceType *string `pulumi:"sourceType"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type SourceControlState struct {
	// The auto sync of the source control. Default is false.
	AutoSync pulumi.BoolPtrInput
	// The repo branch of the source control. Include branch as empty string for VsoTfvc.
	Branch pulumi.StringPtrInput
	// The creation time.
	CreationTime pulumi.StringPtrInput
	// The description.
	Description pulumi.StringPtrInput
	// The folder path of the source control.
	FolderPath pulumi.StringPtrInput
	// The last modified time.
	LastModifiedTime pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The auto publish of the source control. Default is true.
	PublishRunbook pulumi.BoolPtrInput
	// The repo url of the source control.
	RepoUrl pulumi.StringPtrInput
	// The source type. Must be one of VsoGit, VsoTfvc, GitHub.
	SourceType pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (SourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlState)(nil)).Elem()
}

type sourceControlArgs struct {
	// The auto async of the source control. Default is false.
	AutoSync *bool `pulumi:"autoSync"`
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The repo branch of the source control. Include branch as empty string for VsoTfvc.
	Branch *string `pulumi:"branch"`
	// The user description of the source control.
	Description *string `pulumi:"description"`
	// The folder path of the source control. Path must be relative.
	FolderPath *string `pulumi:"folderPath"`
	// The auto publish of the source control. Default is true.
	PublishRunbook *bool `pulumi:"publishRunbook"`
	// The repo url of the source control.
	RepoUrl *string `pulumi:"repoUrl"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The authorization token for the repo of the source control.
	SecurityToken *SourceControlSecurityTokenProperties `pulumi:"securityToken"`
	// The source control name.
	SourceControlName string `pulumi:"sourceControlName"`
	// The source type. Must be one of VsoGit, VsoTfvc, GitHub, case sensitive.
	SourceType *string `pulumi:"sourceType"`
}

// The set of arguments for constructing a SourceControl resource.
type SourceControlArgs struct {
	// The auto async of the source control. Default is false.
	AutoSync pulumi.BoolPtrInput
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The repo branch of the source control. Include branch as empty string for VsoTfvc.
	Branch pulumi.StringPtrInput
	// The user description of the source control.
	Description pulumi.StringPtrInput
	// The folder path of the source control. Path must be relative.
	FolderPath pulumi.StringPtrInput
	// The auto publish of the source control. Default is true.
	PublishRunbook pulumi.BoolPtrInput
	// The repo url of the source control.
	RepoUrl pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The authorization token for the repo of the source control.
	SecurityToken SourceControlSecurityTokenPropertiesPtrInput
	// The source control name.
	SourceControlName pulumi.StringInput
	// The source type. Must be one of VsoGit, VsoTfvc, GitHub, case sensitive.
	SourceType pulumi.StringPtrInput
}

func (SourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlArgs)(nil)).Elem()
}
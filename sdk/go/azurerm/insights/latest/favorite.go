// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Properties that define a favorite that is associated to an Application Insights component.
type Favorite struct {
	pulumi.CustomResourceState

	// Favorite category, as defined by the user at creation time.
	Category pulumi.StringPtrOutput `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config pulumi.StringPtrOutput `pulumi:"config"`
	// Internally assigned unique id of the favorite definition.
	FavoriteId pulumi.StringOutput `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType pulumi.StringPtrOutput `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate pulumi.BoolPtrOutput `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The source of the favorite definition.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this favorite definition.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Unique user id of the specific user that owns this favorite.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewFavorite registers a new resource with the given unique name, arguments, and options.
func NewFavorite(ctx *pulumi.Context,
	name string, args *FavoriteArgs, opts ...pulumi.ResourceOption) (*Favorite, error) {
	if args == nil || args.FavoriteId == nil {
		return nil, errors.New("missing required argument 'FavoriteId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &FavoriteArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:insights/v20150501:Favorite"),
		},
	})
	opts = append(opts, aliases)
	var resource Favorite
	err := ctx.RegisterResource("azurerm:insights/latest:Favorite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFavorite gets an existing Favorite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFavorite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FavoriteState, opts ...pulumi.ResourceOption) (*Favorite, error) {
	var resource Favorite
	err := ctx.ReadResource("azurerm:insights/latest:Favorite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Favorite resources.
type favoriteState struct {
	// Favorite category, as defined by the user at creation time.
	Category *string `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config *string `pulumi:"config"`
	// Internally assigned unique id of the favorite definition.
	FavoriteId *string `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType *string `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate *bool `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name *string `pulumi:"name"`
	// The source of the favorite definition.
	SourceType *string `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags []string `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this favorite definition.
	TimeModified *string `pulumi:"timeModified"`
	// Unique user id of the specific user that owns this favorite.
	UserId *string `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version *string `pulumi:"version"`
}

type FavoriteState struct {
	// Favorite category, as defined by the user at creation time.
	Category pulumi.StringPtrInput
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config pulumi.StringPtrInput
	// Internally assigned unique id of the favorite definition.
	FavoriteId pulumi.StringPtrInput
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType pulumi.StringPtrInput
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate pulumi.BoolPtrInput
	// The user-defined name of the favorite.
	Name pulumi.StringPtrInput
	// The source of the favorite definition.
	SourceType pulumi.StringPtrInput
	// A list of 0 or more tags that are associated with this favorite definition
	Tags pulumi.StringArrayInput
	// Date and time in UTC of the last modification that was made to this favorite definition.
	TimeModified pulumi.StringPtrInput
	// Unique user id of the specific user that owns this favorite.
	UserId pulumi.StringPtrInput
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version pulumi.StringPtrInput
}

func (FavoriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteState)(nil)).Elem()
}

type favoriteArgs struct {
	// Favorite category, as defined by the user at creation time.
	Category *string `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config *string `pulumi:"config"`
	// The Id of a specific favorite defined in the Application Insights component
	FavoriteId string `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType *string `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate *bool `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
	// The source of the favorite definition.
	SourceType *string `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags []string `pulumi:"tags"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Favorite resource.
type FavoriteArgs struct {
	// Favorite category, as defined by the user at creation time.
	Category pulumi.StringPtrInput
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config pulumi.StringPtrInput
	// The Id of a specific favorite defined in the Application Insights component
	FavoriteId pulumi.StringInput
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType pulumi.StringPtrInput
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate pulumi.BoolPtrInput
	// The user-defined name of the favorite.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput
	// The source of the favorite definition.
	SourceType pulumi.StringPtrInput
	// A list of 0 or more tags that are associated with this favorite definition
	Tags pulumi.StringArrayInput
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version pulumi.StringPtrInput
}

func (FavoriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteArgs)(nil)).Elem()
}
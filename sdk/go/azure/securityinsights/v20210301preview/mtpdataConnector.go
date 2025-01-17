// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents MTP (Microsoft Threat Protection) data connector.
type MTPDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes MTPDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatProtection'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMTPDataConnector registers a new resource with the given unique name, arguments, and options.
func NewMTPDataConnector(ctx *pulumi.Context,
	name string, args *MTPDataConnectorArgs, opts ...pulumi.ResourceOption) (*MTPDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("MicrosoftThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:MTPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MTPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:MTPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMTPDataConnector gets an existing MTPDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMTPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MTPDataConnectorState, opts ...pulumi.ResourceOption) (*MTPDataConnector, error) {
	var resource MTPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:MTPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MTPDataConnector resources.
type mtpdataConnectorState struct {
}

type MTPDataConnectorState struct {
}

func (MTPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mtpdataConnectorState)(nil)).Elem()
}

type mtpdataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes MTPDataConnectorDataTypes `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatProtection'.
	Kind string `pulumi:"kind"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MTPDataConnector resource.
type MTPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes MTPDataConnectorDataTypesInput
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatProtection'.
	Kind pulumi.StringInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MTPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mtpdataConnectorArgs)(nil)).Elem()
}

type MTPDataConnectorInput interface {
	pulumi.Input

	ToMTPDataConnectorOutput() MTPDataConnectorOutput
	ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput
}

func (*MTPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnector)(nil))
}

func (i *MTPDataConnector) ToMTPDataConnectorOutput() MTPDataConnectorOutput {
	return i.ToMTPDataConnectorOutputWithContext(context.Background())
}

func (i *MTPDataConnector) ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorOutput)
}

type MTPDataConnectorOutput struct {
	*pulumi.OutputState
}

func (MTPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnector)(nil))
}

func (o MTPDataConnectorOutput) ToMTPDataConnectorOutput() MTPDataConnectorOutput {
	return o
}

func (o MTPDataConnectorOutput) ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MTPDataConnectorOutput{})
}

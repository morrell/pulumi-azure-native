// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EventGrid Domain.
type Domain struct {
	pulumi.CustomResourceState

	// This Boolean is used to specify the creation mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
	// In this context, creation of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
	// When this property is null or set to true, Event Grid is responsible of automatically creating the domain topic when the first event subscription is
	// created at the scope of the domain topic. If this property is set to false, then creating the first event subscription will require creating a domain topic
	// by the user. The self-management mode can be used if the user wants full control of when the domain topic is created, while auto-managed mode provides the
	// flexibility to perform less operations and manage fewer resources by the user. Also, note that in auto-managed creation mode, user is allowed to create the
	// domain topic on demand if needed.
	AutoCreateTopicWithFirstSubscription pulumi.BoolPtrOutput `pulumi:"autoCreateTopicWithFirstSubscription"`
	// This Boolean is used to specify the deletion mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
	// In this context, deletion of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
	// When this property is set to true, Event Grid is responsible of automatically deleting the domain topic when the last event subscription at the scope
	// of the domain topic is deleted. If this property is set to false, then the user needs to manually delete the domain topic when it is no longer needed
	// (e.g., when last event subscription is deleted and the resource needs to be cleaned up). The self-management mode can be used if the user wants full
	// control of when the domain topic needs to be deleted, while auto-managed mode provides the flexibility to perform less operations and manage fewer
	// resources by the user.
	AutoDeleteTopicWithLastSubscription pulumi.BoolPtrOutput `pulumi:"autoDeleteTopicWithLastSubscription"`
	// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the domain.
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// Endpoint for the Event Grid Domain Resource which is used for publishing the events.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Identity information for the Event Grid Domain resource.
	Identity IdentityInfoResponsePtrOutput `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleResponseArrayOutput `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the Event Grid Domain Resource.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrOutput `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Metric resource id for the Event Grid Domain Resource.
	MetricResourceId pulumi.StringOutput `pulumi:"metricResourceId"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the Event Grid Domain Resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The Sku pricing tier for the Event Grid Domain resource.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// The system metadata relating to the Event Grid Domain resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AutoCreateTopicWithFirstSubscription == nil {
		args.AutoCreateTopicWithFirstSubscription = pulumi.BoolPtr(true)
	}
	if args.AutoDeleteTopicWithLastSubscription == nil {
		args.AutoDeleteTopicWithLastSubscription = pulumi.BoolPtr(true)
	}
	if args.DisableLocalAuth == nil {
		args.DisableLocalAuth = pulumi.BoolPtr(false)
	}
	if args.InputSchema == nil {
		args.InputSchema = pulumi.StringPtr("EventGridSchema")
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20210601preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20180915preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190201preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190601:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200101preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200401preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200601:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20201015preview:Domain"),
		},
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azure-native:eventgrid/v20210601preview:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:eventgrid/v20210601preview:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// This Boolean is used to specify the creation mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
	// In this context, creation of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
	// When this property is null or set to true, Event Grid is responsible of automatically creating the domain topic when the first event subscription is
	// created at the scope of the domain topic. If this property is set to false, then creating the first event subscription will require creating a domain topic
	// by the user. The self-management mode can be used if the user wants full control of when the domain topic is created, while auto-managed mode provides the
	// flexibility to perform less operations and manage fewer resources by the user. Also, note that in auto-managed creation mode, user is allowed to create the
	// domain topic on demand if needed.
	AutoCreateTopicWithFirstSubscription *bool `pulumi:"autoCreateTopicWithFirstSubscription"`
	// This Boolean is used to specify the deletion mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
	// In this context, deletion of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
	// When this property is set to true, Event Grid is responsible of automatically deleting the domain topic when the last event subscription at the scope
	// of the domain topic is deleted. If this property is set to false, then the user needs to manually delete the domain topic when it is no longer needed
	// (e.g., when last event subscription is deleted and the resource needs to be cleaned up). The self-management mode can be used if the user wants full
	// control of when the domain topic needs to be deleted, while auto-managed mode provides the flexibility to perform less operations and manage fewer
	// resources by the user.
	AutoDeleteTopicWithLastSubscription *bool `pulumi:"autoDeleteTopicWithLastSubscription"`
	// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the domain.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
	// Identity information for the Event Grid Domain resource.
	Identity *IdentityInfo `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the Event Grid Domain Resource.
	InputSchema *string `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Sku pricing tier for the Event Grid Domain resource.
	Sku *ResourceSku `pulumi:"sku"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// This Boolean is used to specify the creation mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
	// In this context, creation of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
	// When this property is null or set to true, Event Grid is responsible of automatically creating the domain topic when the first event subscription is
	// created at the scope of the domain topic. If this property is set to false, then creating the first event subscription will require creating a domain topic
	// by the user. The self-management mode can be used if the user wants full control of when the domain topic is created, while auto-managed mode provides the
	// flexibility to perform less operations and manage fewer resources by the user. Also, note that in auto-managed creation mode, user is allowed to create the
	// domain topic on demand if needed.
	AutoCreateTopicWithFirstSubscription pulumi.BoolPtrInput
	// This Boolean is used to specify the deletion mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.
	// In this context, deletion of domain topic can be auto-managed (when true) or self-managed (when false). The default value for this property is true.
	// When this property is set to true, Event Grid is responsible of automatically deleting the domain topic when the last event subscription at the scope
	// of the domain topic is deleted. If this property is set to false, then the user needs to manually delete the domain topic when it is no longer needed
	// (e.g., when last event subscription is deleted and the resource needs to be cleaned up). The self-management mode can be used if the user wants full
	// control of when the domain topic needs to be deleted, while auto-managed mode provides the flexibility to perform less operations and manage fewer
	// resources by the user.
	AutoDeleteTopicWithLastSubscription pulumi.BoolPtrInput
	// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the domain.
	DisableLocalAuth pulumi.BoolPtrInput
	// Name of the domain.
	DomainName pulumi.StringPtrInput
	// Identity information for the Event Grid Domain resource.
	Identity IdentityInfoPtrInput
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleArrayInput
	// This determines the format that Event Grid should expect for incoming events published to the Event Grid Domain Resource.
	InputSchema pulumi.StringPtrInput
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping JsonInputSchemaMappingPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The Sku pricing tier for the Event Grid Domain resource.
	Sku ResourceSkuPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}

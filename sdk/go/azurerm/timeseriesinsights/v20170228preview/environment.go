// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170228preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource.
type Environment struct {
	pulumi.CustomResourceState

	// The time the resource was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessFqdn pulumi.StringOutput `pulumi:"dataAccessFqdn"`
	// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessId pulumi.StringOutput `pulumi:"dataAccessId"`
	// ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	DataRetentionTime pulumi.StringOutput `pulumi:"dataRetentionTime"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
	StorageLimitExceededBehavior pulumi.StringPtrOutput `pulumi:"storageLimitExceededBehavior"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil || args.DataRetentionTime == nil {
		return nil, errors.New("missing required argument 'DataRetentionTime'")
	}
	if args == nil || args.EnvironmentName == nil {
		return nil, errors.New("missing required argument 'EnvironmentName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &EnvironmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:timeseriesinsights/latest:Environment"),
		},
		{
			Type: pulumi.String("azurerm:timeseriesinsights/v20171115:Environment"),
		},
		{
			Type: pulumi.String("azurerm:timeseriesinsights/v20180815preview:Environment"),
		},
		{
			Type: pulumi.String("azurerm:timeseriesinsights/v20200515:Environment"),
		},
	})
	opts = append(opts, aliases)
	var resource Environment
	err := ctx.RegisterResource("azurerm:timeseriesinsights/v20170228preview:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("azurerm:timeseriesinsights/v20170228preview:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// The time the resource was created.
	CreationTime *string `pulumi:"creationTime"`
	// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessFqdn *string `pulumi:"dataAccessFqdn"`
	// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessId *string `pulumi:"dataAccessId"`
	// ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	DataRetentionTime *string `pulumi:"dataRetentionTime"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
	Sku *SkuResponse `pulumi:"sku"`
	// The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
	StorageLimitExceededBehavior *string `pulumi:"storageLimitExceededBehavior"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type EnvironmentState struct {
	// The time the resource was created.
	CreationTime pulumi.StringPtrInput
	// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessFqdn pulumi.StringPtrInput
	// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessId pulumi.StringPtrInput
	// ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	DataRetentionTime pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
	Sku SkuResponsePtrInput
	// The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
	StorageLimitExceededBehavior pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	DataRetentionTime string `pulumi:"dataRetentionTime"`
	// Name of the environment
	EnvironmentName string `pulumi:"environmentName"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
	Sku Sku `pulumi:"sku"`
	// The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
	StorageLimitExceededBehavior *string `pulumi:"storageLimitExceededBehavior"`
	// Key-value pairs of additional properties for the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	DataRetentionTime pulumi.StringInput
	// Name of the environment
	EnvironmentName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
	Sku SkuInput
	// The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
	StorageLimitExceededBehavior pulumi.StringPtrInput
	// Key-value pairs of additional properties for the resource.
	Tags pulumi.StringMapInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}
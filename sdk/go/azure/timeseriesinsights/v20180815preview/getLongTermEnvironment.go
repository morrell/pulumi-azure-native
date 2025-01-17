// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180815preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource. LongTerm environments do not have set data retention limits.
func LookupLongTermEnvironment(ctx *pulumi.Context, args *LookupLongTermEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupLongTermEnvironmentResult, error) {
	var rv LookupLongTermEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20180815preview:getLongTermEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLongTermEnvironmentArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// Setting $expand=status will include the status of the internal services of the environment in the Time Series Insights service.
	Expand *string `pulumi:"expand"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource. LongTerm environments do not have set data retention limits.
type LookupLongTermEnvironmentResult struct {
	// The time the resource was created.
	CreationTime string `pulumi:"creationTime"`
	// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessFqdn string `pulumi:"dataAccessFqdn"`
	// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessId string `pulumi:"dataAccessId"`
	// Resource Id
	Id string `pulumi:"id"`
	// The kind of the environment.
	// Expected value is 'LongTerm'.
	Kind string `pulumi:"kind"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sku determines the type of environment, either standard (S1 or S2) or long-term (L1). For standard environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
	Sku SkuResponse `pulumi:"sku"`
	// An object that represents the status of the environment, and its internal state in the Time Series Insights service.
	Status EnvironmentStatusResponse `pulumi:"status"`
	// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
	StorageConfiguration LongTermStorageConfigurationOutputResponse `pulumi:"storageConfiguration"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The list of event properties which will be used to define the environment's time series id.
	TimeSeriesIdProperties []TimeSeriesIdPropertyResponse `pulumi:"timeSeriesIdProperties"`
	// Resource type
	Type string `pulumi:"type"`
	// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
	WarmStoreConfiguration *WarmStoreConfigurationPropertiesResponse `pulumi:"warmStoreConfiguration"`
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200820preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A class representing a CommunicationService resource.
type CommunicationService struct {
	pulumi.CustomResourceState

	// The location where the communication service stores its data at rest.
	DataLocation pulumi.StringOutput `pulumi:"dataLocation"`
	// FQDN of the CommunicationService instance.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The immutable resource Id of the communication service.
	ImmutableResourceId pulumi.StringOutput `pulumi:"immutableResourceId"`
	// The Azure location where the CommunicationService is running.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource ID of an Azure Notification Hub linked to this resource.
	NotificationHubId pulumi.StringOutput `pulumi:"notificationHubId"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the service - e.g. "Microsoft.Communication/CommunicationServices"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the CommunicationService resource. Probably you need the same or higher version of client SDKs.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewCommunicationService registers a new resource with the given unique name, arguments, and options.
func NewCommunicationService(ctx *pulumi.Context,
	name string, args *CommunicationServiceArgs, opts ...pulumi.ResourceOption) (*CommunicationService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataLocation == nil {
		return nil, errors.New("invalid value for required argument 'DataLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:communication/v20200820preview:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-nextgen:communication:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20200820:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-nextgen:communication/v20200820:CommunicationService"),
		},
	})
	opts = append(opts, aliases)
	var resource CommunicationService
	err := ctx.RegisterResource("azure-native:communication/v20200820preview:CommunicationService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommunicationService gets an existing CommunicationService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommunicationService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommunicationServiceState, opts ...pulumi.ResourceOption) (*CommunicationService, error) {
	var resource CommunicationService
	err := ctx.ReadResource("azure-native:communication/v20200820preview:CommunicationService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommunicationService resources.
type communicationServiceState struct {
}

type CommunicationServiceState struct {
}

func (CommunicationServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationServiceState)(nil)).Elem()
}

type communicationServiceArgs struct {
	// The name of the CommunicationService resource.
	CommunicationServiceName *string `pulumi:"communicationServiceName"`
	// The location where the communication service stores its data at rest.
	DataLocation string `pulumi:"dataLocation"`
	// The Azure location where the CommunicationService is running.
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CommunicationService resource.
type CommunicationServiceArgs struct {
	// The name of the CommunicationService resource.
	CommunicationServiceName pulumi.StringPtrInput
	// The location where the communication service stores its data at rest.
	DataLocation pulumi.StringInput
	// The Azure location where the CommunicationService is running.
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
}

func (CommunicationServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationServiceArgs)(nil)).Elem()
}

type CommunicationServiceInput interface {
	pulumi.Input

	ToCommunicationServiceOutput() CommunicationServiceOutput
	ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput
}

func (*CommunicationService) ElementType() reflect.Type {
	return reflect.TypeOf((*CommunicationService)(nil))
}

func (i *CommunicationService) ToCommunicationServiceOutput() CommunicationServiceOutput {
	return i.ToCommunicationServiceOutputWithContext(context.Background())
}

func (i *CommunicationService) ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunicationServiceOutput)
}

type CommunicationServiceOutput struct {
	*pulumi.OutputState
}

func (CommunicationServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommunicationService)(nil))
}

func (o CommunicationServiceOutput) ToCommunicationServiceOutput() CommunicationServiceOutput {
	return o
}

func (o CommunicationServiceOutput) ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CommunicationServiceOutput{})
}

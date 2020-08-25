// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account agreement.
type Agreement struct {
	pulumi.CustomResourceState

	// The agreement type.
	AgreementType pulumi.StringOutput `pulumi:"agreementType"`
	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The agreement content.
	Content AgreementContentResponseOutput `pulumi:"content"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentityResponseOutput `pulumi:"guestIdentity"`
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner pulumi.StringOutput `pulumi:"guestPartner"`
	// The business identity of the host partner.
	HostIdentity BusinessIdentityResponseOutput `pulumi:"hostIdentity"`
	// The integration account partner that is set as host partner for this agreement.
	HostPartner pulumi.StringOutput `pulumi:"hostPartner"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAgreement registers a new resource with the given unique name, arguments, and options.
func NewAgreement(ctx *pulumi.Context,
	name string, args *AgreementArgs, opts ...pulumi.ResourceOption) (*Agreement, error) {
	if args == nil || args.AgreementType == nil {
		return nil, errors.New("missing required argument 'AgreementType'")
	}
	if args == nil || args.Content == nil {
		return nil, errors.New("missing required argument 'Content'")
	}
	if args == nil || args.GuestIdentity == nil {
		return nil, errors.New("missing required argument 'GuestIdentity'")
	}
	if args == nil || args.GuestPartner == nil {
		return nil, errors.New("missing required argument 'GuestPartner'")
	}
	if args == nil || args.HostIdentity == nil {
		return nil, errors.New("missing required argument 'HostIdentity'")
	}
	if args == nil || args.HostPartner == nil {
		return nil, errors.New("missing required argument 'HostPartner'")
	}
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AgreementArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/v20190501:Agreement"),
		},
	})
	opts = append(opts, aliases)
	var resource Agreement
	err := ctx.RegisterResource("azurerm:logic/v20160601:Agreement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgreement gets an existing Agreement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgreement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgreementState, opts ...pulumi.ResourceOption) (*Agreement, error) {
	var resource Agreement
	err := ctx.ReadResource("azurerm:logic/v20160601:Agreement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Agreement resources.
type agreementState struct {
	// The agreement type.
	AgreementType *string `pulumi:"agreementType"`
	// The changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// The agreement content.
	Content *AgreementContentResponse `pulumi:"content"`
	// The created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The business identity of the guest partner.
	GuestIdentity *BusinessIdentityResponse `pulumi:"guestIdentity"`
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner *string `pulumi:"guestPartner"`
	// The business identity of the host partner.
	HostIdentity *BusinessIdentityResponse `pulumi:"hostIdentity"`
	// The integration account partner that is set as host partner for this agreement.
	HostPartner *string `pulumi:"hostPartner"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type AgreementState struct {
	// The agreement type.
	AgreementType pulumi.StringPtrInput
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The agreement content.
	Content AgreementContentResponsePtrInput
	// The created time.
	CreatedTime pulumi.StringPtrInput
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentityResponsePtrInput
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner pulumi.StringPtrInput
	// The business identity of the host partner.
	HostIdentity BusinessIdentityResponsePtrInput
	// The integration account partner that is set as host partner for this agreement.
	HostPartner pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.MapInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (AgreementState) ElementType() reflect.Type {
	return reflect.TypeOf((*agreementState)(nil)).Elem()
}

type agreementArgs struct {
	// The agreement type.
	AgreementType string `pulumi:"agreementType"`
	// The agreement content.
	Content AgreementContent `pulumi:"content"`
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentity `pulumi:"guestIdentity"`
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner string `pulumi:"guestPartner"`
	// The business identity of the host partner.
	HostIdentity BusinessIdentity `pulumi:"hostIdentity"`
	// The integration account partner that is set as host partner for this agreement.
	HostPartner string `pulumi:"hostPartner"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The integration account agreement name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Agreement resource.
type AgreementArgs struct {
	// The agreement type.
	AgreementType pulumi.StringInput
	// The agreement content.
	Content AgreementContentInput
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentityInput
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner pulumi.StringInput
	// The business identity of the host partner.
	HostIdentity BusinessIdentityInput
	// The integration account partner that is set as host partner for this agreement.
	HostPartner pulumi.StringInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.MapInput
	// The integration account agreement name.
	Name pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (AgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agreementArgs)(nil)).Elem()
}
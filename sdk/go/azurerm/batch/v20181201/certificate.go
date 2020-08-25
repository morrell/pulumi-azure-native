// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Contains information about a certificate.
type Certificate struct {
	pulumi.CustomResourceState

	// This is only returned when the certificate provisioningState is 'Failed'.
	DeleteCertificateError DeleteCertificateErrorResponseOutput `pulumi:"deleteCertificateError"`
	// The ETag of the resource, used for concurrency statements.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The previous provisioned state of the resource
	PreviousProvisioningState               pulumi.StringOutput `pulumi:"previousProvisioningState"`
	PreviousProvisioningStateTransitionTime pulumi.StringOutput `pulumi:"previousProvisioningStateTransitionTime"`
	ProvisioningState                       pulumi.StringOutput `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime         pulumi.StringOutput `pulumi:"provisioningStateTransitionTime"`
	// The public key of the certificate.
	PublicData pulumi.StringOutput `pulumi:"publicData"`
	// This must match the thumbprint from the name.
	Thumbprint pulumi.StringPtrOutput `pulumi:"thumbprint"`
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm pulumi.StringPtrOutput `pulumi:"thumbprintAlgorithm"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Data == nil {
		return nil, errors.New("missing required argument 'Data'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CertificateArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:batch/v20170901:Certificate"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20190401:Certificate"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20190801:Certificate"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20200301:Certificate"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20200501:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azurerm:batch/v20181201:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azurerm:batch/v20181201:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// This is only returned when the certificate provisioningState is 'Failed'.
	DeleteCertificateError *DeleteCertificateErrorResponse `pulumi:"deleteCertificateError"`
	// The ETag of the resource, used for concurrency statements.
	Etag *string `pulumi:"etag"`
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format *string `pulumi:"format"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The previous provisioned state of the resource
	PreviousProvisioningState               *string `pulumi:"previousProvisioningState"`
	PreviousProvisioningStateTransitionTime *string `pulumi:"previousProvisioningStateTransitionTime"`
	ProvisioningState                       *string `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime         *string `pulumi:"provisioningStateTransitionTime"`
	// The public key of the certificate.
	PublicData *string `pulumi:"publicData"`
	// This must match the thumbprint from the name.
	Thumbprint *string `pulumi:"thumbprint"`
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm *string `pulumi:"thumbprintAlgorithm"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type CertificateState struct {
	// This is only returned when the certificate provisioningState is 'Failed'.
	DeleteCertificateError DeleteCertificateErrorResponsePtrInput
	// The ETag of the resource, used for concurrency statements.
	Etag pulumi.StringPtrInput
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The previous provisioned state of the resource
	PreviousProvisioningState               pulumi.StringPtrInput
	PreviousProvisioningStateTransitionTime pulumi.StringPtrInput
	ProvisioningState                       pulumi.StringPtrInput
	ProvisioningStateTransitionTime         pulumi.StringPtrInput
	// The public key of the certificate.
	PublicData pulumi.StringPtrInput
	// This must match the thumbprint from the name.
	Thumbprint pulumi.StringPtrInput
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The maximum size is 10KB.
	Data string `pulumi:"data"`
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format *string `pulumi:"format"`
	// The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
	Name string `pulumi:"name"`
	// This is required if the certificate format is pfx and must be omitted if the certificate format is cer.
	Password *string `pulumi:"password"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// This must match the thumbprint from the name.
	Thumbprint *string `pulumi:"thumbprint"`
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm *string `pulumi:"thumbprintAlgorithm"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The name of the Batch account.
	AccountName pulumi.StringInput
	// The maximum size is 10KB.
	Data pulumi.StringInput
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format pulumi.StringPtrInput
	// The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
	Name pulumi.StringInput
	// This is required if the certificate format is pfx and must be omitted if the certificate format is cer.
	Password pulumi.StringPtrInput
	// The name of the resource group that contains the Batch account.
	ResourceGroupName pulumi.StringInput
	// This must match the thumbprint from the name.
	Thumbprint pulumi.StringPtrInput
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}
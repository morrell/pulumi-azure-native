// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azurerm:batch/v20181201:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Contains information about a certificate.
type LookupCertificateResult struct {
	// This is only returned when the certificate provisioningState is 'Failed'.
	DeleteCertificateError DeleteCertificateErrorResponse `pulumi:"deleteCertificateError"`
	// The ETag of the resource, used for concurrency statements.
	Etag string `pulumi:"etag"`
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format *string `pulumi:"format"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The previous provisioned state of the resource
	PreviousProvisioningState               string `pulumi:"previousProvisioningState"`
	PreviousProvisioningStateTransitionTime string `pulumi:"previousProvisioningStateTransitionTime"`
	ProvisioningState                       string `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime         string `pulumi:"provisioningStateTransitionTime"`
	// The public key of the certificate.
	PublicData string `pulumi:"publicData"`
	// This must match the thumbprint from the name.
	Thumbprint *string `pulumi:"thumbprint"`
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm *string `pulumi:"thumbprintAlgorithm"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Virtual machine guest diagnostic settings resource.
func LookupGuestDiagnosticsSettingsAssociation(ctx *pulumi.Context, args *LookupGuestDiagnosticsSettingsAssociationArgs, opts ...pulumi.InvokeOption) (*LookupGuestDiagnosticsSettingsAssociationResult, error) {
	var rv LookupGuestDiagnosticsSettingsAssociationResult
	err := ctx.Invoke("azure-native:insights/v20180601preview:getGuestDiagnosticsSettingsAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestDiagnosticsSettingsAssociationArgs struct {
	// The name of the diagnostic settings association.
	AssociationName string `pulumi:"associationName"`
	// The fully qualified ID of the resource, including the resource name and resource type.
	ResourceUri string `pulumi:"resourceUri"`
}

// Virtual machine guest diagnostic settings resource.
type LookupGuestDiagnosticsSettingsAssociationResult struct {
	// The guest diagnostic settings name.
	GuestDiagnosticSettingsName string `pulumi:"guestDiagnosticSettingsName"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type string `pulumi:"type"`
}

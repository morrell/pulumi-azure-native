// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * EventGrid Partner Namespace.
 */
export class PartnerNamespace extends pulumi.CustomResource {
    /**
     * Get an existing PartnerNamespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PartnerNamespace {
        return new PartnerNamespace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:eventgrid/v20201015preview:PartnerNamespace';

    /**
     * Returns true if the given object is an instance of PartnerNamespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PartnerNamespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PartnerNamespace.__pulumiType;
    }

    /**
     * Endpoint for the partner namespace.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
     * /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
     */
    public readonly partnerRegistrationFullyQualifiedId!: pulumi.Output<string | undefined>;
    /**
     * Provisioning state of the partner namespace.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The system metadata relating to Partner Namespace resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.eventgrid.v20201015preview.SystemDataResponse>;
    /**
     * Tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PartnerNamespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PartnerNamespaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["partnerNamespaceName"] = args ? args.partnerNamespaceName : undefined;
            inputs["partnerRegistrationFullyQualifiedId"] = args ? args.partnerRegistrationFullyQualifiedId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["endpoint"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["endpoint"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partnerRegistrationFullyQualifiedId"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:eventgrid/v20201015preview:PartnerNamespace" }, { type: "azure-native:eventgrid:PartnerNamespace" }, { type: "azure-nextgen:eventgrid:PartnerNamespace" }, { type: "azure-native:eventgrid/v20200401preview:PartnerNamespace" }, { type: "azure-nextgen:eventgrid/v20200401preview:PartnerNamespace" }, { type: "azure-native:eventgrid/v20210601preview:PartnerNamespace" }, { type: "azure-nextgen:eventgrid/v20210601preview:PartnerNamespace" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PartnerNamespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PartnerNamespace resource.
 */
export interface PartnerNamespaceArgs {
    /**
     * Location of the resource.
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the partner namespace.
     */
    partnerNamespaceName?: pulumi.Input<string>;
    /**
     * The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
     * /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
     */
    partnerRegistrationFullyQualifiedId?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Tags of the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

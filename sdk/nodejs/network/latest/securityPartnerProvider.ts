// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Security Partner Provider resource.
 */
export class SecurityPartnerProvider extends pulumi.CustomResource {
    /**
     * Get an existing SecurityPartnerProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecurityPartnerProvider {
        return new SecurityPartnerProvider(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/latest:SecurityPartnerProvider';

    /**
     * Returns true if the given object is an instance of SecurityPartnerProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityPartnerProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityPartnerProvider.__pulumiType;
    }

    /**
     * The connection status with the Security Partner Provider.
     */
    public /*out*/ readonly connectionStatus!: pulumi.Output<string>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the Security Partner Provider resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The security provider name.
     */
    public readonly securityProviderName!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The virtualHub to which the Security Partner Provider belongs.
     */
    public readonly virtualHub!: pulumi.Output<outputs.network.latest.SubResourceResponse | undefined>;

    /**
     * Create a SecurityPartnerProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityPartnerProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityPartnerProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SecurityPartnerProviderArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.securityPartnerProviderName === undefined) {
                throw new Error("Missing required property 'securityPartnerProviderName'");
            }
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["securityPartnerProviderName"] = args ? args.securityPartnerProviderName : undefined;
            inputs["securityProviderName"] = args ? args.securityProviderName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualHub"] = args ? args.virtualHub : undefined;
            inputs["connectionStatus"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/v20200301:SecurityPartnerProvider" }, { type: "azurerm:network/v20200401:SecurityPartnerProvider" }, { type: "azurerm:network/v20200501:SecurityPartnerProvider" }, { type: "azurerm:network/v20200601:SecurityPartnerProvider" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SecurityPartnerProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecurityPartnerProvider resource.
 */
export interface SecurityPartnerProviderArgs {
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Security Partner Provider.
     */
    readonly securityPartnerProviderName: pulumi.Input<string>;
    /**
     * The security provider name.
     */
    readonly securityProviderName?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The virtualHub to which the Security Partner Provider belongs.
     */
    readonly virtualHub?: pulumi.Input<inputs.network.latest.SubResource>;
}
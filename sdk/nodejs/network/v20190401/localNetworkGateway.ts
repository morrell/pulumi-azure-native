// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A common class for general resource information.
 */
export class LocalNetworkGateway extends pulumi.CustomResource {
    /**
     * Get an existing LocalNetworkGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LocalNetworkGateway {
        return new LocalNetworkGateway(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190401:LocalNetworkGateway';

    /**
     * Returns true if the given object is an instance of LocalNetworkGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocalNetworkGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocalNetworkGateway.__pulumiType;
    }

    /**
     * Local network gateway's BGP speaker settings.
     */
    public readonly bgpSettings!: pulumi.Output<outputs.network.v20190401.BgpSettingsResponse | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * IP address of local network gateway.
     */
    public readonly gatewayIpAddress!: pulumi.Output<string | undefined>;
    /**
     * Local network site address space.
     */
    public readonly localNetworkAddressSpace!: pulumi.Output<outputs.network.v20190401.AddressSpaceResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the LocalNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource GUID property of the LocalNetworkGateway resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LocalNetworkGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocalNetworkGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocalNetworkGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as LocalNetworkGatewayArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["bgpSettings"] = args ? args.bgpSettings : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["gatewayIpAddress"] = args ? args.gatewayIpAddress : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["localNetworkAddressSpace"] = args ? args.localNetworkAddressSpace : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/v20150615:LocalNetworkGateway" }, { type: "azurerm:network/v20160330:LocalNetworkGateway" }, { type: "azurerm:network/v20160601:LocalNetworkGateway" }, { type: "azurerm:network/v20160901:LocalNetworkGateway" }, { type: "azurerm:network/v20161201:LocalNetworkGateway" }, { type: "azurerm:network/v20170301:LocalNetworkGateway" }, { type: "azurerm:network/v20170601:LocalNetworkGateway" }, { type: "azurerm:network/v20170801:LocalNetworkGateway" }, { type: "azurerm:network/v20170901:LocalNetworkGateway" }, { type: "azurerm:network/v20171001:LocalNetworkGateway" }, { type: "azurerm:network/v20171101:LocalNetworkGateway" }, { type: "azurerm:network/v20180101:LocalNetworkGateway" }, { type: "azurerm:network/v20180201:LocalNetworkGateway" }, { type: "azurerm:network/v20180401:LocalNetworkGateway" }, { type: "azurerm:network/v20180601:LocalNetworkGateway" }, { type: "azurerm:network/v20180701:LocalNetworkGateway" }, { type: "azurerm:network/v20180801:LocalNetworkGateway" }, { type: "azurerm:network/v20181001:LocalNetworkGateway" }, { type: "azurerm:network/v20181101:LocalNetworkGateway" }, { type: "azurerm:network/v20181201:LocalNetworkGateway" }, { type: "azurerm:network/v20190201:LocalNetworkGateway" }, { type: "azurerm:network/v20190601:LocalNetworkGateway" }, { type: "azurerm:network/v20190701:LocalNetworkGateway" }, { type: "azurerm:network/v20190801:LocalNetworkGateway" }, { type: "azurerm:network/v20190901:LocalNetworkGateway" }, { type: "azurerm:network/v20191101:LocalNetworkGateway" }, { type: "azurerm:network/v20191201:LocalNetworkGateway" }, { type: "azurerm:network/v20200301:LocalNetworkGateway" }, { type: "azurerm:network/v20200401:LocalNetworkGateway" }, { type: "azurerm:network/v20200501:LocalNetworkGateway" }, { type: "azurerm:network/v20200601:LocalNetworkGateway" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LocalNetworkGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LocalNetworkGateway resource.
 */
export interface LocalNetworkGatewayArgs {
    /**
     * Local network gateway's BGP speaker settings.
     */
    readonly bgpSettings?: pulumi.Input<inputs.network.v20190401.BgpSettings>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * IP address of local network gateway.
     */
    readonly gatewayIpAddress?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Local network site address space.
     */
    readonly localNetworkAddressSpace?: pulumi.Input<inputs.network.v20190401.AddressSpace>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the local network gateway.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource GUID property of the LocalNetworkGateway resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A network interface in a resource group.
 */
export class NetworkInterface extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkInterface {
        return new NetworkInterface(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20180401:NetworkInterface';

    /**
     * Returns true if the given object is an instance of NetworkInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterface.__pulumiType;
    }

    /**
     * The DNS settings in network interface.
     */
    public readonly dnsSettings!: pulumi.Output<outputs.network.v20180401.NetworkInterfaceDnsSettingsResponse | undefined>;
    /**
     * If the network interface is accelerated networking enabled.
     */
    public readonly enableAcceleratedNetworking!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether IP forwarding is enabled on this network interface.
     */
    public readonly enableIPForwarding!: pulumi.Output<boolean | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * A list of IPConfigurations of the network interface.
     */
    public readonly ipConfigurations!: pulumi.Output<outputs.network.v20180401.NetworkInterfaceIPConfigurationResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The MAC address of the network interface.
     */
    public readonly macAddress!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The reference of the NetworkSecurityGroup resource.
     */
    public readonly networkSecurityGroup!: pulumi.Output<outputs.network.v20180401.NetworkSecurityGroupResponse | undefined>;
    /**
     * Gets whether this is a primary network interface on a virtual machine.
     */
    public readonly primary!: pulumi.Output<boolean | undefined>;
    /**
     * The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The resource GUID property of the network interface resource.
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
     * The reference of a virtual machine.
     */
    public readonly virtualMachine!: pulumi.Output<outputs.network.v20180401.SubResourceResponse | undefined>;

    /**
     * Create a NetworkInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as NetworkInterfaceArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dnsSettings"] = args ? args.dnsSettings : undefined;
            inputs["enableAcceleratedNetworking"] = args ? args.enableAcceleratedNetworking : undefined;
            inputs["enableIPForwarding"] = args ? args.enableIPForwarding : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipConfigurations"] = args ? args.ipConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["macAddress"] = args ? args.macAddress : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkSecurityGroup"] = args ? args.networkSecurityGroup : undefined;
            inputs["primary"] = args ? args.primary : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualMachine"] = args ? args.virtualMachine : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/v20150615:NetworkInterface" }, { type: "azurerm:network/v20160330:NetworkInterface" }, { type: "azurerm:network/v20160601:NetworkInterface" }, { type: "azurerm:network/v20160901:NetworkInterface" }, { type: "azurerm:network/v20161201:NetworkInterface" }, { type: "azurerm:network/v20170301:NetworkInterface" }, { type: "azurerm:network/v20170601:NetworkInterface" }, { type: "azurerm:network/v20170801:NetworkInterface" }, { type: "azurerm:network/v20170901:NetworkInterface" }, { type: "azurerm:network/v20171001:NetworkInterface" }, { type: "azurerm:network/v20171101:NetworkInterface" }, { type: "azurerm:network/v20180101:NetworkInterface" }, { type: "azurerm:network/v20180201:NetworkInterface" }, { type: "azurerm:network/v20180601:NetworkInterface" }, { type: "azurerm:network/v20180701:NetworkInterface" }, { type: "azurerm:network/v20180801:NetworkInterface" }, { type: "azurerm:network/v20181001:NetworkInterface" }, { type: "azurerm:network/v20181101:NetworkInterface" }, { type: "azurerm:network/v20181201:NetworkInterface" }, { type: "azurerm:network/v20190201:NetworkInterface" }, { type: "azurerm:network/v20190401:NetworkInterface" }, { type: "azurerm:network/v20190601:NetworkInterface" }, { type: "azurerm:network/v20190701:NetworkInterface" }, { type: "azurerm:network/v20190801:NetworkInterface" }, { type: "azurerm:network/v20190901:NetworkInterface" }, { type: "azurerm:network/v20191101:NetworkInterface" }, { type: "azurerm:network/v20191201:NetworkInterface" }, { type: "azurerm:network/v20200301:NetworkInterface" }, { type: "azurerm:network/v20200401:NetworkInterface" }, { type: "azurerm:network/v20200501:NetworkInterface" }, { type: "azurerm:network/v20200601:NetworkInterface" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NetworkInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkInterface resource.
 */
export interface NetworkInterfaceArgs {
    /**
     * The DNS settings in network interface.
     */
    readonly dnsSettings?: pulumi.Input<inputs.network.v20180401.NetworkInterfaceDnsSettings>;
    /**
     * If the network interface is accelerated networking enabled.
     */
    readonly enableAcceleratedNetworking?: pulumi.Input<boolean>;
    /**
     * Indicates whether IP forwarding is enabled on this network interface.
     */
    readonly enableIPForwarding?: pulumi.Input<boolean>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * A list of IPConfigurations of the network interface.
     */
    readonly ipConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20180401.NetworkInterfaceIPConfiguration>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The MAC address of the network interface.
     */
    readonly macAddress?: pulumi.Input<string>;
    /**
     * The name of the network interface.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The reference of the NetworkSecurityGroup resource.
     */
    readonly networkSecurityGroup?: pulumi.Input<inputs.network.v20180401.NetworkSecurityGroup>;
    /**
     * Gets whether this is a primary network interface on a virtual machine.
     */
    readonly primary?: pulumi.Input<boolean>;
    /**
     * The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource GUID property of the network interface resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The reference of a virtual machine.
     */
    readonly virtualMachine?: pulumi.Input<inputs.network.v20180401.SubResource>;
}
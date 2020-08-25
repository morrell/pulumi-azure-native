// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
 */
export class AvailabilitySet extends pulumi.CustomResource {
    /**
     * Get an existing AvailabilitySet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AvailabilitySet {
        return new AvailabilitySet(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20170330:AvailabilitySet';

    /**
     * Returns true if the given object is an instance of AvailabilitySet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AvailabilitySet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AvailabilitySet.__pulumiType;
    }

    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Fault Domain count.
     */
    public readonly platformFaultDomainCount!: pulumi.Output<number | undefined>;
    /**
     * Update Domain count.
     */
    public readonly platformUpdateDomainCount!: pulumi.Output<number | undefined>;
    /**
     * Sku of the availability set
     */
    public readonly sku!: pulumi.Output<outputs.compute.v20170330.SkuResponse | undefined>;
    /**
     * The resource status information.
     */
    public /*out*/ readonly statuses!: pulumi.Output<outputs.compute.v20170330.InstanceViewStatusResponse[]>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A list of references to all virtual machines in the availability set.
     */
    public readonly virtualMachines!: pulumi.Output<outputs.compute.v20170330.SubResourceResponse[] | undefined>;

    /**
     * Create a AvailabilitySet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AvailabilitySetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AvailabilitySetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AvailabilitySetArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platformFaultDomainCount"] = args ? args.platformFaultDomainCount : undefined;
            inputs["platformUpdateDomainCount"] = args ? args.platformUpdateDomainCount : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualMachines"] = args ? args.virtualMachines : undefined;
            inputs["statuses"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:compute/v20150615:AvailabilitySet" }, { type: "azurerm:compute/v20160330:AvailabilitySet" }, { type: "azurerm:compute/v20171201:AvailabilitySet" }, { type: "azurerm:compute/v20180401:AvailabilitySet" }, { type: "azurerm:compute/v20180601:AvailabilitySet" }, { type: "azurerm:compute/v20181001:AvailabilitySet" }, { type: "azurerm:compute/v20190301:AvailabilitySet" }, { type: "azurerm:compute/v20190701:AvailabilitySet" }, { type: "azurerm:compute/v20191201:AvailabilitySet" }, { type: "azurerm:compute/v20200601:AvailabilitySet" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AvailabilitySet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AvailabilitySet resource.
 */
export interface AvailabilitySetArgs {
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the availability set.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Fault Domain count.
     */
    readonly platformFaultDomainCount?: pulumi.Input<number>;
    /**
     * Update Domain count.
     */
    readonly platformUpdateDomainCount?: pulumi.Input<number>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Sku of the availability set
     */
    readonly sku?: pulumi.Input<inputs.compute.v20170330.Sku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of references to all virtual machines in the availability set.
     */
    readonly virtualMachines?: pulumi.Input<pulumi.Input<inputs.compute.v20170330.SubResource>[]>;
}
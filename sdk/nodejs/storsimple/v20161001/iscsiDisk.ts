// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The iSCSI disk.
 */
export class IscsiDisk extends pulumi.CustomResource {
    /**
     * Get an existing IscsiDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IscsiDisk {
        return new IscsiDisk(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storsimple/v20161001:IscsiDisk';

    /**
     * Returns true if the given object is an instance of IscsiDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IscsiDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IscsiDisk.__pulumiType;
    }

    /**
     * The access control records.
     */
    public readonly accessControlRecords!: pulumi.Output<string[]>;
    /**
     * The data policy.
     */
    public readonly dataPolicy!: pulumi.Output<string>;
    /**
     * The description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The disk status.
     */
    public readonly diskStatus!: pulumi.Output<string>;
    /**
     * The local used capacity in bytes.
     */
    public /*out*/ readonly localUsedCapacityInBytes!: pulumi.Output<number>;
    /**
     * The monitoring.
     */
    public readonly monitoringStatus!: pulumi.Output<string>;
    /**
     * The name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The provisioned capacity in bytes.
     */
    public readonly provisionedCapacityInBytes!: pulumi.Output<number>;
    /**
     * The type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The used capacity in bytes.
     */
    public /*out*/ readonly usedCapacityInBytes!: pulumi.Output<number>;

    /**
     * Create a IscsiDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IscsiDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IscsiDiskArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as IscsiDiskArgs | undefined;
            if (!args || args.accessControlRecords === undefined) {
                throw new Error("Missing required property 'accessControlRecords'");
            }
            if (!args || args.dataPolicy === undefined) {
                throw new Error("Missing required property 'dataPolicy'");
            }
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.diskStatus === undefined) {
                throw new Error("Missing required property 'diskStatus'");
            }
            if (!args || args.iscsiServerName === undefined) {
                throw new Error("Missing required property 'iscsiServerName'");
            }
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.monitoringStatus === undefined) {
                throw new Error("Missing required property 'monitoringStatus'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.provisionedCapacityInBytes === undefined) {
                throw new Error("Missing required property 'provisionedCapacityInBytes'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accessControlRecords"] = args ? args.accessControlRecords : undefined;
            inputs["dataPolicy"] = args ? args.dataPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["diskStatus"] = args ? args.diskStatus : undefined;
            inputs["iscsiServerName"] = args ? args.iscsiServerName : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["monitoringStatus"] = args ? args.monitoringStatus : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["provisionedCapacityInBytes"] = args ? args.provisionedCapacityInBytes : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["localUsedCapacityInBytes"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["usedCapacityInBytes"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IscsiDisk.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IscsiDisk resource.
 */
export interface IscsiDiskArgs {
    /**
     * The access control records.
     */
    readonly accessControlRecords: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The data policy.
     */
    readonly dataPolicy: pulumi.Input<string>;
    /**
     * The description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The device name.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The disk status.
     */
    readonly diskStatus: pulumi.Input<string>;
    /**
     * The iSCSI server name.
     */
    readonly iscsiServerName: pulumi.Input<string>;
    /**
     * The manager name
     */
    readonly managerName: pulumi.Input<string>;
    /**
     * The monitoring.
     */
    readonly monitoringStatus: pulumi.Input<string>;
    /**
     * The disk name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The provisioned capacity in bytes.
     */
    readonly provisionedCapacityInBytes: pulumi.Input<number>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
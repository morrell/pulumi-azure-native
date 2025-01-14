// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The file server.
 */
export class FileServer extends pulumi.CustomResource {
    /**
     * Get an existing FileServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FileServer {
        return new FileServer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:storsimple/v20161001:FileServer';

    /**
     * Returns true if the given object is an instance of FileServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FileServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FileServer.__pulumiType;
    }

    /**
     * The backup policy id.
     */
    public readonly backupScheduleGroupId!: pulumi.Output<string>;
    /**
     * The description of the file server
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Domain of the file server
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The storage domain id.
     */
    public readonly storageDomainId!: pulumi.Output<string>;
    /**
     * The type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a FileServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileServerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.backupScheduleGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupScheduleGroupId'");
            }
            if ((!args || args.deviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceName'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.managerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managerName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.storageDomainId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageDomainId'");
            }
            inputs["backupScheduleGroupId"] = args ? args.backupScheduleGroupId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["fileServerName"] = args ? args.fileServerName : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageDomainId"] = args ? args.storageDomainId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["backupScheduleGroupId"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["domainName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["storageDomainId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:storsimple/v20161001:FileServer" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(FileServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FileServer resource.
 */
export interface FileServerArgs {
    /**
     * The backup policy id.
     */
    backupScheduleGroupId: pulumi.Input<string>;
    /**
     * The description of the file server
     */
    description?: pulumi.Input<string>;
    /**
     * The device name.
     */
    deviceName: pulumi.Input<string>;
    /**
     * Domain of the file server
     */
    domainName: pulumi.Input<string>;
    /**
     * The file server name.
     */
    fileServerName?: pulumi.Input<string>;
    /**
     * The manager name
     */
    managerName: pulumi.Input<string>;
    /**
     * The resource group name
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The storage domain id.
     */
    storageDomainId: pulumi.Input<string>;
}

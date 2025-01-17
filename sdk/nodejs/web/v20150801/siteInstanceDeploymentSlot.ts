// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Represents user credentials used for publishing activity
 */
export class SiteInstanceDeploymentSlot extends pulumi.CustomResource {
    /**
     * Get an existing SiteInstanceDeploymentSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SiteInstanceDeploymentSlot {
        return new SiteInstanceDeploymentSlot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:web/v20150801:SiteInstanceDeploymentSlot';

    /**
     * Returns true if the given object is an instance of SiteInstanceDeploymentSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SiteInstanceDeploymentSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SiteInstanceDeploymentSlot.__pulumiType;
    }

    /**
     * Active
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * Author
     */
    public readonly author!: pulumi.Output<string | undefined>;
    /**
     * AuthorEmail
     */
    public readonly authorEmail!: pulumi.Output<string | undefined>;
    /**
     * Deployer
     */
    public readonly deployer!: pulumi.Output<string | undefined>;
    /**
     * Detail
     */
    public readonly details!: pulumi.Output<string | undefined>;
    /**
     * EndTime
     */
    public readonly endTime!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Message
     */
    public readonly message!: pulumi.Output<string | undefined>;
    /**
     * Resource Name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * StartTime
     */
    public readonly startTime!: pulumi.Output<string | undefined>;
    /**
     * Status
     */
    public readonly status!: pulumi.Output<number | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a SiteInstanceDeploymentSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SiteInstanceDeploymentSlotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.slot === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slot'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["author"] = args ? args.author : undefined;
            inputs["authorEmail"] = args ? args.authorEmail : undefined;
            inputs["deployer"] = args ? args.deployer : undefined;
            inputs["details"] = args ? args.details : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["message"] = args ? args.message : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["slot"] = args ? args.slot : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["active"] = undefined /*out*/;
            inputs["author"] = undefined /*out*/;
            inputs["authorEmail"] = undefined /*out*/;
            inputs["deployer"] = undefined /*out*/;
            inputs["details"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["message"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:web/v20150801:SiteInstanceDeploymentSlot" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SiteInstanceDeploymentSlot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SiteInstanceDeploymentSlot resource.
 */
export interface SiteInstanceDeploymentSlotArgs {
    /**
     * Active
     */
    active?: pulumi.Input<boolean>;
    /**
     * Author
     */
    author?: pulumi.Input<string>;
    /**
     * AuthorEmail
     */
    authorEmail?: pulumi.Input<string>;
    /**
     * Deployer
     */
    deployer?: pulumi.Input<string>;
    /**
     * Detail
     */
    details?: pulumi.Input<string>;
    /**
     * EndTime
     */
    endTime?: pulumi.Input<string>;
    /**
     * Resource Id
     */
    id?: pulumi.Input<string>;
    /**
     * Id of web app instance
     */
    instanceId: pulumi.Input<string>;
    /**
     * Kind of resource
     */
    kind?: pulumi.Input<string>;
    /**
     * Resource Location
     */
    location?: pulumi.Input<string>;
    /**
     * Message
     */
    message?: pulumi.Input<string>;
    /**
     * Resource Name
     */
    name: pulumi.Input<string>;
    /**
     * Name of resource group
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of web app slot. If not specified then will default to production slot.
     */
    slot: pulumi.Input<string>;
    /**
     * StartTime
     */
    startTime?: pulumi.Input<string>;
    /**
     * Status
     */
    status?: pulumi.Input<number>;
    /**
     * Resource tags
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource type
     */
    type?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Represents a disaster recovery configuration.
 */
export class DisasterRecoveryConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing DisasterRecoveryConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DisasterRecoveryConfiguration {
        return new DisasterRecoveryConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:sql/v20140401:DisasterRecoveryConfiguration';

    /**
     * Returns true if the given object is an instance of DisasterRecoveryConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DisasterRecoveryConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DisasterRecoveryConfiguration.__pulumiType;
    }

    /**
     * Whether or not failover can be done automatically.
     */
    public /*out*/ readonly autoFailover!: pulumi.Output<string>;
    /**
     * How aggressive the automatic failover should be.
     */
    public /*out*/ readonly failoverPolicy!: pulumi.Output<string>;
    /**
     * Location of the server that contains this disaster recovery configuration.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * Logical name of the server.
     */
    public /*out*/ readonly logicalServerName!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Logical name of the partner server.
     */
    public /*out*/ readonly partnerLogicalServerName!: pulumi.Output<string>;
    /**
     * Id of the partner server.
     */
    public /*out*/ readonly partnerServerId!: pulumi.Output<string>;
    /**
     * The role of the current server in the disaster recovery configuration.
     */
    public /*out*/ readonly role!: pulumi.Output<string>;
    /**
     * The status of the disaster recovery configuration.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DisasterRecoveryConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DisasterRecoveryConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            inputs["disasterRecoveryConfigurationName"] = args ? args.disasterRecoveryConfigurationName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["autoFailover"] = undefined /*out*/;
            inputs["failoverPolicy"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["logicalServerName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partnerLogicalServerName"] = undefined /*out*/;
            inputs["partnerServerId"] = undefined /*out*/;
            inputs["role"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["autoFailover"] = undefined /*out*/;
            inputs["failoverPolicy"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["logicalServerName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partnerLogicalServerName"] = undefined /*out*/;
            inputs["partnerServerId"] = undefined /*out*/;
            inputs["role"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:sql/v20140401:DisasterRecoveryConfiguration" }, { type: "azure-native:sql:DisasterRecoveryConfiguration" }, { type: "azure-nextgen:sql:DisasterRecoveryConfiguration" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DisasterRecoveryConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DisasterRecoveryConfiguration resource.
 */
export interface DisasterRecoveryConfigurationArgs {
    /**
     * The name of the disaster recovery configuration to be created/updated.
     */
    disasterRecoveryConfigurationName?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
}

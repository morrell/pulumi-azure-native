// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * On-premise IoT sensor
 */
export class OnPremiseIotSensor extends pulumi.CustomResource {
    /**
     * Get an existing OnPremiseIotSensor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OnPremiseIotSensor {
        return new OnPremiseIotSensor(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:security/v20200806preview:OnPremiseIotSensor';

    /**
     * Returns true if the given object is an instance of OnPremiseIotSensor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OnPremiseIotSensor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OnPremiseIotSensor.__pulumiType;
    }

    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a OnPremiseIotSensor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OnPremiseIotSensorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["onPremiseIotSensorName"] = args ? args.onPremiseIotSensorName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:security/v20200806preview:OnPremiseIotSensor" }, { type: "azure-native:security:OnPremiseIotSensor" }, { type: "azure-nextgen:security:OnPremiseIotSensor" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(OnPremiseIotSensor.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OnPremiseIotSensor resource.
 */
export interface OnPremiseIotSensorArgs {
    /**
     * Name of the on-premise IoT sensor
     */
    onPremiseIotSensorName?: pulumi.Input<string>;
}

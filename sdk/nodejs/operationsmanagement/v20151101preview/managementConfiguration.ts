// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The container for solution.
 */
export class ManagementConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing ManagementConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagementConfiguration {
        return new ManagementConfiguration(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:operationsmanagement/v20151101preview:ManagementConfiguration';

    /**
     * Returns true if the given object is an instance of ManagementConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagementConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagementConfiguration.__pulumiType;
    }

    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties for ManagementConfiguration object supported by the OperationsManagement resource provider.
     */
    public readonly properties!: pulumi.Output<outputs.operationsmanagement.v20151101preview.ManagementConfigurationPropertiesResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagementConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagementConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagementConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ManagementConfigurationArgs | undefined;
            if (!args || args.managementConfigurationName === undefined) {
                throw new Error("Missing required property 'managementConfigurationName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["managementConfigurationName"] = args ? args.managementConfigurationName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ManagementConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagementConfiguration resource.
 */
export interface ManagementConfigurationArgs {
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * User Management Configuration Name.
     */
    readonly managementConfigurationName: pulumi.Input<string>;
    /**
     * Properties for ManagementConfiguration object supported by the OperationsManagement resource provider.
     */
    readonly properties?: pulumi.Input<inputs.operationsmanagement.v20151101preview.ManagementConfigurationProperties>;
    /**
     * The name of the resource group to get. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
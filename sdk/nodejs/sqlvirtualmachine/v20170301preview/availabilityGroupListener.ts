// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A SQL Server availability group listener.
 */
export class AvailabilityGroupListener extends pulumi.CustomResource {
    /**
     * Get an existing AvailabilityGroupListener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AvailabilityGroupListener {
        return new AvailabilityGroupListener(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sqlvirtualmachine/v20170301preview:AvailabilityGroupListener';

    /**
     * Returns true if the given object is an instance of AvailabilityGroupListener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AvailabilityGroupListener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AvailabilityGroupListener.__pulumiType;
    }

    /**
     * Name of the availability group.
     */
    public readonly availabilityGroupName!: pulumi.Output<string | undefined>;
    /**
     * Create a default availability group if it does not exist.
     */
    public readonly createDefaultAvailabilityGroupIfNotExist!: pulumi.Output<boolean | undefined>;
    /**
     * List of load balancer configurations for an availability group listener.
     */
    public readonly loadBalancerConfigurations!: pulumi.Output<outputs.sqlvirtualmachine.v20170301preview.LoadBalancerConfigurationResponse[] | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Listener port.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * Provisioning state to track the async operation status.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AvailabilityGroupListener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AvailabilityGroupListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AvailabilityGroupListenerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AvailabilityGroupListenerArgs | undefined;
            if (!args || args.availabilityGroupListenerName === undefined) {
                throw new Error("Missing required property 'availabilityGroupListenerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sqlVirtualMachineGroupName === undefined) {
                throw new Error("Missing required property 'sqlVirtualMachineGroupName'");
            }
            inputs["availabilityGroupListenerName"] = args ? args.availabilityGroupListenerName : undefined;
            inputs["availabilityGroupName"] = args ? args.availabilityGroupName : undefined;
            inputs["createDefaultAvailabilityGroupIfNotExist"] = args ? args.createDefaultAvailabilityGroupIfNotExist : undefined;
            inputs["loadBalancerConfigurations"] = args ? args.loadBalancerConfigurations : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sqlVirtualMachineGroupName"] = args ? args.sqlVirtualMachineGroupName : undefined;
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
        super(AvailabilityGroupListener.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AvailabilityGroupListener resource.
 */
export interface AvailabilityGroupListenerArgs {
    /**
     * Name of the availability group listener.
     */
    readonly availabilityGroupListenerName: pulumi.Input<string>;
    /**
     * Name of the availability group.
     */
    readonly availabilityGroupName?: pulumi.Input<string>;
    /**
     * Create a default availability group if it does not exist.
     */
    readonly createDefaultAvailabilityGroupIfNotExist?: pulumi.Input<boolean>;
    /**
     * List of load balancer configurations for an availability group listener.
     */
    readonly loadBalancerConfigurations?: pulumi.Input<pulumi.Input<inputs.sqlvirtualmachine.v20170301preview.LoadBalancerConfiguration>[]>;
    /**
     * Listener port.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the SQL virtual machine group.
     */
    readonly sqlVirtualMachineGroupName: pulumi.Input<string>;
}
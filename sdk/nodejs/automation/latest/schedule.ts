// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Definition of the schedule.
 */
export class Schedule extends pulumi.CustomResource {
    /**
     * Get an existing Schedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Schedule {
        return new Schedule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:automation/latest:Schedule';

    /**
     * Returns true if the given object is an instance of Schedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schedule.__pulumiType;
    }

    /**
     * Gets or sets the advanced schedule.
     */
    public readonly advancedSchedule!: pulumi.Output<outputs.automation.latest.AdvancedScheduleResponse | undefined>;
    /**
     * Gets or sets the creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the end time of the schedule.
     */
    public readonly expiryTime!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the expiry time's offset in minutes.
     */
    public /*out*/ readonly expiryTimeOffsetMinutes!: pulumi.Output<number | undefined>;
    /**
     * Gets or sets the frequency of the schedule.
     */
    public readonly frequency!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the interval of the schedule.
     */
    public readonly interval!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Gets or sets a value indicating whether this schedule is enabled.
     */
    public /*out*/ readonly isEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Gets or sets the last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Gets or sets the next run time of the schedule.
     */
    public /*out*/ readonly nextRun!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the next run time's offset in minutes.
     */
    public /*out*/ readonly nextRunOffsetMinutes!: pulumi.Output<number | undefined>;
    /**
     * Gets or sets the start time of the schedule.
     */
    public readonly startTime!: pulumi.Output<string | undefined>;
    /**
     * Gets the start time's offset in minutes.
     */
    public /*out*/ readonly startTimeOffsetMinutes!: pulumi.Output<number>;
    /**
     * Gets or sets the time zone of the schedule.
     */
    public readonly timeZone!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Schedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ScheduleArgs | undefined;
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.frequency === undefined) {
                throw new Error("Missing required property 'frequency'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.scheduleName === undefined) {
                throw new Error("Missing required property 'scheduleName'");
            }
            if (!args || args.startTime === undefined) {
                throw new Error("Missing required property 'startTime'");
            }
            inputs["advancedSchedule"] = args ? args.advancedSchedule : undefined;
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["expiryTime"] = args ? args.expiryTime : undefined;
            inputs["frequency"] = args ? args.frequency : undefined;
            inputs["interval"] = args ? args.interval : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scheduleName"] = args ? args.scheduleName : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["timeZone"] = args ? args.timeZone : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["expiryTimeOffsetMinutes"] = undefined /*out*/;
            inputs["isEnabled"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["nextRun"] = undefined /*out*/;
            inputs["nextRunOffsetMinutes"] = undefined /*out*/;
            inputs["startTimeOffsetMinutes"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:automation/v20151031:Schedule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Schedule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Schedule resource.
 */
export interface ScheduleArgs {
    /**
     * Gets or sets the AdvancedSchedule.
     */
    readonly advancedSchedule?: pulumi.Input<inputs.automation.latest.AdvancedSchedule>;
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * Gets or sets the description of the schedule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Gets or sets the end time of the schedule.
     */
    readonly expiryTime?: pulumi.Input<string>;
    /**
     * Gets or sets the frequency of the schedule.
     */
    readonly frequency: pulumi.Input<string>;
    /**
     * Gets or sets the interval of the schedule.
     */
    readonly interval?: pulumi.Input<{[key: string]: any}>;
    /**
     * Gets or sets the name of the Schedule.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The schedule name.
     */
    readonly scheduleName: pulumi.Input<string>;
    /**
     * Gets or sets the start time of the schedule.
     */
    readonly startTime: pulumi.Input<string>;
    /**
     * Gets or sets the time zone of the schedule.
     */
    readonly timeZone?: pulumi.Input<string>;
}
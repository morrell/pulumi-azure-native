// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Response to put/get patch schedules for Redis cache.
 */
export function getPatchSchedule(args: GetPatchScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetPatchScheduleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:cache/v20170201:getPatchSchedule", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPatchScheduleArgs {
    /**
     * The name of the redis cache.
     */
    name: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * Response to put/get patch schedules for Redis cache.
 */
export interface GetPatchScheduleResult {
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * List of patch schedules for a Redis cache.
     */
    readonly scheduleEntries: outputs.cache.v20170201.ScheduleEntryResponse[];
    /**
     * Resource type.
     */
    readonly type: string;
}

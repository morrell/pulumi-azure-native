// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getsnapshotPolicy(args: GetsnapshotPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetsnapshotPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:netapp/v20200601:getsnapshotPolicy", {
        "accountName": args.accountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetsnapshotPolicyArgs {
    /**
     * The name of the NetApp account
     */
    readonly accountName: string;
    /**
     * The name of the snapshot policy target
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Snapshot policy information
 */
export interface GetsnapshotPolicyResult {
    /**
     * Schedule for daily snapshots
     */
    readonly dailySchedule?: {[key: string]: any};
    /**
     * The property to decide policy is enabled or not
     */
    readonly enabled?: boolean;
    /**
     * Schedule for hourly snapshots
     */
    readonly hourlySchedule?: {[key: string]: any};
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Schedule for monthly snapshots
     */
    readonly monthlySchedule?: {[key: string]: any};
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Schedule for weekly snapshots
     */
    readonly weeklySchedule?: {[key: string]: any};
}
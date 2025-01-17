// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Trigger details.
 */
/** @deprecated Please use one of the variants: FileEventTrigger, PeriodicTimerEventTrigger. */
export function getTrigger(args: GetTriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetTriggerResult> {
    pulumi.log.warn("getTrigger is deprecated: Please use one of the variants: FileEventTrigger, PeriodicTimerEventTrigger.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:databoxedge/v20190701:getTrigger", {
        "deviceName": args.deviceName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetTriggerArgs {
    /**
     * The device name.
     */
    deviceName: string;
    /**
     * The trigger name.
     */
    name: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
}

/**
 * Trigger details.
 */
export interface GetTriggerResult {
    /**
     * The path ID that uniquely identifies the object.
     */
    readonly id: string;
    /**
     * Trigger Kind.
     */
    readonly kind: string;
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Defines the response of a trigger subscription operation.
 */
export function getTriggerEventSubscriptionStatus(args: GetTriggerEventSubscriptionStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetTriggerEventSubscriptionStatusResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:datafactory/v20180601:getTriggerEventSubscriptionStatus", {
        "factoryName": args.factoryName,
        "resourceGroupName": args.resourceGroupName,
        "triggerName": args.triggerName,
    }, opts);
}

export interface GetTriggerEventSubscriptionStatusArgs {
    /**
     * The factory name.
     */
    factoryName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
    /**
     * The trigger name.
     */
    triggerName: string;
}

/**
 * Defines the response of a trigger subscription operation.
 */
export interface GetTriggerEventSubscriptionStatusResult {
    /**
     * Event Subscription Status.
     */
    readonly status: string;
    /**
     * Trigger name.
     */
    readonly triggerName: string;
}

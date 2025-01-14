// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The list of all devices in a resource and their eligibility status as a failover target device.
 */
export function listDeviceFailoverTars(args: ListDeviceFailoverTarsArgs, opts?: pulumi.InvokeOptions): Promise<ListDeviceFailoverTarsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storsimple/v20170601:listDeviceFailoverTars", {
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
        "sourceDeviceName": args.sourceDeviceName,
        "volumeContainers": args.volumeContainers,
    }, opts);
}

export interface ListDeviceFailoverTarsArgs {
    /**
     * The manager name
     */
    managerName: string;
    /**
     * The resource group name
     */
    resourceGroupName: string;
    /**
     * The source device name on which failover is performed.
     */
    sourceDeviceName: string;
    /**
     * The list of path IDs of the volume containers that needs to be failed-over, for which we want to fetch the eligible targets.
     */
    volumeContainers?: string[];
}

/**
 * The list of all devices in a resource and their eligibility status as a failover target device.
 */
export interface ListDeviceFailoverTarsResult {
    /**
     * The list of all the failover targets.
     */
    readonly value?: outputs.storsimple.v20170601.FailoverTargetResponse[];
}

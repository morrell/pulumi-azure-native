// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The volume.
 * API Version: 2017-06-01.
 */
export function getVolume(args: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storsimple:getVolume", {
        "deviceName": args.deviceName,
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
        "volumeContainerName": args.volumeContainerName,
        "volumeName": args.volumeName,
    }, opts);
}

export interface GetVolumeArgs {
    /**
     * The device name
     */
    deviceName: string;
    /**
     * The manager name
     */
    managerName: string;
    /**
     * The resource group name
     */
    resourceGroupName: string;
    /**
     * The volume container name.
     */
    volumeContainerName: string;
    /**
     * The volume name.
     */
    volumeName: string;
}

/**
 * The volume.
 */
export interface GetVolumeResult {
    /**
     * The IDs of the access control records, associated with the volume.
     */
    readonly accessControlRecordIds: string[];
    /**
     * The IDs of the backup policies, in which this volume is part of.
     */
    readonly backupPolicyIds: string[];
    /**
     * The backup status of the volume.
     */
    readonly backupStatus: string;
    /**
     * The path ID that uniquely identifies the object.
     */
    readonly id: string;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    readonly kind?: string;
    /**
     * The monitoring status of the volume.
     */
    readonly monitoringStatus: string;
    /**
     * The name of the object.
     */
    readonly name: string;
    /**
     * The operation status on the volume.
     */
    readonly operationStatus: string;
    /**
     * The size of the volume in bytes.
     */
    readonly sizeInBytes: number;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
    /**
     * The ID of the volume container, in which this volume is created.
     */
    readonly volumeContainerId: string;
    /**
     * The volume status.
     */
    readonly volumeStatus: string;
    /**
     * The type of the volume.
     */
    readonly volumeType: string;
}

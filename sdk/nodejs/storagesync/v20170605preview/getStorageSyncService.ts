// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Storage Sync Service object.
 */
export function getStorageSyncService(args: GetStorageSyncServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageSyncServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storagesync/v20170605preview:getStorageSyncService", {
        "resourceGroupName": args.resourceGroupName,
        "storageSyncServiceName": args.storageSyncServiceName,
    }, opts);
}

export interface GetStorageSyncServiceArgs {
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Name of Storage Sync Service resource.
     */
    storageSyncServiceName: string;
}

/**
 * Storage Sync Service object.
 */
export interface GetStorageSyncServiceResult {
    /**
     * The id of the resource.
     */
    readonly id: string;
    /**
     * The location of the resource.
     */
    readonly location?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * Storage Sync service status.
     */
    readonly storageSyncServiceStatus: number;
    /**
     * Storage Sync service Uid
     */
    readonly storageSyncServiceUid: string;
    /**
     * The tags of the resource.
     */
    readonly tags?: any;
    /**
     * The type of the resource
     */
    readonly type: string;
}

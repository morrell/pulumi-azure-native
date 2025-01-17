// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Restore Point details.
 */
export function getRestorePoint(args: GetRestorePointArgs, opts?: pulumi.InvokeOptions): Promise<GetRestorePointResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:compute/v20210401:getRestorePoint", {
        "resourceGroupName": args.resourceGroupName,
        "restorePointCollectionName": args.restorePointCollectionName,
        "restorePointName": args.restorePointName,
    }, opts);
}

export interface GetRestorePointArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * The name of the restore point collection.
     */
    restorePointCollectionName: string;
    /**
     * The name of the restore point.
     */
    restorePointName: string;
}

/**
 * Restore Point details.
 */
export interface GetRestorePointResult {
    /**
     * Gets the consistency mode for the restore point. Please refer to https://aka.ms/RestorePoints for more details.
     */
    readonly consistencyMode: string;
    /**
     * List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
     */
    readonly excludeDisks?: outputs.compute.v20210401.ApiEntityReferenceResponse[];
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Gets the provisioning details set by the server during Create restore point operation.
     */
    readonly provisioningDetails: outputs.compute.v20210401.RestorePointProvisioningDetailsResponse;
    /**
     * Gets the provisioning state of the restore point.
     */
    readonly provisioningState: string;
    /**
     * Gets the details of the VM captured at the time of the restore point creation.
     */
    readonly sourceMetadata: outputs.compute.v20210401.RestorePointSourceMetadataResponse;
    /**
     * Resource type
     */
    readonly type: string;
}

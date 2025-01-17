// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The access control record.
 */
export function getAccessControlRecord(args: GetAccessControlRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessControlRecordResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storsimple/v20170601:getAccessControlRecord", {
        "accessControlRecordName": args.accessControlRecordName,
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAccessControlRecordArgs {
    /**
     * Name of access control record to be fetched.
     */
    accessControlRecordName: string;
    /**
     * The manager name
     */
    managerName: string;
    /**
     * The resource group name
     */
    resourceGroupName: string;
}

/**
 * The access control record.
 */
export interface GetAccessControlRecordResult {
    /**
     * The path ID that uniquely identifies the object.
     */
    readonly id: string;
    /**
     * The iSCSI initiator name (IQN).
     */
    readonly initiatorName: string;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    readonly kind?: string;
    /**
     * The name of the object.
     */
    readonly name: string;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
    /**
     * The number of volumes using the access control record.
     */
    readonly volumeCount: number;
}

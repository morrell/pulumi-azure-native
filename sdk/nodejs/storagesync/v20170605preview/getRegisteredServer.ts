// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Registered Server resource.
 */
export function getRegisteredServer(args: GetRegisteredServerArgs, opts?: pulumi.InvokeOptions): Promise<GetRegisteredServerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storagesync/v20170605preview:getRegisteredServer", {
        "resourceGroupName": args.resourceGroupName,
        "serverId": args.serverId,
        "storageSyncServiceName": args.storageSyncServiceName,
    }, opts);
}

export interface GetRegisteredServerArgs {
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * GUID identifying the on-premises server.
     */
    serverId: string;
    /**
     * Name of Storage Sync Service resource.
     */
    storageSyncServiceName: string;
}

/**
 * Registered Server resource.
 */
export interface GetRegisteredServerResult {
    /**
     * Registered Server Agent Version
     */
    readonly agentVersion?: string;
    /**
     * Registered Server clusterId
     */
    readonly clusterId?: string;
    /**
     * Registered Server clusterName
     */
    readonly clusterName?: string;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Registered Server last heart beat
     */
    readonly lastHeartBeat?: string;
    /**
     * Registered Server lastWorkflowId
     */
    readonly lastWorkflowId?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Registered Server Provisioning State
     */
    readonly provisioningState?: string;
    /**
     * Registered Server Certificate
     */
    readonly serverCertificate?: string;
    /**
     * Registered Server serverId
     */
    readonly serverId?: string;
    /**
     * Registered Server Management Error Code
     */
    readonly serverManagementtErrorCode?: number;
    /**
     * Registered Server OS Version
     */
    readonly serverOSVersion?: string;
    /**
     * Registered Server serverRole
     */
    readonly serverRole?: string;
    /**
     * Registered Server storageSyncServiceUid
     */
    readonly storageSyncServiceUid?: string;
    /**
     * Resource type
     */
    readonly type: string;
}

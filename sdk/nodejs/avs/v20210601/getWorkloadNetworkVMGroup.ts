// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * NSX VM Group
 */
export function getWorkloadNetworkVMGroup(args: GetWorkloadNetworkVMGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkloadNetworkVMGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:avs/v20210601:getWorkloadNetworkVMGroup", {
        "privateCloudName": args.privateCloudName,
        "resourceGroupName": args.resourceGroupName,
        "vmGroupId": args.vmGroupId,
    }, opts);
}

export interface GetWorkloadNetworkVMGroupArgs {
    /**
     * Name of the private cloud
     */
    privateCloudName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * NSX VM Group identifier. Generally the same as the VM Group's display name
     */
    vmGroupId: string;
}

/**
 * NSX VM Group
 */
export interface GetWorkloadNetworkVMGroupResult {
    /**
     * Display name of the VM group.
     */
    readonly displayName?: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Virtual machine members of this group.
     */
    readonly members?: string[];
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state
     */
    readonly provisioningState: string;
    /**
     * NSX revision number.
     */
    readonly revision?: number;
    /**
     * VM Group status.
     */
    readonly status: string;
    /**
     * Resource type.
     */
    readonly type: string;
}

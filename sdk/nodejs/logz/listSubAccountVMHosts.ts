// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Response of a list VM Host Update Operation.
 * API Version: 2020-10-01-preview.
 */
export function listSubAccountVMHosts(args: ListSubAccountVMHostsArgs, opts?: pulumi.InvokeOptions): Promise<ListSubAccountVMHostsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:logz:listSubAccountVMHosts", {
        "monitorName": args.monitorName,
        "resourceGroupName": args.resourceGroupName,
        "subAccountName": args.subAccountName,
    }, opts);
}

export interface ListSubAccountVMHostsArgs {
    /**
     * Monitor resource name
     */
    monitorName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Sub Account resource name
     */
    subAccountName: string;
}

/**
 * Response of a list VM Host Update Operation.
 */
export interface ListSubAccountVMHostsResult {
    /**
     * Link to the next set of results, if any.
     */
    readonly nextLink?: string;
    /**
     * Response of a list vm host update operation.
     */
    readonly value?: outputs.logz.VMResourcesResponse[];
}

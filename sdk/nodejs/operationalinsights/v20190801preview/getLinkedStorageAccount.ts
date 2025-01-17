// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Linked storage accounts top level resource container.
 */
export function getLinkedStorageAccount(args: GetLinkedStorageAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetLinkedStorageAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:operationalinsights/v20190801preview:getLinkedStorageAccount", {
        "dataSourceType": args.dataSourceType,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetLinkedStorageAccountArgs {
    /**
     * Linked storage accounts type.
     */
    dataSourceType: string;
    /**
     * The name of the resource group to get. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Name of the Log Analytics Workspace that will contain the resource.
     */
    workspaceName: string;
}

/**
 * Linked storage accounts top level resource container.
 */
export interface GetLinkedStorageAccountResult {
    /**
     * Linked storage accounts type.
     */
    readonly dataSourceType: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Linked storage accounts resources ids.
     */
    readonly storageAccountIds?: string[];
    /**
     * Resource type.
     */
    readonly type: string;
}

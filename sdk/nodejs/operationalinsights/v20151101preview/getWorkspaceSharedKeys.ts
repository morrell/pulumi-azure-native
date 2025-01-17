// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The shared keys for a workspace.
 */
export function getWorkspaceSharedKeys(args: GetWorkspaceSharedKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceSharedKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:operationalinsights/v20151101preview:getWorkspaceSharedKeys", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceSharedKeysArgs {
    /**
     * The name of the resource group to get. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Name of the Log Analytics Workspace.
     */
    workspaceName: string;
}

/**
 * The shared keys for a workspace.
 */
export interface GetWorkspaceSharedKeysResult {
    /**
     * The primary shared key of a workspace.
     */
    readonly primarySharedKey?: string;
    /**
     * The secondary shared key of a workspace.
     */
    readonly secondarySharedKey?: string;
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The list of intermediate cluster code versions for an upgrade or downgrade. Or minimum and maximum upgradable version if no target was given
 */
export function listClusterUpgradableVersions(args: ListClusterUpgradableVersionsArgs, opts?: pulumi.InvokeOptions): Promise<ListClusterUpgradableVersionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:servicefabric/v20210601:listClusterUpgradableVersions", {
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
        "targetVersion": args.targetVersion,
    }, opts);
}

export interface ListClusterUpgradableVersionsArgs {
    /**
     * The name of the cluster resource.
     */
    clusterName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * The target code version.
     */
    targetVersion: string;
}

/**
 * The list of intermediate cluster code versions for an upgrade or downgrade. Or minimum and maximum upgradable version if no target was given
 */
export interface ListClusterUpgradableVersionsResult {
    readonly supportedPath?: string[];
}

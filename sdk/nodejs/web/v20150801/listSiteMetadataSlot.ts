// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * String dictionary resource
 */
export function listSiteMetadataSlot(args: ListSiteMetadataSlotArgs, opts?: pulumi.InvokeOptions): Promise<ListSiteMetadataSlotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:web/v20150801:listSiteMetadataSlot", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "slot": args.slot,
    }, opts);
}

export interface ListSiteMetadataSlotArgs {
    /**
     * Name of web app
     */
    name: string;
    /**
     * Name of resource group
     */
    resourceGroupName: string;
    /**
     * Name of web app slot. If not specified then will default to production slot.
     */
    slot: string;
}

/**
 * String dictionary resource
 */
export interface ListSiteMetadataSlotResult {
    /**
     * Resource Id
     */
    readonly id?: string;
    /**
     * Kind of resource
     */
    readonly kind?: string;
    /**
     * Resource Location
     */
    readonly location: string;
    /**
     * Resource Name
     */
    readonly name?: string;
    /**
     * Settings
     */
    readonly properties: {[key: string]: string};
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type?: string;
}

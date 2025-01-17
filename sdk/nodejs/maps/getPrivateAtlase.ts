// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An Azure resource which represents which will provision the ability to create private location data.
 * API Version: 2020-02-01-preview.
 */
export function getPrivateAtlase(args: GetPrivateAtlaseArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateAtlaseResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:maps:getPrivateAtlase", {
        "accountName": args.accountName,
        "privateAtlasName": args.privateAtlasName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateAtlaseArgs {
    /**
     * The name of the Maps Account.
     */
    accountName: string;
    /**
     * The name of the Private Atlas instance.
     */
    privateAtlasName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * An Azure resource which represents which will provision the ability to create private location data.
 */
export interface GetPrivateAtlaseResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The Private Atlas resource properties.
     */
    readonly properties: outputs.maps.PrivateAtlasPropertiesResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}

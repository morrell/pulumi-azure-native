// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Resource information.
 */
export function getResource(args: GetResourceArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:resources/v20160201:getResource", {
        "parentResourcePath": args.parentResourcePath,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
        "resourceProviderNamespace": args.resourceProviderNamespace,
        "resourceType": args.resourceType,
    }, opts);
}

export interface GetResourceArgs {
    /**
     * Resource identity.
     */
    parentResourcePath: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Resource identity.
     */
    resourceName: string;
    /**
     * Resource identity.
     */
    resourceProviderNamespace: string;
    /**
     * Resource identity.
     */
    resourceType: string;
}

/**
 * Resource information.
 */
export interface GetResourceResult {
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.resources.v20160201.IdentityResponse;
    /**
     * The kind of the resource.
     */
    readonly kind?: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * Id of the resource that manages this resource.
     */
    readonly managedBy?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The plan of the resource.
     */
    readonly plan?: outputs.resources.v20160201.PlanResponse;
    /**
     * The resource properties.
     */
    readonly properties: any;
    /**
     * The sku of the resource.
     */
    readonly sku?: outputs.resources.v20160201.SkuResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}

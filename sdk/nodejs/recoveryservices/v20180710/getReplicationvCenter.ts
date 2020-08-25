// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getReplicationvCenter(args: GetReplicationvCenterArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationvCenterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:recoveryservices/v20180710:getReplicationvCenter", {
        "fabricName": args.fabricName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetReplicationvCenterArgs {
    /**
     * Fabric name.
     */
    readonly fabricName: string;
    /**
     * vCenter name.
     */
    readonly name: string;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the recovery services vault.
     */
    readonly resourceName: string;
}

/**
 * vCenter definition.
 */
export interface GetReplicationvCenterResult {
    /**
     * Resource Location
     */
    readonly location?: string;
    /**
     * Resource Name
     */
    readonly name: string;
    /**
     * VCenter related data.
     */
    readonly properties: outputs.recoveryservices.v20180710.VCenterPropertiesResponse;
    /**
     * Resource Type
     */
    readonly type: string;
}
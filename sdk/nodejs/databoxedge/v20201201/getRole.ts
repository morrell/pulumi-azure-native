// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Compute role.
 */
/** @deprecated Please use one of the variants: CloudEdgeManagementRole, IoTRole, KubernetesRole, MECRole. */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    pulumi.log.warn("getRole is deprecated: Please use one of the variants: CloudEdgeManagementRole, IoTRole, KubernetesRole, MECRole.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:databoxedge/v20201201:getRole", {
        "deviceName": args.deviceName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRoleArgs {
    /**
     * The device name.
     */
    deviceName: string;
    /**
     * The role name.
     */
    name: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
}

/**
 * Compute role.
 */
export interface GetRoleResult {
    /**
     * The path ID that uniquely identifies the object.
     */
    readonly id: string;
    /**
     * Role type.
     */
    readonly kind: string;
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * Role configured on ASE resource
     */
    readonly systemData: outputs.databoxedge.v20201201.SystemDataResponse;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}

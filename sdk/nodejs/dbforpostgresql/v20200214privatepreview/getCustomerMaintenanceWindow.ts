// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getCustomerMaintenanceWindow(args: GetCustomerMaintenanceWindowArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomerMaintenanceWindowResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:dbforpostgresql/v20200214privatepreview:getCustomerMaintenanceWindow", {
        "maintenanceWindowName": args.maintenanceWindowName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetCustomerMaintenanceWindowArgs {
    /**
     * The name of the maintenance window.
     */
    readonly maintenanceWindowName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
}

/**
 * Represents a server firewall rule.
 */
export interface GetCustomerMaintenanceWindowResult {
    /**
     * The day of week of the customer maintenance window to start
     */
    readonly dayOfWeek: number;
    /**
     * The duration of the customer maintenance window to run.
     */
    readonly durationInMinutes?: number;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The starting hour of the customer maintenance window.
     */
    readonly startHour: number;
    /**
     * The starting minutes of the customer maintenance window.
     */
    readonly startMinute: number;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
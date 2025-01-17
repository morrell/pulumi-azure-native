// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Response for list of user's role for Logz.io account.
 * API Version: 2020-10-01-preview.
 */
export function listMonitorUserRoles(args: ListMonitorUserRolesArgs, opts?: pulumi.InvokeOptions): Promise<ListMonitorUserRolesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:logz:listMonitorUserRoles", {
        "emailAddress": args.emailAddress,
        "monitorName": args.monitorName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListMonitorUserRolesArgs {
    /**
     * Email of the user used by Logz for contacting them if needed
     */
    emailAddress?: string;
    /**
     * Monitor resource name
     */
    monitorName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Response for list of user's role for Logz.io account.
 */
export interface ListMonitorUserRolesResult {
    /**
     * Link to the next set of results, if any.
     */
    readonly nextLink?: string;
    /**
     * List of user roles for Logz.io account.
     */
    readonly value?: outputs.logz.UserRoleResponseResponse[];
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getBatchAccount(args: GetBatchAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetBatchAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:batch/v20151201:getBatchAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetBatchAccountArgs {
    /**
     * The name of the account.
     */
    readonly name: string;
    /**
     * The name of the resource group that contains the Batch account.
     */
    readonly resourceGroupName: string;
}

/**
 * Contains information about an Azure Batch account.
 */
export interface GetBatchAccountResult {
    /**
     * The endpoint used by this account to interact with the Batch services.
     */
    readonly accountEndpoint: string;
    /**
     * The active job and job schedule quota for this Batch account.
     */
    readonly activeJobAndJobScheduleQuota: number;
    /**
     * The properties and status of any auto storage account associated with the account.
     */
    readonly autoStorage?: outputs.batch.v20151201.AutoStoragePropertiesResponse;
    /**
     * The core quota for this Batch account.
     */
    readonly coreQuota: number;
    /**
     * The location of the resource
     */
    readonly location?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The pool quota for this Batch account.
     */
    readonly poolQuota: number;
    /**
     * The provisioned state of the resource
     */
    readonly provisioningState?: string;
    /**
     * The tags of the resource
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource
     */
    readonly type: string;
}
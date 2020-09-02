// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getQuery(args: GetQueryArgs, opts?: pulumi.InvokeOptions): Promise<GetQueryResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:insights/v20190901preview:getQuery", {
        "id": args.id,
        "queryPackName": args.queryPackName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetQueryArgs {
    /**
     * The id of a specific query defined in the Log Analytics QueryPack
     */
    readonly id: string;
    /**
     * The name of the Log Analytics QueryPack resource.
     */
    readonly queryPackName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * A Log Analytics QueryPack-Query definition.
 */
export interface GetQueryResult {
    /**
     * Object Id of user creating the query.
     */
    readonly author: string;
    /**
     * Body of the query.
     */
    readonly body: string;
    /**
     * Description of the query.
     */
    readonly description?: string;
    /**
     * Unique display name for your query within the Query Pack.
     */
    readonly displayName: string;
    /**
     * Azure resource name
     */
    readonly name: string;
    /**
     * Additional properties that can be set for the query.
     */
    readonly properties: {[key: string]: any};
    /**
     * The related metadata items for the function.
     */
    readonly related?: outputs.insights.v20190901preview.LogAnalyticsQueryPackQueryPropertiesResponseRelated;
    /**
     * Read only system data
     */
    readonly systemData: outputs.insights.v20190901preview.SystemDataResponse;
    /**
     * Tags associated with the query.
     */
    readonly tags?: {[key: string]: string[]};
    /**
     * Creation Date for the Log Analytics Query, in ISO 8601 format.
     */
    readonly timeCreated: string;
    /**
     * Last modified date of the Log Analytics Query, in ISO 8601 format.
     */
    readonly timeModified: string;
    /**
     * Azure resource type
     */
    readonly type: string;
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Represents AATP (Azure Advanced Threat Protection) data connector.
 */
export function getAATPDataConnector(args: GetAATPDataConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetAATPDataConnectorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:securityinsights/v20200101:getAATPDataConnector", {
        "dataConnectorId": args.dataConnectorId,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetAATPDataConnectorArgs {
    /**
     * Connector ID
     */
    dataConnectorId: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the workspace.
     */
    workspaceName: string;
}

/**
 * Represents AATP (Azure Advanced Threat Protection) data connector.
 */
export interface GetAATPDataConnectorResult {
    /**
     * The available data types for the connector.
     */
    readonly dataTypes?: outputs.securityinsights.v20200101.AlertsDataTypeOfDataConnectorResponse;
    /**
     * Etag of the azure resource
     */
    readonly etag?: string;
    /**
     * Azure resource Id
     */
    readonly id: string;
    /**
     * The kind of the data connector
     * Expected value is 'AzureAdvancedThreatProtection'.
     */
    readonly kind: "AzureAdvancedThreatProtection";
    /**
     * Azure resource name
     */
    readonly name: string;
    /**
     * The tenant id to connect to, and get the data from.
     */
    readonly tenantId?: string;
    /**
     * Azure resource type
     */
    readonly type: string;
}

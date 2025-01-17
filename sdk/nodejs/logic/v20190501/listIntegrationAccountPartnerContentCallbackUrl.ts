// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The workflow trigger callback URL.
 */
export function listIntegrationAccountPartnerContentCallbackUrl(args: ListIntegrationAccountPartnerContentCallbackUrlArgs, opts?: pulumi.InvokeOptions): Promise<ListIntegrationAccountPartnerContentCallbackUrlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:logic/v20190501:listIntegrationAccountPartnerContentCallbackUrl", {
        "integrationAccountName": args.integrationAccountName,
        "keyType": args.keyType,
        "notAfter": args.notAfter,
        "partnerName": args.partnerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListIntegrationAccountPartnerContentCallbackUrlArgs {
    /**
     * The integration account name.
     */
    integrationAccountName: string;
    /**
     * The key type.
     */
    keyType?: string | enums.logic.v20190501.KeyType;
    /**
     * The expiry time.
     */
    notAfter?: string;
    /**
     * The integration account partner name.
     */
    partnerName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
}

/**
 * The workflow trigger callback URL.
 */
export interface ListIntegrationAccountPartnerContentCallbackUrlResult {
    /**
     * Gets the workflow trigger callback URL base path.
     */
    readonly basePath: string;
    /**
     * Gets the workflow trigger callback URL HTTP method.
     */
    readonly method: string;
    /**
     * Gets the workflow trigger callback URL query parameters.
     */
    readonly queries?: outputs.logic.v20190501.WorkflowTriggerListCallbackUrlQueriesResponse;
    /**
     * Gets the workflow trigger callback URL relative path.
     */
    readonly relativePath: string;
    /**
     * Gets the workflow trigger callback URL relative path parameters.
     */
    readonly relativePathParameters?: string[];
    /**
     * Gets the workflow trigger callback URL.
     */
    readonly value: string;
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Base definition for datastore secrets.
 * API Version: 2021-03-01-preview.
 */
export function listDatastoreSecrets(args: ListDatastoreSecretsArgs, opts?: pulumi.InvokeOptions): Promise<ListDatastoreSecretsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:machinelearningservices:listDatastoreSecrets", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListDatastoreSecretsArgs {
    /**
     * Datastore name.
     */
    name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Name of Azure Machine Learning workspace.
     */
    workspaceName: string;
}

/**
 * Base definition for datastore secrets.
 */
export interface ListDatastoreSecretsResult {
    /**
     * Credential type used to authentication with storage.
     */
    readonly secretsType: string;
}

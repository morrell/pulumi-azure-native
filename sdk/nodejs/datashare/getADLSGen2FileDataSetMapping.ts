// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An ADLS Gen2 file data set mapping.
 * API Version: 2020-09-01.
 */
export function getADLSGen2FileDataSetMapping(args: GetADLSGen2FileDataSetMappingArgs, opts?: pulumi.InvokeOptions): Promise<GetADLSGen2FileDataSetMappingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:datashare:getADLSGen2FileDataSetMapping", {
        "accountName": args.accountName,
        "dataSetMappingName": args.dataSetMappingName,
        "resourceGroupName": args.resourceGroupName,
        "shareSubscriptionName": args.shareSubscriptionName,
    }, opts);
}

export interface GetADLSGen2FileDataSetMappingArgs {
    /**
     * The name of the share account.
     */
    accountName: string;
    /**
     * The name of the dataSetMapping.
     */
    dataSetMappingName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
    /**
     * The name of the shareSubscription.
     */
    shareSubscriptionName: string;
}

/**
 * An ADLS Gen2 file data set mapping.
 */
export interface GetADLSGen2FileDataSetMappingResult {
    /**
     * The id of the source data set.
     */
    readonly dataSetId: string;
    /**
     * Gets the status of the data set mapping.
     */
    readonly dataSetMappingStatus: string;
    /**
     * File path within the file system.
     */
    readonly filePath: string;
    /**
     * File system to which the file belongs.
     */
    readonly fileSystem: string;
    /**
     * The resource id of the azure resource
     */
    readonly id: string;
    /**
     * Kind of data set mapping.
     * Expected value is 'AdlsGen2File'.
     */
    readonly kind: "AdlsGen2File";
    /**
     * Name of the azure resource
     */
    readonly name: string;
    /**
     * Type of output file
     */
    readonly outputType?: string;
    /**
     * Provisioning state of the data set mapping.
     */
    readonly provisioningState: string;
    /**
     * Resource group of storage account.
     */
    readonly resourceGroup: string;
    /**
     * Storage account name of the source data set.
     */
    readonly storageAccountName: string;
    /**
     * Subscription id of storage account.
     */
    readonly subscriptionId: string;
    /**
     * System Data of the Azure resource.
     */
    readonly systemData: outputs.datashare.SystemDataResponse;
    /**
     * Type of the azure resource
     */
    readonly type: string;
}

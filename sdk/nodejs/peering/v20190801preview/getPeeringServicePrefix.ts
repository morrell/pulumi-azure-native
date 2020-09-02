// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getPeeringServicePrefix(args: GetPeeringServicePrefixArgs, opts?: pulumi.InvokeOptions): Promise<GetPeeringServicePrefixResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:peering/v20190801preview:getPeeringServicePrefix", {
        "peeringServiceName": args.peeringServiceName,
        "prefixName": args.prefixName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPeeringServicePrefixArgs {
    /**
     * The peering service name.
     */
    readonly peeringServiceName: string;
    /**
     * The prefix name.
     */
    readonly prefixName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * The peering service prefix class.
 */
export interface GetPeeringServicePrefixResult {
    /**
     * The prefix learned type
     */
    readonly learnedType?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * Valid route prefix
     */
    readonly prefix?: string;
    /**
     * The prefix validation state
     */
    readonly prefixValidationState?: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
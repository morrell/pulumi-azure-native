// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Description of topic resource.
 */
export function getNamespaceNetworkRuleSet(args: GetNamespaceNetworkRuleSetArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceNetworkRuleSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:eventhub/v20180101preview:getNamespaceNetworkRuleSet", {
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNamespaceNetworkRuleSetArgs {
    /**
     * The Namespace name
     */
    namespaceName: string;
    /**
     * Name of the resource group within the azure subscription.
     */
    resourceGroupName: string;
}

/**
 * Description of topic resource.
 */
export interface GetNamespaceNetworkRuleSetResult {
    /**
     * Default Action for Network Rule Set
     */
    readonly defaultAction?: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * List of IpRules
     */
    readonly ipRules?: outputs.eventhub.v20180101preview.NWRuleSetIpRulesResponse[];
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Value that indicates whether Trusted Service Access is Enabled or not.
     */
    readonly trustedServiceAccessEnabled?: boolean;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * List VirtualNetwork Rules
     */
    readonly virtualNetworkRules?: outputs.eventhub.v20180101preview.NWRuleSetVirtualNetworkRulesResponse[];
}

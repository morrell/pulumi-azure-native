// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualNetworkRule(args: GetVirtualNetworkRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:dbforpostgresql/v20171201preview:getVirtualNetworkRule", {
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
        "virtualNetworkRuleName": args.virtualNetworkRuleName,
    }, opts);
}

export interface GetVirtualNetworkRuleArgs {
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
    /**
     * The name of the virtual network rule.
     */
    readonly virtualNetworkRuleName: string;
}

/**
 * A virtual network rule.
 */
export interface GetVirtualNetworkRuleResult {
    /**
     * Create firewall rule before the virtual network has vnet service endpoint enabled.
     */
    readonly ignoreMissingVnetServiceEndpoint?: boolean;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Virtual Network Rule State
     */
    readonly state: string;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The ARM resource id of the virtual network subnet.
     */
    readonly virtualNetworkSubnetId: string;
}
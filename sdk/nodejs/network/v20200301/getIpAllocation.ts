// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getIpAllocation(args: GetIpAllocationArgs, opts?: pulumi.InvokeOptions): Promise<GetIpAllocationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20200301:getIpAllocation", {
        "expand": args.expand,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetIpAllocationArgs {
    /**
     * Expands referenced resources.
     */
    readonly expand?: string;
    /**
     * The name of the IpAllocation.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * IpAllocation resource.
 */
export interface GetIpAllocationResult {
    /**
     * IpAllocation tags.
     */
    readonly allocationTags?: {[key: string]: string};
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * The IPAM allocation ID.
     */
    readonly ipamAllocationId?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The address prefix for the IpAllocation.
     */
    readonly prefix?: string;
    /**
     * The address prefix length for the IpAllocation.
     */
    readonly prefixLength?: number;
    /**
     * The address prefix Type for the IpAllocation.
     */
    readonly prefixType?: string;
    /**
     * The Subnet that using the prefix of this IpAllocation resource.
     */
    readonly subnet: outputs.network.v20200301.SubResourceResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The VirtualNetwork that using the prefix of this IpAllocation resource.
     */
    readonly virtualNetwork: outputs.network.v20200301.SubResourceResponse;
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualHub(args: GetVirtualHubArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualHubResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20180701:getVirtualHub", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetVirtualHubArgs {
    /**
     * The name of the VirtualHub.
     */
    readonly name: string;
    /**
     * The resource group name of the VirtualHub.
     */
    readonly resourceGroupName: string;
}

/**
 * VirtualHub Resource.
 */
export interface GetVirtualHubResult {
    /**
     * Address-prefix for this VirtualHub.
     */
    readonly addressPrefix?: string;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * list of all vnet connections with this VirtualHub.
     */
    readonly hubVirtualNetworkConnections?: outputs.network.v20180701.HubVirtualNetworkConnectionResponse[];
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The VirtualWAN to which the VirtualHub belongs
     */
    readonly virtualWan?: outputs.network.v20180701.SubResourceResponse;
}
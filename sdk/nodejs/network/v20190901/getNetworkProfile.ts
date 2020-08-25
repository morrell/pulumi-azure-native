// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getNetworkProfile(args: GetNetworkProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkProfileResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20190901:getNetworkProfile", {
        "expand": args.expand,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNetworkProfileArgs {
    /**
     * Expands referenced resources.
     */
    readonly expand?: string;
    /**
     * The name of the public IP prefix.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Network profile resource.
 */
export interface GetNetworkProfileResult {
    /**
     * List of chid container network interface configurations.
     */
    readonly containerNetworkInterfaceConfigurations?: outputs.network.v20190901.ContainerNetworkInterfaceConfigurationResponse[];
    /**
     * List of child container network interfaces.
     */
    readonly containerNetworkInterfaces: outputs.network.v20190901.ContainerNetworkInterfaceResponse[];
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the network profile resource.
     */
    readonly provisioningState: string;
    /**
     * The resource GUID property of the network profile resource.
     */
    readonly resourceGuid: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
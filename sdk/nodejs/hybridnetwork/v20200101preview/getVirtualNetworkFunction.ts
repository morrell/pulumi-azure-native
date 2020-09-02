// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualNetworkFunction(args: GetVirtualNetworkFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkFunctionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:hybridnetwork/v20200101preview:getVirtualNetworkFunction", {
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkFunctionName": args.virtualNetworkFunctionName,
    }, opts);
}

export interface GetVirtualNetworkFunctionArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of hybrid network virtual network function resource.
     */
    readonly virtualNetworkFunctionName: string;
}

/**
 * Hybrid network virtual network function resource response.
 */
export interface GetVirtualNetworkFunctionResult {
    /**
     * The reference to the hybrid network device.
     */
    readonly device?: outputs.hybridnetwork.v20200101preview.SubResourceResponse;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * The resource URI of the managed application.
     */
    readonly managedApplication: outputs.hybridnetwork.v20200101preview.SubResourceResponse;
    /**
     * The parameters for the managed application.
     */
    readonly managedApplicationParameters?: {[key: string]: any};
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the hybrid network virtual network function resource.
     */
    readonly provisioningState: string;
    /**
     * The service key for the virtual network function resource.
     */
    readonly serviceKey: string;
    /**
     * The sku name for the hybrid network virtual network function.
     */
    readonly skuName?: string;
    /**
     * The sku type for the hybrid network virtual network function.
     */
    readonly skuType: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The vendor name for the hybrid network virtual network function.
     */
    readonly vendorName?: string;
    /**
     * The vendor provisioning state for the virtual network function resource.
     */
    readonly vendorProvisioningState: string;
    /**
     * The virtual network function configurations from the user.
     */
    readonly virtualNetworkFunctionUserConfigurations?: outputs.hybridnetwork.v20200101preview.VirtualNetworkFunctionUserConfigurationResponse[];
}
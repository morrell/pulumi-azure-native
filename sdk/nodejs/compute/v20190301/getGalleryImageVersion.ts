// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Specifies information about the gallery Image Version that you want to create or update.
 */
export function getGalleryImageVersion(args: GetGalleryImageVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetGalleryImageVersionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:compute/v20190301:getGalleryImageVersion", {
        "expand": args.expand,
        "galleryImageName": args.galleryImageName,
        "galleryImageVersionName": args.galleryImageVersionName,
        "galleryName": args.galleryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGalleryImageVersionArgs {
    /**
     * The expand expression to apply on the operation.
     */
    expand?: string;
    /**
     * The name of the gallery Image Definition in which the Image Version resides.
     */
    galleryImageName: string;
    /**
     * The name of the gallery Image Version to be retrieved.
     */
    galleryImageVersionName: string;
    /**
     * The name of the Shared Image Gallery in which the Image Definition resides.
     */
    galleryName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * Specifies information about the gallery Image Version that you want to create or update.
 */
export interface GetGalleryImageVersionResult {
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * The publishing profile of a gallery Image Version.
     */
    readonly publishingProfile: outputs.compute.v20190301.GalleryImageVersionPublishingProfileResponse;
    /**
     * This is the replication status of the gallery Image Version.
     */
    readonly replicationStatus: outputs.compute.v20190301.ReplicationStatusResponse;
    /**
     * This is the storage profile of a Gallery Image Version.
     */
    readonly storageProfile: outputs.compute.v20190301.GalleryImageVersionStorageProfileResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}

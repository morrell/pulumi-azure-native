// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getApplicationDefinition(args: GetApplicationDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:solutions/v20170901:getApplicationDefinition", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationDefinitionArgs {
    /**
     * The name of the managed application definition.
     */
    readonly name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Information about managed application definition.
 */
export interface GetApplicationDefinitionResult {
    /**
     * The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
     */
    readonly artifacts?: outputs.solutions.v20170901.ApplicationArtifactResponse[];
    /**
     * The managed application provider authorizations.
     */
    readonly authorizations: outputs.solutions.v20170901.ApplicationProviderAuthorizationResponse[];
    /**
     * The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
     */
    readonly createUiDefinition?: {[key: string]: any};
    /**
     * The managed application definition description.
     */
    readonly description?: string;
    /**
     * The managed application definition display name.
     */
    readonly displayName?: string;
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.solutions.v20170901.IdentityResponse;
    /**
     * A value indicating whether the package is enabled or not.
     */
    readonly isEnabled?: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * The managed application lock level.
     */
    readonly lockLevel: string;
    /**
     * The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
     */
    readonly mainTemplate?: {[key: string]: any};
    /**
     * ID of the resource that manages this resource.
     */
    readonly managedBy?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The managed application definition package file Uri. Use this element
     */
    readonly packageFileUri?: string;
    /**
     * The SKU of the resource.
     */
    readonly sku?: outputs.solutions.v20170901.SkuResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Certificate signing request object
 */
export function getCertificateCsr(args: GetCertificateCsrArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateCsrResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:web/v20150801:getCertificateCsr", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCertificateCsrArgs {
    /**
     * Name of the certificate.
     */
    name: string;
    /**
     * Name of the resource group
     */
    resourceGroupName: string;
}

/**
 * Certificate signing request object
 */
export interface GetCertificateCsrResult {
    /**
     * Actual CSR string created
     */
    readonly csrString?: string;
    /**
     * Distinguished name of certificate to be created
     */
    readonly distinguishedName?: string;
    /**
     * Hosting environment
     */
    readonly hostingEnvironment?: string;
    /**
     * Resource Id
     */
    readonly id?: string;
    /**
     * Kind of resource
     */
    readonly kind?: string;
    /**
     * Resource Location
     */
    readonly location: string;
    /**
     * Resource Name
     */
    readonly name?: string;
    /**
     * PFX password
     */
    readonly password?: string;
    /**
     * PFX certificate of created certificate
     */
    readonly pfxBlob?: string;
    /**
     * Hash of the certificates public key
     */
    readonly publicKeyHash?: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type?: string;
}

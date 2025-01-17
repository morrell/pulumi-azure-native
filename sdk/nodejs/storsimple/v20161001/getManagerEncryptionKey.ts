// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * This class can be used as the Type for any secret entity represented as Value, ValueCertificateThumbprint, EncryptionAlgorithm. In this case, "Value" is a secret and the "valueThumbprint" represents the certificate thumbprint of the value. The algorithm field is mainly for future usage to potentially allow different entities encrypted using different algorithms.
 */
export function getManagerEncryptionKey(args: GetManagerEncryptionKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetManagerEncryptionKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storsimple/v20161001:getManagerEncryptionKey", {
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagerEncryptionKeyArgs {
    /**
     * The manager name
     */
    managerName: string;
    /**
     * The resource group name
     */
    resourceGroupName: string;
}

/**
 * This class can be used as the Type for any secret entity represented as Value, ValueCertificateThumbprint, EncryptionAlgorithm. In this case, "Value" is a secret and the "valueThumbprint" represents the certificate thumbprint of the value. The algorithm field is mainly for future usage to potentially allow different entities encrypted using different algorithms.
 */
export interface GetManagerEncryptionKeyResult {
    /**
     * Algorithm used to encrypt "Value"
     */
    readonly encryptionAlgorithm: string;
    /**
     * The value of the secret itself. If the secret is in plaintext or null then EncryptionAlgorithm will be none
     */
    readonly value: string;
    /**
     * Thumbprint cert that was used to encrypt "Value"
     */
    readonly valueCertificateThumbprint?: string;
}

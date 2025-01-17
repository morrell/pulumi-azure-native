// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Registration definition.
 */
export function getRegistrationDefinition(args: GetRegistrationDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistrationDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:managedservices/v20190401preview:getRegistrationDefinition", {
        "registrationDefinitionId": args.registrationDefinitionId,
        "scope": args.scope,
    }, opts);
}

export interface GetRegistrationDefinitionArgs {
    /**
     * Guid of the registration definition.
     */
    registrationDefinitionId: string;
    /**
     * Scope of the resource.
     */
    scope: string;
}

/**
 * Registration definition.
 */
export interface GetRegistrationDefinitionResult {
    /**
     * Fully qualified path of the registration definition.
     */
    readonly id: string;
    /**
     * Name of the registration definition.
     */
    readonly name: string;
    /**
     * Plan details for the managed services.
     */
    readonly plan?: outputs.managedservices.v20190401preview.PlanResponse;
    /**
     * Properties of a registration definition.
     */
    readonly properties: outputs.managedservices.v20190401preview.RegistrationDefinitionPropertiesResponse;
    /**
     * Type of the resource.
     */
    readonly type: string;
}

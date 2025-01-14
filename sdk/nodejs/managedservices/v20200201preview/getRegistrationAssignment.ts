// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The registration assignment.
 */
export function getRegistrationAssignment(args: GetRegistrationAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistrationAssignmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:managedservices/v20200201preview:getRegistrationAssignment", {
        "expandRegistrationDefinition": args.expandRegistrationDefinition,
        "registrationAssignmentId": args.registrationAssignmentId,
        "scope": args.scope,
    }, opts);
}

export interface GetRegistrationAssignmentArgs {
    /**
     * The flag indicating whether to return the registration definition details along with the registration assignment details.
     */
    expandRegistrationDefinition?: boolean;
    /**
     * The GUID of the registration assignment.
     */
    registrationAssignmentId: string;
    /**
     * The scope of the resource.
     */
    scope: string;
}

/**
 * The registration assignment.
 */
export interface GetRegistrationAssignmentResult {
    /**
     * The fully qualified path of the registration assignment.
     */
    readonly id: string;
    /**
     * The name of the registration assignment.
     */
    readonly name: string;
    /**
     * The properties of a registration assignment.
     */
    readonly properties: outputs.managedservices.v20200201preview.RegistrationAssignmentPropertiesResponse;
    /**
     * The type of the Azure resource (Microsoft.ManagedServices/registrationAssignments).
     */
    readonly type: string;
}

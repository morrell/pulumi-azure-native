// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getAutomation(args: GetAutomationArgs, opts?: pulumi.InvokeOptions): Promise<GetAutomationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:security/v20190101preview:getAutomation", {
        "automationName": args.automationName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAutomationArgs {
    /**
     * The security automation name.
     */
    readonly automationName: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * The security automation resource.
 */
export interface GetAutomationResult {
    /**
     * A collection of the actions which are triggered if all the configured rules evaluations, within at least one rule set, are true.
     */
    readonly actions?: outputs.security.v20190101preview.AutomationActionResponse[];
    /**
     * The security automation description.
     */
    readonly description?: string;
    /**
     * Entity tag is used for comparing two or more entities from the same requested resource.
     */
    readonly etag?: string;
    /**
     * Indicates whether the security automation is enabled.
     */
    readonly isEnabled?: boolean;
    /**
     * Kind of the resource
     */
    readonly kind?: string;
    /**
     * Location where the resource is stored
     */
    readonly location?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * A collection of scopes on which the security automations logic is applied. Supported scopes are the subscription itself or a resource group under that subscription. The automation will only apply on defined scopes.
     */
    readonly scopes?: outputs.security.v20190101preview.AutomationScopeResponse[];
    /**
     * A collection of the source event types which evaluate the security automation set of rules.
     */
    readonly sources?: outputs.security.v20190101preview.AutomationSourceResponse[];
    /**
     * A list of key value pairs that describe the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
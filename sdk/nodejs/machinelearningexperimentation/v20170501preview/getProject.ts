// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:machinelearningexperimentation/v20170501preview:getProject", {
        "accountName": args.accountName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetProjectArgs {
    /**
     * The name of the machine learning team account.
     */
    readonly accountName: string;
    /**
     * The name of the machine learning project under a team account workspace.
     */
    readonly projectName: string;
    /**
     * The name of the resource group to which the machine learning team account belongs.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the machine learning team account workspace.
     */
    readonly workspaceName: string;
}

/**
 * An object that represents a machine learning project.
 */
export interface GetProjectResult {
    /**
     * The immutable id of the team account which contains this project.
     */
    readonly accountId: string;
    /**
     * The creation date of the project in ISO8601 format.
     */
    readonly creationDate: string;
    /**
     * The description of this project.
     */
    readonly description?: string;
    /**
     * The friendly name for this project.
     */
    readonly friendlyName: string;
    /**
     * The reference to git repo for this project.
     */
    readonly gitrepo?: string;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    readonly location: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The immutable id of this project.
     */
    readonly projectId: string;
    /**
     * The current deployment state of project resource. The provisioningState is to indicate states for resource provisioning.
     */
    readonly provisioningState: string;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * The immutable id of the workspace which contains this project.
     */
    readonly workspaceId: string;
}
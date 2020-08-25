// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listTaskDetails(args: ListTaskDetailsArgs, opts?: pulumi.InvokeOptions): Promise<ListTaskDetailsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:containerregistry/v20180901:listTaskDetails", {
        "name": args.name,
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListTaskDetailsArgs {
    /**
     * The name of the container registry task.
     */
    readonly name: string;
    /**
     * The name of the container registry.
     */
    readonly registryName: string;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * The task that has the ARM resource and task properties. 
 * The task will have all information to schedule a run against it.
 */
export interface ListTaskDetailsResult {
    /**
     * The machine configuration of the run agent.
     */
    readonly agentConfiguration?: outputs.containerregistry.v20180901.AgentPropertiesResponse;
    /**
     * The creation date of task.
     */
    readonly creationDate: string;
    /**
     * The properties that describes a set of credentials that will be used when this run is invoked.
     */
    readonly credentials?: outputs.containerregistry.v20180901.CredentialsResponse;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    readonly location: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The platform properties against which the run has to happen.
     */
    readonly platform: outputs.containerregistry.v20180901.PlatformPropertiesResponse;
    /**
     * The provisioning state of the task.
     */
    readonly provisioningState: string;
    /**
     * The current status of task.
     */
    readonly status?: string;
    /**
     * The properties of a task step.
     */
    readonly step: outputs.containerregistry.v20180901.TaskStepPropertiesResponse;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Run timeout in seconds.
     */
    readonly timeout?: number;
    /**
     * The properties that describe all triggers for the task.
     */
    readonly trigger?: outputs.containerregistry.v20180901.TriggerPropertiesResponse;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
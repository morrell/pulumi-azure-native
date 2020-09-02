// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * This type describes an application resource.
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicefabricmesh/v20180901preview:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * Internal - used by Visual Studio to setup the debugging session on the local development environment.
     */
    public readonly debugParams!: pulumi.Output<string | undefined>;
    /**
     * User readable description of the application.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Describes the diagnostics definition and usage for an application resource.
     */
    public readonly diagnostics!: pulumi.Output<outputs.servicefabricmesh.v20180901preview.DiagnosticsDescriptionResponse | undefined>;
    /**
     * Describes the health state of an application resource.
     */
    public /*out*/ readonly healthState!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * State of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Names of the services in the application.
     */
    public /*out*/ readonly serviceNames!: pulumi.Output<string[]>;
    /**
     * Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
     */
    public readonly services!: pulumi.Output<outputs.servicefabricmesh.v20180901preview.ServiceResourceDescriptionResponse[] | undefined>;
    /**
     * Status of the application.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Gives additional information about the current status of the application.
     */
    public /*out*/ readonly statusDetails!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
     */
    public /*out*/ readonly unhealthyEvaluation!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ApplicationArgs | undefined;
            if (!args || args.applicationResourceName === undefined) {
                throw new Error("Missing required property 'applicationResourceName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationResourceName"] = args ? args.applicationResourceName : undefined;
            inputs["debugParams"] = args ? args.debugParams : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["diagnostics"] = args ? args.diagnostics : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["services"] = args ? args.services : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["healthState"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["serviceNames"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusDetails"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["unhealthyEvaluation"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:servicefabricmesh/v20180701preview:Application" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The identity of the application.
     */
    readonly applicationResourceName: pulumi.Input<string>;
    /**
     * Internal - used by Visual Studio to setup the debugging session on the local development environment.
     */
    readonly debugParams?: pulumi.Input<string>;
    /**
     * User readable description of the application.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Describes the diagnostics definition and usage for an application resource.
     */
    readonly diagnostics?: pulumi.Input<inputs.servicefabricmesh.v20180901preview.DiagnosticsDescription>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * Azure resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
     */
    readonly services?: pulumi.Input<pulumi.Input<inputs.servicefabricmesh.v20180901preview.ServiceResourceDescription>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
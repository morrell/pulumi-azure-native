// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * IoT Security solution configuration and resource information.
 */
export class IotSecuritySolution extends pulumi.CustomResource {
    /**
     * Get an existing IotSecuritySolution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IotSecuritySolution {
        return new IotSecuritySolution(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:security/v20190801:IotSecuritySolution';

    /**
     * Returns true if the given object is an instance of IotSecuritySolution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IotSecuritySolution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IotSecuritySolution.__pulumiType;
    }

    /**
     * List of resources that were automatically discovered as relevant to the security solution.
     */
    public /*out*/ readonly autoDiscoveredResources!: pulumi.Output<string[]>;
    /**
     * Disabled data sources. Disabling these data sources compromises the system.
     */
    public readonly disabledDataSources!: pulumi.Output<string[] | undefined>;
    /**
     * Resource display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * List of additional options for exporting to workspace data.
     */
    public readonly export!: pulumi.Output<string[] | undefined>;
    /**
     * IoT Hub resource IDs
     */
    public readonly iotHubs!: pulumi.Output<string[]>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of the configuration status for each recommendation type.
     */
    public readonly recommendationsConfiguration!: pulumi.Output<outputs.security.v20190801.RecommendationConfigurationPropertiesResponse[] | undefined>;
    /**
     * Status of the IoT Security solution.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Unmasked IP address logging status
     */
    public readonly unmaskedIpLoggingStatus!: pulumi.Output<string | undefined>;
    /**
     * Properties of the IoT Security solution's user defined resources.
     */
    public readonly userDefinedResources!: pulumi.Output<outputs.security.v20190801.UserDefinedResourcesPropertiesResponse | undefined>;
    /**
     * Workspace resource ID
     */
    public readonly workspace!: pulumi.Output<string | undefined>;

    /**
     * Create a IotSecuritySolution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IotSecuritySolutionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IotSecuritySolutionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as IotSecuritySolutionArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.iotHubs === undefined) {
                throw new Error("Missing required property 'iotHubs'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["disabledDataSources"] = args ? args.disabledDataSources : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["export"] = args ? args.export : undefined;
            inputs["iotHubs"] = args ? args.iotHubs : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["recommendationsConfiguration"] = args ? args.recommendationsConfiguration : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["unmaskedIpLoggingStatus"] = args ? args.unmaskedIpLoggingStatus : undefined;
            inputs["userDefinedResources"] = args ? args.userDefinedResources : undefined;
            inputs["workspace"] = args ? args.workspace : undefined;
            inputs["autoDiscoveredResources"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IotSecuritySolution.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IotSecuritySolution resource.
 */
export interface IotSecuritySolutionArgs {
    /**
     * Disabled data sources. Disabling these data sources compromises the system.
     */
    readonly disabledDataSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource display name.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * List of additional options for exporting to workspace data.
     */
    readonly export?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IoT Hub resource IDs
     */
    readonly iotHubs: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the IoT Security solution.
     */
    readonly name: pulumi.Input<string>;
    /**
     * List of the configuration status for each recommendation type.
     */
    readonly recommendationsConfiguration?: pulumi.Input<pulumi.Input<inputs.security.v20190801.RecommendationConfigurationProperties>[]>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Status of the IoT Security solution.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Unmasked IP address logging status
     */
    readonly unmaskedIpLoggingStatus?: pulumi.Input<string>;
    /**
     * Properties of the IoT Security solution's user defined resources.
     */
    readonly userDefinedResources?: pulumi.Input<inputs.security.v20190801.UserDefinedResourcesProperties>;
    /**
     * Workspace resource ID
     */
    readonly workspace?: pulumi.Input<string>;
}
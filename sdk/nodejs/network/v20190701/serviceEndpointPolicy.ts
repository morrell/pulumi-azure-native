// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Service End point policy resource.
 */
export class ServiceEndpointPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceEndpointPolicy {
        return new ServiceEndpointPolicy(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190701:ServiceEndpointPolicy';

    /**
     * Returns true if the given object is an instance of ServiceEndpointPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointPolicy.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the service endpoint policy resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource GUID property of the service endpoint policy resource.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * A collection of service endpoint policy definitions of the service endpoint policy.
     */
    public readonly serviceEndpointPolicyDefinitions!: pulumi.Output<outputs.network.v20190701.ServiceEndpointPolicyDefinitionResponse[] | undefined>;
    /**
     * A collection of references to subnets.
     */
    public /*out*/ readonly subnets!: pulumi.Output<outputs.network.v20190701.SubnetResponse[]>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ServiceEndpointPolicyArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceEndpointPolicyDefinitions"] = args ? args.serviceEndpointPolicyDefinitions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/v20180701:ServiceEndpointPolicy" }, { type: "azurerm:network/v20180801:ServiceEndpointPolicy" }, { type: "azurerm:network/v20181001:ServiceEndpointPolicy" }, { type: "azurerm:network/v20181101:ServiceEndpointPolicy" }, { type: "azurerm:network/v20181201:ServiceEndpointPolicy" }, { type: "azurerm:network/v20190201:ServiceEndpointPolicy" }, { type: "azurerm:network/v20190401:ServiceEndpointPolicy" }, { type: "azurerm:network/v20190601:ServiceEndpointPolicy" }, { type: "azurerm:network/v20190801:ServiceEndpointPolicy" }, { type: "azurerm:network/v20190901:ServiceEndpointPolicy" }, { type: "azurerm:network/v20191101:ServiceEndpointPolicy" }, { type: "azurerm:network/v20191201:ServiceEndpointPolicy" }, { type: "azurerm:network/v20200301:ServiceEndpointPolicy" }, { type: "azurerm:network/v20200401:ServiceEndpointPolicy" }, { type: "azurerm:network/v20200501:ServiceEndpointPolicy" }, { type: "azurerm:network/v20200601:ServiceEndpointPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ServiceEndpointPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceEndpointPolicy resource.
 */
export interface ServiceEndpointPolicyArgs {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the service endpoint policy.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A collection of service endpoint policy definitions of the service endpoint policy.
     */
    readonly serviceEndpointPolicyDefinitions?: pulumi.Input<pulumi.Input<inputs.network.v20190701.ServiceEndpointPolicyDefinition>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
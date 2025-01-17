// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Graph Query entity definition.
 */
export class GraphQuery extends pulumi.CustomResource {
    /**
     * Get an existing GraphQuery resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GraphQuery {
        return new GraphQuery(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:resourcegraph/v20200401preview:GraphQuery';

    /**
     * Returns true if the given object is an instance of GraphQuery.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GraphQuery {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GraphQuery.__pulumiType;
    }

    /**
     * The description of a graph query.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * This will be used to handle Optimistic Concurrency.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The location of the resource
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name. This is GUID value. The display name should be assigned within properties field.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * KQL query that will be graph.
     */
    public readonly query!: pulumi.Output<string>;
    /**
     * Enum indicating a type of graph query.
     */
    public /*out*/ readonly resultKind!: pulumi.Output<string>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.resourcegraph.v20200401preview.SystemDataResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Date and time in UTC of the last modification that was made to this graph query definition.
     */
    public /*out*/ readonly timeModified!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GraphQuery resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GraphQueryArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.query === undefined) && !opts.urn) {
                throw new Error("Missing required property 'query'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["query"] = args ? args.query : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["resultKind"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["timeModified"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["query"] = undefined /*out*/;
            inputs["resultKind"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeModified"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:resourcegraph/v20200401preview:GraphQuery" }, { type: "azure-native:resourcegraph:GraphQuery" }, { type: "azure-nextgen:resourcegraph:GraphQuery" }, { type: "azure-native:resourcegraph/v20180901preview:GraphQuery" }, { type: "azure-nextgen:resourcegraph/v20180901preview:GraphQuery" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GraphQuery.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GraphQuery resource.
 */
export interface GraphQueryArgs {
    /**
     * The description of a graph query.
     */
    description?: pulumi.Input<string>;
    /**
     * This will be used to handle Optimistic Concurrency.
     */
    etag?: pulumi.Input<string>;
    /**
     * KQL query that will be graph.
     */
    query: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Graph Query resource.
     */
    resourceName?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Specific entity query.
 *
 * @deprecated Please use one of the variants: ActivityCustomEntityQuery.
 */
export class EntityQuery extends pulumi.CustomResource {
    /**
     * Get an existing EntityQuery resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EntityQuery {
        pulumi.log.warn("EntityQuery is deprecated: Please use one of the variants: ActivityCustomEntityQuery.")
        return new EntityQuery(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:securityinsights/v20210301preview:EntityQuery';

    /**
     * Returns true if the given object is an instance of EntityQuery.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EntityQuery {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EntityQuery.__pulumiType;
    }

    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * the entity query kind
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.securityinsights.v20210301preview.SystemDataResponse>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a EntityQuery resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Please use one of the variants: ActivityCustomEntityQuery. */
    constructor(name: string, args: EntityQueryArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EntityQuery is deprecated: Please use one of the variants: ActivityCustomEntityQuery.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.operationalInsightsResourceProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operationalInsightsResourceProvider'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["entityQueryId"] = args ? args.entityQueryId : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["operationalInsightsResourceProvider"] = args ? args.operationalInsightsResourceProvider : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:securityinsights/v20210301preview:EntityQuery" }, { type: "azure-native:securityinsights:EntityQuery" }, { type: "azure-nextgen:securityinsights:EntityQuery" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(EntityQuery.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EntityQuery resource.
 */
export interface EntityQueryArgs {
    /**
     * entity query ID
     */
    entityQueryId?: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    etag?: pulumi.Input<string>;
    /**
     * the entity query kind
     */
    kind: pulumi.Input<string | enums.securityinsights.v20210301preview.CustomEntityQueryKind>;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    operationalInsightsResourceProvider: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}

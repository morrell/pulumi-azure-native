// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a case relation
 */
export class CaseRelation extends pulumi.CustomResource {
    /**
     * Get an existing CaseRelation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CaseRelation {
        return new CaseRelation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:securityinsights/v20190101preview:CaseRelation';

    /**
     * Returns true if the given object is an instance of CaseRelation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CaseRelation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CaseRelation.__pulumiType;
    }

    /**
     * The case related bookmark id
     */
    public /*out*/ readonly bookmarkId!: pulumi.Output<string>;
    /**
     * The case related bookmark name
     */
    public /*out*/ readonly bookmarkName!: pulumi.Output<string | undefined>;
    /**
     * The case identifier
     */
    public /*out*/ readonly caseIdentifier!: pulumi.Output<string>;
    /**
     * ETag for relation
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The type of relation node
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Name of relation
     */
    public readonly relationName!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CaseRelation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CaseRelationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.caseId === undefined) {
                throw new Error("Missing required property 'caseId'");
            }
            if (!args || args.operationalInsightsResourceProvider === undefined) {
                throw new Error("Missing required property 'operationalInsightsResourceProvider'");
            }
            if (!args || args.relationName === undefined) {
                throw new Error("Missing required property 'relationName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["caseId"] = args ? args.caseId : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["operationalInsightsResourceProvider"] = args ? args.operationalInsightsResourceProvider : undefined;
            inputs["relationName"] = args ? args.relationName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceRelationNode"] = args ? args.sourceRelationNode : undefined;
            inputs["targetRelationNode"] = args ? args.targetRelationNode : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["bookmarkId"] = undefined /*out*/;
            inputs["bookmarkName"] = undefined /*out*/;
            inputs["caseIdentifier"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["bookmarkId"] = undefined /*out*/;
            inputs["bookmarkName"] = undefined /*out*/;
            inputs["caseIdentifier"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["relationName"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CaseRelation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CaseRelation resource.
 */
export interface CaseRelationArgs {
    /**
     * Case ID
     */
    readonly caseId: pulumi.Input<string>;
    /**
     * ETag for relation
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    readonly operationalInsightsResourceProvider: pulumi.Input<string>;
    /**
     * Name of relation
     */
    readonly relationName: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Relation source node
     */
    readonly sourceRelationNode?: pulumi.Input<inputs.securityinsights.v20190101preview.RelationNode>;
    /**
     * Relation target node
     */
    readonly targetRelationNode?: pulumi.Input<inputs.securityinsights.v20190101preview.RelationNode>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
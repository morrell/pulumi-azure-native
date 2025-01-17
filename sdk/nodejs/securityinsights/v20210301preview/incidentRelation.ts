// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Represents a relation between two resources
 */
export class IncidentRelation extends pulumi.CustomResource {
    /**
     * Get an existing IncidentRelation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IncidentRelation {
        return new IncidentRelation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:securityinsights/v20210301preview:IncidentRelation';

    /**
     * Returns true if the given object is an instance of IncidentRelation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IncidentRelation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IncidentRelation.__pulumiType;
    }

    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource ID of the related resource
     */
    public readonly relatedResourceId!: pulumi.Output<string>;
    /**
     * The resource kind of the related resource
     */
    public /*out*/ readonly relatedResourceKind!: pulumi.Output<string>;
    /**
     * The name of the related resource
     */
    public /*out*/ readonly relatedResourceName!: pulumi.Output<string>;
    /**
     * The resource type of the related resource
     */
    public /*out*/ readonly relatedResourceType!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.securityinsights.v20210301preview.SystemDataResponse>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IncidentRelation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IncidentRelationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.incidentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'incidentId'");
            }
            if ((!args || args.operationalInsightsResourceProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operationalInsightsResourceProvider'");
            }
            if ((!args || args.relatedResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'relatedResourceId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["incidentId"] = args ? args.incidentId : undefined;
            inputs["operationalInsightsResourceProvider"] = args ? args.operationalInsightsResourceProvider : undefined;
            inputs["relatedResourceId"] = args ? args.relatedResourceId : undefined;
            inputs["relationName"] = args ? args.relationName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["relatedResourceKind"] = undefined /*out*/;
            inputs["relatedResourceName"] = undefined /*out*/;
            inputs["relatedResourceType"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["relatedResourceId"] = undefined /*out*/;
            inputs["relatedResourceKind"] = undefined /*out*/;
            inputs["relatedResourceName"] = undefined /*out*/;
            inputs["relatedResourceType"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:securityinsights/v20210301preview:IncidentRelation" }, { type: "azure-native:securityinsights:IncidentRelation" }, { type: "azure-nextgen:securityinsights:IncidentRelation" }, { type: "azure-native:securityinsights/v20190101preview:IncidentRelation" }, { type: "azure-nextgen:securityinsights/v20190101preview:IncidentRelation" }, { type: "azure-native:securityinsights/v20210401:IncidentRelation" }, { type: "azure-nextgen:securityinsights/v20210401:IncidentRelation" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(IncidentRelation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IncidentRelation resource.
 */
export interface IncidentRelationArgs {
    /**
     * Etag of the azure resource
     */
    etag?: pulumi.Input<string>;
    /**
     * Incident ID
     */
    incidentId: pulumi.Input<string>;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    operationalInsightsResourceProvider: pulumi.Input<string>;
    /**
     * The resource ID of the related resource
     */
    relatedResourceId: pulumi.Input<string>;
    /**
     * Relation Name
     */
    relationName?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}

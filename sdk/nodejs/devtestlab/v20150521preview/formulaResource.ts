// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * A formula.
 */
export class FormulaResource extends pulumi.CustomResource {
    /**
     * Get an existing FormulaResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FormulaResource {
        return new FormulaResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:devtestlab/v20150521preview:FormulaResource';

    /**
     * Returns true if the given object is an instance of FormulaResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FormulaResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FormulaResource.__pulumiType;
    }

    /**
     * The author of the formula.
     */
    public readonly author!: pulumi.Output<string | undefined>;
    /**
     * The creation date of the formula.
     */
    public readonly creationDate!: pulumi.Output<string | undefined>;
    /**
     * The description of the formula.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The content of the formula.
     */
    public readonly formulaContent!: pulumi.Output<outputs.devtestlab.v20150521preview.LabVirtualMachineResponse | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The OS type of the formula.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * The provisioning status of the resource.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Information about a VM from which a formula is to be created.
     */
    public readonly vm!: pulumi.Output<outputs.devtestlab.v20150521preview.FormulaPropertiesFromVmResponse | undefined>;

    /**
     * Create a FormulaResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FormulaResourceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.labName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'labName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["author"] = args ? args.author : undefined;
            inputs["creationDate"] = args ? args.creationDate : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["formulaContent"] = args ? args.formulaContent : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vm"] = args ? args.vm : undefined;
        } else {
            inputs["author"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["formulaContent"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["osType"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vm"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:devtestlab/v20150521preview:FormulaResource" }, { type: "azure-native:devtestlab:FormulaResource" }, { type: "azure-nextgen:devtestlab:FormulaResource" }, { type: "azure-native:devtestlab/v20160515:FormulaResource" }, { type: "azure-nextgen:devtestlab/v20160515:FormulaResource" }, { type: "azure-native:devtestlab/v20180915:FormulaResource" }, { type: "azure-nextgen:devtestlab/v20180915:FormulaResource" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(FormulaResource.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FormulaResource resource.
 */
export interface FormulaResourceArgs {
    /**
     * The author of the formula.
     */
    author?: pulumi.Input<string>;
    /**
     * The creation date of the formula.
     */
    creationDate?: pulumi.Input<string>;
    /**
     * The description of the formula.
     */
    description?: pulumi.Input<string>;
    /**
     * The content of the formula.
     */
    formulaContent?: pulumi.Input<inputs.devtestlab.v20150521preview.LabVirtualMachineArgs>;
    /**
     * The identifier of the resource.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The OS type of the formula.
     */
    osType?: pulumi.Input<string>;
    /**
     * The provisioning status of the resource.
     */
    provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the resource.
     */
    type?: pulumi.Input<string>;
    /**
     * Information about a VM from which a formula is to be created.
     */
    vm?: pulumi.Input<inputs.devtestlab.v20150521preview.FormulaPropertiesFromVmArgs>;
}

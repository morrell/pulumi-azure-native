// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Specifies information about the Shared Image Gallery that you want to create or update.
 */
export class Gallery extends pulumi.CustomResource {
    /**
     * Get an existing Gallery resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Gallery {
        return new Gallery(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20190301:Gallery';

    /**
     * Returns true if the given object is an instance of Gallery.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gallery {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gallery.__pulumiType;
    }

    /**
     * The description of this Shared Image Gallery resource. This property is updatable.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Describes the gallery unique name.
     */
    public /*out*/ readonly identifier!: pulumi.Output<outputs.compute.v20190301.GalleryIdentifierResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Gallery resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GalleryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GalleryArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as GalleryArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["identifier"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:compute/v20180601:Gallery" }, { type: "azurerm:compute/v20190701:Gallery" }, { type: "azurerm:compute/v20191201:Gallery" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Gallery.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Gallery resource.
 */
export interface GalleryArgs {
    /**
     * The description of this Shared Image Gallery resource. This property is updatable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the Shared Image Gallery. The allowed characters are alphabets and numbers with dots and periods allowed in the middle. The maximum length is 80 characters.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
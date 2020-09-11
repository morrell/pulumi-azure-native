// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Specifies information about the gallery image definition that you want to create or update.
 */
export class GalleryImage extends pulumi.CustomResource {
    /**
     * Get an existing GalleryImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GalleryImage {
        return new GalleryImage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20200930:GalleryImage';

    /**
     * Returns true if the given object is an instance of GalleryImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GalleryImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GalleryImage.__pulumiType;
    }

    /**
     * The description of this gallery image definition resource. This property is updatable.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Describes the disallowed disk types.
     */
    public readonly disallowed!: pulumi.Output<outputs.compute.v20200930.DisallowedResponse | undefined>;
    /**
     * The end of life date of the gallery image definition. This property can be used for decommissioning purposes. This property is updatable.
     */
    public readonly endOfLifeDate!: pulumi.Output<string | undefined>;
    /**
     * The Eula agreement for the gallery image definition.
     */
    public readonly eula!: pulumi.Output<string | undefined>;
    /**
     * A list of gallery image features.
     */
    public readonly features!: pulumi.Output<outputs.compute.v20200930.GalleryImageFeatureResponse[] | undefined>;
    /**
     * The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
     */
    public readonly hyperVGeneration!: pulumi.Output<string | undefined>;
    /**
     * This is the gallery image definition identifier.
     */
    public readonly identifier!: pulumi.Output<outputs.compute.v20200930.GalleryImageIdentifierResponse>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
     */
    public readonly osState!: pulumi.Output<string>;
    /**
     * This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
     */
    public readonly osType!: pulumi.Output<string>;
    /**
     * The privacy statement uri.
     */
    public readonly privacyStatementUri!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Describes the gallery image definition purchase plan. This is used by marketplace images.
     */
    public readonly purchasePlan!: pulumi.Output<outputs.compute.v20200930.ImagePurchasePlanResponse | undefined>;
    /**
     * The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
     */
    public readonly recommended!: pulumi.Output<outputs.compute.v20200930.RecommendedMachineConfigurationResponse | undefined>;
    /**
     * The release note uri.
     */
    public readonly releaseNoteUri!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GalleryImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GalleryImageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.galleryImageName === undefined) {
                throw new Error("Missing required property 'galleryImageName'");
            }
            if (!args || args.galleryName === undefined) {
                throw new Error("Missing required property 'galleryName'");
            }
            if (!args || args.identifier === undefined) {
                throw new Error("Missing required property 'identifier'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.osState === undefined) {
                throw new Error("Missing required property 'osState'");
            }
            if (!args || args.osType === undefined) {
                throw new Error("Missing required property 'osType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["disallowed"] = args ? args.disallowed : undefined;
            inputs["endOfLifeDate"] = args ? args.endOfLifeDate : undefined;
            inputs["eula"] = args ? args.eula : undefined;
            inputs["features"] = args ? args.features : undefined;
            inputs["galleryImageName"] = args ? args.galleryImageName : undefined;
            inputs["galleryName"] = args ? args.galleryName : undefined;
            inputs["hyperVGeneration"] = args ? args.hyperVGeneration : undefined;
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["osState"] = args ? args.osState : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["privacyStatementUri"] = args ? args.privacyStatementUri : undefined;
            inputs["purchasePlan"] = args ? args.purchasePlan : undefined;
            inputs["recommended"] = args ? args.recommended : undefined;
            inputs["releaseNoteUri"] = args ? args.releaseNoteUri : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["disallowed"] = undefined /*out*/;
            inputs["endOfLifeDate"] = undefined /*out*/;
            inputs["eula"] = undefined /*out*/;
            inputs["features"] = undefined /*out*/;
            inputs["hyperVGeneration"] = undefined /*out*/;
            inputs["identifier"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["osState"] = undefined /*out*/;
            inputs["osType"] = undefined /*out*/;
            inputs["privacyStatementUri"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["purchasePlan"] = undefined /*out*/;
            inputs["recommended"] = undefined /*out*/;
            inputs["releaseNoteUri"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:GalleryImage" }, { type: "azurerm:compute/v20180601:GalleryImage" }, { type: "azurerm:compute/v20190301:GalleryImage" }, { type: "azurerm:compute/v20190701:GalleryImage" }, { type: "azurerm:compute/v20191201:GalleryImage" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(GalleryImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GalleryImage resource.
 */
export interface GalleryImageArgs {
    /**
     * The description of this gallery image definition resource. This property is updatable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Describes the disallowed disk types.
     */
    readonly disallowed?: pulumi.Input<inputs.compute.v20200930.Disallowed>;
    /**
     * The end of life date of the gallery image definition. This property can be used for decommissioning purposes. This property is updatable.
     */
    readonly endOfLifeDate?: pulumi.Input<string>;
    /**
     * The Eula agreement for the gallery image definition.
     */
    readonly eula?: pulumi.Input<string>;
    /**
     * A list of gallery image features.
     */
    readonly features?: pulumi.Input<pulumi.Input<inputs.compute.v20200930.GalleryImageFeature>[]>;
    /**
     * The name of the gallery image definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
     */
    readonly galleryImageName: pulumi.Input<string>;
    /**
     * The name of the Shared Image Gallery in which the Image Definition is to be created.
     */
    readonly galleryName: pulumi.Input<string>;
    /**
     * The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
     */
    readonly hyperVGeneration?: pulumi.Input<string>;
    /**
     * This is the gallery image definition identifier.
     */
    readonly identifier: pulumi.Input<inputs.compute.v20200930.GalleryImageIdentifier>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
     */
    readonly osState: pulumi.Input<string>;
    /**
     * This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
     */
    readonly osType: pulumi.Input<string>;
    /**
     * The privacy statement uri.
     */
    readonly privacyStatementUri?: pulumi.Input<string>;
    /**
     * Describes the gallery image definition purchase plan. This is used by marketplace images.
     */
    readonly purchasePlan?: pulumi.Input<inputs.compute.v20200930.ImagePurchasePlan>;
    /**
     * The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
     */
    readonly recommended?: pulumi.Input<inputs.compute.v20200930.RecommendedMachineConfiguration>;
    /**
     * The release note uri.
     */
    readonly releaseNoteUri?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
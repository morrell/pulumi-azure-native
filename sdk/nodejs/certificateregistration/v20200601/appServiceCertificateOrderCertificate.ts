// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Key Vault container ARM resource for a certificate that is purchased through Azure.
 */
export class AppServiceCertificateOrderCertificate extends pulumi.CustomResource {
    /**
     * Get an existing AppServiceCertificateOrderCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppServiceCertificateOrderCertificate {
        return new AppServiceCertificateOrderCertificate(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:certificateregistration/v20200601:AppServiceCertificateOrderCertificate';

    /**
     * Returns true if the given object is an instance of AppServiceCertificateOrderCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppServiceCertificateOrderCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppServiceCertificateOrderCertificate.__pulumiType;
    }

    /**
     * Key Vault resource Id.
     */
    public readonly keyVaultId!: pulumi.Output<string | undefined>;
    /**
     * Key Vault secret name.
     */
    public readonly keyVaultSecretName!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Status of the Key Vault secret.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AppServiceCertificateOrderCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppServiceCertificateOrderCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppServiceCertificateOrderCertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AppServiceCertificateOrderCertificateArgs | undefined;
            if (!args || args.certificateOrderName === undefined) {
                throw new Error("Missing required property 'certificateOrderName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["certificateOrderName"] = args ? args.certificateOrderName : undefined;
            inputs["keyVaultId"] = args ? args.keyVaultId : undefined;
            inputs["keyVaultSecretName"] = args ? args.keyVaultSecretName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:certificateregistration/v20150801:AppServiceCertificateOrderCertificate" }, { type: "azurerm:certificateregistration/v20180201:AppServiceCertificateOrderCertificate" }, { type: "azurerm:certificateregistration/v20190801:AppServiceCertificateOrderCertificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AppServiceCertificateOrderCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppServiceCertificateOrderCertificate resource.
 */
export interface AppServiceCertificateOrderCertificateArgs {
    /**
     * Name of the certificate order.
     */
    readonly certificateOrderName: pulumi.Input<string>;
    /**
     * Key Vault resource Id.
     */
    readonly keyVaultId?: pulumi.Input<string>;
    /**
     * Key Vault secret name.
     */
    readonly keyVaultSecretName?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the certificate.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
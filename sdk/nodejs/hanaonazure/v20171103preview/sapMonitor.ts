// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * SAP monitor info on Azure (ARM properties and SAP monitor properties)
 */
export class SapMonitor extends pulumi.CustomResource {
    /**
     * Get an existing SapMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SapMonitor {
        return new SapMonitor(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:hanaonazure/v20171103preview:SapMonitor';

    /**
     * Returns true if the given object is an instance of SapMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SapMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SapMonitor.__pulumiType;
    }

    /**
     * The value indicating whether to send analytics to Microsoft
     */
    public readonly enableCustomerAnalytics!: pulumi.Output<boolean | undefined>;
    /**
     * MSI ID passed by customer which has access to customer's KeyVault and to be assigned to the Collector VM.
     */
    public readonly hanaDbCredentialsMsiId!: pulumi.Output<string | undefined>;
    /**
     * Database name of the HANA instance.
     */
    public readonly hanaDbName!: pulumi.Output<string | undefined>;
    /**
     * Database password of the HANA instance.
     */
    public readonly hanaDbPassword!: pulumi.Output<string | undefined>;
    /**
     * KeyVault URL link to the password for the HANA database.
     */
    public readonly hanaDbPasswordKeyVaultUrl!: pulumi.Output<string | undefined>;
    /**
     * Database port of the HANA instance.
     */
    public readonly hanaDbSqlPort!: pulumi.Output<number | undefined>;
    /**
     * Database username of the HANA instance.
     */
    public readonly hanaDbUsername!: pulumi.Output<string | undefined>;
    /**
     * Hostname of the HANA instance.
     */
    public readonly hanaHostname!: pulumi.Output<string | undefined>;
    /**
     * Specifies the SAP monitor unique ID.
     */
    public readonly hanaSubnet!: pulumi.Output<string | undefined>;
    /**
     * Key Vault ID containing customer's HANA credentials.
     */
    public readonly keyVaultId!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The ARM ID of the Log Analytics Workspace that is used for monitoring
     */
    public readonly logAnalyticsWorkspaceArmId!: pulumi.Output<string | undefined>;
    /**
     * The workspace ID of the log analytics workspace to be used for monitoring
     */
    public readonly logAnalyticsWorkspaceId!: pulumi.Output<string | undefined>;
    /**
     * The shared key of the log analytics workspace that is used for monitoring
     */
    public readonly logAnalyticsWorkspaceSharedKey!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group the SAP Monitor resources get deployed into.
     */
    public /*out*/ readonly managedResourceGroupName!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * State of provisioning of the HanaInstance
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SapMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SapMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SapMonitorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SapMonitorArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sapMonitorName === undefined) {
                throw new Error("Missing required property 'sapMonitorName'");
            }
            inputs["enableCustomerAnalytics"] = args ? args.enableCustomerAnalytics : undefined;
            inputs["hanaDbCredentialsMsiId"] = args ? args.hanaDbCredentialsMsiId : undefined;
            inputs["hanaDbName"] = args ? args.hanaDbName : undefined;
            inputs["hanaDbPassword"] = args ? args.hanaDbPassword : undefined;
            inputs["hanaDbPasswordKeyVaultUrl"] = args ? args.hanaDbPasswordKeyVaultUrl : undefined;
            inputs["hanaDbSqlPort"] = args ? args.hanaDbSqlPort : undefined;
            inputs["hanaDbUsername"] = args ? args.hanaDbUsername : undefined;
            inputs["hanaHostname"] = args ? args.hanaHostname : undefined;
            inputs["hanaSubnet"] = args ? args.hanaSubnet : undefined;
            inputs["keyVaultId"] = args ? args.keyVaultId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logAnalyticsWorkspaceArmId"] = args ? args.logAnalyticsWorkspaceArmId : undefined;
            inputs["logAnalyticsWorkspaceId"] = args ? args.logAnalyticsWorkspaceId : undefined;
            inputs["logAnalyticsWorkspaceSharedKey"] = args ? args.logAnalyticsWorkspaceSharedKey : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sapMonitorName"] = args ? args.sapMonitorName : undefined;
            inputs["managedResourceGroupName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:hanaonazure/v20200207preview:SapMonitor" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SapMonitor.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SapMonitor resource.
 */
export interface SapMonitorArgs {
    /**
     * The value indicating whether to send analytics to Microsoft
     */
    readonly enableCustomerAnalytics?: pulumi.Input<boolean>;
    /**
     * MSI ID passed by customer which has access to customer's KeyVault and to be assigned to the Collector VM.
     */
    readonly hanaDbCredentialsMsiId?: pulumi.Input<string>;
    /**
     * Database name of the HANA instance.
     */
    readonly hanaDbName?: pulumi.Input<string>;
    /**
     * Database password of the HANA instance.
     */
    readonly hanaDbPassword?: pulumi.Input<string>;
    /**
     * KeyVault URL link to the password for the HANA database.
     */
    readonly hanaDbPasswordKeyVaultUrl?: pulumi.Input<string>;
    /**
     * Database port of the HANA instance.
     */
    readonly hanaDbSqlPort?: pulumi.Input<number>;
    /**
     * Database username of the HANA instance.
     */
    readonly hanaDbUsername?: pulumi.Input<string>;
    /**
     * Hostname of the HANA instance.
     */
    readonly hanaHostname?: pulumi.Input<string>;
    /**
     * Specifies the SAP monitor unique ID.
     */
    readonly hanaSubnet?: pulumi.Input<string>;
    /**
     * Key Vault ID containing customer's HANA credentials.
     */
    readonly keyVaultId?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The ARM ID of the Log Analytics Workspace that is used for monitoring
     */
    readonly logAnalyticsWorkspaceArmId?: pulumi.Input<string>;
    /**
     * The workspace ID of the log analytics workspace to be used for monitoring
     */
    readonly logAnalyticsWorkspaceId?: pulumi.Input<string>;
    /**
     * The shared key of the log analytics workspace that is used for monitoring
     */
    readonly logAnalyticsWorkspaceSharedKey?: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the SAP monitor resource.
     */
    readonly sapMonitorName: pulumi.Input<string>;
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * A managed database resource.
 */
export class ManagedDatabase extends pulumi.CustomResource {
    /**
     * Get an existing ManagedDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagedDatabase {
        return new ManagedDatabase(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sql/v20180601preview:ManagedDatabase';

    /**
     * Returns true if the given object is an instance of ManagedDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedDatabase.__pulumiType;
    }

    /**
     * Collation of the metadata catalog.
     */
    public readonly catalogCollation!: pulumi.Output<string | undefined>;
    /**
     * Collation of the managed database.
     */
    public readonly collation!: pulumi.Output<string | undefined>;
    /**
     * Managed database create mode. PointInTimeRestore: Create a database by restoring a point in time backup of an existing database. SourceDatabaseName, SourceManagedInstanceName and PointInTime must be specified. RestoreExternalBackup: Create a database by restoring from external backup files. Collation, StorageContainerUri and StorageContainerSasToken must be specified. Recovery: Creates a database by restoring a geo-replicated backup. RecoverableDatabaseId must be specified as the recoverable database resource ID to restore.
     */
    public readonly createMode!: pulumi.Output<string | undefined>;
    /**
     * Creation date of the database.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Geo paired region.
     */
    public /*out*/ readonly defaultSecondaryLocation!: pulumi.Output<string>;
    /**
     * Earliest restore point in time for point in time restore.
     */
    public /*out*/ readonly earliestRestorePoint!: pulumi.Output<string>;
    /**
     * Instance Failover Group resource identifier that this managed database belongs to.
     */
    public /*out*/ readonly failoverGroupId!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Long Term Retention backup to be used for restore of this managed database.
     */
    public readonly longTermRetentionBackupResourceId!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource identifier of the recoverable database associated with create operation of this database.
     */
    public readonly recoverableDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * The restorable dropped database resource id to restore when creating this database.
     */
    public readonly restorableDroppedDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * Conditional. If createMode is PointInTimeRestore, this value is required. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
     */
    public readonly restorePointInTime!: pulumi.Output<string | undefined>;
    /**
     * The resource identifier of the source database associated with create operation of this database.
     */
    public readonly sourceDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * Status of the database.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Conditional. If createMode is RestoreExternalBackup, this value is required. Specifies the storage container sas token.
     */
    public readonly storageContainerSasToken!: pulumi.Output<string | undefined>;
    /**
     * Conditional. If createMode is RestoreExternalBackup, this value is required. Specifies the uri of the storage container where backups for this restore are stored.
     */
    public readonly storageContainerUri!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagedDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedDatabaseArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ManagedDatabaseArgs | undefined;
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.managedInstanceName === undefined) {
                throw new Error("Missing required property 'managedInstanceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["catalogCollation"] = args ? args.catalogCollation : undefined;
            inputs["collation"] = args ? args.collation : undefined;
            inputs["createMode"] = args ? args.createMode : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["longTermRetentionBackupResourceId"] = args ? args.longTermRetentionBackupResourceId : undefined;
            inputs["managedInstanceName"] = args ? args.managedInstanceName : undefined;
            inputs["recoverableDatabaseId"] = args ? args.recoverableDatabaseId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restorableDroppedDatabaseId"] = args ? args.restorableDroppedDatabaseId : undefined;
            inputs["restorePointInTime"] = args ? args.restorePointInTime : undefined;
            inputs["sourceDatabaseId"] = args ? args.sourceDatabaseId : undefined;
            inputs["storageContainerSasToken"] = args ? args.storageContainerSasToken : undefined;
            inputs["storageContainerUri"] = args ? args.storageContainerUri : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["defaultSecondaryLocation"] = undefined /*out*/;
            inputs["earliestRestorePoint"] = undefined /*out*/;
            inputs["failoverGroupId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:sql/v20170301preview:ManagedDatabase" }, { type: "azurerm:sql/v20190601preview:ManagedDatabase" }, { type: "azurerm:sql/v20200202preview:ManagedDatabase" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagedDatabase.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedDatabase resource.
 */
export interface ManagedDatabaseArgs {
    /**
     * Collation of the metadata catalog.
     */
    readonly catalogCollation?: pulumi.Input<string>;
    /**
     * Collation of the managed database.
     */
    readonly collation?: pulumi.Input<string>;
    /**
     * Managed database create mode. PointInTimeRestore: Create a database by restoring a point in time backup of an existing database. SourceDatabaseName, SourceManagedInstanceName and PointInTime must be specified. RestoreExternalBackup: Create a database by restoring from external backup files. Collation, StorageContainerUri and StorageContainerSasToken must be specified. Recovery: Creates a database by restoring a geo-replicated backup. RecoverableDatabaseId must be specified as the recoverable database resource ID to restore.
     */
    readonly createMode?: pulumi.Input<string>;
    /**
     * The name of the database.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the Long Term Retention backup to be used for restore of this managed database.
     */
    readonly longTermRetentionBackupResourceId?: pulumi.Input<string>;
    /**
     * The name of the managed instance.
     */
    readonly managedInstanceName: pulumi.Input<string>;
    /**
     * The resource identifier of the recoverable database associated with create operation of this database.
     */
    readonly recoverableDatabaseId?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The restorable dropped database resource id to restore when creating this database.
     */
    readonly restorableDroppedDatabaseId?: pulumi.Input<string>;
    /**
     * Conditional. If createMode is PointInTimeRestore, this value is required. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
     */
    readonly restorePointInTime?: pulumi.Input<string>;
    /**
     * The resource identifier of the source database associated with create operation of this database.
     */
    readonly sourceDatabaseId?: pulumi.Input<string>;
    /**
     * Conditional. If createMode is RestoreExternalBackup, this value is required. Specifies the storage container sas token.
     */
    readonly storageContainerSasToken?: pulumi.Input<string>;
    /**
     * Conditional. If createMode is RestoreExternalBackup, this value is required. Specifies the uri of the storage container where backups for this restore are stored.
     */
    readonly storageContainerUri?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
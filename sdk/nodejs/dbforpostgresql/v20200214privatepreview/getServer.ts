// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:dbforpostgresql/v20200214privatepreview:getServer", {
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
}

/**
 * Represents a server.
 */
export interface GetServerResult {
    /**
     * The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
     */
    readonly administratorLogin?: string;
    /**
     * The administrator login password (required for server creation).
     */
    readonly administratorLoginPassword?: string;
    /**
     * availability Zone information of the server.
     */
    readonly availabilityZone?: string;
    /**
     * The mode to create a new PostgreSQL server.
     */
    readonly createMode?: string;
    /**
     * The display name of a server.
     */
    readonly displayName?: string;
    /**
     * The fully qualified domain name of a server.
     */
    readonly fullyQualifiedDomainName: string;
    /**
     * A state of a HA server that is visible to user.
     */
    readonly haState: string;
    /**
     * The Azure Active Directory identity of the server.
     */
    readonly identity?: outputs.dbforpostgresql.v20200214privatepreview.IdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Restore point creation time (ISO8601 format), specifying the time to restore from.
     */
    readonly pointInTimeUTC?: string;
    /**
     * public network access is enabled or not
     */
    readonly publicNetworkAccess?: string;
    /**
     * The SKU (pricing tier) of the server.
     */
    readonly sku?: outputs.dbforpostgresql.v20200214privatepreview.SkuResponse;
    /**
     * The source PostgreSQL server name to restore from.
     */
    readonly sourceServerName?: string;
    /**
     * stand by count value can be either 0 or 1
     */
    readonly standbyCount?: number;
    /**
     * A state of a server that is visible to user.
     */
    readonly state: string;
    /**
     * Storage profile of a server.
     */
    readonly storageProfile?: outputs.dbforpostgresql.v20200214privatepreview.StorageProfileResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
    /**
     * PostgreSQL Server version.
     */
    readonly version?: string;
    readonly vnetInjArgs?: outputs.dbforpostgresql.v20200214privatepreview.ServerPropertiesResponseVnetInjArgs;
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDomainService(args: GetDomainServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:aad/v20170601:getDomainService", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDomainServiceArgs {
    /**
     * The name of the domain service.
     */
    readonly name: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Domain service.
 */
export interface GetDomainServiceResult {
    /**
     * List of Domain Controller IP Address
     */
    readonly domainControllerIpAddress: string[];
    /**
     * The name of the Azure domain that the user would like to deploy Domain Services to.
     */
    readonly domainName?: string;
    /**
     * DomainSecurity Settings
     */
    readonly domainSecuritySettings?: outputs.aad.v20170601.DomainSecuritySettingsResponse;
    /**
     * Resource etag
     */
    readonly etag?: string;
    /**
     * Enabled or Disabled flag to turn on Group-based filtered sync
     */
    readonly filteredSync?: string;
    /**
     * List of Domain Health Alerts
     */
    readonly healthAlerts: outputs.aad.v20170601.HealthAlertResponse[];
    /**
     * Last domain evaluation run DateTime
     */
    readonly healthLastEvaluated: string;
    /**
     * List of Domain Health Monitors
     */
    readonly healthMonitors: outputs.aad.v20170601.HealthMonitorResponse[];
    /**
     * Secure LDAP Settings
     */
    readonly ldapsSettings?: outputs.aad.v20170601.LdapsSettingsResponse;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Notification Settings
     */
    readonly notificationSettings?: outputs.aad.v20170601.NotificationSettingsResponse;
    /**
     * the current deployment or provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Status of Domain Service instance
     */
    readonly serviceStatus: string;
    /**
     * The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
     */
    readonly subnetId?: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Azure Active Directory tenant id
     */
    readonly tenantId: string;
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Virtual network site id
     */
    readonly vnetSiteId: string;
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getApplicationGateway(args: GetApplicationGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20180601:getApplicationGateway", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationGatewayArgs {
    /**
     * The name of the application gateway.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Application gateway resource
 */
export interface GetApplicationGatewayResult {
    /**
     * Authentication certificates of the application gateway resource.
     */
    readonly authenticationCertificates?: outputs.network.v20180601.ApplicationGatewayAuthenticationCertificateResponse[];
    /**
     * Autoscale Configuration.
     */
    readonly autoscaleConfiguration?: outputs.network.v20180601.ApplicationGatewayAutoscaleConfigurationResponse;
    /**
     * Backend address pool of the application gateway resource.
     */
    readonly backendAddressPools?: outputs.network.v20180601.ApplicationGatewayBackendAddressPoolResponse[];
    /**
     * Backend http settings of the application gateway resource.
     */
    readonly backendHttpSettingsCollection?: outputs.network.v20180601.ApplicationGatewayBackendHttpSettingsResponse[];
    /**
     * Whether FIPS is enabled on the application gateway resource.
     */
    readonly enableFips?: boolean;
    /**
     * Whether HTTP2 is enabled on the application gateway resource.
     */
    readonly enableHttp2?: boolean;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: string;
    /**
     * Frontend IP addresses of the application gateway resource.
     */
    readonly frontendIPConfigurations?: outputs.network.v20180601.ApplicationGatewayFrontendIPConfigurationResponse[];
    /**
     * Frontend ports of the application gateway resource.
     */
    readonly frontendPorts?: outputs.network.v20180601.ApplicationGatewayFrontendPortResponse[];
    /**
     * Subnets of application the gateway resource.
     */
    readonly gatewayIPConfigurations?: outputs.network.v20180601.ApplicationGatewayIPConfigurationResponse[];
    /**
     * Http listeners of the application gateway resource.
     */
    readonly httpListeners?: outputs.network.v20180601.ApplicationGatewayHttpListenerResponse[];
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Operational state of the application gateway resource.
     */
    readonly operationalState: string;
    /**
     * Probes of the application gateway resource.
     */
    readonly probes?: outputs.network.v20180601.ApplicationGatewayProbeResponse[];
    /**
     * Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: string;
    /**
     * Redirect configurations of the application gateway resource.
     */
    readonly redirectConfigurations?: outputs.network.v20180601.ApplicationGatewayRedirectConfigurationResponse[];
    /**
     * Request routing rules of the application gateway resource.
     */
    readonly requestRoutingRules?: outputs.network.v20180601.ApplicationGatewayRequestRoutingRuleResponse[];
    /**
     * Resource GUID property of the application gateway resource.
     */
    readonly resourceGuid?: string;
    /**
     * SKU of the application gateway resource.
     */
    readonly sku?: outputs.network.v20180601.ApplicationGatewaySkuResponse;
    /**
     * SSL certificates of the application gateway resource.
     */
    readonly sslCertificates?: outputs.network.v20180601.ApplicationGatewaySslCertificateResponse[];
    /**
     * SSL policy of the application gateway resource.
     */
    readonly sslPolicy?: outputs.network.v20180601.ApplicationGatewaySslPolicyResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * URL path map of the application gateway resource.
     */
    readonly urlPathMaps?: outputs.network.v20180601.ApplicationGatewayUrlPathMapResponse[];
    /**
     * Web application firewall configuration.
     */
    readonly webApplicationFirewallConfiguration?: outputs.network.v20180601.ApplicationGatewayWebApplicationFirewallConfigurationResponse;
    /**
     * A list of availability zones denoting where the resource needs to come from.
     */
    readonly zones?: string[];
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./configuration";
export * from "./database";
export * from "./firewallRule";
export * from "./getConfiguration";
export * from "./getDatabase";
export * from "./getFirewallRule";
export * from "./getServer";
export * from "./getServerAdministrator";
export * from "./getServerSecurityAlertPolicy";
export * from "./getVirtualNetworkRule";
export * from "./server";
export * from "./serverAdministrator";
export * from "./serverSecurityAlertPolicy";
export * from "./virtualNetworkRule";

// Export enums:
export * from "../../types/enums/dbforpostgresql/v20171201";

// Import resources to register:
import { Configuration } from "./configuration";
import { Database } from "./database";
import { FirewallRule } from "./firewallRule";
import { Server } from "./server";
import { ServerAdministrator } from "./serverAdministrator";
import { ServerSecurityAlertPolicy } from "./serverSecurityAlertPolicy";
import { VirtualNetworkRule } from "./virtualNetworkRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:dbforpostgresql/v20171201:Configuration":
                return new Configuration(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20171201:Database":
                return new Database(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20171201:FirewallRule":
                return new FirewallRule(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20171201:Server":
                return new Server(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20171201:ServerAdministrator":
                return new ServerAdministrator(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20171201:ServerSecurityAlertPolicy":
                return new ServerSecurityAlertPolicy(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20171201:VirtualNetworkRule":
                return new VirtualNetworkRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "dbforpostgresql/v20171201", _module)

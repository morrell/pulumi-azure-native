// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./configuration";
export * from "./firewallRule";
export * from "./getConfiguration";
export * from "./getFirewallRule";
export * from "./getServer";
export * from "./getServerKey";
export * from "./server";
export * from "./serverKey";

// Export enums:
export * from "../../types/enums/dbforpostgresql/v20200214privatepreview";

// Import resources to register:
import { Configuration } from "./configuration";
import { FirewallRule } from "./firewallRule";
import { Server } from "./server";
import { ServerKey } from "./serverKey";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:dbforpostgresql/v20200214privatepreview:Configuration":
                return new Configuration(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20200214privatepreview:FirewallRule":
                return new FirewallRule(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20200214privatepreview:Server":
                return new Server(name, <any>undefined, { urn })
            case "azure-native:dbforpostgresql/v20200214privatepreview:ServerKey":
                return new ServerKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "dbforpostgresql/v20200214privatepreview", _module)

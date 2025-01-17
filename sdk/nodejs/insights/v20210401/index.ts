// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./dataCollectionEndpoint";
export * from "./dataCollectionRule";
export * from "./dataCollectionRuleAssociation";
export * from "./getDataCollectionEndpoint";
export * from "./getDataCollectionRule";
export * from "./getDataCollectionRuleAssociation";

// Export enums:
export * from "../../types/enums/insights/v20210401";

// Import resources to register:
import { DataCollectionEndpoint } from "./dataCollectionEndpoint";
import { DataCollectionRule } from "./dataCollectionRule";
import { DataCollectionRuleAssociation } from "./dataCollectionRuleAssociation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:insights/v20210401:DataCollectionEndpoint":
                return new DataCollectionEndpoint(name, <any>undefined, { urn })
            case "azure-native:insights/v20210401:DataCollectionRule":
                return new DataCollectionRule(name, <any>undefined, { urn })
            case "azure-native:insights/v20210401:DataCollectionRuleAssociation":
                return new DataCollectionRuleAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "insights/v20210401", _module)

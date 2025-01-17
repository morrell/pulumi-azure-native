// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const Kind = {
    Sdk: "sdk",
    Designer: "designer",
    Bot: "bot",
    Function: "function",
    Azurebot: "azurebot",
} as const;

/**
 * Required. Gets or sets the Kind of the resource.
 */
export type Kind = (typeof Kind)[keyof typeof Kind];

export const SkuName = {
    F0: "F0",
    S1: "S1",
} as const;

/**
 * The sku name
 */
export type SkuName = (typeof SkuName)[keyof typeof SkuName];

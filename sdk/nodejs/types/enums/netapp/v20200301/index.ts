// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EndpointType = {
    Src: "src",
    Dst: "dst",
} as const;

/**
 * Indicates whether the local volume is the source or destination for the Volume Replication
 */
export type EndpointType = (typeof EndpointType)[keyof typeof EndpointType];

export const ReplicationSchedule = {
    ReplicationSchedule_10minutely: "_10minutely",
    Hourly: "hourly",
    Daily: "daily",
    Weekly: "weekly",
    Monthly: "monthly",
} as const;

/**
 * Schedule
 */
export type ReplicationSchedule = (typeof ReplicationSchedule)[keyof typeof ReplicationSchedule];

export const ServiceLevel = {
    /**
     * Standard service level
     */
    Standard: "Standard",
    /**
     * Premium service level
     */
    Premium: "Premium",
    /**
     * Ultra service level
     */
    Ultra: "Ultra",
} as const;

/**
 * The service level of the file system
 */
export type ServiceLevel = (typeof ServiceLevel)[keyof typeof ServiceLevel];

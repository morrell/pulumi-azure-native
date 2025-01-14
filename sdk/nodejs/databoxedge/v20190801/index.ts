// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./bandwidthSchedule";
export * from "./container";
export * from "./device";
export * from "./fileEventTrigger";
export * from "./getBandwidthSchedule";
export * from "./getContainer";
export * from "./getDevice";
export * from "./getDeviceExtendedInformation";
export * from "./getFileEventTrigger";
export * from "./getIoTRole";
export * from "./getOrder";
export * from "./getPeriodicTimerEventTrigger";
export * from "./getRole";
export * from "./getShare";
export * from "./getStorageAccount";
export * from "./getStorageAccountCredential";
export * from "./getTrigger";
export * from "./getUser";
export * from "./ioTRole";
export * from "./order";
export * from "./periodicTimerEventTrigger";
export * from "./role";
export * from "./share";
export * from "./storageAccount";
export * from "./storageAccountCredential";
export * from "./trigger";
export * from "./user";

// Export enums:
export * from "../../types/enums/databoxedge/v20190801";

// Import resources to register:
import { BandwidthSchedule } from "./bandwidthSchedule";
import { Container } from "./container";
import { Device } from "./device";
import { FileEventTrigger } from "./fileEventTrigger";
import { IoTRole } from "./ioTRole";
import { Order } from "./order";
import { PeriodicTimerEventTrigger } from "./periodicTimerEventTrigger";
import { Role } from "./role";
import { Share } from "./share";
import { StorageAccount } from "./storageAccount";
import { StorageAccountCredential } from "./storageAccountCredential";
import { Trigger } from "./trigger";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:databoxedge/v20190801:BandwidthSchedule":
                return new BandwidthSchedule(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:Container":
                return new Container(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:Device":
                return new Device(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:FileEventTrigger":
                return new FileEventTrigger(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:IoTRole":
                return new IoTRole(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:Order":
                return new Order(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:PeriodicTimerEventTrigger":
                return new PeriodicTimerEventTrigger(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:Role":
                return new Role(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:Share":
                return new Share(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:StorageAccount":
                return new StorageAccount(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:StorageAccountCredential":
                return new StorageAccountCredential(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:Trigger":
                return new Trigger(name, <any>undefined, { urn })
            case "azure-native:databoxedge/v20190801:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "databoxedge/v20190801", _module)

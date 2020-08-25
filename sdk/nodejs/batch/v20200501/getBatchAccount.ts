// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getBatchAccount(args: GetBatchAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetBatchAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:batch/v20200501:getBatchAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetBatchAccountArgs {
    /**
     * The name of the Batch account.
     */
    readonly name: string;
    /**
     * The name of the resource group that contains the Batch account.
     */
    readonly resourceGroupName: string;
}

/**
 * Contains information about an Azure Batch account.
 */
export interface GetBatchAccountResult {
    /**
     * The account endpoint used to interact with the Batch service.
     */
    readonly accountEndpoint: string;
    readonly activeJobAndJobScheduleQuota: number;
    /**
     * Contains information about the auto-storage account associated with a Batch account.
     */
    readonly autoStorage: outputs.batch.v20200501.AutoStoragePropertiesResponse;
    /**
     * For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
     */
    readonly dedicatedCoreQuota: number;
    /**
     * A list of the dedicated core quota per Virtual Machine family for the Batch account. For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
     */
    readonly dedicatedCoreQuotaPerVMFamily: outputs.batch.v20200501.VirtualMachineFamilyCoreQuotaResponse[];
    /**
     * Batch is transitioning its core quota system for dedicated cores to be enforced per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not yet be enforced. If this flag is false, dedicated core quota is enforced via the old dedicatedCoreQuota property on the account and does not consider Virtual Machine family. If this flag is true, dedicated core quota is enforced via the dedicatedCoreQuotaPerVMFamily property on the account, and the old dedicatedCoreQuota does not apply.
     */
    readonly dedicatedCoreQuotaPerVMFamilyEnforced: boolean;
    /**
     * Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using a Microsoft managed key. For additional control, a customer-managed key can be used instead.
     */
    readonly encryption: outputs.batch.v20200501.EncryptionPropertiesResponse;
    /**
     * The identity of the Batch account.
     */
    readonly identity?: outputs.batch.v20200501.BatchAccountIdentityResponse;
    /**
     * Identifies the Azure key vault associated with a Batch account.
     */
    readonly keyVaultReference: outputs.batch.v20200501.KeyVaultReferenceResponse;
    /**
     * The location of the resource.
     */
    readonly location: string;
    /**
     * For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
     */
    readonly lowPriorityCoreQuota: number;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The allocation mode for creating pools in the Batch account.
     */
    readonly poolAllocationMode: string;
    readonly poolQuota: number;
    /**
     * List of private endpoint connections associated with the Batch account
     */
    readonly privateEndpointConnections: outputs.batch.v20200501.PrivateEndpointConnectionResponse[];
    /**
     * The provisioned state of the resource
     */
    readonly provisioningState: string;
    /**
     * If not specified, the default value is 'enabled'.
     */
    readonly publicNetworkAccess: string;
    /**
     * The tags of the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
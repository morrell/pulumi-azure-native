// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TimeSeriesInsights.Outputs
{

    [OutputType]
    public sealed class Gen2StorageConfigurationOutputResponse
    {
        /// <summary>
        /// The name of the storage account that will hold the environment's Gen2 data.
        /// </summary>
        public readonly string AccountName;

        [OutputConstructor]
        private Gen2StorageConfigurationOutputResponse(string accountName)
        {
            AccountName = accountName;
        }
    }
}

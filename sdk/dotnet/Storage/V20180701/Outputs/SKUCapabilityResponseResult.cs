// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20180701.Outputs
{

    [OutputType]
    public sealed class SKUCapabilityResponseResult
    {
        /// <summary>
        /// The name of capability, The capability information in the specified SKU, including file encryption, network ACLs, change notification, etc.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A string value to indicate states of given capability. Possibly 'true' or 'false'.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private SKUCapabilityResponseResult(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NotificationHubs.V20140901.Outputs
{

    [OutputType]
    public sealed class WnsCredentialResponseResult
    {
        /// <summary>
        /// Gets or sets properties of NotificationHub WnsCredential.
        /// </summary>
        public readonly Outputs.WnsCredentialPropertiesResponseResult? Properties;

        [OutputConstructor]
        private WnsCredentialResponseResult(Outputs.WnsCredentialPropertiesResponseResult? properties)
        {
            Properties = properties;
        }
    }
}
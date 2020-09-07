// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class LinkedIntegrationRuntimeKeyResponseResult
    {
        /// <summary>
        /// Type of the secret.
        /// </summary>
        public readonly string AuthorizationType;
        /// <summary>
        /// Type of the secret.
        /// </summary>
        public readonly Outputs.SecureStringResponseResult Key;

        [OutputConstructor]
        private LinkedIntegrationRuntimeKeyResponseResult(
            string authorizationType,

            Outputs.SecureStringResponseResult key)
        {
            AuthorizationType = authorizationType;
            Key = key;
        }
    }
}
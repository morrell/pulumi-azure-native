// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20150615.Outputs
{

    [OutputType]
    public sealed class CustomDomainResponseResult
    {
        /// <summary>
        /// The custom domain name. Name is the CNAME source.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates
        /// </summary>
        public readonly bool? UseSubDomainName;

        [OutputConstructor]
        private CustomDomainResponseResult(
            string name,

            bool? useSubDomainName)
        {
            Name = name;
            UseSubDomainName = useSubDomainName;
        }
    }
}
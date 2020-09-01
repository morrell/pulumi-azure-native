// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppPlatform.Latest.Outputs
{

    [OutputType]
    public sealed class CustomDomainPropertiesResponseResult
    {
        /// <summary>
        /// The app name of domain.
        /// </summary>
        public readonly string AppName;
        /// <summary>
        /// The bound certificate name of domain.
        /// </summary>
        public readonly string? CertName;
        /// <summary>
        /// The thumbprint of bound certificate.
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private CustomDomainPropertiesResponseResult(
            string appName,

            string? certName,

            string? thumbprint)
        {
            AppName = appName;
            CertName = certName;
            Thumbprint = thumbprint;
        }
    }
}
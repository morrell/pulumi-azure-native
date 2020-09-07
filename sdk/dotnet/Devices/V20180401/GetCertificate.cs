// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20180401
{
    public static class GetCertificate
    {
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("azurerm:devices/v20180401:getCertificate", args ?? new GetCertificateArgs(), options.WithVersion());
    }


    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the certificate
        /// </summary>
        [Input("certificateName", required: true)]
        public string CertificateName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the IoT hub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the IoT hub.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetCertificateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// The entity tag.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The name of the certificate.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The description of an X509 CA Certificate.
        /// </summary>
        public readonly Outputs.CertificatePropertiesResponseResult Properties;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCertificateResult(
            string etag,

            string name,

            Outputs.CertificatePropertiesResponseResult properties,

            string type)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
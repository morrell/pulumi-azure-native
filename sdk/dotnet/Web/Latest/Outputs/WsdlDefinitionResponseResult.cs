// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.Latest.Outputs
{

    [OutputType]
    public sealed class WsdlDefinitionResponseResult
    {
        /// <summary>
        /// The WSDL content
        /// </summary>
        public readonly string? Content;
        /// <summary>
        /// The WSDL import method
        /// </summary>
        public readonly string? ImportMethod;
        /// <summary>
        /// The service with name and endpoint names
        /// </summary>
        public readonly Outputs.WsdlServiceResponseResult? Service;
        /// <summary>
        /// The WSDL URL
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private WsdlDefinitionResponseResult(
            string? content,

            string? importMethod,

            Outputs.WsdlServiceResponseResult? service,

            string? url)
        {
            Content = content;
            ImportMethod = importMethod;
            Service = service;
            Url = url;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20180601.Outputs
{

    [OutputType]
    public sealed class CustomDataSourceLinkedServiceResponseResult
    {
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponseResult? ConnectVia;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? Parameters;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private CustomDataSourceLinkedServiceResponseResult(
            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            Outputs.IntegrationRuntimeReferenceResponseResult? connectVia,

            string? description,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            string type)
        {
            Annotations = annotations;
            ConnectVia = connectVia;
            Description = description;
            Parameters = parameters;
            Type = type;
        }
    }
}
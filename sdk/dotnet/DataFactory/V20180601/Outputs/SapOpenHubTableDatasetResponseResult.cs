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
    public sealed class SapOpenHubTableDatasetResponseResult
    {
        /// <summary>
        /// List of tags that can be used for describing the Dataset.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// The ID of request for delta loading. Once it is set, only data with requestId larger than the value of this property will be retrieved. The default value is 0. Type: integer (or Expression with resultType integer ).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? BaseRequestId;
        /// <summary>
        /// Dataset description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Whether to exclude the records of the last request. The default value is true. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? ExcludeLastRequest;
        /// <summary>
        /// The folder that this Dataset is in. If not specified, Dataset will appear at the root level.
        /// </summary>
        public readonly Outputs.DatasetResponseFolderResult? Folder;
        /// <summary>
        /// Linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponseResult LinkedServiceName;
        /// <summary>
        /// The name of the Open Hub Destination with destination type as Database Table. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object> OpenHubDestinationName;
        /// <summary>
        /// Parameters for dataset.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? Parameters;
        /// <summary>
        /// Columns that define the physical type schema of the dataset. Type: array (or Expression with resultType array), itemType: DatasetSchemaDataElement.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Schema;
        /// <summary>
        /// Columns that define the structure of the dataset. Type: array (or Expression with resultType array), itemType: DatasetDataElement.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Structure;
        /// <summary>
        /// Type of dataset.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SapOpenHubTableDatasetResponseResult(
            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            ImmutableDictionary<string, object>? baseRequestId,

            string? description,

            ImmutableDictionary<string, object>? excludeLastRequest,

            Outputs.DatasetResponseFolderResult? folder,

            Outputs.LinkedServiceReferenceResponseResult linkedServiceName,

            ImmutableDictionary<string, object> openHubDestinationName,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            ImmutableDictionary<string, object>? schema,

            ImmutableDictionary<string, object>? structure,

            string type)
        {
            Annotations = annotations;
            BaseRequestId = baseRequestId;
            Description = description;
            ExcludeLastRequest = excludeLastRequest;
            Folder = folder;
            LinkedServiceName = linkedServiceName;
            OpenHubDestinationName = openHubDestinationName;
            Parameters = parameters;
            Schema = schema;
            Structure = structure;
            Type = type;
        }
    }
}
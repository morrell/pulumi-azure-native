// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.Latest.Outputs
{

    [OutputType]
    public sealed class AzureDataLakeStoreDatasetResponseResult
    {
        /// <summary>
        /// List of tags that can be used for describing the Dataset.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// The data compression method used for the item(s) in the Azure Data Lake Store.
        /// </summary>
        public readonly Union<Outputs.DatasetBZip2CompressionResponseResult, Union<Outputs.DatasetDeflateCompressionResponseResult, Union<Outputs.DatasetGZipCompressionResponseResult, Union<Outputs.DatasetTarCompressionResponseResult, Union<Outputs.DatasetTarGZipCompressionResponseResult, Outputs.DatasetZipDeflateCompressionResponseResult>>>>>? Compression;
        /// <summary>
        /// Dataset description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The name of the file in the Azure Data Lake Store. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? FileName;
        /// <summary>
        /// The folder that this Dataset is in. If not specified, Dataset will appear at the root level.
        /// </summary>
        public readonly Outputs.DatasetResponseFolderResult? Folder;
        /// <summary>
        /// Path to the folder in the Azure Data Lake Store. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? FolderPath;
        /// <summary>
        /// The format of the Data Lake Store.
        /// </summary>
        public readonly Union<Outputs.AvroFormatResponseResult, Union<Outputs.JsonFormatResponseResult, Union<Outputs.OrcFormatResponseResult, Union<Outputs.ParquetFormatResponseResult, Outputs.TextFormatResponseResult>>>>? Format;
        /// <summary>
        /// Linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponseResult LinkedServiceName;
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
        private AzureDataLakeStoreDatasetResponseResult(
            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            Union<Outputs.DatasetBZip2CompressionResponseResult, Union<Outputs.DatasetDeflateCompressionResponseResult, Union<Outputs.DatasetGZipCompressionResponseResult, Union<Outputs.DatasetTarCompressionResponseResult, Union<Outputs.DatasetTarGZipCompressionResponseResult, Outputs.DatasetZipDeflateCompressionResponseResult>>>>>? compression,

            string? description,

            ImmutableDictionary<string, object>? fileName,

            Outputs.DatasetResponseFolderResult? folder,

            ImmutableDictionary<string, object>? folderPath,

            Union<Outputs.AvroFormatResponseResult, Union<Outputs.JsonFormatResponseResult, Union<Outputs.OrcFormatResponseResult, Union<Outputs.ParquetFormatResponseResult, Outputs.TextFormatResponseResult>>>>? format,

            Outputs.LinkedServiceReferenceResponseResult linkedServiceName,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponseResult>? parameters,

            ImmutableDictionary<string, object>? schema,

            ImmutableDictionary<string, object>? structure,

            string type)
        {
            Annotations = annotations;
            Compression = compression;
            Description = description;
            FileName = fileName;
            Folder = folder;
            FolderPath = folderPath;
            Format = format;
            LinkedServiceName = linkedServiceName;
            Parameters = parameters;
            Schema = schema;
            Structure = structure;
            Type = type;
        }
    }
}
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
    public sealed class GoogleCloudStorageLocationResponseResult
    {
        /// <summary>
        /// Specify the bucketName of Google Cloud Storage. Type: string (or Expression with resultType string)
        /// </summary>
        public readonly ImmutableDictionary<string, object>? BucketName;
        /// <summary>
        /// Specify the file name of dataset. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? FileName;
        /// <summary>
        /// Specify the folder path of dataset. Type: string (or Expression with resultType string)
        /// </summary>
        public readonly ImmutableDictionary<string, object>? FolderPath;
        /// <summary>
        /// Type of dataset storage location.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specify the version of Google Cloud Storage. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Version;

        [OutputConstructor]
        private GoogleCloudStorageLocationResponseResult(
            ImmutableDictionary<string, object>? bucketName,

            ImmutableDictionary<string, object>? fileName,

            ImmutableDictionary<string, object>? folderPath,

            string type,

            ImmutableDictionary<string, object>? version)
        {
            BucketName = bucketName;
            FileName = fileName;
            FolderPath = folderPath;
            Type = type;
            Version = version;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataShare.Latest.Outputs
{

    [OutputType]
    public sealed class SynchronizationDetailsResponseResult
    {
        /// <summary>
        /// Id of data set
        /// </summary>
        public readonly string DataSetId;
        /// <summary>
        /// Type of the data set
        /// </summary>
        public readonly string DataSetType;
        /// <summary>
        /// Duration of data set level copy
        /// </summary>
        public readonly int DurationMs;
        /// <summary>
        /// End time of data set level copy
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The number of files read from the source data set
        /// </summary>
        public readonly int FilesRead;
        /// <summary>
        /// The number of files written into the sink data set
        /// </summary>
        public readonly int FilesWritten;
        /// <summary>
        /// Error message if any
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Name of the data set
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The number of files copied into the sink data set
        /// </summary>
        public readonly int RowsCopied;
        /// <summary>
        /// The number of rows read from the source data set.
        /// </summary>
        public readonly int RowsRead;
        /// <summary>
        /// The size of the data read from the source data set in bytes
        /// </summary>
        public readonly int SizeRead;
        /// <summary>
        /// The size of the data written into the sink data set in bytes
        /// </summary>
        public readonly int SizeWritten;
        /// <summary>
        /// Start time of data set level copy
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Raw Status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The vCore units consumed for the data set synchronization
        /// </summary>
        public readonly int VCore;

        [OutputConstructor]
        private SynchronizationDetailsResponseResult(
            string dataSetId,

            string dataSetType,

            int durationMs,

            string endTime,

            int filesRead,

            int filesWritten,

            string message,

            string name,

            int rowsCopied,

            int rowsRead,

            int sizeRead,

            int sizeWritten,

            string startTime,

            string status,

            int vCore)
        {
            DataSetId = dataSetId;
            DataSetType = dataSetType;
            DurationMs = durationMs;
            EndTime = endTime;
            FilesRead = filesRead;
            FilesWritten = filesWritten;
            Message = message;
            Name = name;
            RowsCopied = rowsCopied;
            RowsRead = rowsRead;
            SizeRead = sizeRead;
            SizeWritten = sizeWritten;
            StartTime = startTime;
            Status = status;
            VCore = vCore;
        }
    }
}
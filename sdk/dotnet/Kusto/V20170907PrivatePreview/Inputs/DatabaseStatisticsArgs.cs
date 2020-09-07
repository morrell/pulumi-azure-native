// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20170907PrivatePreview.Inputs
{

    public sealed class DatabaseStatisticsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database size - the total size of compressed data and index in bytes.
        /// </summary>
        [Input("size")]
        public Input<double>? Size { get; set; }

        public DatabaseStatisticsArgs()
        {
        }
    }
}
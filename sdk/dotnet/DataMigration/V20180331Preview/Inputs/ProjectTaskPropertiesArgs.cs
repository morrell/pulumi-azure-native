// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180331Preview.Inputs
{

    /// <summary>
    /// Base class for all types of DMS task properties. If task is not supported by current client, this object is returned.
    /// </summary>
    public sealed class ProjectTaskPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Task type.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        public ProjectTaskPropertiesArgs()
        {
        }
    }
}
// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class ResourceRequestsResponseResult
    {
        /// <summary>
        /// Requested number of CPU cores. At present, only full cores are supported.
        /// </summary>
        public readonly double Cpu;
        /// <summary>
        /// The memory request in GB for this container.
        /// </summary>
        public readonly double MemoryInGB;

        [OutputConstructor]
        private ResourceRequestsResponseResult(
            double cpu,

            double memoryInGB)
        {
            Cpu = cpu;
            MemoryInGB = memoryInGB;
        }
    }
}
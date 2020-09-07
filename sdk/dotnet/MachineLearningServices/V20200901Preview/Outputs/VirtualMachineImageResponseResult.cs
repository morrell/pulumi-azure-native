// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200901Preview.Outputs
{

    [OutputType]
    public sealed class VirtualMachineImageResponseResult
    {
        /// <summary>
        /// Virtual Machine image path
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private VirtualMachineImageResponseResult(string id)
        {
            Id = id;
        }
    }
}
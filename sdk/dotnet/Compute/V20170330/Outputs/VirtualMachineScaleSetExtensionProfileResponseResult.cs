// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20170330.Outputs
{

    [OutputType]
    public sealed class VirtualMachineScaleSetExtensionProfileResponseResult
    {
        /// <summary>
        /// The virtual machine scale set child extension resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineScaleSetExtensionResponseResult> Extensions;

        [OutputConstructor]
        private VirtualMachineScaleSetExtensionProfileResponseResult(ImmutableArray<Outputs.VirtualMachineScaleSetExtensionResponseResult> extensions)
        {
            Extensions = extensions;
        }
    }
}